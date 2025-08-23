package services

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type TourService struct {
	db *gorm.DB
}

func NewTourService(db *gorm.DB) *TourService {
	return &TourService{db: db}
}

func (ts *TourService) GetAll() ([]models.Tour, error) {
	var tours []models.Tour
	err := ts.db.Find(&tours).Error
	return tours, err
}

func (ts *TourService) GetByID(id string) (*models.Tour, error) {
	var tour models.Tour
	err := ts.db.First(&tour, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &tour, err
}

func (ts *TourService) Create(ctx context.Context, tour *models.Tour) (*audit.Event, error) {
	if err := ts.db.WithContext(ctx).Create(tour).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "tour.created",
		Entity:   "tour",
		EntityID: tour.ID.String(),
		At:       time.Now(),
		Before:   audit.MustMarshal(nil),
		After:    audit.MustMarshal(tour),
	}
	return evt, nil
}

func (ts *TourService) Update(ctx context.Context, tour *models.Tour) (*audit.Event, error) {
	var before models.Tour
	if err := ts.db.WithContext(ctx).First(&before, "id = ?", tour.ID).Error; err != nil {
		return nil, err
	}
	if err := ts.db.WithContext(ctx).Save(tour).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "tour.updated",
		Entity:   "tour",
		EntityID: tour.ID.String(),
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(tour),
	}
	return evt, nil
}

func (ts *TourService) Delete(ctx context.Context, id string) (*audit.Event, error) {
	var before models.Tour
	if err := ts.db.WithContext(ctx).First(&before, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := ts.db.WithContext(ctx).Delete(&models.Tour{}, "id = ?", id).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "tour.deleted",
		Entity:   "tour",
		EntityID: id,
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(nil),
	}
	return evt, nil
}

func (ts *TourService) GetProviders() ([]models.Provider, error) {
	var providers []models.Provider
	if err := ts.db.Find(&providers).Error; err != nil {
		return nil, err
	}
	return providers, nil
}
