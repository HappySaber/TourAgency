package controllers

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	service     *services.ClientService
	auditLogger audit.Logger
}

func NewClientController(service *services.ClientService, al audit.Logger) *ClientController {
	return &ClientController{
		service:     service,
		auditLogger: al,
	}
}

// GET /clients
func (cc *ClientController) List(c *gin.Context) {
	clients, err := cc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки клиентов"})
		return
	}
	c.JSON(http.StatusOK, clients)
}

// GET /clients/:id
func (cc *ClientController) GetByID(c *gin.Context) {
	id := c.Param("id")
	client, err := cc.service.GetByID(id)
	if err != nil || client == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Клиент не найден"})
		return
	}
	c.JSON(http.StatusOK, client)
}

// POST /clients
func (cc *ClientController) Create(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	evt, err := cc.service.Create(c.Request.Context(), &client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании клиента"})
		return
	}

	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := cc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusCreated, client)
}

// PUT /clients/:id
// PUT /clients/:id
func (cc *ClientController) Update(c *gin.Context) {
	id := c.Param("id")
	client, err := cc.service.GetByID(id)
	if err != nil || client == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Клиент не найден"})
		return
	}

	if err := c.ShouldBindJSON(client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	evt, err := cc.service.Update(c.Request.Context(), client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении клиента"})
		return
	}

	// Заполняем дополнительные поля события
	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := cc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusOK, client)
}

// DELETE /clients/:id
func (cc *ClientController) Delete(c *gin.Context) {
	id := c.Param("id")
	var evt *audit.Event
	var err error
	if evt, err = cc.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении клиента"})
		return
	}

	evt = &audit.Event{
		ActorID:       c.GetString("user_id"),
		CorrelationID: c.GetHeader("X-Request-ID"),
		IP:            c.ClientIP(),
		UserAgent:     c.Request.UserAgent(),
	}

	if err := cc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Клиент удалён"})
}
