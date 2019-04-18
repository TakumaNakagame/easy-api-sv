package handler

import (
    "net/http"
    "github.com/labstack/echo"
    // "fmt"
)

var healthStatus = "OK"

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.String(http.StatusOK, "CheckStatus: /health\nChangeStatsu: /update")
    }
}

func Health() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.String(http.StatusOK, healthStatus)
    }
}

func Update() echo.HandlerFunc {
    return func(c echo.Context) error {
        if healthStatus == "OK"{
            healthStatus = "NG"
        }else if healthStatus == "NG"{
            healthStatus = "OK"
        }else{
            healthStatus = "NG"
        }
        return c.String(http.StatusOK, "update status => " + healthStatus)
    }
}