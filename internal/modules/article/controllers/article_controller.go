package controllers

import (
	ArticleService "blog/internal/modules/article/services"
	"blog/pkg/html"
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
	c.JSON(http.StatusFound, gin.H{"title": "Create Article"})
}
