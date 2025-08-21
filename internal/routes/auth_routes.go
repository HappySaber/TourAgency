package routes

import (
	"TurAgency/internal/controllers"
	middleware "TurAgency/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initAuthRoutes(r *gin.RouterGroup, authCtrl *controllers.AuthController, emplCtrl *controllers.EmployeeController) {
	r.POST("/login", authCtrl.Login)
	auth := r.Group("/employee").Use(middleware.IsAuthorized(), middleware.IsAdmin())
	{
		auth.GET("/", emplCtrl.List)
		auth.GET("/employeeEdit/:id", emplCtrl.Edit)
		auth.PUT("/employeeEdit/:id", emplCtrl.Update)
		auth.POST("/employeeCreate", authCtrl.CreateNewEmployee)
		auth.POST("/logout", authCtrl.Logout)
	}
}
