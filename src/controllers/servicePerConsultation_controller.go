package controllers

import (
	"TurAgency/src/models"
	"TurAgency/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ServiceWithCheck struct {
	ID       uint
	Name     string
	Discount string
	Quantity string
	Checked  bool
}

type ServicePerConsultationController struct {
	service *services.ServicePerConsultationService
	allSvc  *services.ServService // для получения всех услуг
}

func NewServicePerConsultationController(svc *services.ServicePerConsultationService, allSvc *services.ServService) *ServicePerConsultationController {
	return &ServicePerConsultationController{svc, allSvc}
}
func (c *ServicePerConsultationController) EditPage(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error", gin.H{"error": "Неверный ID консультации"})
		return
	}

	current, _ := c.service.GetByConsultationID(id) // []models.ServicePerConsultation
	all, _ := c.allSvc.GetAll()

	// Создаём мапу для быстрого доступа по ID
	currentMap := make(map[uint]models.ServicePerConsultation)
	for _, item := range current {
		currentMap[item.ServiceID] = item
	}

	var servicesWithCheck []ServiceWithCheck
	for _, s := range all {
		item, checked := currentMap[s.ID]

		servicesWithCheck = append(servicesWithCheck, ServiceWithCheck{
			ID:       s.ID,
			Name:     s.Name,
			Checked:  checked,
			Discount: item.Discount, // пустая строка если не найдено
			Quantity: item.Quantity, // пустая строка если не найдено
		})
	}

	ctx.HTML(http.StatusOK, "consultation/consultation_services_add", gin.H{
		"ConsultationID": id,
		"AllServices":    servicesWithCheck,
	})
}

func (c *ServicePerConsultationController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	consultationID, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID консультации"})
		return
	}

	form := ctx.PostFormArray("services")
	var ids []uint
	for _, s := range form {
		id, err := strconv.ParseUint(s, 10, 64)
		if err == nil {
			ids = append(ids, uint(id))
		}
	}

	err = c.service.UpdateServicesForConsultation(consultationID, ids, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/consultation")
}
