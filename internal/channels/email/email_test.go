package email_test

import (
	"notification-service/internal/channels/email"
	"notification-service/internal/config"
	"testing"
)

func TestEmailNotifier_Send_InvalidServer(t *testing.T) {
	cfg := &config.SMTPConfig{
		Server:   "invalid.local",
		Port:     "2525",
		Username: "user",
		Password: "pass",
		From:     "from@example.com",
	}
	notifier := email.EmailNotifier{Config: cfg}
	err := notifier.Send([]string{"to@example.com"}, "subject", "body")
	if err == nil {
		t.Error("expected error for invalid SMTP server, got nil")
	}
}
