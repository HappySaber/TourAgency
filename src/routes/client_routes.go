package routes

import (
	"TurAgency/src/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initClientRoutes(r *gin.Engine, clientCtrl *controllers.ClientController, db *gorm.DB) {
	clients := r.Group("/client")
	{
		clients.GET("/", clientCtrl.List)
		clients.GET("/new", clientCtrl.New)
		clients.POST("/new", clientCtrl.Create)
		clients.GET("/edit/:id", clientCtrl.Edit)
		clients.PUT("/edit/:id", clientCtrl.Update)
		clients.DELETE("/:id", clientCtrl.Delete)
	}
}
