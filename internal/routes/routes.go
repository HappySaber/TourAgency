package routes

import (
	"TurAgency/internal/controllers"
	"TurAgency/internal/services"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TourAgencyRoutes(r *gin.Engine, db *gorm.DB) {

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

	// Устанавливаем кастомный рендерер
	r.HTMLRender = renderTemplates()

	// Подключение статики
	r.Static("/internal/static", "./static")

	// Главная страница
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"Title": "Главная",
		})
	})

	initAuthRoutes(r, authController, employeeController)
	initTourRoutes(r, tourController)
	initProviderRoutes(r, providerController)
	initConsultationRoutes(r, consulatationController, servicePerConsultationController, tourPerConsultationController)
	initClientRoutes(r, clientController)
	initServiceRoutes(r, serviceController)
	initPositionRoutes(r, positionController)
}

func renderTemplates() multitemplate.Renderer {
	render := multitemplate.NewRenderer()

	templatesDir := "internal/templates"
	baseLayout := filepath.Join(templatesDir, "base.html")

	err := filepath.Walk(templatesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Пропускаем директории и base.html
		if info.IsDir() || filepath.Base(path) == "base.html" {
			return nil
		}

		if filepath.Ext(path) == ".html" {
			// Уникальное имя шаблона по относительному пути (например, "providers/list")
			relativePath, err := filepath.Rel(templatesDir, path)
			if err != nil {
				return err
			}

			// Убираем .html и заменяем \ на /
			name := strings.TrimSuffix(relativePath, filepath.Ext(relativePath))
			name = filepath.ToSlash(name)

			render.AddFromFiles(name, baseLayout, path)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Ошибка при загрузке шаблонов: %v", err)
	}
	return render
}
