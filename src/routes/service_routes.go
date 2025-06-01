package routes

import (
	"TurAgency/src/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initServiceRoutes(r *gin.Engine, serviceCntr *controllers.ServiceController, db *gorm.DB) {
	service := r.Group("/service")
	{
		service.GET("/", serviceCntr.List)
		service.GET("/new", serviceCntr.New)
		service.POST("/new", serviceCntr.Create)
		service.GET("/edit/:id", serviceCntr.Edit)
		service.PUT("/edit/:id", serviceCntr.Update)
		service.DELETE("/:id", serviceCntr.Delete)
	}
}
