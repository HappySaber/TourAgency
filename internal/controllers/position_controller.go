package controllers

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PositionController struct {
	service     *services.PositionService
	auditLogger audit.Logger
}

func NewPositionController(service *services.PositionService, al audit.Logger) *PositionController {
	return &PositionController{
		service:     service,
		auditLogger: al,
	}
}

// GET /positions
func (cc *PositionController) GetAll(c *gin.Context) {
	positions, err := cc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки должностей"})
		return
	}
	c.JSON(http.StatusOK, positions)
}

// GET /positions/:id
func (cc *PositionController) GetByID(c *gin.Context) {
	id := c.Param("id")
	position, err := cc.service.GetByID(id)
	if err != nil || position == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Должность не найдена"})
		return
	}
	c.JSON(http.StatusOK, position)
}

// POST /positions
func (cc *PositionController) Create(c *gin.Context) {
	var position models.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	evt, err := cc.service.Create(c.Request.Context(), &position)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании должности"})
		return
	}

	// Заполняем данные для аудита
	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := cc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusCreated, position)
}

// PUT /positions/:id
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

	evt, err := cc.service.Update(c.Request.Context(), position)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении должности"})
		return
	}

	// Заполняем данные для аудита
	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := cc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusOK, position)
}

// DELETE /positions/:id
func (cc *PositionController) Delete(c *gin.Context) {
	id := c.Param("id")
	evt, err := cc.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении должности"})
		return
	}

	// Если сервис не вернул событие, создаем пустое и заполняем поля
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

	c.JSON(http.StatusOK, gin.H{"success": "Должность удалена"})
}
