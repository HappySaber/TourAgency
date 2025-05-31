package controllers

import (
	"TurAgency/src/models"
	"TurAgency/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProviderController struct {
	service *services.ProviderService
}

func NewProviderController(service *services.ProviderService) *ProviderController {
	return &ProviderController{service}
}

// List отображает список поставщиков в HTML
func (pc *ProviderController) List(c *gin.Context) {
	providers, err := pc.service.GetAll()
	if err != nil {
		c.Set("Error", err)
		return
	}
	c.HTML(http.StatusOK, "providers/providers", gin.H{
		"Title":     "Список поставщиков",
		"Providers": providers,
	})
}

// New отображает форму создания нового поставщика
func (pc *ProviderController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "providers/provider_new", gin.H{
		"Title": "Создание нового поставщика",
	})
}

func (pc *ProviderController) GetAll(c *gin.Context) {
	providers, err := pc.service.GetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка загрузки поставщиков"})
		return
	}

	c.HTML(http.StatusOK, "providers", gin.H{
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

func (pc *ProviderController) Create(c *gin.Context) {
	var provider models.Provider
	if err := c.ShouldBind(&provider); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка формы"})
		return
	}

	if err := pc.service.Create(&provider); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при создании поставщика"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/providers")
}

// Edit отображает форму редактирования поставщика
func (pc *ProviderController) Edit(c *gin.Context) {
	id := c.Param("id")
	provider, err := pc.service.GetByID(id)
	if err != nil {
		c.Set("Error", err)
		return
	}
	c.HTML(http.StatusOK, "providers/provider_edit", gin.H{
		"Title":    "Редактирование поставщика",
		"Provider": provider,
	})
}

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

	if err := pc.service.Update(provider); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при обновлении поставщика"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/providers")
}

func (pc *ProviderController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := pc.service.Delete(id); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при удалении поставщика"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/providers")
}
