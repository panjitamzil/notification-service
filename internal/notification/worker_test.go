package notification_test

import (
	"testing"
	"time"

	"notification-service/internal/models"
	"notification-service/internal/notification"
)

func TestStartWorkers_DoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("StartWorkers panicked: %v", r)
		}
	}()
	notification.StartWorkers()
}

func TestAddTask_DoesNotBlock(t *testing.T) {
	task := notification.NotificationTask{Type: "unknown", Request: models.NotificationRequest{}}
	done := make(chan struct{})
	go func() {
		notification.AddTask(task)
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Error("AddTask blocked")
	}
}
