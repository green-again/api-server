package server

import (
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"

	article "api-server/internal/app/article/presentation/api"
)

func InitRoutes(router *echo.Echo, articleController *article.ArticleController) {
	registerV1APIRoutes(router, articleController)
	registerSwagger(router)
}

func registerV1APIRoutes(router *echo.Echo, articleController *article.ArticleController) {
	v1 := router.Group("api/v1")
	{
		v1.POST("/articles", articleController.PostArticle)
		v1.GET("/articles/:id", articleController.GetArticle)
	}
}

func registerSwagger(router *echo.Echo) {
	router.GET("/swagger/*", echoSwagger.WrapHandler)
}
