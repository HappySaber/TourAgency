package controllers

import (
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TourWithCheck struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Country  string    `json:"country"`
	Discount string    `json:"discount"`
	Quantity string    `json:"quantity"`
	Checked  bool      `json:"checked"`
}

type TourPerConsultationController struct {
	service  *services.TourPerConsultationService
	allTours *services.TourService
}

func NewTourPerConsultationController(svc *services.TourPerConsultationService, allTours *services.TourService) *TourPerConsultationController {
	return &TourPerConsultationController{svc, allTours}
}

// GetForConsultation возвращает все туры с отметкой выбора для консультации
func (c *TourPerConsultationController) GetForConsultation(ctx *gin.Context) {
	idStr := ctx.Param("id")
	consultationID, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID консультации"})
		return
	}

	current, err := c.service.GetByConsultationID(consultationID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении туров"})
		return
	}

	all, err := c.allTours.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении всех туров"})
		return
	}

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

	ctx.JSON(http.StatusOK, gin.H{
		"consultation_id": consultationID,
		"tours":           toursWithCheck,
	})
}

// Update обновляет туры для консультации через JSON
func (c *TourPerConsultationController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	consultationID, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID консультации"})
		return
	}

	var tours []models.TourPerConsultation
	if err := ctx.ShouldBindJSON(&tours); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные", "details": err.Error()})
		return
	}

	// Принудительно присваиваем consultationID для безопасности
	for i := range tours {
		tours[i].ConsultationID = consultationID
	}

	events, err := c.service.UpdateToursWithData(ctx.Request.Context(), consultationID, tours)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении туров"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "Туры для консультации обновлены",
		"events":  events, // можно убрать, если не нужен клиенту
	})
}
