package api

import (
	"time"

	"api-server/internal/app/article/domain"
)

type Article struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Source string `json:"source"`

	Body string `json:"body"`

	PublishedDate time.Time `json:"publishedDate"`
}

func MapArticleResponse(entity *domain.Article) Article {
	return Article{
		ID:            entity.ID,
		Title:         entity.Title,
		Author:        entity.Author,
		Source:        entity.Source,
		Body:          entity.Body,
		PublishedDate: entity.PublishedDate,
	}
}
