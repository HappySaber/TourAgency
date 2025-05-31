package main

import (
	"TurAgency/src/database"
	"TurAgency/src/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println(os.Getwd())
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error while loading .env file: %v", err)
	}
	// Создание сервиса и контроллера
	// Создаем экземпляр DBConfig
	db := database.Init()

	port := "8080"
	router := gin.New()
	routes.TourAgencyRoutes(router, db)
	router.Run(":" + port)
}
