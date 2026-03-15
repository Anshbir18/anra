# Task API

Minimal REST API for managing tasks written in Go.

## Tech Stack

- Go 1.21+
- Gin
- Swagger
- In-memory storage

## Run Service

Install dependencies:

go mod tidy

Generate swagger docs:

swag init -g cmd/server/main.go

Run server:

go run cmd/server/main.go

Server runs on:

http://localhost:8080

Swagger docs:

http://localhost:8080/swagger/index.html

## API Endpoints

POST /tasks
Create a task

GET /tasks
List tasks

## Run Tests

go test ./...

## Future Improvements

- Persistent storage (mysql)
- Pagination for task listing
- Structured logging
- Docker support