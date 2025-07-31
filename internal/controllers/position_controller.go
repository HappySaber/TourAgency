package controllers

import (
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PositionController struct {
	service *services.PositionService
}

func NewPositionController(service *services.PositionService) *PositionController {
	return &PositionController{service}
}

func (cc *PositionController) List(c *gin.Context) {
	positions, err := cc.service.GetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка загрузки клиентов"})
		return
	}
	c.HTML(http.StatusOK, "position/position", gin.H{
		"Title":     "Список должностей",
		"Positions": positions,
	})
}

// New отображает форму создания нового поставщика
func (cc *PositionController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "position/position_new", gin.H{
		"Title": "Создание нового клиента",
	})
}

func (cc *PositionController) GetAll(c *gin.Context) {
	position, err := cc.service.GetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка загрузки клиентов"})
		return
	}

	c.HTML(http.StatusOK, "position", gin.H{
		"Title":     "Список должностей",
		"Positions": position,
	})
}

func (cc *PositionController) GetByID(c *gin.Context) {
	id := c.Param("id")
	position, err := cc.service.GetByID(id)
	if err != nil || position == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Клиент не найден"})
		return
	}

	c.HTML(http.StatusOK, "position_detail", gin.H{
		"Title":    "Детали клиента",
		"Position": position,
	})
}

func (cc *PositionController) Create(c *gin.Context) {
	var position models.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка формы"})
		return
	}

	if err := cc.service.Create(&position); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании клиента"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Клиент создан"})

}

// Edit отображает форму редактирования поставщика
func (cc *PositionController) Edit(c *gin.Context) {
	id := c.Param("id")
	position, err := cc.service.GetByID(id)
	if err != nil {
		c.Set("Error", err)
		return
	}
	c.HTML(http.StatusOK, "position/position_edit", gin.H{
		"Title":    "Редактирование поставщика",
		"Position": position,
	})
}

func (cc *PositionController) Update(c *gin.Context) {
	id := c.Param("id")
	position, err := cc.service.GetByID(id)
	if err != nil || position == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Клиент не найден"})
		return
	}

	if err := c.ShouldBind(position); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка формы"})
		return
	}

	if err := cc.service.Update(position); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при обновлении поставщика"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Position updated successfully"})
}

func (cc *PositionController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := cc.service.Delete(id); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при удалении поставщика"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Position deleted successfully"})
}
