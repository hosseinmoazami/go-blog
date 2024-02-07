package routes

import (
	"blog/pkg/html"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		html.Render(c, 200, "modules/home/html/home", gin.H{
			"title": "Home Page",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		html.Render(c, 200, "modules/home/html/about", gin.H{
			"title": "About Page",
		})
	})
}
