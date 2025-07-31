package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type LocalTime struct {
	time.Time
}

// Парсит из БД
func (lt *LocalTime) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot convert %v to time", value)
	}
	t, err := time.Parse("15:04:05", str)
	if err != nil {
		return err
	}
	lt.Time = t
	return nil
}

// Для вставки в БД
func (lt LocalTime) Value() (driver.Value, error) {
	return lt.Format("15:04:05"), nil
}

// Для JSON
func (lt LocalTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", lt.Format("15:04"))), nil
}

func (lt *LocalTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("15:04", s)
	if err != nil {
		return err
	}
	lt.Time = t
	return nil
}
