package html

import (
	"blog/internal/providers/view"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	data = view.WithGlobalData(c, data)
	fmt.Println(data)
	c.HTML(code, name, data)
}
