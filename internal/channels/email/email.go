package email

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"notification-service/internal/config"
)

type EmailNotifier struct {
	Config *config.SMTPConfig
}

func (e *EmailNotifier) Send(to []string, subject, body string) error {
	auth := smtp.PlainAuth("", e.Config.Username, e.Config.Password, e.Config.Server)
	address := fmt.Sprintf("%s:%s", e.Config.Server, e.Config.Port)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		e.Config.From, strings.Join(to, ","), subject, body)

	err := smtp.SendMail(address, auth, e.Config.From, to, []byte(msg))
	if err != nil {
		log.Printf("Failed to send email to %v: %v", to, err)
		return err
	}

	log.Printf("Email sent successfully to %v", to)
	return nil
}
