package persistence

import (
	"fmt"
	"gorm.io/gorm"

	"api-server/internal/app/article/domain"
)

type ArticleRepository struct {
	source *gorm.DB
}

func (r *ArticleRepository) GetArticleByID(articleID string) (*domain.Article, error) {
	id, err := marshalID(articleID)
	if err != nil {
		return nil, fmt.Errorf("[persistence.ArticleRepository.GetArticleByID()] err: %w", err)
	}

	var model Article
	err = r.source.Where("id = ?", string(id)).Find(&model).Error
	if err != nil {
		return nil, fmt.Errorf("[persistence.ArticleRepository.GetArticleByID()] err: %w", err)
	}

	ret, err := model.toDomain()
	if err != nil {
		return nil, fmt.Errorf("[persistence.ArticleRepository.GetArticleByID()] err: %w", err)
	}

	return ret, nil
}

func (r *ArticleRepository) SaveArticle(article *domain.Article) error {
	model, err := NewArticleModel(article)
	if err != nil {
		return fmt.Errorf("[persistence.ArticleRepository.GetArticleByID()] err: %w", err)
	}
	return r.source.Save(model).Error
}

func NewArticleRepository(source *gorm.DB) *ArticleRepository {
	return &ArticleRepository{source: source}
}
