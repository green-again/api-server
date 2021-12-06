package server

import (
	"api-server/internal/pkg/db"
	"github.com/labstack/echo/v4"

	articleService "api-server/internal/app/article/application"
	articleRepo "api-server/internal/app/article/infrastructure/persistence"
	articleAPI "api-server/internal/app/article/presentation/api"
	"api-server/internal/pkg/config"
	"api-server/internal/pkg/http"
)

func NewHTTPServer(cfg *config.Config) *echo.Echo {
	e := echo.New()
	e.Validator = http.NewRequestValidator()

	var (
		dbEngine = db.ConnectDatabase(cfg)

		articleRepository = articleRepo.NewArticleRepository(dbEngine)
		articleHandler = articleService.NewArticleHandler(articleRepository)
		articleValidator = articleAPI.NewGetArticleRequestValidator()
		articleController = articleAPI.NewController(articleHandler, articleValidator)
	)

	BuildRouter(e, articleController)

	return e
}
