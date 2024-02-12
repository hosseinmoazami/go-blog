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
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error converting the ID"})
		return
	}
	// Find the article from the DB
	article, err := controller.articleService.GetSingleArticle(id)

	// If the article not found, show error page
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	// IF the article found, render article template
	html.Render(c, http.StatusOK, "modules/article/html/article", gin.H{
		"title":   "Article Page",
		"article": article,
	})
	// c.JSON(http.StatusOK, gin.H{"article": article})

}
