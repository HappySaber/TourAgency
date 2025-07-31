package services

import (
	"TurAgency/internal/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientService struct {
	db *gorm.DB
}

func NewClientService(db *gorm.DB) *ClientService {
	return &ClientService{db: db}
}

func (cs *ClientService) Create(client *models.Client) error {
	client.ID = uuid.New()
	return cs.db.Create(client).Error
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

func (cs *ClientService) Update(updated *models.Client) error {
	return cs.db.Save(updated).Error
}

func (cs *ClientService) Delete(id string) error {
	return cs.db.Delete(&models.Client{}, "id = ?", id).Error
}
