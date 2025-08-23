package routes

import (
	"TurAgency/internal/controllers"
	midlleware "TurAgency/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initClientRoutes(r *gin.RouterGroup, clientCtrl *controllers.ClientController) {
	clients := r.Group("/client").Use(midlleware.IsAuthorized())
	{
		clients.GET("/", clientCtrl.List)
		clients.POST("/new", clientCtrl.Create)
		clients.PUT("/edit/:id", clientCtrl.Update)
		clients.DELETE("/:id", clientCtrl.Delete)
	}
}
