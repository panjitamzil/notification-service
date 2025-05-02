package config

import (
	"os"
)

var Config *Cfg

type Cfg struct {
	SMTPConfig SMTPConfig
	SMSConfig  SMSConfig
}

type SMTPConfig struct {
	Server   string
	Port     string
	Username string
	Password string
	From     string
}

type SMSConfig struct {
	APIKey string `json:"SMS_API_KEY"`
}

func Load() {
	Config = &Cfg{
		SMTPConfig: SMTPConfig{
			Server:   os.Getenv("SMTP_SERVER"),
			Port:     os.Getenv("SMTP_PORT"),
			Username: os.Getenv("SMTP_USERNAME"),
			Password: os.Getenv("SMTP_PASSWORD"),
			From:     os.Getenv("SMTP_SENDER"),
		},
		SMSConfig: SMSConfig{
			APIKey: os.Getenv("SMS_API_KEY"),
		},
	}
}

func GetSMTPConfig() *SMTPConfig {
	return &Config.SMTPConfig
}

func GetSMSAPIKey() string {
	return Config.SMSConfig.APIKey
}
