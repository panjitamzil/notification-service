package notification

import (
	"context"
	"log"
	"notification-service/internal/config"
	"notification-service/internal/models"
	"sync"
	"time"
)

var (
	taskQueue       = make(chan NotificationTask, 100)
	numberOfWorkers = 5
	wg              sync.WaitGroup
)

type NotificationTask struct {
	Type    string
	Request models.NotificationRequest
}

func StartWorkers() {
	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}
}

func worker(id int) {
	defer wg.Done()

	for task := range taskQueue {
		log.Printf("Worker %d started processing %s for %s", id, task.Type, task.Request.To)

		var notifier Notifier
		var err error
		switch task.Type {
		case "email":
			smtpConfig := config.GetSMTPConfig()
			notifier, err = NewNotifier("email", smtpConfig)
		case "sms":
			apiKey := config.GetSMSAPIKey()
			notifier, err = NewNotifier("sms", apiKey)
		default:
			log.Printf("Worker %d: Unsupported notification type: %s", id, task.Type)
			continue
		}

		if err != nil {
			log.Printf("Worker %d: Failed to create Notifier: %v", id, err)
			continue
		}

		success := false

		for attempt := 1; attempt <= 4; attempt++ {
			log.Printf("Attempt %d for %s (%s)", attempt, task.Request.To, task.Type)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			result := make(chan error, 1)

			go func() {
				err := notifier.Send(task.Request.To, task.Request.Subject, task.Request.Body)
				result <- err
			}()

			select {
			case err := <-result:
				if err == nil {
					log.Printf("Worker %d successfully sent %s to %s on attempt %d", id, task.Type, task.Request.To, attempt)
					success = true
					break
				} else {
					log.Printf("Worker %d failed to send %s to %s on attempt %d: %v", id, task.Type, task.Request.To, attempt, err)
				}
			case <-ctx.Done():
				log.Printf("Worker %d timeout for %s (%s) on attempt %d", id, task.Request.To, task.Type, attempt)
			}

			if !success && attempt < 4 {
				log.Printf("Worker %d will retry after 2 seconds", id)
				time.Sleep(2 * time.Second)
			}
		}

		if !success {
			log.Printf("Worker %d failed to send %s to %s after 4 attempts", id, task.Type, task.Request.To)
		}
	}
}

func AddTask(task NotificationTask) {
	taskQueue <- task
}
