package services

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServicePerConsultationService struct {
	db *gorm.DB
}

func NewServicePerConsultationService(db *gorm.DB) *ServicePerConsultationService {
	return &ServicePerConsultationService{db}
}

func (s *ServicePerConsultationService) GetByConsultationID(consultationID uuid.UUID) ([]models.ServicePerConsultation, error) {
	var services []models.ServicePerConsultation
	err := s.db.Where("consultation_id = ?", consultationID).Find(&services).Error
	return services, err
}

func (s *ServicePerConsultationService) UpdateServicesForConsultation(ctx context.Context, consultationID uuid.UUID, serviceIDs []uint, formData map[string]string) ([]*audit.Event, error) {
	var events []*audit.Event

	// Получаем старые связи для аудита
	var oldServices []models.ServicePerConsultation
	if err := s.db.Where("consultation_id = ?", consultationID).Find(&oldServices).Error; err != nil {
		return nil, err
	}

	// Удаляем старые связи
	if err := s.db.Where("consultation_id = ?", consultationID).Delete(&models.ServicePerConsultation{}).Error; err != nil {
		return nil, err
	}

	for _, id := range serviceIDs {
		discount := formData[fmt.Sprintf("discount_%d", id)]
		quantity := formData[fmt.Sprintf("quantity_%d", id)]

		entry := models.ServicePerConsultation{
			ConsultationID: consultationID,
			ServiceID:      id,
			Discount:       discount,
			Quantity:       quantity,
		}

		if err := s.db.WithContext(ctx).Create(&entry).Error; err != nil {
			return nil, err
		}

		evt := &audit.Event{
			Event:    "service_per_consultation.updated",
			Entity:   "service_per_consultation",
			EntityID: consultationID.String() + "_" + strconv.FormatUint(uint64(id), 10),
			At:       time.Now(),
			Before:   audit.MustMarshal(nil), // Можно расширить, чтобы хранить старое значение конкретной связи
			After:    audit.MustMarshal(&entry),
		}
		events = append(events, evt)
	}

	// Если нужно, можно добавить один общий evt для удаления всех старых связей

	return events, nil
}
