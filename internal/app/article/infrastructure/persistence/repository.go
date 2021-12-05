package persistence

import (
	"gorm.io/gorm"

	"api-server/internal/app/article/domain"
)

type ArticleRepository struct {
	source *gorm.DB
}

func (r *ArticleRepository) GetArticleByID(id string) (*domain.Article, error) {
	return nil, nil
}

func (r *ArticleRepository) SaveArticle(article *domain.Article) error {
	model := NewArticleModel(article)
	return r.source.Save(model).Error
}

func NewArticleRepository(source *gorm.DB) *ArticleRepository {
	return &ArticleRepository{source: source}
}
