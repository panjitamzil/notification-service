package models

import "context"

type NotificationRequest struct {
	To      []string `json:"to" validate:"required,min=1,dive,email_or_phone" example:"test@gmail.com"`
	Subject string   `json:"subject" validate:"required" example:"This is subject"`
	Body    string   `json:"body" validate:"required" example:"This is body message"`
	Type    string   `json:"type" validate:"required,oneof=email sms" example:"email or sms"`
}

type NotificationTask struct {
	Request    NotificationRequest
	RetryCount int
	MaxRetries int
	Ctx        context.Context
}
