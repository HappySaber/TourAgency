package services

import (
	"TurAgency/internal/models"
	"errors"

	"gorm.io/gorm"
)

type TourService struct {
	db *gorm.DB
}

func NewTourService(db *gorm.DB) *TourService {
	return &TourService{db: db}
}

func (ts *TourService) GetAll() ([]*models.Tour, error) {
	var tours []*models.Tour
	err := ts.db.Find(&tours).Error
	return tours, err
}

func (ts *TourService) GetByID(id string) (*models.Tour, error) {
	var tour models.Tour
	err := ts.db.First(&tour, "id = ?", id).Error // Используем оператор сравнения
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &tour, err
}

func (ts *TourService) Create(tour *models.Tour) error {
	return ts.db.Create(tour).Error
}

func (ss *TourService) Update(tour *models.Tour) error {
	return ss.db.Save(tour).Error
}

func (ts *TourService) Delete(id string) error {
	res := ts.db.Delete(&models.Tour{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func (s *TourService) GetProviders() ([]models.Provider, error) {
	var providers []models.Provider
	if err := s.db.Find(&providers).Error; err != nil {
		return nil, err
	}
	return providers, nil
}
