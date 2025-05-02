package api

import (
	"net/http"
	"notification-service/internal/config"
	"notification-service/internal/models"
	"notification-service/internal/notification"
	"notification-service/internal/pkg/pkg_error"
	"notification-service/internal/pkg/response"

	"github.com/labstack/echo/v4"
)

func SendNotificationHandler(c echo.Context) error {
	var req models.NotificationRequest
	err := c.Bind(&req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, pkg_error.INVALID_PAYLOAD)
	}

	var deps []interface{}
	if req.Type == "email" {
		deps = append(deps, config.GetSMTPConfig())
	} else if req.Type == "sms" {
		deps = append(deps, config.GetSMSAPIKey())
	}

	notifier, err := notification.NewNotifier(req.Type, deps...)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	err = notifier.Send(req.To, req.Subject, req.Body)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, nil)
}
