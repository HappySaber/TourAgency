package routes

import (
	"TurAgency/internal/controllers"
	midlleware "TurAgency/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initServiceRoutes(r *gin.RouterGroup, serviceCntr *controllers.ServiceController) {
	service := r.Group("/service").Use(midlleware.IsAuthorized())
	{
		service.GET("/", serviceCntr.List)
		service.POST("/new", serviceCntr.Create)
		service.PUT("/edit/:id", serviceCntr.Update)
		service.DELETE("/:id", serviceCntr.Delete)
	}
}
