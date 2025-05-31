package routes

import (
	"TurAgency/src/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initTourRoutes(r *gin.Engine, tourCntr *controllers.TourController, db *gorm.DB) {
	tours := r.Group("/tours")
	{
		tours.GET("/", tourCntr.List)
		tours.GET("/new", tourCntr.New)
		tours.POST("/new", tourCntr.Create)
		tours.GET("/edit/:id", tourCntr.Edit)
		tours.PUT("/edit/:id", tourCntr.Update)
		tours.DELETE("/:id", tourCntr.Delete)
	}
}
