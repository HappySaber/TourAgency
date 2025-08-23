package services

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PositionService struct {
	db *gorm.DB
}

func NewPositionService(db *gorm.DB) *PositionService {
	return &PositionService{db: db}
}

func (ps *PositionService) Create(ctx context.Context, position *models.Position) (*audit.Event, error) {
	position.ID = uuid.New()
	if err := ps.db.WithContext(ctx).Create(position).Error; err != nil {
		return nil, err
	}

	evt := &audit.Event{
		Event:    "position.created",
		Entity:   "position",
		EntityID: position.ID.String(),
		At:       time.Now(),
		Before:   audit.MustMarshal(nil),
		After:    audit.MustMarshal(position),
	}
	return evt, nil

}

func (ps *PositionService) GetAll() ([]models.Position, error) {
	var Position []models.Position
	err := ps.db.Find(&Position).Error
	return Position, err
}

func (ps *PositionService) GetByID(id string) (*models.Position, error) {
	var Position models.Position
	err := ps.db.First(&Position, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &Position, err
}

func (ps *PositionService) Update(ctx context.Context, updated *models.Position) (*audit.Event, error) {
	var before models.Position
	if err := ps.db.WithContext(ctx).First(&before, "id = ?", updated.ID).Error; err != nil {
		return nil, err
	}

	if err := ps.db.WithContext(ctx).Save(updated).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "position.updated",
		Entity:   "position",
		EntityID: updated.ID.String(),
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(updated),
	}
	return evt, nil
}

func (ps *PositionService) Delete(ctx context.Context, id string) (*audit.Event, error) {
	var before models.Employee
	if err := ps.db.WithContext(ctx).First(&before, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := ps.db.WithContext(ctx).Delete(&models.Position{}, "id = ?", id).Error; err != nil {
		return nil, err
	}

	evt := &audit.Event{
		Event:    "position.deleted",
		Entity:   "position",
		EntityID: id,
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(nil),
	}
	return evt, nil
}
