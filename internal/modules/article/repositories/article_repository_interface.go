package repositories

import ArticleModel "blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
}
