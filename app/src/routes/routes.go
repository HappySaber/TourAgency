package routes

import (
	"TurAgency/src/controllers"
	"TurAgency/src/services"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func TourAgencyRoutes(r *gin.Engine, authController *controllers.AuthController) {
	// Создаем экземпляр рендерера
	render := multitemplate.NewRenderer()

	// Путь к директории с шаблонами
	templatesDir := "src/templates"

	// Обход всех файлов в директории
	err := filepath.Walk(templatesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Проверяем, что это файл и он имеет расширение .html
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			// Извлекаем имя файла без расширения
			name := info.Name()[:len(info.Name())-len(filepath.Ext(info.Name()))]
			// Добавляем шаблон
			render.AddFromFiles(name, "src/templates/base.html", path)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Ошибка при обходе директории: %v", err)
	}

	// Устанавливаем кастомный рендерер
	r.HTMLRender = render

	// Подключение статики
	r.Static("/src/static", "./static")

	// Главная страница
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"Title": "Главная",
		})
	})

	// Страница логина
	r.GET("/login", func(c *gin.Context) {
		service := &services.AuthService{}
		positions, err := service.GetPositions()
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error", gin.H{
				"Message": "Ошибка при получении должностей",
			})
			return
		}
		c.HTML(http.StatusOK, "login", gin.H{
			"Title":     "Вход",
			"Positions": positions,
		})
	})

	// Обработчик для входа
	r.POST("/login", authController.Login)

	r.GET("/create_new_empl", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create_new_empl", gin.H{"Title": "Создание нового работника"})
	})

	r.POST("/create_new_empl", authController.CreateNewEmployee)

	r.POST("/logout", authController.Logout)

	r.GET("/tours", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tours", gin.H{"Title": "Создание нового работника"})
	})

	r.GET("/tours/edit", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tour_edit", gin.H{"Title": "Создание нового работника"})
	})

	r.GET("/tours/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tour_new", gin.H{
			"Title": "Создание нового тура",
		})
	})

}
