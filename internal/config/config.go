package config

import (
	"os"
)

var Config *Cfg

type Cfg struct {
	SMTPConfig SMTPConfig
}

type SMTPConfig struct {
	Server   string
	Port     string
	Username string
	Password string
	From     string
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
	}
}

func GetSMTPConfig() *SMTPConfig {
	return &Config.SMTPConfig
}
