package api

import (
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"

	"api-server/internal/app/article/application"
	article "api-server/internal/app/article/presentation/api"
)

func InitRoutes(router *echo.Echo) {
	registerV1APIRoutes(router)
	registerSwagger(router)
}

func registerV1APIRoutes(router *echo.Echo) {
	v1 := router.Group("api/v1")
	{
		registerArticleAPIs(v1)
	}
}

func registerSwagger(router *echo.Echo) {
	router.GET("/swagger/*", echoSwagger.WrapHandler)
}

func registerArticleAPIs(rg *echo.Group) {
	con := article.NewController(
		application.NewArticleHandler(),
		article.NewGetArticleRequestValidator(),
	)
	rg.GET("/articles/:id", con.GetArticle)
}
