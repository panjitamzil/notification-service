package api

import (
	"notification-service/internal/pkg/validator"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetupRouter() *echo.Echo {
	e := echo.New()
	e.Use(LoggingMiddleware)
	e.Validator = validator.New()

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/send", SendNotificationHandler)

	return e
}
