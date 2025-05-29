package models

import "github.com/google/uuid"

type TourPerConsultation struct {
	TourID         uuid.UUID `json:"tour_id" gorm:"type:uuid;primary_key"`         // Идентификатор тура
	ConsultationID uuid.UUID `json:"consultation_id" gorm:"type:uuid;primary_key"` // Идентификатор консультации
}
