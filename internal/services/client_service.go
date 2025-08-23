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

type ClientService struct {
	db *gorm.DB
}

func NewClientService(db *gorm.DB) *ClientService {
	return &ClientService{
		db: db,
	}
}

func (cs *ClientService) Create(ctx context.Context, client *models.Client) (*audit.Event, error) {
	client.ID = uuid.New()
	if err := cs.db.WithContext(ctx).Create(client).Error; err != nil {
		return nil, err
	}

	evt := &audit.Event{
		Event:    "client.created",
		Entity:   "client",
		EntityID: client.ID.String(),
		At:       time.Now(),
		After:    audit.MustMarshal(client),
	}
	return evt, nil
}

func (cs *ClientService) GetAll() ([]models.Client, error) {
	var Client []models.Client
	err := cs.db.Find(&Client).Error
	return Client, err
}

func (cs *ClientService) GetByID(id string) (*models.Client, error) {
	var Client models.Client
	err := cs.db.First(&Client, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &Client, err
}

func (cs *ClientService) Update(ctx context.Context, updated *models.Client) (*audit.Event, error) {
	var before models.Client
	if err := cs.db.WithContext(ctx).First(&before, "id = ?", updated.ID).Error; err != nil {
		return nil, err
	}

	if err := cs.db.WithContext(ctx).Save(updated).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "client.updated",
		Entity:   "client",
		EntityID: updated.ID.String(),
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(updated),
	}
	return evt, nil
}

func (cs *ClientService) Delete(ctx context.Context, id string) (*audit.Event, error) {
	var before models.Client
	if err := cs.db.WithContext(ctx).First(&before, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := cs.db.WithContext(ctx).Delete(&models.Client{}, "id = ?", id).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "client.deleted",
		Entity:   "client",
		EntityID: id,
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(nil),
	}
	return evt, nil
}
