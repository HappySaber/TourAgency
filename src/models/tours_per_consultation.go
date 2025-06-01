package models

import "github.com/google/uuid"

type TourPerConsultation struct {
	TourID         uuid.UUID `json:"tour_id" gorm:"type:uuid;primary_key;column:tour_id"`                 // Идентификатор тура
	ConsultationID uuid.UUID `json:"consultation_id" gorm:"type:uuid;primary_key;column:consultation_id"` // Идентификатор консультации
}
