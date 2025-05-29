package models

import "github.com/google/uuid"

// Tour представляет структуру для таблицы tours
type Tour struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"` // UUID идентификатор
	Name       string    `json:"name" gorm:"type:varchar(32);not null"`                      // Название тура
	Rating     string    `json:"rating" gorm:"type:varchar(8);"`                             // Рейтинг
	Hotel      string    `json:"hotel" gorm:"type:varchar(64);"`                             // Отель
	Nutrition  string    `json:"nutrition" gorm:"type:varchar(64);"`                         // Питание
	City       string    `json:"city" gorm:"type:varchar(64);"`                              // Город
	Country    string    `json:"country" gorm:"type:varchar(64);"`                           // Страна
	ProviderID uuid.UUID `json:"provider" gorm:"type:uuid;"`                                 // Идентификатор провайдера
}
