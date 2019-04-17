package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"./handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.Hello())
	e.GET("/hello", handler.Hello())
	e.GET("/health",handler.Health())
	e.GET("/update",handler.Update())

	e.Start(":8080")
}