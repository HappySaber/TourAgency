package routes

import (
	"TurAgency/internal/controllers"
	"TurAgency/internal/services"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TourAgencyRoutes(r *gin.Engine, db *gorm.DB) {
	// CORS для React dev-сервера
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Vite порт по умолчанию
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authService := services.NewAuthService(db)
	authController := controllers.NewAuthController(authService)

	tourService := services.NewTourService(db)
	tourController := controllers.NewTourController(tourService)

	providerService := services.NewProviderService(db)
	providerController := controllers.NewProviderController(providerService)

	consulatationService := services.NewConsultationService(db)
	consulatationController := controllers.NewConsultationController(consulatationService)

	clientService := services.NewClientService(db)
	clientController := controllers.NewClientController(clientService)

	serviceService := services.NewServService(db)
	serviceController := controllers.NewServiceController(serviceService)

	positionService := services.NewPositionService(db)
	positionController := controllers.NewPositionController(positionService)

	employeeService := services.NewEmployeeService(db)
	employeeController := controllers.NewEmployeeController(employeeService)

	servicePerConsultationService := services.NewServicePerConsultationService(db)
	tourPerConsultationService := services.NewTourPerConsultationService(db)
	servicePerConsultationController := controllers.NewServicePerConsultationController(servicePerConsultationService, serviceService)
	tourPerConsultationController := controllers.NewTourPerConsultationController(tourPerConsultationService, tourService)

	// === API ===
	api := r.Group("/api")
	{
		initAuthRoutes(api, authController, employeeController)
		initTourRoutes(api, tourController)
		initProviderRoutes(api, providerController)
		initConsultationRoutes(api, consulatationController, servicePerConsultationController, tourPerConsultationController)
		initClientRoutes(api, clientController)
		initServiceRoutes(api, serviceController)
		initPositionRoutes(api, positionController)
	}

	// === React build ===
	// Отдаём статические файлы из папки React build
	r.Static("/assets", filepath.Join("web", "frontend", "dist", "assets"))

	// Обработка всех остальных запросов через index.html
	r.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join("web", "frontend", "dist", "index.html"))
	})
}
