package controllers

import (
	"TurAgency/internal/audit"
	"TurAgency/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service     *services.EmployeeService
	auditLogger audit.Logger
}

func NewEmployeeController(service *services.EmployeeService, al audit.Logger) *EmployeeController {
	return &EmployeeController{
		service:     service,
		auditLogger: al,
	}
}

// GET /employees
func (ec *EmployeeController) List(c *gin.Context) {
	employees, err := ec.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки сотрудников"})
		return
	}

	positions, err := ec.service.GetPositions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки позиций"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"employees": employees,
		"positions": positions,
	})
}

// GET /employees/:id
func (ec *EmployeeController) GetByID(c *gin.Context) {
	id := c.Param("id")
	employee, err := ec.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при поиске сотрудника"})
		return
	}
	if employee == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Сотрудник не найден"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// PUT /employees/:id
func (ec *EmployeeController) Update(c *gin.Context) {
	id := c.Param("id")
	employee, err := ec.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при поиске сотрудника"})
		return
	}
	if employee == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Сотрудник не найден"})
		return
	}

	if err := c.ShouldBindJSON(employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	evt, err := ec.service.Update(c.Request.Context(), employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении сотрудника"})
		return
	}

	evt.ActorID = c.GetString("user_id")
	evt.CorrelationID = c.GetHeader("X-Request-ID")
	evt.IP = c.ClientIP()
	evt.UserAgent = c.Request.UserAgent()

	if err := ec.auditLogger.Log(c.Request.Context(), *evt); err != nil {
		log.Printf("failed to write audit log: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"success": "Сотрудник обновлён"})
}
