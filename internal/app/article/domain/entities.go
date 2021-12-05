package domain

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	id     string
	title  string
	author string
	source string

	body string

	publishedDate *time.Time
}

func (a *Article) ID() string {
	return a.id
}

func (a *Article) Title() string {
	return a.title
}

func (a *Article) Author() string {
	return a.author
}

func (a *Article) Source() string {
	return a.source
}

func (a *Article) Body() string {
	return a.body
}

func (a *Article) PublishedDate() *time.Time {
	return a.publishedDate
}

func NewArticle(title, author, source, body string) Article {
	return Article{
		id: uuid.NewString(),
		title: title,
		author: author,
		source: source,
		body: body,
	}
}
