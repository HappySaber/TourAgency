package routes

import (
	"TurAgency/internal/controllers"
	middleware "TurAgency/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initAuthRoutes(r *gin.Engine, authCtrl *controllers.AuthController, emplCtrl *controllers.EmployeeController) {
	r.GET("/login", authCtrl.LoginPage)
	r.POST("/login", authCtrl.Login)
	auth := r.Group("/employee").Use(middleware.IsAuthorized(), middleware.IsAdmin())
	{
		auth.GET("/", emplCtrl.List)
		auth.GET("/edit/:id", emplCtrl.Edit)
		auth.PUT("/edit/:id", emplCtrl.Update)
		auth.GET("/new", authCtrl.CreateEmployeePage)
		auth.POST("/new", authCtrl.CreateNewEmployee)
		auth.POST("/logout", authCtrl.Logout)
	}
}
