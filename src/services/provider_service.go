package services

import (
	"TurAgency/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProviderService struct {
	db *gorm.DB
}

func NewProviderService(db *gorm.DB) *ProviderService {
	return &ProviderService{db: db}
}

// GET /providers
func (ps *ProviderService) GetAllProviders(c *gin.Context) {
	var providers []models.Provider

	if err := ps.db.Find(&providers).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка загрузки провайдеров"})
		return
	}

	c.HTML(http.StatusOK, "providers", gin.H{
		"Title":     "Список провайдеров",
		"Providers": providers,
	})
}

// GET /providers/:id
func (ps *ProviderService) GetProviderByID(c *gin.Context) {
	id := c.Param("id")
	var provider models.Provider

	if err := ps.db.First(&provider, "id = ?", id).Error; err != nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Провайдер не найден"})
		return
	}

	c.HTML(http.StatusOK, "provider_detail", gin.H{
		"Title":    "Детали провайдера",
		"Provider": provider,
	})
}

// POST /providers
func (ps *ProviderService) CreateProvider(c *gin.Context) {
	var provider models.Provider

	if err := c.ShouldBind(&provider); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка формы"})
		return
	}

	provider.ID = uuid.New()

	if err := ps.db.Create(&provider).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при создании провайдера"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/providers")
}

// POST /providers/edit/:id
func (ps *ProviderService) UpdateProvider(c *gin.Context) {
	id := c.Param("id")
	var provider models.Provider

	if err := ps.db.First(&provider, "id = ?", id).Error; err != nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Провайдер не найден"})
		return
	}

	if err := c.ShouldBind(&provider); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка формы"})
		return
	}

	if err := ps.db.Save(&provider).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при обновлении провайдера"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/providers")
}

// POST /providers/delete/:id
func (ps *ProviderService) DeleteProvider(c *gin.Context) {
	id := c.Param("id")

	if err := ps.db.Delete(&models.Provider{}, "id = ?", id).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при удалении провайдера"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/providers")
}
