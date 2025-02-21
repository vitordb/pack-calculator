
# Pack Calculator

This application calculates the minimum number of packages required to fulfill a requested amount of items. The solution may exceed the requested amount if necessary (since packages cannot be split). The project is built using Go with the Echo framework, includes unit tests, and is containerized with Docker.

---

## Project Structure

```
pack-calculator/
├── cmd/
│   └── main.go                           # Application entry point
├── internal/
│   ├── domain/
│   │   ├── calculator.go                 # Business logic (pure domain) for calculating packages
│   │   ├── calculator_test.go            # Unit tests for domain logic
│   │   └── ports/
│   │       └── dbInterface.go            # Domain contracts: DBInterface and Result type
│   ├── application/
│   │   ├── calculatorService.go         # Service to orchestrate domain logic and persistence
│   │   └── calculatorService_test.go    # Unit tests for service layer
│   └── infra/
│       ├── handlers/
│       │   └── http.go                   # HTTP handlers using Echo
│       └── repositories/
│           └── inMemoryRepository.go     # In-memory repository implementation (adapter)
├── ui/
│   ├── index.html                        # Simple front-end demo
│   └── style.css                         # CSS file for UI styling
├── Dockerfile                            # Docker build file
├── Makefile                              # Build and run commands
├── go.mod                                # Go module file
├── go.sum                                # Go dependencies file
└── README.md                             # Project documentation (this file)

```
---

## Prerequisites

Make sure you have the following tools installed on your local machine:

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [Make](https://www.gnu.org/software/make/)

---

## Running the Application

### Using the Makefile

The Makefile contains predefined commands to build and run the application. Here are the available commands:

- **Run the application locally** (start the app):
  ```bash
  make run
  ```

- **Build the Go binary** (compile the Go code):
  ```bash
  make build
  ```

- **Build the Docker image**:
  ```bash
  make docker-build
  ```

- **Run the Docker container** (after building the image):
  ```bash
  make docker-run
  ```

- **Check the logs of the running Docker container**:
  ```bash
  make logs
  ```

- **Clean up Docker containers and images**:
  ```bash
  make clean
  ```

- **Run unit tests**:
  ```bash
  make test
  ```

### Running Locally (Without Makefile)

If you don't want to use the Makefile, you can also run the application and tests manually:

1. **Run the application locally:**

   First, build and run the Go application using the following commands:

   ```bash
   go run ./cmd/main.go
   ```

2. **Open the application in your browser**:

   Once the application is running, open your browser and visit `http://localhost:8080/ui` to access the app ui.

---

### Running with Docker (Without Makefile)

If you prefer to run the application using Docker directly, follow these steps:

1. **Build the Docker image**:
   ```bash
   docker build -t pack-calculator .
   ```

2. **Run the Docker container**:
   ```bash
   docker run -p 8080:8080 pack-calculator
   ```

   The application will now be running on `http://localhost:8080`.

---

### Testing

To run unit tests, simply execute the following command:

```bash
go test ./...
```

This will run all the unit tests in the project.

---
