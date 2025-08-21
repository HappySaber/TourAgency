package routes

import (
	"TurAgency/internal/controllers"
	midlleware "TurAgency/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initProviderRoutes(r *gin.RouterGroup, providerCtrl *controllers.ProviderController) {
	providers := r.Group("/provider").Use(midlleware.IsAuthorized())
	{
		providers.GET("/", providerCtrl.List)
		providers.GET("/new", providerCtrl.New)
		providers.POST("/new", providerCtrl.Create)
		providers.GET("/edit/:id", providerCtrl.Edit)
		providers.PUT("/edit/:id", providerCtrl.Update)
		providers.DELETE("/:id", providerCtrl.Delete)
	}
}
