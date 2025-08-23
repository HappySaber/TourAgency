package controllers

import (
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{service}
}

func (ac *AuthController) Login(c *gin.Context) {
	var userReq models.EmployeeRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := ac.service.Login(userReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", tokenString, 7200, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"success": "User logged in"})
}

func (ac *AuthController) CreateNewEmployee(c *gin.Context) {
	var user models.Employee
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.service.Signup(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "User created successfully"})
}

func (ac *AuthController) DummyLoginController(c *gin.Context) {
	var user models.Employee

	// Читаем JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Вызываем сервис
	tokenString, err := ac.service.DummyLoginService(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ставим куку
	expirationTime := time.Now().Add(30 * time.Minute)
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

	// Ответ клиенту
	c.JSON(http.StatusOK, gin.H{"success": "user logged in by dummyLogin"})
}

func (ac *AuthController) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)

	if c.GetHeader("Accept") != "application/json" {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "User logged out"})
}
