package controllers

import (
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TourController struct {
	service *services.TourService
}

func NewTourController(service *services.TourService) *TourController {
	return &TourController{service}
}

// List отображает список туров
func (tc *TourController) List(c *gin.Context) {
	tours, err := tc.service.GetAll()
	if err != nil {
		c.Set("Error", err)
		return
	}
	c.HTML(http.StatusOK, "tours/tours", gin.H{
		"Title": "Туры",
		"Tours": tours,
	})
}

func (tc *TourController) GetByID(c *gin.Context) {
	id := c.Param("id")

	tour, err := tc.service.GetByID(id)
	if err != nil || tour == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Тур не найден"})
		return
	}

	c.HTML(http.StatusOK, "tour_detail", gin.H{
		"Title":   "Детали тура",
		"Service": tour,
	})
}

func (tc *TourController) GetAll(c *gin.Context) {
	tours, err := tc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении туров"})
		return
	}

	c.JSON(http.StatusOK, tours)
}

// New отображает форму создания нового тура
func (tc *TourController) New(c *gin.Context) {
	providers, err := tc.service.GetProviders() // Получаем провайдеров для формы
	if err != nil {
		c.Set("Error", err)
		return
	}
	c.HTML(http.StatusOK, "tours/tour_new", gin.H{
		"Title":     "Создание нового тура",
		"Providers": providers,
	})
}

func (tc *TourController) Create(c *gin.Context) {
	var tour models.Tour
	if err := c.ShouldBindJSON(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ввод"})
		return
	}

	if err := tc.service.Create(&tour); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании тура"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Тур успешно создан"})
}

// Edit отображает форму редактирования тура
func (tc *TourController) Edit(c *gin.Context) {
	id := c.Param("id")
	tour, err := tc.service.GetByID(id)
	if err != nil {
		c.Set("Error", err)
		return
	}
	providers, err := tc.service.GetProviders() // Получаем провайдеров для формы
	if err != nil {
		c.Set("Error", err)
		return
	}
	c.HTML(http.StatusOK, "tours/tour_edit", gin.H{
		"Title":     "Редактирование тура",
		"Tour":      tour,
		"Providers": providers,
	})
}

func (tc *TourController) Update(c *gin.Context) {
	id := c.Param("id")

	// Получение существующего тура
	tour, err := tc.service.GetByID(id)
	if err != nil || tour == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Тур не найден"})
		return
	}

	// Привязка данных формы к новой переменной, не затирая существующую
	var updatedTour models.Tour
	if err := c.ShouldBind(&updatedTour); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка в форме"})
		return
	}

	// Сохраняем ID из пути, т.к. он может не прийти в форме
	updatedTour.ID = tour.ID

	// Выполнение обновления
	if err := tc.service.Update(&updatedTour); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при обновлении тура"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Tour created successfully"})
}

func (tc *TourController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := tc.service.Delete(id); err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Тур не найден"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении тура"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Tour deleted successfully"})
}
