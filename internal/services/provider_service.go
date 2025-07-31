package services

import (
	"TurAgency/internal/models"
	"errors"

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

func (ps *ProviderService) Create(provider *models.Provider) error {
	provider.ID = uuid.New()
	return ps.db.Create(provider).Error
}

func (ps *ProviderService) Update(provider *models.Provider) error {
	return ps.db.Save(provider).Error
}

func (ps *ProviderService) Delete(id string) error {
	return ps.db.Delete(&models.Provider{}, "id = ?", id).Error
}
