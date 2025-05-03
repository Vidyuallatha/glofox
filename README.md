
# Glofox Class Booking API

This project provides a simple API to manage fitness classes and member bookings for a studio. The API supports creating classes and making bookings for those classes. All data is stored in memory.

## Features

- **Create a Class**: Allows the studio owner to create new classes with basic details (class name, start date, end date, and capacity).
- **Book a Class**: Allows a member to book a class by providing their name and the date they wish to attend.

## Prerequisites

- Go 1.18 or higher
- `go mod` for managing dependencies

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/Vidyuallatha/glofox.git
    cd glofox
    ```

2. Install the dependencies:

    ```bash
    go mod tidy
    ```

## Running the Server

To start the server, run the following command:

```bash
go run src/main.go
```

The server will start on `http://localhost:9000`.

## API Endpoints

### 1. **Create a Class**
- **Endpoint**: `POST /classes`
- **Request Body**:
    ```json
    {
        "class_name": "Yoga",
        "start_date": "2025-05-03T10:00:00Z",
        "end_date": "2025-05-03T11:00:00Z",
        "capacity": 20
    }
    ```
- **Response**:
  - Status Code: `201 Created`
  - Response Body:
      ```json
      {
          "class_name": "Yoga",
          "start_date": "2025-05-03T10:00:00Z",
          "end_date": "2025-05-03T11:00:00Z",
          "capacity": 20
      }
      ```

### 2. **Book a Class**
- **Endpoint**: `POST /bookings`
- **Request Body**:
    ```json
    {
        "name": "John Doe",
        "date": "2025-05-03T10:00:00Z"
    }
    ```
- **Response**:
  - Status Code: `201 Created`
  - Response Body:
      ```json
      {
          "name": "John Doe",
          "date": "2025-05-03T10:00:00Z"
      }
      ```

## Running Tests

To run tests for the project, use the following command:

```bash
go test ./...
```

This will execute all the tests in the `tests` folder and provide a detailed output for each test case.

## Test Coverage

To generate a test coverage report, run:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

This will generate a code coverage report in HTML format that you can view in your browser.

## Directory Structure

```
glofox/
├── go.mod                          # Go module definition file
├── openapi.yml                     # OpenAPI Specification for the API
└── src/
    ├── components/                 # Contains the API business logic
    │   ├── tests/
    │   │   ├── bookings_test.go    # Tests for booking-related components
    │   │   └── classes_test.go     # Tests for class-related components
    │   ├── bookings.go             # Logic for creating bookings
    │   └── classes.go              # Logic for creating classes
    ├── controllers/
    │   ├── bookings_controller.go  # Handles API requests related to bookings
    │   └── classes_controller.go   # Handles API requests related to classes
    ├── entities/                   # Contains the data models and repositories
    │   ├── tests/
    │   │   ├── bookings_test.go    # Tests for booking entities/models
    │   │   └── classes_test.go     # Tests for class entities/models
    │   ├── bookings.go             # Booking-related entities and repository
    │   └── classes.go              # Class-related entities and repository
    └── utils/
    │    └── utils.go               # Utility functions used across the application
    └── main.go                     # The main entry point of the API

```
