package persistence

import (
	"time"

	"api-server/internal/app/article/domain"
)

type Article struct {
	ID     string `gorm:"primary_key;index:cursor,priority:2"`
	Title  string
	Author string
	Source string

	Body string

	Status int

	PublishedDate *time.Time `gorm:"index:cursor,priority:1,sort:desc"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewArticleModel(article *domain.Article) Article {
	return Article{
		ID: article.ID(),
		Title: article.Title(),
		Author: article.Author(),
		Source: article.Source(),
		Body: article.Body(),
		Status: article.Status(),
		PublishedDate: article.PublishedDate(),
	}
}
