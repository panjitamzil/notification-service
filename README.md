# Notification Service
Notification Service is a system that enables the sending of notifications through various channels such as email and SMS. It is designed to be scalable, extensible, and easily integrated with other systems. By leveraging Go's concurrency features, the service efficiently handles multiple requests in parallel.

## Features
- Send notifications via email and SMS.
- Supports adding other notification channels in the future with minimal refactoring.
- Uses a worker pool to handle notifications in parallel.

## Requirements
- Go 1.22 or later.
- An SMTP account (e.g., Brevo) for sending emails.
- An API Key for the SMS service.

## Installation
1. Clone or download this repository.
2. Navigate to the project directory:
```
cd notification-service
```
3. Install Dependencies:
```
go mod tidy
```
4. Configure Environment: Rename `.env.copy` to `.env` in the root directory and fill the SMTP and SMS configurations (e.g., SMTP host, port, credentials).
5. Build the Application using the provided Makefile:
```
make build
```
6. Run the Application using:
```
make run
```

## Docker Setup
1. Ensure Docker and Docker Compose are installed on your system.
2. Use the Makefile to build the Docker image:
```
make docker-build
```
3. Launch the containerized application:
```
make docker-run
```

## How to Send a Test Email
To send a test email, use the API endpoint provided by the service.
| Method (Endpoint)                 | Description                              |
|-----------------------------------|------------------------------------------|
| `POST /send`                      | Send a notification                      |

Sample Payload:
```
{
  "to": ["recipient@example.com"],
  "subject": "Test Subject",
  "body": "This is a test notification",
  "type": "email"
}
```

Sample cURL:
```
curl -X POST http://localhost:8080/send \
     -H "Content-Type: application/json" \
     -d '{"to": ["recipient@example.com"], "subject": "Test Subject", "body": "This is a test notification", "type": "email"}'
```

Response:
```
{
  "status": 200,
  "message": "Success"
}
```

## Swagger Documentation
Access the interactive API documentation at `http://localhost:8080/swagger/index.html`.

## Design Considerations and Decisions
### Architecture
The architecture implements the Strategy and Factory patterns to provide flexibility. This structure makes it easy to add new notification channels, such as switching from email to Slack, without altering the core code. The benefit is easier future development. The trade-off is that it requires more initial effort to set up, but it is worth it for long-term growth. Additionally, a Worker Pool is used to process notifications in parallel, like multiple workers handling tasks at once. This speeds up the process but requires careful management to avoid conflicts. Go's tools help handle this.

### Concurrency
To handle multiple tasks simultaneously, the system leverages Go's Goroutines and Channels. This is like opening multiple work lanes so notifications can be sent at the same time without waiting. The advantage is faster performance, even with many requests. The risk is potential data conflicts if not managed well, but Go's features make it manageable.

### Validation
Input validation is performed using tags in the code to check incoming data. This is similar to checking tickets before entering a cinema to prevent issues inside. The benefit is a more secure and stable program. The trade-off is a slight decrease in processing speed, but the security gained is far greater.

### Documentation
API documentation is created with Swagger, which serves as an interactive guide. This makes it easy for users to understand and try the program. The benefit is a better user experience. The challenge is the need to keep the documentation updated, but it is highly helpful overall.

### Containerization
Docker is used to package the program in a container, so it can run the same anywhere. This is like carrying a mini-house that can be used anywhere. The advantage is no compatibility issues across systems. The trade-off is the need for additional knowledge about Docker, but it is very useful for production deployment.