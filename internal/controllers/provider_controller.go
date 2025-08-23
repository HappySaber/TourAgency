package controllers

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProviderController struct {
	service     *services.ProviderService
	auditLogger audit.Logger
}

func NewProviderController(service *services.ProviderService, al audit.Logger) *ProviderController {
	return &ProviderController{
		service:     service,
		auditLogger: al,
	}
}

// List отображает список поставщиков в HTML
func (pc *ProviderController) List(c *gin.Context) {
	providers, err := pc.service.GetAll()
	if err != nil {
		c.Set("Error", err)
		return
	}
	c.HTML(http.StatusOK, "provider/provider", gin.H{
		"Title":     "Список поставщиков",
		"Providers": providers,
	})
}

func (pc *ProviderController) GetAll(c *gin.Context) {
	providers, err := pc.service.GetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка загрузки поставщиков"})
		return
	}

	c.HTML(http.StatusOK, "provider", gin.H{
		"Title":     "Список поставщиков",
		"Providers": providers,
	})
}

func (pc *ProviderController) GetByID(c *gin.Context) {
	id := c.Param("id")
	provider, err := pc.service.GetByID(id)
	if err != nil || provider == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Поставщик не найден"})
		return
	}

	c.HTML(http.StatusOK, "provider_detail", gin.H{
		"Title":    "Детали поставщика",
		"Provider": provider,
	})
}

// POST /providers
func (pc *ProviderController) Create(c *gin.Context) {
	var provider models.Provider
	if err := c.ShouldBindJSON(&provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка формы"})
		return
	}

	evt, err := pc.service.Create(c.Request.Context(), &provider)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при создании поставщика"})
		return
	}

	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := pc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusCreated, gin.H{"success": "Provider created successfully"})
}

// PUT /providers/:id
func (pc *ProviderController) Update(c *gin.Context) {
	id := c.Param("id")
	provider, err := pc.service.GetByID(id)
	if err != nil || provider == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Поставщик не найден"})
		return
	}

	if err := c.ShouldBind(provider); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка формы"})
		return
	}

	evt, err := pc.service.Update(c.Request.Context(), provider)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при обновлении поставщика"})
		return
	}

	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := pc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"success": "Provider updated successfully"})
}

// DELETE /providers/:id
func (pc *ProviderController) Delete(c *gin.Context) {
	id := c.Param("id")
	evt, err := pc.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при удалении поставщика"})
		return
	}

	if evt == nil {
		evt = &audit.Event{}
	}
	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := pc.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"success": "Provider deleted successfully"})
}
