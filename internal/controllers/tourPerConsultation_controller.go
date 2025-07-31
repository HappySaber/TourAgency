package controllers

import (
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TourPerConsultationController struct {
	service  *services.TourPerConsultationService
	allTours *services.TourService
}

func NewTourPerConsultationController(svc *services.TourPerConsultationService, allTours *services.TourService) *TourPerConsultationController {
	return &TourPerConsultationController{svc, allTours}
}

type TourWithCheck struct {
	ID       uuid.UUID
	Name     string
	Country  string
	Discount string
	Quantity string
	Checked  bool
}

func (c *TourPerConsultationController) EditPage(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error", gin.H{"error": "Неверный ID консультации"})
		return
	}

	current, _ := c.service.GetByConsultationID(id)
	all, _ := c.allTours.GetAll()

	currentMap := make(map[uuid.UUID]models.TourPerConsultation)
	for _, t := range current {
		currentMap[t.TourID] = t
	}

	var toursWithCheck []TourWithCheck
	for _, t := range all {
		item, checked := currentMap[t.ID]
		toursWithCheck = append(toursWithCheck, TourWithCheck{
			ID:       t.ID,
			Name:     t.Name,
			Country:  t.Country,
			Checked:  checked,
			Discount: item.Discount,
			Quantity: item.Quantity,
		})
	}

	ctx.HTML(http.StatusOK, "consultation/consultation_tours_add", gin.H{
		"ConsultationID": id,
		"Tours":          toursWithCheck,
	})
}
func (c *TourPerConsultationController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	consultationID, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID консультации"})
		return
	}

	form := ctx.PostFormArray("tours")
	var tours []models.TourPerConsultation

	for _, s := range form {
		tourID, err := uuid.Parse(s)
		if err != nil {
			continue
		}
		discount := ctx.PostForm("discount_" + s)
		quantity := ctx.PostForm("quantity_" + s)

		tours = append(tours, models.TourPerConsultation{
			TourID:         tourID,
			ConsultationID: consultationID,
			Discount:       discount,
			Quantity:       quantity,
		})
	}

	if err := c.service.UpdateToursWithData(consultationID, tours); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении туров"})
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/consultation")
}
