package controllers

import (
	"TurAgency/src/models"
	"TurAgency/src/services"
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

func (sc *ServiceController) GetAll(c *gin.Context) {
	services, err := sc.service.GetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Не удалось загрузить услуги"})
		return
	}

	c.HTML(http.StatusOK, "services", gin.H{
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

func (sc *ServiceController) Create(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBind(&service); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка в форме"})
		return
	}

	if err := sc.service.Create(&service); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Не удалось создать услугу"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/services")
}

func (sc *ServiceController) Update(c *gin.Context) {
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

	if err := c.ShouldBind(service); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка в форме"})
		return
	}

	if err := sc.service.Update(service); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при обновлении услуги"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/services")
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

	c.Redirect(http.StatusSeeOther, "/services")
}
