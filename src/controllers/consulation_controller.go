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

// List отображает список поставщиков в HTML
func (cc *ConsultationController) List(c *gin.Context) {
	consultation, err := cc.service.GetAll()
	if err != nil {
		c.Set("Error", err)
		return
	}
	c.HTML(http.StatusOK, "consultation/consultation", gin.H{
		"Title":     "Список консультаций",
		"Providers": consultation,
	})
}

// New отображает форму создания нового поставщика
func (cc *ConsultationController) New(c *gin.Context) {
	clients, err := cc.service.GetAllClients() // Получите всех клиентов
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка получения клиентов"})
		return
	}

	employees, err := cc.service.GetAllEmployees() // Получите всех сотрудников
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка получения сотрудников"})
		return
	}

	c.HTML(http.StatusOK, "consultation/consultation_new", gin.H{
		"Title":     "Создание новой консультации",
		"Clients":   clients,
		"Employees": employees,
	})
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

	c.JSON(http.StatusOK, gin.H{"success": "Consultation created successfully"})
}

func (cc *ConsultationController) GetAll(c *gin.Context) {
	consultation, err := cc.service.GetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка получения данных"})
		return
	}

	c.HTML(http.StatusOK, "consultation", gin.H{
		"Title":         "Список консультаций",
		"Consultations": consultation,
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

	c.JSON(http.StatusOK, gin.H{"success": "Consultation updated successfully"})
}

// Edit отображает форму редактирования поставщика
func (cc *ConsultationController) Edit(c *gin.Context) {
	id := c.Param("id")
	client, err := cc.service.GetByID(id)
	if err != nil {
		c.Set("Error", err)
		return
	}
	c.HTML(http.StatusOK, "consultation/consultation_edit", gin.H{
		"Title":    "Редактирование поставщика",
		"Provider": client,
	})
}

func (cc *ConsultationController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := cc.service.Delete(id); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка удаления консультации"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Consultation deleted successfully"})
}
