package services

import (
	"TurAgency/internal/database"
	"TurAgency/internal/models"
	"TurAgency/internal/utils"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct {
	db     *gorm.DB
	jwtKey []byte
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db:     db,
		jwtKey: []byte(os.Getenv("JWTKEY")),
	}
}

func (as *AuthService) Login(userReq models.EmployeeRequest) (*models.TokenPair, error) {

	var user models.Employee

	// Поиск пользователя по email
	if err := as.db.Where("email = ?", userReq.Email).First(&user).Error; err != nil {
		return nil, errors.New("Неверный email или пароль")
	}

	// Проверка пароля
	if !utils.CompareHashPassword(userReq.Password, user.Password) {
		return nil, errors.New("Неверный email или пароль")
	}

	// Получение роли
	var position models.Position
	if err := as.db.First(&position, "id = ?", user.PositionID).Error; err != nil {
		return nil, errors.New("Не удалось получить должность пользователя")
	}

	accessExp := time.Now().Add(30 * time.Minute)
	accessClaims := &models.Claims{
		Role: position.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(accessExp),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenStr, err := accessToken.SignedString(as.jwtKey)
	if err != nil {
		return nil, errors.New("Не удалось создать access токен")
	}

	// Refresh token (живет дольше, напр. 7 дней)
	refreshExp := time.Now().Add(7 * 24 * time.Hour)
	refreshClaims := &models.Claims{
		Role: position.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(refreshExp),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenStr, err := refreshToken.SignedString(as.jwtKey)
	if err != nil {
		return nil, errors.New("Не удалось создать refresh токен")
	}

	// Сохраняем refresh в Redis
	err = database.RedisDB.Set(
		database.Ctx,
		refreshTokenStr,
		"valid",
		time.Until(refreshExp),
	).Err()
	if err != nil {
		return nil, fmt.Errorf("Ошибка сохранения refresh токена в Redis: %w", err)
	}

	return &models.TokenPair{
		AccessToken:  accessTokenStr,
		RefreshToken: refreshTokenStr,
	}, nil
}

func (as *AuthService) Signup(user *models.Employee) error {
	if !govalidator.IsEmail(user.Email) {
		return errors.New("invalid email address")
	}

	var position models.Position
	if err := as.db.First(&position, "id = ?", user.PositionID).Error; err != nil {
		return errors.New("such role doesn't exist")
	}

	var existing models.Employee
	if err := as.db.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		return errors.New("user with this email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("database error")
	}

	hashedPassword, err := utils.GenerateHashPassword(user.Password)
	if err != nil {
		return errors.New("could not hash password")
	}
	user.Password = hashedPassword
	user.ID = uuid.New()
	user.DateOfHiring = time.Now()

	if err := as.db.Create(user).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}
func (as *AuthService) Logout(tokenString string) error {
	return database.RedisDB.Del(database.Ctx, tokenString).Err()
}

func (as *AuthService) GetPositions() ([]models.Position, error) {
	var positions []models.Position
	if err := as.db.Find(&positions).Error; err != nil {
		return nil, err
	}
	return positions, nil
}

func (as *AuthService) ValidateToken(tokenString string) (*models.Claims, error) {
	val, err := database.RedisDB.Get(database.Ctx, tokenString).Result()
	if err != nil || val != "valid" {
		return nil, errors.New("токен не найден или отозван")
	}

	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return as.jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("невалидный токен")
	}

	return claims, nil
}

func (as *AuthService) GenerateTokens(user models.Employee) (*models.TokenPair, error) {
	// Access Token
	accessExp := time.Now().Add(30 * time.Minute)
	accessClaims := &models.Claims{
		UserID: user.ID,
		Role:   user.Position.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExp),
			Subject:   user.ID.String(),
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(as.jwtKey)
	if err != nil {
		return nil, err
	}

	// Refresh Token
	refreshExp := time.Now().Add(7 * 24 * time.Hour)
	refreshClaims := &models.Claims{
		UserID: user.ID,
		Role:   user.Position.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExp),
			Subject:   user.ID.String(),
		},
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(as.jwtKey)
	if err != nil {
		return nil, err
	}

	// Сохраняем refresh token в Redis
	key := fmt.Sprintf("refresh:%s", user.ID.String())
	if err := database.RedisDB.Set(database.Ctx, key, refreshToken, 7*24*time.Hour).Err(); err != nil {
		return nil, fmt.Errorf("ошибка сохранения токена в Redis: %w", err)
	}

	return &models.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
