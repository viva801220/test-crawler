package router

import (
	"test-crawler/controller"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func Routers(e *echo.Echo) {

	e.GET("/", controller.Hello)
	go func() {
		if err := e.Start(viper.GetString("HTTP_PORT")); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()
	return
}
