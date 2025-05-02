package sms_test

import (
	"notification-service/internal/channels/sms"
	"testing"
)

func TestSMSNotifier_Send_NoRecipients(t *testing.T) {
	notifier := sms.SMSNotifier{APIKey: "dummy-api-key"}
	err := notifier.Send([]string{}, "subject", "body")
	if err != nil {
		t.Errorf("expected no error when no recipients, got %v", err)
	}
}
