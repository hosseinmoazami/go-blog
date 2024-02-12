package repositories

import ArticleModel "blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	Item(ID int) ArticleModel.Article
	List(limit int) []ArticleModel.Article
}
