package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "api-server/docs"
)

// @title Green Again API server
// @version 1.0
// @description This is a green again backend api server
// BasePath /v1
func main() {
	e := echo.New()
	e.GET("/", HelloWorld)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8000"))
}

// HelloWorld godoc
// @Summary Print HelloWorld
// @Success 200 {string} result "Hello, World"
// @Router / [get]
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Green-Again!")
}
