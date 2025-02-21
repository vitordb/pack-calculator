.PHONY: all run test build docker-build docker-run clean

# Default target: run the application
all: run

# Run the application locally
run:
	go run ./cmd/main.go

# Run unit tests for all packages
test:
	go test ./...

# Build the Go binary
build:
	go build -o pack-calculator ./cmd/main.go

# Build the Docker image
docker-build:
	docker build -t pack-calculator .

# Run the Docker container
up:
	docker run -d -p 8080:8080 --name pack-calculator pack-calculator

# Stop and remove the Docker container
down:
	docker stop pack-calculator && docker rm pack-calculator

# Tail the logs of the container
logs:
	docker logs -f pack-calculator

# Clean up the built binary
clean:
	rm -f pack-calculator