package routes

import (
	homeRoutes "blog/internal/modules/home/routes"
	articleRoutes "blog/internal/modules/article/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
}
