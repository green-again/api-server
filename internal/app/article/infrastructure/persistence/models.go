package persistence

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"api-server/internal/app/article/domain"
)

type Article struct {
	ID     []byte `gorm:"primary_key;size:16"`
	Title  string `gorm:"type:varchar(255);not null"`
	Author string `gorm:"type:varchar(255);not null"`
	Source string `gorm:"type:varchar(255)"`

	Body string `gorm:"type:longtext;not null"`

	Status int `gorm:"type:tinyint;not null"`

	PublishedDate *time.Time `gorm:"index:,priority:1,sort:desc"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Article) setID(articleID string) error {
	id, err := marshalID(articleID)
	if err != nil {
		return fmt.Errorf("[persistence.Article.setID()] err: %w", err)
	}
	a.ID = id
	return nil
}

func (a *Article) toDomain() (*domain.Article, error) {
	uid, err := uuid.FromBytes(a.ID)
	if err != nil {
		return nil, fmt.Errorf("[persistence.Article.toDomain()] article id: %s, err: %w", a.ID, UnmarshalBinaryError)
	}

	ret := domain.NewArticle(uid.String(), a.Title, a.Author, a.Source, a.Body, a.Status)
	return &ret, nil
}

func marshalID(articleID string) ([]byte, error) {
	id, err := uuid.Parse(articleID)
	if err != nil {
		return nil, fmt.Errorf("[persistence.marshalID()] err: %w", ParseUUIDError)
	}

	return id.MarshalBinary()
}

func NewArticleModel(article *domain.Article) (*Article, error) {
	ret := &Article{
		Title: article.Title(),
		Author: article.Author(),
		Source: article.Source(),
		Body: article.Body(),
		Status: article.Status(),
		PublishedDate: article.PublishedDate(),
	}

	err := ret.setID(article.ID())
	if err != nil {
		return nil, fmt.Errorf("[persistence.NewArticleModel()] err: %w", err)
	}

	return ret, nil
}
