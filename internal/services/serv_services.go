package services

import (
	"TurAgency/internal/models"
	"errors"

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

func (ss *ServService) Create(service *models.Service) error {
	return ss.db.Create(service).Error
}

func (ss *ServService) Update(service *models.Service) error {
	return ss.db.Save(service).Error
}

func (ss *ServService) Delete(id int) error {
	return ss.db.Delete(&models.Service{}, id).Error
}
