package notification

import (
	"errors"
	"notification-service/internal/channels/email"
	"notification-service/internal/channels/sms"
	"notification-service/internal/config"
	"notification-service/internal/pkg/pkg_error"
)

type Notifier interface {
	Send(to []string, subject, body string) error
}

func NewNotifier(notificationType string, dependencies ...interface{}) (Notifier, error) {
	switch notificationType {
	case "email":
		if len(dependencies) < 1 {
			return nil, errors.New(pkg_error.MISSING_CONFIG)
		}
		smtpConfig, ok := dependencies[0].(*config.SMTPConfig)
		if !ok {
			return nil, errors.New(pkg_error.INVALID_CONFIG)
		}
		return &email.EmailNotifier{Config: smtpConfig}, nil

	case "sms":
		if len(dependencies) < 1 {
			return nil, errors.New(pkg_error.MISSING_CONFIG)
		}

		apiKey, ok := dependencies[0].(string)
		if !ok {
			return nil, errors.New(pkg_error.INVALID_CONFIG)
		}
		return &sms.SMSNotifier{APIKey: apiKey}, nil

	default:
		return nil, errors.New(pkg_error.UNSUPPORTED_TYPE)
	}
}
