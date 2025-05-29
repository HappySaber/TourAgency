package models

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Firstname   string    `gorm:"type:varchar(32);not null" json:"firstname"`
	Lastname    string    `gorm:"type:varchar(32);not null" json:"lastname"`
	Middlename  string    `gorm:"type:varchar(32)" json:"middlename,omitempty"`
	Address     string    `gorm:"type:varchar(255);not null" json:"address"`
	Phonenumber string    `gorm:"type:varchar(10)" json:"phonenumber,omitempty"`
	Dateofbirth time.Time `gorm:"type:date" json:"dateofbirth"`
}
