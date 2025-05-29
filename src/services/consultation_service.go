package services

import (
	"TurAgency/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConsultationService struct {
	db *gorm.DB
}

func NewConsultationService(db *gorm.DB) *ConsultationService {
	return &ConsultationService{db: db}
}

// Create (POST /consultations)
func (cs *ConsultationService) CreateConsultation(c *gin.Context) {
	var consultation models.Consultation

	if err := c.ShouldBind(&consultation); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка при разборе формы"})
		return
	}

	consultation.ID = uuid.New()

	if err := cs.db.Create(&consultation).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка сохранения в БД"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/consultations")
}

// Read All (GET /consultations)
func (cs *ConsultationService) GetAllConsultations(c *gin.Context) {
	var consultations []models.Consultation
	if err := cs.db.Find(&consultations).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка получения данных"})
		return
	}

	c.HTML(http.StatusOK, "consultations", gin.H{
		"Title":         "Список консультаций",
		"Consultations": consultations,
	})
}

// Read One (GET /consultations/:id)
func (cs *ConsultationService) GetConsultationById(c *gin.Context) {
	id := c.Param("id")
	var consultation models.Consultation

	if err := cs.db.First(&consultation, "id = ?", id).Error; err != nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Консультация не найдена"})
		return
	}

	c.HTML(http.StatusOK, "consultation_detail", gin.H{
		"Title":        "Детали консультации",
		"Consultation": consultation,
	})
}

// Update (POST /consultations/edit/:id)
func (cs *ConsultationService) UpdateConsultation(c *gin.Context) {
	id := c.Param("id")
	var consultation models.Consultation

	if err := cs.db.First(&consultation, "id = ?", id).Error; err != nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Консультация не найдена"})
		return
	}

	if err := c.ShouldBind(&consultation); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка при разборе формы"})
		return
	}

	if err := cs.db.Save(&consultation).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Не удалось сохранить изменения"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/consultations")
}

// Delete (POST /consultations/delete/:id)
func (cs *ConsultationService) DeleteConsultation(c *gin.Context) {
	id := c.Param("id")

	if err := cs.db.Delete(&models.Consultation{}, "id = ?", id).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка удаления консультации"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/consultations")
}
