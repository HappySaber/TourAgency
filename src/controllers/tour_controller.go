package controllers

import (
	"TurAgency/src/models"
	"TurAgency/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TourController struct {
	service *services.TourService
}

func NewTourController(service *services.TourService) *TourController {
	return &TourController{service}
}

func (tc *TourController) GetAll(c *gin.Context) {
	tours, err := tc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении туров"})
		return
	}

	c.JSON(http.StatusOK, tours)
}

func (tc *TourController) Create(c *gin.Context) {
	var tour models.Tour
	if err := c.ShouldBind(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ввод"})
		return
	}

	if err := tc.service.Create(&tour); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании тура"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Тур успешно создан"})
}

func (tc *TourController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := tc.service.Delete(id); err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Тур не найден"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении тура"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Тур успешно удалён"})
}
