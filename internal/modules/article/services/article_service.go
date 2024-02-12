package services

import (
	ArticleRepository "blog/internal/modules/article/repositories"
	ArticleResponse "blog/internal/modules/article/responses"
	"errors"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (articleService *ArticleService) GetSingleArticle(ID int) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article
	article := articleService.articleRepository.Item(ID)
	if article.ID == 0 {
		return response, errors.New("article not found")
	}
	return ArticleResponse.ToArticle(article), nil
}

func (articleService *ArticleService) GetFeaturedArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(4)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(6)
	return ArticleResponse.ToArticles(articles)
}
