package application

import (
	"fmt"

	"api-server/internal/app/article/domain"
)

type Handler interface {
	GetArticleByID(id string) (*domain.Article, error)
	CreateArticle(title, author, source, body string) (*domain.Article, error)
}

type articleHandler struct{
	repo domain.ArticleRepository
}

func (h *articleHandler) GetArticleByID(id string) (*domain.Article, error) {
	article, err := h.repo.GetArticleByID(id)
	if err != nil {
		return nil, fmt.Errorf("GetArticleByID error. %w", err)
	}

	return article, nil
}

func (h *articleHandler) CreateArticle(title, author, source, body string) (*domain.Article, error) {
	article := domain.NewArticle("", title, author, source, body)
	article.GenerateID()

	err := h.repo.SaveArticle(&article)
	if err != nil {
		return nil, fmt.Errorf("CreateArticle error. %w", err)
	}

	return &article, nil
}

func NewArticleHandler() Handler {
	return &articleHandler{}
}
