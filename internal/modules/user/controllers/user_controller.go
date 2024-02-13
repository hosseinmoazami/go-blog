package controllers

import (
	"blog/internal/modules/user/requests/auth"
	UserServices "blog/internal/modules/user/services"
	"blog/pkg/converters"
	"blog/pkg/errors"
	"blog/pkg/html"
	"blog/pkg/keepFormData"
	"blog/pkg/sessions"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserServices.UserServiceInterface
}
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New() *Controller {
	return &Controller{
		userService: UserServices.New(),
	}
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
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		keepFormData.Init()
		keepFormData.SetFromData(c)
		sessions.Set(c, "formData", converters.UrlValuesToString(keepFormData.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	// create the user
	email := c.Request.PostForm.Get("email")
	userExist := controller.userService.CheckUserExist(email)

	if userExist {
		errors.Init()
		errors.Add("Email", "Email address already exists")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		keepFormData.Init()
		keepFormData.SetFromData(c)
		sessions.Set(c, "formData", converters.UrlValuesToString(keepFormData.Get()))

		c.Redirect(http.StatusFound, "/register")
		c.Redirect(http.StatusFound, "/register")
		return

	} else {
		user, err := controller.userService.Create(registerRequest)

		// check if there is any error on the user creation
		if err != nil {
			c.Redirect(http.StatusFound, "/register")
			return
		}

		sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

		// after creating the user > redirect user to home page
		log.Printf("The user created successfully with the name: %s", user.Name)
		c.Redirect(http.StatusFound, "/")
	}
}
