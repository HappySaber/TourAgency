package main

import (
	"TurAgency/internal/database"
	"TurAgency/internal/kafka"
	"TurAgency/internal/routes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	kafkago "github.com/segmentio/kafka-go"
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

	kafkaConsumer := kafka.NewKafkaConsumer(
		[]string{os.Getenv("KAFKA_BROKER")},
		os.Getenv("KAFKA_TOPIC"),
		os.Getenv("KAFKA_GROUP"),
	)

	ctx := context.Background()
	kafkaConsumer.Start(ctx, func(msg kafkago.Message) {
		fmt.Printf("Received message: %s = %s\n", string(msg.Key), string(msg.Value))
		// Здесь можно вызвать audit/handler или сервисы
	})

	port := "8080"
	router := gin.New()
	routes.TourAgencyRoutes(router, db)

	router.Run(":" + port)
}
