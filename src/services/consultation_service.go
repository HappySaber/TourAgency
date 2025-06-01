package services

import (
	"TurAgency/src/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConsultationService struct {
	db *gorm.DB
}

func NewConsultationService(db *gorm.DB) *ConsultationService {
	return &ConsultationService{db: db}
}

func (cs *ConsultationService) Create(consultation *models.Consultation) error {
	consultation.ID = uuid.New()
	return cs.db.Create(consultation).Error
}

func (cs *ConsultationService) GetAll() ([]models.Consultation, error) {
	var consultations []models.Consultation
	err := cs.db.Find(&consultations).Error
	return consultations, err
}

func (cs *ConsultationService) GetByID(id string) (*models.Consultation, error) {
	var consultation models.Consultation
	err := cs.db.First(&consultation, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &consultation, err
}

func (cs *ConsultationService) Update(updated *models.Consultation) error {
	return cs.db.Save(updated).Error
}

func (cs *ConsultationService) Delete(id string) error {
	return cs.db.Delete(&models.Consultation{}, "id = ?", id).Error
}

func (cs *ConsultationService) GetAllClients() ([]models.Client, error) {
	var clients []models.Client
	if err := cs.db.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func (cs *ConsultationService) GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	if err := cs.db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}
