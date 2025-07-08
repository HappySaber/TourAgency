package routes

import (
	"TurAgency/src/controllers"
	midlleware "TurAgency/src/middleware"

	"github.com/gin-gonic/gin"
)

func initPositionRoutes(r *gin.Engine, positionCtrl *controllers.PositionController) {
	positions1 := r.Group("/position").Use(midlleware.IsAuthorized())
	{
		positions1.GET("/", positionCtrl.List)
	}
	positions := r.Group("/position").Use(midlleware.IsAuthorized(), midlleware.IsAdmin())
	{
		positions.GET("/new", positionCtrl.New)
		positions.POST("/new", positionCtrl.Create)
		positions.GET("/edit/:id", positionCtrl.Edit)
		positions.PUT("/edit/:id", positionCtrl.Update)
		positions.DELETE("/:id", positionCtrl.Delete)
	}
}
