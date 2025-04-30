package api

import (
	"net/http"
	"notification-service/internal/api/models"
	"notification-service/internal/channels/email"

	"github.com/labstack/echo/v4"
)

func SendNotificationHandler(c echo.Context) error {
	var req models.NotificationRequest
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid payload")
	}

	err = email.SendEmail(req.To, req.Subject, req.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to send email")
	}

	return c.String(http.StatusOK, "Success")
}
