package controllers

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConsultationController struct {
	service     *services.ConsultationService
	auditLogger audit.Logger
}

func NewConsultationController(service *services.ConsultationService, al audit.Logger) *ConsultationController {
	return &ConsultationController{
		service:     service,
		auditLogger: al,
	}
}

// GET /consultations
func (cc *ConsultationController) List(c *gin.Context) {
	consultations, err := cc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки консультаций"})
		return
	}
	c.JSON(http.StatusOK, consultations)
}

// GET /consultations/:id
func (cc *ConsultationController) GetByID(c *gin.Context) {
	id := c.Param("id")
	consultation, err := cc.service.GetByID(id)
	if err != nil || consultation == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Консультация не найдена"})
		return
	}
	c.JSON(http.StatusOK, consultation)
}

// POST /consultations
func (cc *ConsultationController) Create(c *gin.Context) {
	var consultation models.Consultation
	if err := c.ShouldBindJSON(&consultation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	evt, err := cc.service.Create(c.Request.Context(), &consultation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании консультации"})
		return
	}

	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := cc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusCreated, consultation)
}

// PUT /consultations/:id
func (cc *ConsultationController) Update(c *gin.Context) {
	id := c.Param("id")
	consultation, err := cc.service.GetByID(id)
	if err != nil || consultation == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Консультация не найдена"})
		return
	}

	if err := c.ShouldBindJSON(consultation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	evt, err := cc.service.Update(c.Request.Context(), consultation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении консультации"})
		return
	}

	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := cc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusOK, consultation)
}

// DELETE /consultations/:id
func (cc *ConsultationController) Delete(c *gin.Context) {
	id := c.Param("id")
	evt, err := cc.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении консультации"})
		return
	}

	// Дополняем поля события, если сервис их не заполняет
	if evt == nil {
		evt = &audit.Event{}
	}
	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := cc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Консультация удалена"})
}
