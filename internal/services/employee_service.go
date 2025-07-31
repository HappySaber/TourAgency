package services

import (
	"TurAgency/internal/models"
	"errors"

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

func (es *EmployeeService) Update(updated *models.Employee) error {
	return es.db.Save(updated).Error
}

func (as *EmployeeService) GetPositions() ([]models.Position, error) {
	var positions []models.Position
	if err := as.db.Find(&positions).Error; err != nil {
		return nil, err
	}
	return positions, nil
}
