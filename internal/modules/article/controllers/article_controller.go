package controllers

import (
	"blog/internal/modules/article/requests/articles"
	ArticleService "blog/internal/modules/article/services"
	UserHelpers "blog/internal/modules/user/helpers"
	ProviderHelpers "blog/internal/providers/helpers"
	"blog/pkg/converters"
	"blog/pkg/errors"
	"blog/pkg/html"
	"blog/pkg/keepFormData"
	"blog/pkg/sessions"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	// Get the article ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError,
			"templates/errors/html/500",
			gin.H{
				"title":   "Server Error",
				"message": "error converting the ID",
			})
		return
	}
	// Find the article from the DB
	article, err := controller.articleService.GetSingleArticle(id)

	// If the article not found, show error page
	if err != nil {
		html.Render(c, http.StatusInternalServerError,
			"templates/errors/html/404",
			gin.H{
				"title":   "Not Found",
				"message": err.Error(),
			})
		return
	}

	// IF the article found, render article template
	html.Render(c, http.StatusOK, "modules/article/html/item", gin.H{
		"title":   "Show Article",
		"article": article,
	})
}

func (controller *Controller) Create(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/article/html/create", gin.H{
		"title": "Create Article Page",
	})
}

func (controller *Controller) CreateHandle(c *gin.Context) {
	var request articles.CreateRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		keepFormData.Init()
		keepFormData.SetFromData(c)
		sessions.Set(c, "formData", converters.UrlValuesToString(keepFormData.Get()))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	imgName, err := ProviderHelpers.SaveUploadFile(c, "assets/img/demopic/")

	if err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		keepFormData.Init()
		keepFormData.SetFromData(c)
		sessions.Set(c, "formData", converters.UrlValuesToString(keepFormData.Get()))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	user := UserHelpers.Auth(c)
	article, err := controller.articleService.CreateArticle(request, user, imgName)

	if err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	log.Printf("The Article \"%s\" created successfully\n", article.Title)
	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}
