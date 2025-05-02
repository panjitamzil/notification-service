package api

import (
	"fmt"
	"net/http"
	"notification-service/internal/config"
	"notification-service/internal/models"
	"notification-service/internal/notification"
	"notification-service/internal/pkg/response"
	"notification-service/internal/pkg/validator"

	"github.com/labstack/echo/v4"
)

// @Summary Send a notification
// @Description Send a notification via specified channel
// @Accept json
// @Produce json
// @Param notification body models.NotificationRequest true "Notification request payload"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /send [post]
func SendNotificationHandler(c echo.Context) error {
	var req models.NotificationRequest
	err := validator.BindRequestAndValidate(c, &req)
	if err != nil {
		if he, ok := err.(*echo.HTTPError); ok {
			return response.Error(c, he.Code, fmt.Sprintf("%v", he.Message))
		}

		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	var deps []interface{}
	if req.Type == "email" {
		deps = append(deps, config.GetSMTPConfig())
	} else if req.Type == "sms" {
		deps = append(deps, config.GetSMSAPIKey())
	}

	notifier, err := notification.NewNotifier(req.Type, deps...)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	err = notifier.Send(req.To, req.Subject, req.Body)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.OK(c, nil)
}
