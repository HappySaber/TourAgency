package audit

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	SaveEvent(ctx context.Context, evt Event) error
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (gr *GormRepository) SaveEvent(ctx context.Context, evt Event) error {
	auditLog := AuditLog{
		EventID:       evt.EventID,
		Event:         evt.Event,
		Entity:        evt.Entity,
		EntityID:      evt.EntityID,
		ActorID:       evt.ActorID,
		CorrelationID: evt.CorrelationID,
		IP:            evt.IP,
		UserAgent:     evt.UserAgent,
		At:            evt.At,
		BeforeJSON:    string(evt.Before),
		AfterJSON:     string(evt.After),
	}
	return gr.db.WithContext(ctx).Create(&auditLog).Error
}
