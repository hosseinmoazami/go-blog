package responses

import (
	articleModel "blog/internal/modules/article/models"
	userResponse "blog/internal/modules/user/responses"
	"fmt"
)

type Article struct {
	ID        uint
	Title     string
	Content   string
	Image     string
	CreatedAt string
	User      userResponse.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article articleModel.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Image:     fmt.Sprintf("/assets/img/demopic/%d.jpg", article.ID),
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      userResponse.ToUser(article.User),
	}
}

func ToArticles(articles []articleModel.Article) Articles {
	var response Articles
	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}
	return response
}
