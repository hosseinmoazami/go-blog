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
		"title": "Register Page",
	})
}

func (controller *Controller) RegisterHandle(c *gin.Context) {
	var request auth.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		keepFormData.Init()
		keepFormData.SetFromData(c)
		sessions.Set(c, "formData", converters.UrlValuesToString(keepFormData.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

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

	}

	user, err := controller.userService.Create(request)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))
	log.Printf("The user created successfully with the name: %s", user.Name)
	c.Redirect(http.StatusFound, "/")

}

func (controller *Controller) Login(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login Page",
	})
}

func (controller *Controller) LoginHandle(c *gin.Context) {
	var request auth.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		keepFormData.Init()
		keepFormData.SetFromData(c)
		sessions.Set(c, "formData", converters.UrlValuesToString(keepFormData.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := controller.userService.HandleUserLogin(request)

	if err != nil {
		errors.Init()
		errors.Add("email", err.Error())
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))
	log.Printf("The user logged in successfully with the name: %s", user.Name)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) LogoutHandle(c *gin.Context) {
	sessions.Remove(c, "auth")
	log.Println("The user logged out successfully")
	c.Redirect(http.StatusFound, "/login")
}
