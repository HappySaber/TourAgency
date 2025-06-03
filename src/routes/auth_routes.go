package routes

import (
	"TurAgency/src/controllers"
	middleware "TurAgency/src/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initAuthRoutes(r *gin.Engine, authCtrl *controllers.AuthController, emplCtrl *controllers.EmployeeController, db *gorm.DB) {
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
