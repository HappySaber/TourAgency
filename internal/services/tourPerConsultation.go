package services

import (
	"TurAgency/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourPerConsultationService struct {
	db *gorm.DB
}

func NewTourPerConsultationService(db *gorm.DB) *TourPerConsultationService {
	return &TourPerConsultationService{db}
}

func (t *TourPerConsultationService) GetByConsultationID(consultationID uuid.UUID) ([]models.TourPerConsultation, error) {
	var tours []models.TourPerConsultation
	err := t.db.Where("consultation_id = ?", consultationID).Find(&tours).Error
	return tours, err
}

func (t *TourPerConsultationService) UpdateToursForConsultation(consultationID uuid.UUID, tourIDs []uuid.UUID) error {
	if err := t.db.Where("consultation_id = ?", consultationID).Delete(&models.TourPerConsultation{}).Error; err != nil {
		return err
	}

	for _, id := range tourIDs {
		entry := models.TourPerConsultation{
			ConsultationID: consultationID,
			TourID:         id,
		}
		if err := t.db.Create(&entry).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *TourPerConsultationService) UpdateToursWithData(consultationID uuid.UUID, tours []models.TourPerConsultation) error {
	// Удалим старые записи
	if err := s.db.Where("consultation_id = ?", consultationID).Delete(&models.TourPerConsultation{}).Error; err != nil {
		return err
	}

	// Сохраним новые
	for _, t := range tours {
		if err := s.db.Create(&t).Error; err != nil {
			return err
		}
	}

	return nil
}
