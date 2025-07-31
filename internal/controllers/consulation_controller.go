package controllers

import (
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		"Title":         "Список консультаций",
		"Consultations": consultation,
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
	// Промежуточная структура (DTO) для парсинга JSON
	type ConsultationInput struct {
		DateOfConsultation string `json:"dateofconsultation"`
		TimeOfConsultation string `json:"timeofconsultation"`
		ClientID           string `json:"client"`
		EmployeeID         string `json:"employee"`
		Notes              string `json:"notes"`
	}

	var input ConsultationInput

	// Пробуем распарсить JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидные входные данные"})
		return
	}

	// Парсим дату
	date, err := time.Parse("2006-01-02", input.DateOfConsultation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат даты (нужно YYYY-MM-DD)"})
		return
	}

	// Парсим время
	timePart, err := time.Parse("15:04", input.TimeOfConsultation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат времени (нужно HH:MM)"})
		return
	}

	// Парсим UUID клиента
	clientUUID, err := uuid.Parse(input.ClientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный UUID клиента"})
		return
	}

	// Парсим UUID сотрудника
	employeeUUID, err := uuid.Parse(input.EmployeeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный UUID сотрудника"})
		return
	}

	// Формируем объект Consultation
	consultation := models.Consultation{
		DateOfConsultation: date,
		TimeOfConsultation: models.LocalTime{Time: timePart},
		ClientID:           clientUUID,
		EmployeeID:         employeeUUID,
		Notes:              input.Notes,
	}

	// Сохраняем
	if err := cc.service.Create(&consultation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении консультации"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Консультация успешно создана"})
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

	// Получаем консультацию по ID
	consultation, err := cc.service.GetByID(id)
	if err != nil || consultation == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Консультация не найдена"})
		return
	}

	// Промежуточная структура (DTO)
	type ConsultationInput struct {
		DateOfConsultation string `json:"dateofconsultation"`
		TimeOfConsultation string `json:"timeofconsultation"`
		ClientID           string `json:"client"`
		EmployeeID         string `json:"employee"`
		Notes              string `json:"notes"`
	}

	var input ConsultationInput

	// Пробуем распарсить JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидные входные данные"})
		return
	}

	// Парсим дату
	date, err := time.Parse("2006-01-02", input.DateOfConsultation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат даты (нужно YYYY-MM-DD)"})
		return
	}

	// Парсим время
	timePart, err := time.Parse("15:04", input.TimeOfConsultation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат времени (нужно HH:MM)"})
		return
	}

	// Парсим UUID клиента
	clientUUID, err := uuid.Parse(input.ClientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный UUID клиента"})
		return
	}

	// Парсим UUID сотрудника
	employeeUUID, err := uuid.Parse(input.EmployeeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный UUID сотрудника"})
		return
	}

	// Обновляем поля существующей консультации
	consultation.DateOfConsultation = date
	consultation.TimeOfConsultation = models.LocalTime{Time: timePart}
	consultation.ClientID = clientUUID
	consultation.EmployeeID = employeeUUID
	consultation.Notes = input.Notes

	// Сохраняем обновления
	if err := cc.service.Update(consultation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить изменения"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Консультация успешно обновлена"})
}

func (cc *ConsultationController) Edit(c *gin.Context) {
	idStr := c.Param("id")

	consultation, err := cc.service.GetByID(idStr)
	if err != nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Консультация не найдена"})
		return
	}

	// Получаем клиентов и сотрудников
	clients, err := cc.service.GetAllClients()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка загрузки клиентов"})
		return
	}

	employees, err := cc.service.GetAllEmployees()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка загрузки сотрудников"})
		return
	}

	c.HTML(http.StatusOK, "consultation/consultation_edit", gin.H{
		"Title":        "Редактировать консультацию",
		"Consultation": consultation,
		"Clients":      clients,
		"Employees":    employees,
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
