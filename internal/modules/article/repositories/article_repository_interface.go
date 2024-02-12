package repositories

import articleModel "blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []articleModel.Article
}
