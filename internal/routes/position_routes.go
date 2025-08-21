package routes

import (
	"TurAgency/internal/controllers"
	midlleware "TurAgency/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initPositionRoutes(r *gin.RouterGroup, positionCtrl *controllers.PositionController) {
	positions1 := r.Group("/position").Use(midlleware.IsAuthorized())
	{
		positions1.GET("/", positionCtrl.GetAll)
	}
	positions := r.Group("/position").Use(midlleware.IsAuthorized(), midlleware.IsAdmin())
	{
		positions.POST("/new", positionCtrl.Create)
		positions.PUT("/edit/:id", positionCtrl.Update)
		positions.DELETE("/:id", positionCtrl.Delete)
	}
}
