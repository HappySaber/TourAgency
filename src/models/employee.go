package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Employee представляет структуру для таблицы employees
type Employee struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"` // UUID идентификатор
	Email        string    `json:"email" gorm:"type:varchar(100);unique;not null"`             // Почта
	FirstName    string    `json:"firstname" gorm:"type:varchar(32);not null"`                 // Имя
	LastName     string    `json:"lastname" gorm:"type:varchar(32);not null"`                  // Фамилия
	MiddleName   string    `json:"middlename,omitempty" gorm:"type:varchar(32);"`              // Отчество
	Address      string    `json:"address" gorm:"type:varchar(255);not null"`                  // Адрес
	PhoneNumber  string    `json:"phonenumber,omitempty" gorm:"type:varchar(10);"`             // Номер телефона
	DateOfBirth  time.Time `json:"dateofbirth,omitempty" gorm:"type:date"`                     // Дата рождения
	DateOfHiring time.Time `json:"dateofhiring,omitempty" gorm:"type:date"`                    // Дата приема на работу
	PositionID   uuid.UUID `json:"position,omitempty" gorm:"type:uuid"`                        // Идентификатор должности
	Password     string    `json:"password" gorm:"type:varchar(16);not null"`
}

type EmployeeRequest struct {
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"password" gorm:"-"`
}

type Claims struct {
	Role string `json:"Role"`
	jwt.RegisteredClaims
}
