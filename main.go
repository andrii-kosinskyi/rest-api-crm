package main

import (
	"context"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "CRM/docs"

	"CRM/internal"
)

// @title Your API
// @version 1.0
// @description This is a sample server.

// @host localhost:1323
// @BasePath /
func main() {
	e := echo.New()
	ctx := context.Background()
	handlers := internal.NewHandler(ctx)
	internal.InitRoute(e, handlers)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
