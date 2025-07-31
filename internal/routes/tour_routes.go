package routes

import (
	"TurAgency/internal/controllers"
	midlleware "TurAgency/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initTourRoutes(r *gin.Engine, tourCntr *controllers.TourController) {
	tours := r.Group("/tours").Use(midlleware.IsAuthorized())
	{
		tours.GET("/", tourCntr.List)
		tours.GET("/new", tourCntr.New)
		tours.POST("/new", tourCntr.Create)
		tours.GET("/edit/:id", tourCntr.Edit)
		tours.PUT("/edit/:id", tourCntr.Update)
		tours.DELETE("/:id", tourCntr.Delete)
	}
}
