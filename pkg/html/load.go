package html

import "github.com/gin-gonic/gin"

func LoadHtml(router *gin.Engine) {
	// internal/modules/modulesName/html/view.tmpl
	router.LoadHTMLGlob("internal/**/**/**/*tmpl")
}
