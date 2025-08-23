package services

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"context"
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type ServService struct {
	db *gorm.DB
}

func NewServService(db *gorm.DB) *ServService {
	return &ServService{db: db}
}

func (ss *ServService) GetAll() ([]models.Service, error) {
	var services []models.Service
	err := ss.db.Find(&services).Error
	return services, err
}

func (ss *ServService) GetByID(id int) (*models.Service, error) {
	var service models.Service
	err := ss.db.First(&service, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &service, err
}

func (ss *ServService) Create(ctx context.Context, service *models.Service) (*audit.Event, error) {
	if err := ss.db.WithContext(ctx).Create(service).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "service.created",
		Entity:   "service",
		EntityID: strconv.FormatUint(uint64(service.ID), 10),
		At:       time.Now(),
		Before:   audit.MustMarshal(nil),
		After:    audit.MustMarshal(service),
	}
	return evt, nil
}

func (ss *ServService) Update(ctx context.Context, service *models.Service) (*audit.Event, error) {
	var before models.Service
	if err := ss.db.WithContext(ctx).First(&before, "id = ?", service.ID).Error; err != nil {
		return nil, err
	}
	if err := ss.db.WithContext(ctx).Save(service).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "service.updated",
		Entity:   "service",
		EntityID: strconv.FormatUint(uint64(service.ID), 10),
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(service),
	}
	return evt, nil
}

func (ss *ServService) Delete(ctx context.Context, id int) (*audit.Event, error) {
	var before models.Service
	if err := ss.db.WithContext(ctx).First(&before, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := ss.db.WithContext(ctx).Delete(&models.Service{}, id).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "service.deleted",
		Entity:   "service",
		EntityID: strconv.Itoa(id),
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(nil),
	}
	return evt, nil
}
