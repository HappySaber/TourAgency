package models

import "github.com/google/uuid"

// Tour представляет структуру для таблицы tours
type Tour struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:id"` // UUID идентификатор
	Name       string    `json:"name" gorm:"type:varchar(32);not null;column:name"`                    // Название тура
	Rating     string    `json:"rating" gorm:"type:varchar(8);column:rating"`                          // Рейтинг
	Hotel      string    `json:"hotel" gorm:"type:varchar(64);column:hotel"`                           // Отель
	Nutrition  string    `json:"nutrition" gorm:"type:varchar(64);column:nutrition"`                   // Питание
	City       string    `json:"city" gorm:"type:varchar(64);column:city"`                             // Город
	Country    string    `json:"country" gorm:"type:varchar(64);column:country"`                       // Страна
	ProviderID uuid.UUID `json:"provider" gorm:"type:uuid;column:providerid"`                          // Идентификатор провайдера
}
