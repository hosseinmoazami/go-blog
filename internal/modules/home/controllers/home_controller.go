package controllers

import (
	ArticleService "blog/internal/modules/article/services"
	"blog/pkg/html"
	"net/http"

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

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home Page",
		"featured": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),
	})

	// c.JSON(http.StatusOK, controller.articleRepository.List(1))

	// c.JSON(http.StatusOK, gin.H{
	// 	"featured": controller.articleService.GetFeaturedArticles(),
	// 	"stories":  controller.articleService.GetStoriesArticles(),
	// })
}
