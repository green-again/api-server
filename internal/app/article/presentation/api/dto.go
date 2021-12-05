package api

import (
	"time"

	"api-server/internal/app/article/domain"
)

type Article struct {
	ID     string `json:"id" form:"id"`
	Title  string `json:"title" form:"title" validate:"required"`
	Author string `json:"author" form:"author" validate:"required"`
	Source string `json:"source" form:"source" validate:"required"`

	Body string `json:"body" form:"body" validate:"required"`

	Status int `json:"status" form:"status" validate:"required"`

	PublishedDate *time.Time `json:"publishedDate"`
}

func (a *Article) MapDomain() domain.Article {
	return domain.NewArticle(a.ID, a.Title, a.Author, a.Source, a.Body)
}

func MapArticleResponse(entity *domain.Article) Article {
	return Article{
		ID:            entity.ID(),
		Title:         entity.Title(),
		Author:        entity.Author(),
		Source:        entity.Source(),
		Body:          entity.Body(),
		PublishedDate: entity.PublishedDate(),
	}
}
