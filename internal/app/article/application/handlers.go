package application

import (
	"api-server/internal/app/article/domain"
	"fmt"
)

type Handler interface {
	GetArticleByID(id string) (*domain.Article, error)
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

func NewArticleHandler() Handler {
	return &articleHandler{}
}
