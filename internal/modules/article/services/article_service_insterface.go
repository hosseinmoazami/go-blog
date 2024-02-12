package services

import (
	ArticleResponse "blog/internal/modules/article/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	GetSingleArticle(ID int) (ArticleResponse.Article, error)
}
