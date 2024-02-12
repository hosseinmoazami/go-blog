package controllers

import (
	articleRepository "blog/internal/modules/article/repositories"
	"blog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleRepository articleRepository.ArticleRepositoryInterface
}

func New() *Controller {
	return &Controller{
		articleRepository: articleRepository.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title": "Home Page",
	})
	// c.JSON(http.StatusOK, controller.articleRepository.List(1))
}
