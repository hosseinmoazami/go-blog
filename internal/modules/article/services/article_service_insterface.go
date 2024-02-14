package services

import (
	"blog/internal/modules/article/requests/articles"
	ArticleResponse "blog/internal/modules/article/responses"
	UserResponse "blog/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	CreateArticle(request articles.CreateRequest, user UserResponse.User, imgName string) (ArticleResponse.Article, error)
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	GetSingleArticle(ID int) (ArticleResponse.Article, error)
}
