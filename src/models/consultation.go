package models

import (
	"time"

	"github.com/google/uuid"
)

// Consultation представляет структуру для таблицы consultations
type Consultation struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"` // UUID идентификатор
	DateOfConsultation time.Time `json:"dateofconsultation" gorm:"type:date;not null"`               // Дата консультации
	DateOf             time.Time `json:"dateof" gorm:"type:date;not null"`                           // Дата
	ClientID           uuid.UUID `json:"client" gorm:"type:uuid;not null"`                           // Идентификатор клиента
	EmployeeID         uuid.UUID `json:"employee" gorm:"type:uuid;not null"`                         // Идентификатор сотрудника
}
