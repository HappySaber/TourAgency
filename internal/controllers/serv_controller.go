package controllers

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	service     *services.ServService
	auditLogger audit.Logger
}

func NewServiceController(service *services.ServService, al audit.Logger) *ServiceController {
	return &ServiceController{
		service:     service,
		auditLogger: al,
	}
}

// List возвращает все услуги
func (sc *ServiceController) List(c *gin.Context) {
	services, err := sc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось загрузить услуги"})
		return
	}

	c.JSON(http.StatusOK, services)
}

// GetByID возвращает услугу по ID
func (sc *ServiceController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	service, err := sc.service.GetByID(id)
	if err != nil || service == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Услуга не найдена"})
		return
	}

	c.JSON(http.StatusOK, service)
}

// Create создаёт новую услугу
func (sc *ServiceController) Create(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка в форме"})
		return
	}

	evt, err := sc.service.Create(c.Request.Context(), &service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать услугу"})
		return
	}

	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := sc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusCreated, gin.H{"success": "Услуга успешно создана", "service": service})
}

// Update обновляет существующую услугу
func (sc *ServiceController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	service, err := sc.service.GetByID(id)
	if err != nil || service == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Услуга не найдена"})
		return
	}

	var updatedService models.Service
	if err := c.ShouldBindJSON(&updatedService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка в форме", "details": err.Error()})
		return
	}

	updatedService.ID = service.ID

	evt, err := sc.service.Update(c.Request.Context(), &updatedService)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении услуги"})
		return
	}

	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := sc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"success": "Услуга успешно обновлена", "service": updatedService})
}

// Delete удаляет услугу
func (sc *ServiceController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	evt, err := sc.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении услуги"})
		return
	}

	if evt == nil {
		evt = &audit.Event{}
	}
	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := sc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"success": "Услуга успешно удалена"})
}
