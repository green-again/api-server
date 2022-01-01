package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
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

func (a *Article) setStatus(status int) error {
	if status != statusDraft && status != statusPublished {
		return fmt.Errorf("[domain.Article.setStatus()] err: %w", InvalidStatusError)
	}
	if a.isPublished() && status == statusDraft {
		return fmt.Errorf("[domain.Article.setStatus()] err: %w", AlreadyPublishedError)
	}
	a.status = status

	return nil
}

func (a *Article) PublishedDate() *time.Time {
	return a.publishedDate
}

func (a *Article) GenerateID() {
	if a.id == "" {
		id, _ := uuid.NewUUID()
		a.id = id.String()
	}
}

func (a *Article) isPublished() bool {
	return a.status == statusPublished
}

func (a *Article) Update(newTitle, newAuthor, newSource, newBody string, newStatus int) error {
	a.title = newTitle
	a.author = newAuthor
	a.source = newSource
	a.body = newBody

	err := a.setStatus(newStatus)
	if err != nil {
		return fmt.Errorf("[domain.Article.Update()] err: %w", err)
	}

	return nil
}

func NewArticle(id, title, author, source, body string, status int) Article {
	if id == "" {
		ret := Article{
			title:  title,
			author: author,
			source: source,
			status: status,
			body:   body,
		}
		ret.GenerateID()
		return ret
	}

	return Article{
		id:     id,
		title:  title,
		author: author,
		source: source,
		status: status,
		body:   body,
	}
}
