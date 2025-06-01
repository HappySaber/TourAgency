package models

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();column:id" json:"id"`
	Firstname   string    `gorm:"type:varchar(32);not null;column:firstname" json:"firstname"`
	Lastname    string    `gorm:"type:varchar(32);not null;column:lastname" json:"lastname"`
	Middlename  string    `gorm:"type:varchar(32);column:middlename" json:"middlename,omitempty"`
	Address     string    `gorm:"type:varchar(255);not null;column:address" json:"address"`
	Phonenumber string    `gorm:"type:varchar(10);column:phonenumber" json:"phonenumber,omitempty"`
	Dateofbirth time.Time `gorm:"type:date;column:dateofbirth" json:"dateofbirth"`
}
