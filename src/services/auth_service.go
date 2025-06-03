package services

import (
	"TurAgency/src/models"
	"TurAgency/src/utils"
	"errors"
	"log"
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
func (as *AuthService) Login(userReq models.EmployeeRequest) (string, error) {

	var user models.Employee

	// Поиск пользователя по email
	if err := as.db.Where("email = ?", userReq.Email).First(&user).Error; err != nil {
		return "", errors.New("Неверный email или пароль")
	}

	// Проверка пароля
	if !utils.CompareHashPassword(userReq.Password, user.Password) {
		return "", errors.New("Неверный email или пароль")
	}

	// Получение роли
	var position models.Position
	if err := as.db.First(&position, "id = ?", user.PositionID).Error; err != nil {
		return "", errors.New("Не удалось получить должность пользователя")
	}

	// Создание JWT-токена
	expirationTime := time.Now().Add(8 * time.Hour)
	claims := &models.Claims{
		Role: position.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	log.Println(claims.Role)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(as.jwtKey)
	if err != nil {
		return "", errors.New("Не удалось создать токен")
	}

	return tokenString, nil
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

func (as *AuthService) GetPositions() ([]models.Position, error) {
	var positions []models.Position
	if err := as.db.Find(&positions).Error; err != nil {
		return nil, err
	}
	return positions, nil
}
