package services

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type ConsultationService struct {
	db *gorm.DB
}

func NewConsultationService(db *gorm.DB) *ConsultationService {
	return &ConsultationService{db: db}
}

func (cs *ConsultationService) Create(ctx context.Context, consultation *models.Consultation) (*audit.Event, error) {
	if err := cs.db.WithContext(ctx).Create(consultation).Error; err != nil {
		return nil, err
	}

	evt := &audit.Event{
		Event:    "consultation.created",
		Entity:   "client",
		EntityID: consultation.ID.String(),
		At:       time.Now(),
		After:    audit.MustMarshal(consultation),
	}
	return evt, nil
}

func (cs *ConsultationService) GetAll() ([]models.Consultation, error) {

	var consultations []models.Consultation
	err := cs.db.
		Preload("Client").
		Preload("Employee").
		Find(&consultations).Error

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

func (cs *ConsultationService) Update(ctx context.Context, updated *models.Consultation) (*audit.Event, error) {
	var before models.Consultation
	if err := cs.db.WithContext(ctx).First(&before, "id = ?", updated.ID).Error; err != nil {
		return nil, err
	}

	if err := cs.db.WithContext(ctx).Save(updated).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "consultation.updated",
		Entity:   "consultation",
		EntityID: updated.ID.String(),
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(updated),
	}
	return evt, nil
}

func (cs *ConsultationService) Delete(ctx context.Context, id string) (*audit.Event, error) {
	var before models.Consultation
	if err := cs.db.WithContext(ctx).First(&before, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := cs.db.WithContext(ctx).Delete(&models.Consultation{}, "id = ?", id).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "consultation.deleted",
		Entity:   "consultation",
		EntityID: id,
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(nil),
	}
	return evt, nil
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
