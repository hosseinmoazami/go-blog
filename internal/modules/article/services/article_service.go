package services

import (
	ArticleModel "blog/internal/modules/article/models"
	ArticleRepository "blog/internal/modules/article/repositories"
	"blog/internal/modules/article/requests/articles"
	ArticleResponse "blog/internal/modules/article/responses"
	UserResponse "blog/internal/modules/user/responses"
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

func (articleService *ArticleService) CreateArticle(request articles.CreateRequest, user UserResponse.User) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article
	var article ArticleModel.Article

	article.Title = request.Title
	article.Content = request.Content
	article.UserID = user.ID

	newArticle := articleService.articleRepository.Create(article)

	if newArticle.ID == 0 {
		return response, errors.New("error on creating the article")
	}
	return ArticleResponse.ToArticle(newArticle), nil
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
