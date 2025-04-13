# JSON Storage API

A RESTful service for storing and retrieving JSON documents, developed in Go with the Echo framework and Docker.

## Table of Contents

- [Technologies](#-technologies)
- [Features](#-features)
- [Prerequisites](#-prerequisites)
- [Installation](#-installation)
- [API Endpoints](#-api-endpoints)
- [Usage Examples](#-usage-examples)
- [Development](#-development)

## Technologies

- **Go** (1.24.2+) - Main programming language
- **Echo** - Fast and minimalist web framework
- **Docker** - Containerization
- **Makefile** - Task automation

## Features

- Storage of JSON documents
- Retrieval of documents by ID
- Automatic validation of JSON format
- RESTful API with intuitive endpoints
- Service health monitoring
- Containerization with Docker

## Prerequisites

- Go 1.24.2+
- Docker (optional)
- Make (optional, for automation)

## Installation

### With Docker (recommended)

1. Clone the repository:
```bash
git clone https://github.com/marcusbiava/json-post-box.git
cd json-post-box
```

2. Start the container:
```bash
make docker-up
```

The service will be available at `http://localhost:8030`

### Without Docker

1. Clone the repository:
```bash
git clone https://github.com/marcusbiava/json-post-box.git
cd json-post-box
```

2. Compile and run:
```bash
make build && ./bin/github.com/marcusbiava/json-post-box
```

## API Endpoints

### Store JSON Data (`POST /api/v1/jsons`)

Stores a JSON document and returns a unique ID.

```bash
curl -X POST http://localhost:8030/api/v1/jsons \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "age": 30,
    "email": "john@example.com",
    "preferences": {
      "theme": "dark",
      "notifications": true
    }
  }'
```

Response (201 Created):
```json
{
  "id": "generated-id",
  "data": {
    "name": "John Doe",
    "age": 30,
    "email": "john@example.com",
    "preferences": {
      "theme": "dark",
      "notifications": true
    }
  }
}
```

### Retrieve JSON Data (`GET /api/v1/jsons/:id`)

Retrieves a JSON document by its ID.

```bash
curl http://localhost:8030/api/v1/jsons/generated-id
```

Response (200 OK):
```json
{
  "name": "John Doe",
  "age": 30,
  "email": "john@example.com",
  "preferences": {
    "theme": "dark",
    "notifications": true
  }
}
```

Response (404 Not Found):
```json
{
  "message": "JSON not found"
}
```

### Health Check (`GET /api/v1/health`)

Checks the health status of the service.

```bash
curl http://localhost:8030/api/v1/health
```

Response (200 OK):
```json
{
  "status": "healthy"
}
```

## Development

### Project Structure

```
.
├── cmd/
│   └── main.go          # Entry point of the application
├── internal/
│   ├── adapter/         # Adapters (handlers, repositories)
│   ├── domain/          # Domain entities and interfaces
│   ├── infrastructure/  # Configuration and initialization
│   └── usecase/        # Business logic
├── Dockerfile           # Docker configuration
├── Makefile             # Automation commands
└── README.md            # Documentation
```

### Makefile Commands

```bash
make build           # Compiles the project
make test            # Runs tests
make docker-build     # Builds the Docker image
make docker-up        # Starts the container
make docker-down      # Stops the container
```

## Notes

- The service uses in-memory storage by default
- In production, it is recommended to implement persistent storage
- IDs are automatically generated for each document
- All endpoints return responses in JSON format