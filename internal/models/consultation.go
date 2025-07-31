package models

import (
	"time"

	"github.com/google/uuid"
)

// Consultation представляет структуру для таблицы consultations
type Consultation struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	DateOfConsultation time.Time `json:"dateofconsultation" gorm:"type:date;not null;column:dateofconsultation"`
	TimeOfConsultation LocalTime `json:"timeofconsultation" gorm:"type:time;not null;column:timeofconsultation"`

	ClientID uuid.UUID `json:"client" gorm:"type:uuid;not null;column:client"`
	Client   Client    `gorm:"foreignKey:ClientID;references:ID"`

	EmployeeID uuid.UUID `json:"employee" gorm:"type:uuid;not null;column:employee"`
	Employee   Employee  `gorm:"foreignKey:EmployeeID;references:ID"`

	Notes string `json:"notes" gorm:"type:varchar(512);column:notes"`
}
