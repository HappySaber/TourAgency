package services

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourPerConsultationService struct {
	db *gorm.DB
}

func NewTourPerConsultationService(db *gorm.DB) *TourPerConsultationService {
	return &TourPerConsultationService{db}
}

func (s *TourPerConsultationService) GetByConsultationID(consultationID uuid.UUID) ([]models.TourPerConsultation, error) {
	var tours []models.TourPerConsultation
	err := s.db.Where("consultation_id = ?", consultationID).Find(&tours).Error
	return tours, err
}

// Обновляем туры для консультации с аудитом
func (s *TourPerConsultationService) UpdateToursForConsultation(ctx context.Context, consultationID uuid.UUID, tourIDs []uuid.UUID) ([]*audit.Event, error) {
	var events []*audit.Event

	// Сохраняем старые записи для аудита
	var oldTours []models.TourPerConsultation
	if err := s.db.Where("consultation_id = ?", consultationID).Find(&oldTours).Error; err != nil {
		return nil, err
	}

	// Удаляем старые
	if err := s.db.Where("consultation_id = ?", consultationID).Delete(&models.TourPerConsultation{}).Error; err != nil {
		return nil, err
	}

	for _, id := range tourIDs {
		entry := models.TourPerConsultation{
			ConsultationID: consultationID,
			TourID:         id,
		}

		if err := s.db.WithContext(ctx).Create(&entry).Error; err != nil {
			return nil, err
		}

		evt := &audit.Event{
			Event:    "tour_per_consultation.updated",
			Entity:   "tour_per_consultation",
			EntityID: consultationID.String() + "_" + id.String(),
			At:       time.Now(),
			Before:   audit.MustMarshal(nil), // Можно расширить для конкретного старого значения
			After:    audit.MustMarshal(&entry),
		}
		events = append(events, evt)
	}

	return events, nil
}

// Обновляем с массивом структур (можно использовать для POST с JSON)
func (s *TourPerConsultationService) UpdateToursWithData(ctx context.Context, consultationID uuid.UUID, tours []models.TourPerConsultation) ([]*audit.Event, error) {
	var events []*audit.Event

	// Сохраняем старые записи для аудита
	var oldTours []models.TourPerConsultation
	if err := s.db.Where("consultation_id = ?", consultationID).Find(&oldTours).Error; err != nil {
		return nil, err
	}

	// Удаляем старые
	if err := s.db.Where("consultation_id = ?", consultationID).Delete(&models.TourPerConsultation{}).Error; err != nil {
		return nil, err
	}

	// Сохраняем новые
	for _, t := range tours {
		if err := s.db.WithContext(ctx).Create(&t).Error; err != nil {
			return nil, err
		}

		evt := &audit.Event{
			Event:    "tour_per_consultation.updated",
			Entity:   "tour_per_consultation",
			EntityID: t.ConsultationID.String() + "_" + t.TourID.String(),
			At:       time.Now(),
			Before:   audit.MustMarshal(nil),
			After:    audit.MustMarshal(&t),
		}
		events = append(events, evt)
	}

	return events, nil
}
