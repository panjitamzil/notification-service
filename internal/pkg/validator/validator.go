package validator

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"notification-service/internal/pkg/pkg_error"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var phoneRegex = regexp.MustCompile(`^\+?[0-9]{10,15}$`)

type CustomValidator struct {
	Validator *validator.Validate
}

func New() *CustomValidator {
	v := validator.New()
	_ = v.RegisterValidation("email_or_phone", validateEmailOrPhone)
	return &CustomValidator{Validator: v}
}

func validateEmailOrPhone(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if strings.Contains(val, "@") {
		err := validator.New().Var(val, "email")
		return err == nil
	}
	return phoneRegex.MatchString(val)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err == nil {
		return nil
	}

	var sb strings.Builder
	if ves, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ves {
			field := fe.Field()
			tag := fe.Tag()
			param := fe.Param()
			var detail string
			switch tag {
			case "required":
				detail = "is required"
			case "min":
				detail = fmt.Sprintf("must be at least %s characters", param)
			case "max":
				detail = fmt.Sprintf("must be at most %s characters", param)
			case "len":
				detail = fmt.Sprintf("must be %s characters long", param)
			case "email":
				detail = "must be a valid email address"
			case "email_or_phone":
				detail = "must be a valid email or phone number"
			case "oneof":
				detail = fmt.Sprintf("must be one of [%s]", param)
			default:
				detail = fmt.Sprintf("failed '%s' validation", tag)
			}
			if sb.Len() > 0 {
				sb.WriteString("; ")
			}
			sb.WriteString(fmt.Sprintf("%s %s", field, detail))
		}
		msg := sb.String()
		return echo.NewHTTPError(http.StatusBadRequest, msg)
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func BindRequestAndValidate(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, pkg_error.BAD_REQUEST)
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	return nil
}
