package domain

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Article struct {
	id     string
	title  string
	author string
	source string

	body string

	status int

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

func (a *Article) Status() int {
	return a.status
}

func (a *Article) PublishedDate() *time.Time {
	return a.publishedDate
}

func (a *Article) GenerateID() error {
	if a.id != "" {
		return fmt.Errorf("id is already exists")
	}
	id, _ := uuid.NewUUID()
	a.id = id.String()
	return nil
}

func NewArticle(id, title, author, source, body string, status int) Article {
	return Article{
		id: id,
		title: title,
		author: author,
		source: source,
		status: status,
		body: body,
	}
}
