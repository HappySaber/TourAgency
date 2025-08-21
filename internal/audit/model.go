package audit

import "time"

type AuditLog struct {
	ID            uint      `gorm:"primaryKey"`
	EventID       string    `gorm:"uniqueIndex;size:64"`
	Event         string    `gorm:"index;size:64"`
	Entity        string    `gorm:"index;size:64"`
	EntityID      string    `gorm:"index;size:64"`
	ActorID       string    `gorm:"size:64"`
	CorrelationID string    `gorm:"size:64"`
	IP            string    `gorm:"size:64"`
	UserAgent     string    `gorm:"size:255"`
	At            time.Time `gorm:"index"`
	BeforeJSON    string    `gorm:"type:text"`
	AfterJSON     string    `gorm:"type:text"`
	CreatedAt     time.Time
}
