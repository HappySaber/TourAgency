package services

import (
	"TurAgency/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServService struct {
	db *gorm.DB
}

func NewServService(db *gorm.DB) *ServService {
	return &ServService{db: db}
}

// GET /services
func (ss *ServService) GetAllServices(c *gin.Context) {
	var services []models.Service

	if err := ss.db.Find(&services).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Не удалось загрузить услуги"})
		return
	}

	c.HTML(http.StatusOK, "services", gin.H{
		"Title":    "Список услуг",
		"Services": services,
	})
}

// GET /services/:id
func (ss *ServService) GetServiceByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Некорректный ID"})
		return
	}

	var service models.Service
	if err := ss.db.First(&service, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Услуга не найдена"})
		return
	}

	c.HTML(http.StatusOK, "service_detail", gin.H{
		"Title":   "Детали услуги",
		"Service": service,
	})
}

// POST /services
func (ss *ServService) CreateService(c *gin.Context) {
	var service models.Service

	if err := c.ShouldBind(&service); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка в форме"})
		return
	}

	if err := ss.db.Create(&service).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Не удалось создать услугу"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/services")
}

// POST /services/edit/:id
func (ss *ServService) UpdateService(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Некорректный ID"})
		return
	}

	var service models.Service
	if err := ss.db.First(&service, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Услуга не найдена"})
		return
	}

	if err := c.ShouldBind(&service); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка в форме"})
		return
	}

	if err := ss.db.Save(&service).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при обновлении услуги"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/services")
}

// POST /services/delete/:id
func (ss *ServService) DeleteService(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Некорректный ID"})
		return
	}

	if err := ss.db.Delete(&models.Service{}, id).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при удалении услуги"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/services")
}
