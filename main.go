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

	e.GET("/hello", handler.MainPage())

	e.Start(":8080")
}
