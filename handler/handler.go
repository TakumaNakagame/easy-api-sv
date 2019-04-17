package handler

import (
    "net/http"
    "github.com/labstack/echo"
    // "fmt"
)

var healthStatus = false  
var healthMsg = "NG"

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello World")
    }
}

func Health() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.String(http.StatusOK, healthMsg)
    }
}

func Update() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.String(http.StatusOK, "update status => " + healthMsg + " " )
    }
}