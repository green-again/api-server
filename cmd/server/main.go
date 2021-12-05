package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	_ "api-server/docs"
	"api-server/internal/app/article/application"
	"api-server/internal/app/article/infrastructure/persistence"
	"api-server/internal/app/article/presentation/api"
	"api-server/internal/pkg/config"
	"api-server/internal/pkg/db"
	"api-server/internal/pkg/server"
)

// @title Green Again API server
// @version 1.0
// @description This is a green again backend api server
// BasePath /v1
func main() {
	config.InitConfig()
	var (
		dbEngine = db.ConnectDB()

		articleRepository = persistence.NewArticleRepository(dbEngine)
		articleHandler = application.NewArticleHandler(articleRepository)
		articleValidator = api.NewGetArticleRequestValidator()
		articleController = api.NewController(articleHandler, articleValidator)

		server = server.NewHTTPServer(articleController)
	)

	server.GET("/", HelloWorld)
	server.Logger.Fatal(server.Start(":8000"))
}

// HelloWorld godoc
// @Summary Print HelloWorld
// @Success 200 {string} result "Hello, World"
// @Router / [get]
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Green-Again! Let's Start!")
}
