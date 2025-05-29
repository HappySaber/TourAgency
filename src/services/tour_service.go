package services

import (
	"TurAgency/src/models"
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

func (ts *TourService) Create(tour *models.Tour) error {
	return ts.db.Create(tour).Error
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
