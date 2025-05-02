package api

import (
	"log"

	"github.com/labstack/echo/v4"
)

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Request received:", c.Request().URL)
		return next(c)
	}
}
