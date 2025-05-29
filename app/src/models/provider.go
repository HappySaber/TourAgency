package models

import "github.com/google/uuid"

type Provider struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"` // UUID идентификатор
	Name        string    `json:"name" gorm:"type:varchar(32);not null"`                      // Название провайдера
	Addressto   string    `json:"addressto" gorm:"type:varchar(64);"`                         // Адрес назначения
	Address     string    `json:"address" gorm:"type:varchar(255);not null"`                  // Адрес
	Email       string    `json:"email" gorm:"type:varchar(64);"`                             // Электронная почта
	PhoneNumber string    `json:"phonenumber" gorm:"type:varchar(16);"`                       // Номер телефона
}
