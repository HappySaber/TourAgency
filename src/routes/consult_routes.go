package routes

import (
	"TurAgency/src/controllers"
	midlleware "TurAgency/src/middleware"

	"github.com/gin-gonic/gin"
)

func initConsultationRoutes(
	r *gin.Engine,
	consultationCtrl *controllers.ConsultationController,
	servicePerConsultationCtrl *controllers.ServicePerConsultationController,
	tourPerConsultationCtrl *controllers.TourPerConsultationController,
) {
	consultations := r.Group("/consultation").Use(midlleware.IsAuthorized())
	{
		// Стандартные маршруты консультаций
		consultations.GET("/", consultationCtrl.List)
		consultations.GET("/new", consultationCtrl.New)
		consultations.POST("/new", consultationCtrl.Create)
		consultations.GET("/edit/:id", consultationCtrl.Edit)
		consultations.PUT("/edit/:id", consultationCtrl.Update)
		consultations.DELETE("/:id", consultationCtrl.Delete)

		// Редактирование услуг, привязанных к консультации
		consultations.GET("/:id/services", servicePerConsultationCtrl.EditPage)
		consultations.POST("/:id/services", servicePerConsultationCtrl.Update)

		// Редактирование туров, привязанных к консультации
		consultations.GET("/:id/tours", tourPerConsultationCtrl.EditPage)
		consultations.POST("/:id/tours", tourPerConsultationCtrl.Update)
	}
}
