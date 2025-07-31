package services

import (
	"TurAgency/internal/models"
	"fmt"

	"github.com/gin-gonic/gin"
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

func (s *ServicePerConsultationService) UpdateServicesForConsultation(consultationID uuid.UUID, serviceIDs []uint, ctx *gin.Context) error {
	// Удалим старые связи
	if err := s.db.Where("consultation_id = ?", consultationID).Delete(&models.ServicePerConsultation{}).Error; err != nil {
		return err
	}

	for _, id := range serviceIDs {
		discount := ctx.PostForm(fmt.Sprintf("discount_%d", id))
		quanity := ctx.PostForm(fmt.Sprintf("quanity_%d", id))

		entry := models.ServicePerConsultation{
			ConsultationID: consultationID,
			ServiceID:      id,
			Discount:       discount,
			Quantity:       quanity,
		}

		if err := s.db.Create(&entry).Error; err != nil {
			return err
		}
	}

	return nil
}
