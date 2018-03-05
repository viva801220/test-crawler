package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {

	return c.String(http.StatusOK, "ok")
}

func Crash(c echo.Context) error {

	panic("test")
}
