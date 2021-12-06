package persistence

import (
	"gorm.io/gorm"

	"api-server/internal/app/article/domain"
)

type ArticleRepository struct {
	source *gorm.DB
}

func (r *ArticleRepository) GetArticleByID(articleID string) (*domain.Article, error) {
	id, err := marshalID(articleID)
	if err != nil {
		return nil, err
	}

	var model Article
	err = r.source.Where("id = ?", string(id)).Find(&model).Error
	if err != nil {
		return nil, err
	}

	ret, err := model.toDomain()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (r *ArticleRepository) SaveArticle(article *domain.Article) error {
	model, err := NewArticleModel(article)
	if err != nil {
		return err
	}
	return r.source.Save(model).Error
}

func NewArticleRepository(source *gorm.DB) *ArticleRepository {
	return &ArticleRepository{source: source}
}
