package view

import (
	"blog/internal/modules/user/helpers"
	"blog/pkg/converters"
	"blog/pkg/sessions"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("App.Name")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(c, "errors"))
	data["FORM_DATA"] = converters.StringToUrlValues(sessions.Flash(c, "formData"))
	data["AUTH"] = helpers.Auth(c)
	return data
}
