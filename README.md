# Employee CRUD API

A RESTful API for managing employee records, built with Go following Test-Driven Development (TDD) and a layered architecture.

## Features

- Full CRUD operations for employees (Create, Read, Update, Delete)
- SQLite database for data persistence
- Validation for all required fields (full name, job title, country)
- Positive salary validation
- Layered architecture: DAO → Service → Handler
- Comprehensive unit tests for all layers
- RESTful HTTP endpoints

## Architecture

The application follows a strict layered architecture:

1. **Model Layer** (`internal/model`): Defines the Employee entity
2. **DAO Layer** (`internal/dao`): Handles database operations
3. **Service Layer** (`internal/service`): Contains business logic and validation
4. **Handler Layer** (`internal/handler`): Manages HTTP request/response handling
5. **Main Application** (`cmd/api`): Wires everything together and starts the server

## Requirements

- Go 1.16 or higher
- CGO enabled (for SQLite support)

## Installation

```bash
go mod download
```

## Building

```bash
CGO_ENABLED=1 go build -o bin/api ./cmd/api
```

## Running

```bash
./bin/api
```

The server will start on port 8080. A SQLite database file `employees.db` will be created in the current directory.

## Running Tests

Run all tests:
```bash
CGO_ENABLED=1 go test ./...
```

Run tests with verbose output:
```bash
CGO_ENABLED=1 go test ./... -v
```

Run tests for a specific layer:
```bash
CGO_ENABLED=1 go test ./internal/dao -v
CGO_ENABLED=1 go test ./internal/service -v
CGO_ENABLED=1 go test ./internal/handler -v
```

## API Endpoints

### Create Employee
```bash
POST /employees
Content-Type: application/json

{
  "full_name": "John Doe",
  "job_title": "Software Engineer",
  "country": "USA",
  "salary": 75000
}
```

### Get All Employees
```bash
GET /employees
```

### Get Employee by ID
```bash
GET /employees/{id}
```

### Update Employee
```bash
PUT /employees/{id}
Content-Type: application/json

{
  "full_name": "John Doe",
  "job_title": "Senior Software Engineer",
  "country": "USA",
  "salary": 95000
}
```

### Delete Employee
```bash
DELETE /employees/{id}
```

## Example Usage

```bash
# Create an employee
curl -X POST http://localhost:8080/employees \
  -H "Content-Type: application/json" \
  -d '{"full_name": "Jane Smith", "job_title": "Product Manager", "country": "UK", "salary": 85000}'

# Get all employees
curl http://localhost:8080/employees

# Get a specific employee
curl http://localhost:8080/employees/1

# Update an employee
curl -X PUT http://localhost:8080/employees/1 \
  -H "Content-Type: application/json" \
  -d '{"full_name": "Jane Smith", "job_title": "Senior Product Manager", "country": "UK", "salary": 95000}'

# Delete an employee
curl -X DELETE http://localhost:8080/employees/1
```

## Validation Rules

The API enforces the following validation rules:

- **Full Name**: Required, cannot be empty
- **Job Title**: Required, cannot be empty
- **Country**: Required, cannot be empty
- **Salary**: Required, must be positive (> 0)

Validation errors return HTTP 400 Bad Request with a descriptive error message.

## Development Approach

This project was developed using **Test-Driven Development (TDD)**:

1. Write failing tests for each layer
2. Implement minimal code to make tests pass
3. Refactor if needed

The commit history reflects this incremental TDD approach at each layer:
- DAO layer tests and implementation
- Service layer tests and implementation (with validation)
- Handler layer tests and implementation (with HTTP endpoints)

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── dao/
│   │   ├── employee_dao.go      # Database operations
│   │   └── employee_dao_test.go # DAO tests
│   ├── handler/
│   │   ├── employee_handler.go      # HTTP handlers
│   │   └── employee_handler_test.go # Handler tests
│   ├── model/
│   │   └── employee.go          # Employee model
│   └── service/
│       ├── employee_service.go      # Business logic
│       └── employee_service_test.go # Service tests
├── go.mod
├── go.sum
└── README.md
```
