package controllers

import (
	"TurAgency/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service *services.EmployeeService
}

func NewEmployeeController(service *services.EmployeeService) *EmployeeController {
	return &EmployeeController{service}
}

// List отображает список поставщиков в HTML
func (pc *EmployeeController) List(c *gin.Context) {
	employees, err := pc.service.GetAll()
	if err != nil {
		c.Set("Error", err)
		return
	}
	positions, err := pc.service.GetPositions()
	if err != nil {
		c.Set("Error", err)
		return
	}
	c.HTML(http.StatusOK, "auth/employee", gin.H{
		"Title":     "Список сотрудников",
		"Employees": employees,
		"Positions": positions,
	})
}

func (pc *EmployeeController) GetAll(c *gin.Context) {
	employees, err := pc.service.GetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка загрузки поставщиков"})
		return
	}

	c.HTML(http.StatusOK, "employee", gin.H{
		"Title":     "Список сотрудников",
		"Employees": employees,
	})
}

func (pc *EmployeeController) GetByID(c *gin.Context) {
	id := c.Param("id")
	employee, err := pc.service.GetByID(id)
	if err != nil || employee == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Сотрудник не найден"})
		return
	}

	c.HTML(http.StatusOK, "employee_detail", gin.H{
		"Title":    "Детали поставщика",
		"Employee": employee,
	})
}

// Edit отображает форму редактирования поставщика
func (pc *EmployeeController) Edit(c *gin.Context) {
	id := c.Param("id")
	employee, err := pc.service.GetByID(id)
	if err != nil {
		c.Set("Error", err)
		return
	}
	positions, err := pc.service.GetPositions()
	if err != nil {
		c.Set("Error", err)
		return
	}

	c.HTML(http.StatusOK, "auth/employee_edit", gin.H{
		"Title":     "Редактирование поставщика",
		"Employee":  employee,
		"Positions": positions,
	})
}

func (pc *EmployeeController) Update(c *gin.Context) {
	id := c.Param("id")
	employee, err := pc.service.GetByID(id)
	if err != nil || employee == nil {
		c.HTML(http.StatusNotFound, "error", gin.H{"error": "Поставщик не найден"})
		return
	}

	if err := c.ShouldBind(employee); err != nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Ошибка формы"})
		return
	}

	if err := pc.service.Update(employee); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Ошибка при обновлении поставщика"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Employee updated successfully"})
}
