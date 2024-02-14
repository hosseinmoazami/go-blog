package routes

import (
	userCtrl "blog/internal/modules/user/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	userController := userCtrl.New()
	router.GET("/register", userController.Register)
	router.POST("/register", userController.RegisterHandle)

	router.GET("/login", userController.Login)
	router.POST("/login", userController.LoginHandle)
}
