package sms

import (
	"context"
	"fmt"
	"log"

	"github.com/getbrevo/brevo-go/lib"
)

type SMSNotifier struct {
	APIKey string
}

func (s *SMSNotifier) Send(to []string, subject, body string) error {
	var (
		sendErrors []error
		ctx        = context.Background()
	)

	config := lib.NewConfiguration()
	config.AddDefaultHeader("api-key", s.APIKey)
	client := lib.NewAPIClient(config)

	for _, recipient := range to {
		sendTransacSms := lib.SendTransacSms{
			Sender:    "MuhammadTest",
			Recipient: recipient,
			Content:   fmt.Sprintf("%s - %s", subject, body),
			Type_:     "transactional",
		}

		_, _, err := client.TransactionalSMSApi.SendTransacSms(ctx, sendTransacSms)
		if err != nil {
			log.Printf("Failed to send SMS to %s: %v", recipient, err)
			sendErrors = append(sendErrors, err)
			continue
		}

		log.Printf("SM sent successfully to %s", recipient)
	}

	if len(sendErrors) > 0 {
		return fmt.Errorf("Failed to send SMS to %d recipient(s)", len(sendErrors))
	}

	return nil
}
