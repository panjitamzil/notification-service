package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIResponse struct {
	Status  int         `json:"status" example:"200"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data,omitempty" swaggerignore:"true"`
}

func Success(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func OK(c echo.Context, data interface{}) error {
	return Success(c, http.StatusOK, "success", data)
}

func Created(c echo.Context, data interface{}) error {
	return Success(c, http.StatusCreated, "created", data)
}

func Accepted(c echo.Context) error {
	return c.NoContent(http.StatusAccepted)
}

func Error(c echo.Context, status int, errMsg string) error {
	return c.JSON(status, APIResponse{
		Status:  status,
		Message: errMsg,
	})
}
