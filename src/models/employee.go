package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Employee представляет структуру для таблицы employees
type Employee struct {
	ID           uuid.UUID `json:"id" gorm:"column:id;type:uuid;default:uuid_generate_v4();primary_key"` // UUID идентификатор
	Email        string    `json:"email" gorm:"column:email;type:varchar(100);unique;not null"`          // Почта
	FirstName    string    `json:"firstname" gorm:"column:firstname;type:varchar(32);not null"`          // Имя
	LastName     string    `json:"lastname" gorm:"column:lastname;type:varchar(32);not null"`            // Фамилия
	MiddleName   string    `json:"middlename,omitempty" gorm:"column:middlename;type:varchar(32)"`       // Отчество
	Address      string    `json:"address" gorm:"column:address;type:varchar(255);not null"`             // Адрес
	PhoneNumber  string    `json:"phonenumber,omitempty" gorm:"column:phonenumber;type:varchar(10)"`     // Номер телефона
	DateOfBirth  time.Time `json:"dateofbirth,omitempty" gorm:"column:dateofbirth;type:timestamp"`       // Дата рождения
	DateOfHiring time.Time `json:"dateofhiring,omitempty" gorm:"column:dateofhiring;type:timestamp"`     // Дата приема на работу
	PositionID   uuid.UUID `json:"position,omitempty" gorm:"column:position;type:uuid"`                  // Идентификатор должности
	Password     string    `json:"password" gorm:"column:password;type:varchar(100);not null"`           // Пароль
}

type EmployeeRequest struct {
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"password" gorm:"-"`
}

type Claims struct {
	Role string `json:"Role"`
	jwt.RegisteredClaims
}
