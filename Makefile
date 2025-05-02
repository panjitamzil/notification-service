all: build run

build:
	@echo "Building notification-service..."
	@go build -o notification-service cmd/main.go

run: build
	@echo "Running notification-service..."
	@./notification-service

test:
	@echo "Running unit tests..."
	@go test ./... -v

docker-build:
	@echo "Building Docker image..."
	@docker build -t notification-service .

docker-run: docker-build
	@echo "Running Docker container..."
	@docker run -p 8080:8080 notification-service

clean:
	@echo "Cleaning up..."
	@rm -f notification-service
	@docker rmi notification-service || true

help:
	@echo "Makefile commands:"
	@echo "  make build       - Build the application"
	@echo "  make run         - Build and run the application"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run  - Build and run Docker container"
	@echo "  make clean       - Clean up build artifacts"
	@echo "  make help        - Show this help message"