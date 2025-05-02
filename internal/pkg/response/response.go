package response

import "github.com/labstack/echo/v4"

type response struct {
	Status  int         `json:"status" example:"200"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data,omitempty" swaggerignore:"true"`
}

func SuccessResponse(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, response{
		Status:  status,
		Message: "success",
	})
}

func SuccessResponseWithCustomMessage(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, status int, message string) error {
	return c.JSON(status, response{
		Status:  status,
		Message: message,
		Data:    nil,
	})
}
