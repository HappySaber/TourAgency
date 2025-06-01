package routes

import (
	"TurAgency/src/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initConsultationRoutes(r *gin.Engine, consultationCtrl *controllers.ConsultationController, db *gorm.DB) {
	providers := r.Group("/consultation")
	{
		providers.GET("/", consultationCtrl.List)
		providers.GET("/new", consultationCtrl.New)
		providers.POST("/new", consultationCtrl.Create)
		providers.GET("/edit/:id", consultationCtrl.Edit)
		providers.PUT("/edit/:id", consultationCtrl.Update)
		providers.DELETE("/:id", consultationCtrl.Delete)
	}
}
