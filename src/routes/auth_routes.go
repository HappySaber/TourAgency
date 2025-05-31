package routes

import (
	"TurAgency/src/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initAuthRoutes(r *gin.Engine, authCtrl *controllers.AuthController, db *gorm.DB) {
	auth := r.Group("/auth")
	{
		auth.GET("/login", authCtrl.LoginPage)
		auth.POST("/login", authCtrl.Login)
		auth.GET("/create-employee", authCtrl.CreateEmployeePage)
		auth.POST("/create-employee", authCtrl.CreateNewEmployee)
		auth.POST("/logout", authCtrl.Logout)
	}
}
