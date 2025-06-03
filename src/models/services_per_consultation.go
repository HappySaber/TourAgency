package models

import "github.com/google/uuid"

// ServicePerConsultation представляет структуру для таблицы services_per_consultation
type ServicePerConsultation struct {
	ServiceID      uint      `json:"service_id" gorm:"primary_key;column:service_id"`           // Идентификатор услуги
	ConsultationID uuid.UUID `json:"consultation_id" gorm:"primary_key;column:consultation_id"` // Идентификатор консультации
	Discount       string    `json:"discount" gorm:"type:varchar(64);column:discount"`          // Скидка на услугу
	Quantity       string    `json:"quantity" gorm:"type:varchar(64);column:quanity"`           // Количество услуги
}

func (ServicePerConsultation) TableName() string {
	return "services_per_consultation"
}
