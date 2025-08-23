package controllers

import (
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ServiceWithCheck struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Discount string `json:"discount"`
	Quantity string `json:"quantity"`
	Checked  bool   `json:"checked"`
}

type ServicePerConsultationController struct {
	service *services.ServicePerConsultationService
	allSvc  *services.ServService // для получения всех услуг
}

func NewServicePerConsultationController(svc *services.ServicePerConsultationService, allSvc *services.ServService) *ServicePerConsultationController {
	return &ServicePerConsultationController{svc, allSvc}
}

// GetForConsultation возвращает все услуги для консультации с флагом выбора
func (c *ServicePerConsultationController) GetForConsultation(ctx *gin.Context) {
	idStr := ctx.Param("id")
	consultationID, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID консультации"})
		return
	}

	current, err := c.service.GetByConsultationID(consultationID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении услуг"})
		return
	}

	all, err := c.allSvc.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении всех услуг"})
		return
	}

	// Мапа для быстрого поиска
	currentMap := make(map[uint]models.ServicePerConsultation)
	for _, s := range current {
		currentMap[s.ServiceID] = s
	}

	var servicesWithCheck []ServiceWithCheck
	for _, s := range all {
		item, checked := currentMap[s.ID]
		servicesWithCheck = append(servicesWithCheck, ServiceWithCheck{
			ID:       s.ID,
			Name:     s.Name,
			Checked:  checked,
			Discount: item.Discount,
			Quantity: item.Quantity,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"consultation_id": consultationID,
		"services":        servicesWithCheck,
	})
}

// Update обновляет выбранные услуги для консультации
func (c *ServicePerConsultationController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	consultationID, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID консультации"})
		return
	}

	// Читаем данные формы
	form := ctx.Request.PostForm
	serviceIDs := form["services"]

	formData := make(map[string]string)
	for k, v := range form {
		if len(v) > 0 {
			formData[k] = v[0]
		}
	}

	var ids []uint
	for _, s := range serviceIDs {
		id, err := strconv.ParseUint(s, 10, 64)
		if err == nil {
			ids = append(ids, uint(id))
		}
	}

	events, err := c.service.UpdateServicesForConsultation(ctx.Request.Context(), consultationID, ids, formData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении услуг"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "Услуги для консультации обновлены",
		"events":  events, // можно вернуть события аудита, если нужно
	})
}
