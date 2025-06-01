package controllers

import (
	"TurAgency/src/models"
	"TurAgency/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	service *services.ClientService
}

func NewClientController(service *services.ClientService) *ClientController {
	return &ClientController{service}
}

func (cc *ClientController) List(c *gin.Context) {
	clients, err := cc.service.GetAll() // Изменено на clients
	if err != nil {
		c.Set("Error", err)
		return
	}
	c.HTML(http.StatusOK, "client/client", gin.H{
		"Title":   "Список клиентов",
		"Clients": clients, // Изменено на "Clients"
	})
}

// New отображает форму создания нового поставщика
func (cc *ClientController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "client/client_new", gin.H{
		"Title": "Создание нового клиента",
	})
}

func (cc *ClientController) GetAll(c *gin.Context) {
	client, err := cc.service.GetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка загрузки поставщиков"})
		return
	}

	c.HTML(http.StatusOK, "client", gin.H{
		"Title":   "Список клиентов",
		"Clients": client,
	})
}

func (cc *ClientController) GetByID(c *gin.Context) {
	id := c.Param("id")
	client, err := cc.service.GetByID(id)
	if err != nil || client == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Клиент не найден"})
		return
	}

	c.HTML(http.StatusOK, "client_detail", gin.H{
		"Title":  "Детали поставщика",
		"Client": client,
	})
}

func (cc *ClientController) Create(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка формы"})
		return
	}

	if err := cc.service.Create(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании клиента"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Клиент создан"})

}

// Edit отображает форму редактирования поставщика
func (cc *ClientController) Edit(c *gin.Context) {
	id := c.Param("id")
	client, err := cc.service.GetByID(id)
	if err != nil {
		c.Set("Error", err)
		return
	}
	c.HTML(http.StatusOK, "client/client_edit", gin.H{
		"Title":  "Редактирование поставщика",
		"Client": client,
	})
}

func (cc *ClientController) Update(c *gin.Context) {
	id := c.Param("id")
	client, err := cc.service.GetByID(id)
	if err != nil || client == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Поставщик не найден"})
		return
	}

	if err := c.ShouldBind(client); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка формы"})
		return
	}

	if err := cc.service.Update(client); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при обновлении поставщика"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Client updated successfully"})
	//c.Redirect(http.StatusSeeOther, "/client")
}

func (cc *ClientController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := cc.service.Delete(id); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при удалении поставщика"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Client deleted successfully"})
}
