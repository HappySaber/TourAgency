package controllers

import (
	"TurAgency/internal/models"
	"TurAgency/internal/services"
	"net/http"
	"strings"

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

	tokens, err := ac.service.Login(userReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", tokens.AccessToken, 1800, "/", "", false, true)

	c.SetCookie("refresh_token", tokens.RefreshToken, 7*24*3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
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

func (ac *AuthController) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "токен обязателен"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if err := ac.service.Logout(token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось выйти"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "User logged out"})
}
