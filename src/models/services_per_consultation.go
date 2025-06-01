package models

import "github.com/google/uuid"

// ServicePerConsultation представляет структуру для таблицы services_per_consultation
type ServicePerConsultation struct {
	ServiceID      uint      `json:"service_id" gorm:"primary_key;column:service_id"`           // Идентификатор услуги
	ConsultationID uuid.UUID `json:"consultation_id" gorm:"primary_key;column:consultation_id"` // Идентификатор консультации
}
