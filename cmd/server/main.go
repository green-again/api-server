package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	_ "api-server/docs"
	"api-server/internal/pkg/config"
	"api-server/internal/pkg/server"
)

// @title Green Again API server
// @version 1.0
// @description This is a green again backend http server
// BasePath /v1
func main() {
	cfg := config.Make()

	server := server.NewHTTPServer(&cfg)
	server.GET("/", HelloWorld)
	server.Logger.Fatal(server.Start(":8000"))
}

// HelloWorld godoc
// @Summary Print HelloWorld
// @Success 200 {string} result "Hello, World"
// @BuildRouter / [get]
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Green-Again! Let's Start!")
}
