package routes

import (
	"TurAgency/internal/controllers"
	midlleware "TurAgency/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initConsultationRoutes(
	r *gin.RouterGroup,
	consultationCtrl *controllers.ConsultationController,
	servicePerConsultationCtrl *controllers.ServicePerConsultationController,
	tourPerConsultationCtrl *controllers.TourPerConsultationController,
) {
	consultations := r.Group("/consultation").Use(midlleware.IsAuthorized())
	{
		// Стандартные маршруты консультаций
		consultations.GET("/", consultationCtrl.List)
		consultations.POST("/new", consultationCtrl.Create)
		consultations.PUT("/edit/:id", consultationCtrl.Update)
		consultations.DELETE("/:id", consultationCtrl.Delete)

		// Редактирование услуг, привязанных к консультации
		consultations.POST("/:id/services", servicePerConsultationCtrl.Update)

		// Редактирование туров, привязанных к консультации
		consultations.POST("/:id/tours", tourPerConsultationCtrl.Update)
	}
}
