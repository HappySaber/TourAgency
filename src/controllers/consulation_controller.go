package controllers

import (
	"TurAgency/src/models"
	"TurAgency/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConsultationController struct {
	service *services.ConsultationService
}

func NewConsultationController(service *services.ConsultationService) *ConsultationController {
	return &ConsultationController{service}
}

func (cc *ConsultationController) Create(c *gin.Context) {
	var consultation models.Consultation
	if err := c.ShouldBind(&consultation); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка при разборе формы"})
		return
	}

	if err := cc.service.Create(&consultation); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка сохранения в БД"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/consultations")
}

func (cc *ConsultationController) GetAll(c *gin.Context) {
	consultations, err := cc.service.GetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка получения данных"})
		return
	}

	c.HTML(http.StatusOK, "consultations", gin.H{
		"Title":         "Список консультаций",
		"Consultations": consultations,
	})
}

func (cc *ConsultationController) GetByID(c *gin.Context) {
	id := c.Param("id")
	consultation, err := cc.service.GetByID(id)
	if err != nil || consultation == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Консультация не найдена"})
		return
	}

	c.HTML(http.StatusOK, "consultation_detail", gin.H{
		"Title":        "Детали консультации",
		"Consultation": consultation,
	})
}

func (cc *ConsultationController) Update(c *gin.Context) {
	id := c.Param("id")
	consultation, err := cc.service.GetByID(id)
	if err != nil || consultation == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Консультация не найдена"})
		return
	}

	if err := c.ShouldBind(consultation); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка при разборе формы"})
		return
	}

	if err := cc.service.Update(consultation); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Не удалось сохранить изменения"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/consultations")
}

func (cc *ConsultationController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := cc.service.Delete(id); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка удаления консультации"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/consultations")
}
