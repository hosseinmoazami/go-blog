package services

import (
	ArticleModel "blog/internal/modules/article/models"
	ArticleRepository "blog/internal/modules/article/repositories"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() []ArticleModel.Article {
	return articleService.articleRepository.List(4)
}

func (articleService *ArticleService) GetStoriesArticles() []ArticleModel.Article {
	return articleService.articleRepository.List(6)
}
