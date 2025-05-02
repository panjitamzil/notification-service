package notification_test

import (
	"testing"

	"notification-service/internal/channels/email"
	"notification-service/internal/channels/sms"
	"notification-service/internal/config"
	"notification-service/internal/notification"
	"notification-service/internal/pkg/pkg_error"
)

func TestNewNotifier_Email(t *testing.T) {
	cfg := &config.SMTPConfig{}
	notifier, err := notification.NewNotifier("email", cfg)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if _, ok := notifier.(*email.EmailNotifier); !ok {
		t.Errorf("expected EmailNotifier, got %T", notifier)
	}
}

func TestNewNotifier_SMS(t *testing.T) {
	notifier, err := notification.NewNotifier("sms", "apikey")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if _, ok := notifier.(*sms.SMSNotifier); !ok {
		t.Errorf("expected SMSNotifier, got %T", notifier)
	}
}

func TestNewNotifier_MissingConfig(t *testing.T) {
	_, err := notification.NewNotifier("email")
	if err == nil || err.Error() != pkg_error.MISSING_CONFIG {
		t.Errorf("expected missing config error, got %v", err)
	}
}

func TestNewNotifier_InvalidConfig(t *testing.T) {
	_, err := notification.NewNotifier("email", 123)
	if err == nil || err.Error() != pkg_error.INVALID_CONFIG {
		t.Errorf("expected invalid config error, got %v", err)
	}
}

func TestNewNotifier_Unsupported(t *testing.T) {
	_, err := notification.NewNotifier("push", nil)
	if err == nil || err.Error() != pkg_error.UNSUPPORTED_TYPE {
		t.Errorf("expected unsupported type error, got %v", err)
	}
}
