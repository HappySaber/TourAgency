package controllers

import (
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PositionController struct {
	service *services.PositionService
}

func NewPositionController(service *services.PositionService) *PositionController {
	return &PositionController{service}
}

// Получить все должности
func (cc *PositionController) GetAll(c *gin.Context) {
	positions, err := cc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки должностей"})
		return
	}
	c.JSON(http.StatusOK, positions)
}

// Получить должность по ID
func (cc *PositionController) GetByID(c *gin.Context) {
	id := c.Param("id")
	position, err := cc.service.GetByID(id)
	if err != nil || position == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Должность не найдена"})
		return
	}
	c.JSON(http.StatusOK, position)
}

// Создать новую должность
func (cc *PositionController) Create(c *gin.Context) {
	var position models.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	if err := cc.service.Create(&position); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании должности"})
		return
	}

	c.JSON(http.StatusCreated, position)
}

// Обновить должность
func (cc *PositionController) Update(c *gin.Context) {
	id := c.Param("id")
	position, err := cc.service.GetByID(id)
	if err != nil || position == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Должность не найдена"})
		return
	}

	if err := c.ShouldBindJSON(position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	if err := cc.service.Update(position); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении должности"})
		return
	}

	c.JSON(http.StatusOK, position)
}

// Удалить должность
func (cc *PositionController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := cc.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении должности"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Должность удалена"})
}
