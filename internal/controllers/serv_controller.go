package controllers

import (
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	service *services.ServService
}

func NewServiceController(service *services.ServService) *ServiceController {
	return &ServiceController{service}
}

// List отображает список услуг
func (sc *ServiceController) List(c *gin.Context) {
	services, err := sc.service.GetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Не удалось загрузить услуги"})
		return
	}

	c.HTML(http.StatusOK, "service/service", gin.H{
		"Title":    "Список услуг",
		"Services": services,
	})
}

func (sc *ServiceController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Некорректный ID"})
		return
	}

	service, err := sc.service.GetByID(id)
	if err != nil || service == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Услуга не найдена"})
		return
	}

	c.HTML(http.StatusOK, "service_detail", gin.H{
		"Title":   "Детали услуги",
		"Service": service,
	})
}

// New отображает форму создания новой услуги
func (sc *ServiceController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "service/service_new", gin.H{
		"Title": "Создание новой услуги",
	})
}

func (sc *ServiceController) Create(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка в форме"})
		return
	}

	if err := sc.service.Create(&service); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Не удалось создать услугу"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Услуга успешно создана"})
}

// Edit отображает форму редактирования услуги
func (sc *ServiceController) Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Некорректный ID"})
		return
	}

	service, err := sc.service.GetByID(id)
	if err != nil || service == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Услуга не найдена"})
		return
	}

	c.HTML(http.StatusOK, "service/service_edit", gin.H{
		"Title":   "Редактирование услуги",
		"Service": service,
	})
}

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

	if err := sc.service.Update(&updatedService); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении услуги"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Услуга успешно обновлена"})
}

func (sc *ServiceController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Некорректный ID"})
		return
	}

	if err := sc.service.Delete(id); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при удалении услуги"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Услуга успешно удалена"})
}
