package services

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type EmployeeService struct {
	db *gorm.DB
}

func NewEmployeeService(db *gorm.DB) *EmployeeService {
	return &EmployeeService{db: db}
}

func (s *EmployeeService) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := s.db.Preload("Position").Find(&employees).Error
	return employees, err
}

func (es *EmployeeService) GetByID(id string) (*models.Employee, error) {
	var Employee models.Employee
	err := es.db.First(&Employee, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &Employee, err
}

func (es *EmployeeService) Update(ctx context.Context, updated *models.Employee) (*audit.Event, error) {
	var before models.Employee
	if err := es.db.WithContext(ctx).First(&before, "id = ?", updated.ID).Error; err != nil {
		return nil, err
	}

	if err := es.db.WithContext(ctx).Save(updated).Error; err != nil {
		return nil, err
	}
	evt := &audit.Event{
		Event:    "employee.updated",
		Entity:   "employee",
		EntityID: updated.ID.String(),
		At:       time.Now(),
		Before:   audit.MustMarshal(&before),
		After:    audit.MustMarshal(updated),
	}
	return evt, nil
}

func (as *EmployeeService) GetPositions() ([]models.Position, error) {
	var positions []models.Position
	if err := as.db.Find(&positions).Error; err != nil {
		return nil, err
	}
	return positions, nil
}
