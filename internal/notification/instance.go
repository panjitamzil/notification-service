package notification

import (
	"errors"
	"notification-service/internal/channels/email"
	"notification-service/internal/channels/sms"
	"notification-service/internal/config"
)

type Notifier interface {
	Send(to []string, subject, body string) error
}

func NewNotifier(notificationType string, dependencies ...interface{}) (Notifier, error) {
	switch notificationType {
	case "email":
		if len(dependencies) < 1 {
			return nil, errors.New("missing SMTP config")
		}
		smtpConfig, ok := dependencies[0].(*config.SMTPConfig)
		if !ok {
			return nil, errors.New("invalid SMTP config type")
		}
		return &email.EmailNotifier{Config: smtpConfig}, nil

	case "sms":
		if len(dependencies) < 1 {
			return nil, errors.New("missing API key")
		}

		apiKey, ok := dependencies[0].(string)
		if !ok {
			return nil, errors.New("invalid API key type")
		}
		return &sms.SMSNotifier{APIKey: apiKey}, nil

	default:
		return nil, errors.New("unsupported notification type")
	}
}
