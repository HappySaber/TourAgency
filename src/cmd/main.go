package main

import (
	"TurAgency/src/controllers"
	"TurAgency/src/database"
	"TurAgency/src/routes"
	"TurAgency/src/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error while loading .env file: %v", err)
	}
	// Создание сервиса и контроллера
	// Создаем экземпляр DBConfig
	db := database.Init()

	// DI
	authService := services.NewAuthService(db)
	authController := controllers.NewAuthController(authService)

	port := "8080"
	router := gin.New()
	routes.TourAgencyRoutes(router, authController)
	router.Run(":" + port)
}
