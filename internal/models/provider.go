package models

import "github.com/google/uuid"

type Provider struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:id"` // UUID идентификатор
	Name        string    `json:"name" gorm:"type:varchar(32);not null;column:name"`                    // Название провайдера
	Addressto   string    `json:"addressto" gorm:"type:varchar(64);column:addressto"`                   // Адрес назначения
	Address     string    `json:"address" gorm:"type:varchar(255);not null;column:address"`             // Адрес
	Email       string    `json:"email" gorm:"type:varchar(64);column:email"`                           // Электронная почта
	PhoneNumber string    `json:"phonenumber" gorm:"type:varchar(16);column:phonenumber"`               // Номер телефона
}
