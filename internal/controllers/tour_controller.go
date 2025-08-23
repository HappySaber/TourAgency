package controllers

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TourController struct {
	service     *services.TourService
	auditLogger audit.Logger
}

func NewTourController(service *services.TourService, al audit.Logger) *TourController {
	return &TourController{
		service:     service,
		auditLogger: al,
	}
}

// List возвращает все туры
func (tc *TourController) List(c *gin.Context) {
	tours, err := tc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении туров"})
		return
	}
	c.JSON(http.StatusOK, tours)
}

// GetByID возвращает тур по ID
func (tc *TourController) GetByID(c *gin.Context) {
	id := c.Param("id")
	tour, err := tc.service.GetByID(id)
	if err != nil || tour == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Тур не найден"})
		return
	}
	c.JSON(http.StatusOK, tour)
}

// Create создаёт новый тур
func (tc *TourController) Create(c *gin.Context) {
	var tour models.Tour
	if err := c.ShouldBindJSON(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ввод"})
		return
	}

	evt, err := tc.service.Create(c.Request.Context(), &tour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании тура"})
		return
	}

	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := tc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusCreated, gin.H{"success": "Тур успешно создан", "tour": tour})
}

// Update обновляет существующий тур
func (tc *TourController) Update(c *gin.Context) {
	id := c.Param("id")
	tour, err := tc.service.GetByID(id)
	if err != nil || tour == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Тур не найден"})
		return
	}

	var updatedTour models.Tour
	if err := c.ShouldBindJSON(&updatedTour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка в форме", "details": err.Error()})
		return
	}

	updatedTour.ID = tour.ID

	evt, err := tc.service.Update(c.Request.Context(), &updatedTour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении тура"})
		return
	}

	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := tc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"success": "Тур успешно обновлён", "tour": updatedTour})
}

// Delete удаляет тур
func (tc *TourController) Delete(c *gin.Context) {
	id := c.Param("id")
	evt, err := tc.service.Delete(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Тур не найден"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении тура"})
		}
		return
	}

	if evt == nil {
		evt = &audit.Event{}
	}
	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := tc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"success": "Тур успешно удалён"})
}
