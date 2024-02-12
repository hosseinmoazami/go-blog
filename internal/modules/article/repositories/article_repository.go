package repositories

import (
	articleModel "blog/internal/modules/article/models"
	"blog/pkg/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []articleModel.Article {
	var articles []articleModel.Article
	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)
	return articles
}
