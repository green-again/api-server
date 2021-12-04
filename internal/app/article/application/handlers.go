package application

import (
	"errors"

	"api-server/internal/app/article/domain"
)

type Handler interface {
	GetArticleByID(id string) (*domain.Article, error)
}

type articleHandler struct{}

func (h *articleHandler) GetArticleByID(id string) (*domain.Article, error) {
	return nil, errors.New("not implemented")
}

func NewArticleHandler() Handler {
	return &articleHandler{}
}
