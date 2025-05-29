package services

import (
	"TurAgency/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TourService struct {
	db *gorm.DB
}

func NewTourService(db *gorm.DB) *TourService {
	return &TourService{
		db: db,
	}
}

func (ts *TourService) GetAllTours(c *gin.Context) {
	var tours []*models.Tour
	if err := ts.db.Find(&tours).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении туров"})
		return
	}

	c.JSON(http.StatusOK, tours)
}

func (ts *TourService) CreateTour(c *gin.Context) {
	var tour models.Tour

	if err := c.ShouldBind(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ввод"})
		return
	}

	if err := ts.db.Create(&tour).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании тура"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Тур успешно создан"})
}

func (ts *TourService) DeleteTour(c *gin.Context) {
	id := c.Param("id")
	var tour models.Tour
	res := ts.db.Delete(&tour, id)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении тура"})
		return
	}
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Тур не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Тур успешно удалён"})
}
