package models

import (
	"time"

	"github.com/google/uuid"
)

// Consultation представляет структуру для таблицы consultations
type Consultation struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:id"`   // UUID идентификатор
	DateOfConsultation time.Time `json:"dateofconsultation" gorm:"type:date;not null;column:dateofconsultation"` // Дата консультации
	TimeOfConsultation time.Time `json:"dateof" gorm:"type:time;not null;column:timeofconsultation"`             // Дата
	ClientID           uuid.UUID `json:"client" gorm:"type:uuid;not null;column:client"`                         // Идентификатор клиента
	EmployeeID         uuid.UUID `json:"employee" gorm:"type:uuid;not null;column:employee"`                     // Идентификатор сотрудника
}
