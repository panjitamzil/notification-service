package email

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"notification-service/internal/config"
)

func SendEmail(to []string, subject, body string) error {
	cfg := config.GetSMTPConfig()

	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Server)
	address := fmt.Sprintf("%s:%s", cfg.Server, cfg.Port)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		cfg.From, strings.Join(to, ","), subject, body)

	err := smtp.SendMail(address, auth, cfg.From, to, []byte(msg))
	if err != nil {
		log.Printf("Failed to send email to %v: %v", to, err)
		return err
	}

	log.Printf("Email sent successfully to %v", to)
	return nil
}
