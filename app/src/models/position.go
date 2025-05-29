package models

import "github.com/google/uuid"

// Position представляет структуру для таблицы positions
type Position struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"` // UUID идентификатор
	Name             string    `json:"name" gorm:"type:varchar(32);not null"`                      // Название должности
	Salary           string    `json:"salary" gorm:"type:varchar(64);"`                            // Зарплата
	Responsibilities string    `json:"responsibilities" gorm:"type:varchar(255);"`                 // Обязанности
}
