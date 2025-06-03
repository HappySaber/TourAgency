package models

import "github.com/google/uuid"

type TourPerConsultation struct {
	TourID         uuid.UUID `json:"tour_id" gorm:"type:uuid;primary_key;column:tour_id"`                 // Идентификатор тура
	ConsultationID uuid.UUID `json:"consultation_id" gorm:"type:uuid;primary_key;column:consultation_id"` // Идентификатор консультации
	Discount       string    `json:"discount" gorm:"type:varchar(64);column:discount"`                    // Скидка на услугу
	Quantity       string    `json:"quantity" gorm:"type:varchar(64);column:quanity"`                     // Количество услуги
}

func (TourPerConsultation) TableName() string {
	return "tours_per_consultation"
}
