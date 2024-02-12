package routes

import (
	articleCtrl "blog/internal/modules/article/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articleController := articleCtrl.New()
	router.GET("/articles/:id", articleController.Show)
}
