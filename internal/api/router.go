package api

import (
	"github.com/labstack/echo/v4"
)

func SetupRouter() *echo.Echo {
	e := echo.New()
	e.Use(LoggingMiddleware)

	e.POST("/send", SendNotificationHandler)

	return e
}
