package controllers

import (
	"blog/internal/modules/user/requests/auth"
	"blog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New() *Controller {
	return &Controller{}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"message": "Register Page",
	})
}

func (controller *Controller) RegisterHandle(c *gin.Context) {
	// validate the request
	var registerRequest auth.RegisterRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&registerRequest); err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	// create the user

	// check if there is any error on the user creation

	// after creating the user > redirect user to home page

	c.JSON(http.StatusOK, gin.H{
		"message": "Register Handle",
		"user":    registerRequest,
	})
}
