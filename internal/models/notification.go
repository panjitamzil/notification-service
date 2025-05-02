package models

import "context"

type NotificationRequest struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
	Type    string   `json:"type"`
}

type NotificationTask struct {
	Request    NotificationRequest
	RetryCount int
	MaxRetries int
	Ctx        context.Context
}
