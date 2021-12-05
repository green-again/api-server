package server

import (
	article "api-server/internal/app/article/presentation/api"
	"api-server/internal/pkg/api"
	"github.com/labstack/echo/v4"
)

func NewHTTPServer(articleController *article.ArticleController) *echo.Echo {
	e := echo.New()
	e.Validator = api.NewRequestValidator()

	InitRoutes(e, articleController)

	return e
}
