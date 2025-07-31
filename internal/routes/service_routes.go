package routes

import (
	"TurAgency/internal/controllers"
	midlleware "TurAgency/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initServiceRoutes(r *gin.Engine, serviceCntr *controllers.ServiceController) {
	service := r.Group("/service").Use(midlleware.IsAuthorized())
	{
		service.GET("/", serviceCntr.List)
		service.GET("/new", serviceCntr.New)
		service.POST("/new", serviceCntr.Create)
		service.GET("/edit/:id", serviceCntr.Edit)
		service.PUT("/edit/:id", serviceCntr.Update)
		service.DELETE("/:id", serviceCntr.Delete)
	}
}
