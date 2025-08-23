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

type ProviderService struct {
	db *gorm.DB
}

func NewProviderService(db *gorm.DB) *ProviderService {
	return &ProviderService{db: db}
}

func (ps *ProviderService) GetAll() ([]models.Provider, error) {
	var providers []models.Provider
	err := ps.db.Find(&providers).Error
	return providers, err
}

func (ps *ProviderService) GetByID(id string) (*models.Provider, error) {
	var provider models.Provider
	err := ps.db.First(&provider, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &provider, err
}

func (ps *ProviderService) Create(ctx context.Context, provider *models.Provider) (*audit.Event, error) {
	provider.ID = uuid.New()
	if err := ps.db.WithContext(ctx).Create(provider).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "provider.created",
		Entity:   "provider",
		EntityID: provider.ID.String(),
		At:       time.Now(),
		Before:   audit.MustMarshal(nil),
		After:    audit.MustMarshal(provider),
	}
	return evt, nil
}

func (ps *ProviderService) Update(ctx context.Context, provider *models.Provider) (*audit.Event, error) {
	var before models.Provider
	if err := ps.db.WithContext(ctx).First(&before, "id = ?", provider.ID).Error; err != nil {
		return nil, err
	}

	if err := ps.db.WithContext(ctx).Save(provider).Error; err != nil {
		return nil, err
	}

	evt := &audit.Event{
		Event:    "provider.updated",
		Entity:   "provider",
		EntityID: provider.ID.String(),
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(provider),
	}
	return evt, nil
}

func (ps *ProviderService) Delete(ctx context.Context, id string) (*audit.Event, error) {
	var before models.Provider
	if err := ps.db.WithContext(ctx).First(&before, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := ps.db.WithContext(ctx).Delete(&models.Provider{}, "id = ?", id).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "provider.deleted",
		Entity:   "provider",
		EntityID: id,
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(nil),
	}
	return evt, nil
}
