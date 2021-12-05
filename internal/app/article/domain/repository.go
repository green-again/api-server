package domain

type ArticleRepository interface {
	GetArticleByID(id string) (*Article, error)
	SaveArticle(*Article) error
}
