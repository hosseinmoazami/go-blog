package keepFormData

import (
	"github.com/gin-gonic/gin"
)

var dataList = make(map[string][]string)

func Init() {
	dataList = map[string][]string{}
}

func SetFromData(c *gin.Context) {
	c.Request.ParseForm()

	dataList = c.Request.PostForm
}

func Get() map[string][]string {
	return dataList
}
