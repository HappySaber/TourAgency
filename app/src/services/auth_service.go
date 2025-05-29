package services

import (
	"TurAgency/src/models"
	"TurAgency/src/utils"
	"errors"
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var jwtKey = []byte(os.Getenv("JWTKEY"))

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (as *AuthService) Login(userReq models.EmployeeRequest) (string, error) {
	var user models.Employee
	if err := as.db.Where("email = ?", userReq.Email).First(&user).Error; err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CompareHashPassword(userReq.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	var position models.Position
	if err := as.db.First(&position, "id = ?", user.PositionID).Error; err != nil {
		return "", errors.New("failed to fetch user role")
	}

	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &models.Claims{
		Role: position.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("could not create token")
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
