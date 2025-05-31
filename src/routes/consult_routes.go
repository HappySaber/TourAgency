package routes

import (
	"TurAgency/src/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initConsultrRoutes(r *gin.Engine, consultCntr *controllers.ConsultationController, db *gorm.DB) {
	r.GET("/consultation", func(c *gin.Context) {
		c.HTML(http.StatusOK, "provider/provider", gin.H{"Title": "Создание нового работника"})
	})

	r.GET("/provider/edit/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "provider/provider_edit", gin.H{"Title": "Создание нового работника"})
	})

	r.GET("/provider/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "provider/provider_new", gin.H{
			"Title": "Создание нового тура",
		})
	})
}
