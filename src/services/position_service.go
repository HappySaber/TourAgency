package services

import (
	"TurAgency/src/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PositionService struct {
	db *gorm.DB
}

func NewPositionService(db *gorm.DB) *PositionService {
	return &PositionService{db: db}
}

func (cs *PositionService) Create(position *models.Position) error {
	position.ID = uuid.New()
	return cs.db.Create(position).Error
}

func (cs *PositionService) GetAll() ([]models.Position, error) {
	var Position []models.Position
	err := cs.db.Find(&Position).Error
	return Position, err
}

func (cs *PositionService) GetByID(id string) (*models.Position, error) {
	var Position models.Position
	err := cs.db.First(&Position, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &Position, err
}

func (cs *PositionService) Update(updated *models.Position) error {
	return cs.db.Save(updated).Error
}

func (cs *PositionService) Delete(id string) error {
	return cs.db.Delete(&models.Position{}, "id = ?", id).Error
}
