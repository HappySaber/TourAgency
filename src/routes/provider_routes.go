package routes

import (
	"TurAgency/src/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initProviderRoutes(r *gin.Engine, providerCtrl *controllers.ProviderController, db *gorm.DB) {
	providers := r.Group("/providers")
	{
		providers.GET("/", providerCtrl.List)
		providers.GET("/new", providerCtrl.New)
		providers.POST("/new", providerCtrl.Create)
		providers.GET("/edit/:id", providerCtrl.Edit)
		providers.PUT("/edit/:id", providerCtrl.Update)
		providers.DELETE("/:id", providerCtrl.Delete)
	}
}
