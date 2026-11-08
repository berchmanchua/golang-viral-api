# Viral API - Golang Skeleton

This is just a skeleton Go REST API configured with common dependencies. It leverages Uber Fx for robust dependency injection and integrates with Postgres, Redis, DynamoDB, and Kinesis.

## Features

* **Dependency Injection:** Utilizes `uber-go/fx` for automatic lifecycle management and graceful shutdowns of database connections and HTTP servers.
* **Local AWS Emulation:** Fully functional offline development using `amazon/dynamodb-local` and `kinesis-mock` via Docker.
* **DynamoDB GSIs:** Configured with Global Secondary Indexes for advanced query capabilities (e.g., querying top views across all platforms).
* **Automated DB Initialization:** Uses an ephemeral Docker container to automatically create DynamoDB tables and schemas on startup.

## Tech Stack

* **Language:** Go 1.27.3
* **Router:** [go-chi/chi](https://github.com/go-chi/chi)
* **Dependency Injection:** [uber-go/fx](https://github.com/uber-go/fx)
* **Configuration:** [spf13/viper](https://github.com/spf13/viper)
* **Postgres:** [jackc/pgx](https://github.com/jackc/pgx)
* **Redis:** [redis/go-redis](https://github.com/redis/go-redis)
* **AWS SDK:** Official AWS SDK for Go v2 (DynamoDB, Kinesis)

## Prerequisites

* [Go 1.27.3](https://go.dev/dl/)
* [Docker](https://docs.docker.com/get-docker/) and Docker Compose

## Getting Started

### 1. Spin up the infrastructure

Start the API and all dependencies (Postgres, Redis, DynamoDB, Kinesis) using Docker Compose:
```bash
docker-compose up --build
```
*Note: On startup, the `dynamodb-init` container will automatically run the AWS CLI to create the `ViralItems` table along with its `TopViewsIndex` Global Secondary Index, and then gracefully exit.*

### 2. Run without Docker (Optional)

If you prefer to run the Go application directly on your host machine while keeping the databases in Docker:

1. Start only the data stores:
   ```bash
   docker-compose up -d postgres redis dynamodb-local dynamodb-init kinesis-local
   ```
2. Download Go modules:
   ```bash
   go mod tidy
   ```
3. Run the application:
   ```bash
   go run ./cmd/api
   ```

## Project Structure

```text
.
├── cmd/
│   └── api/
│       └── main.go              # Entrypoint; initializes Uber Fx
├── internal/
│   ├── aws/                     # AWS client initializers (DynamoDB, Kinesis)
│   ├── client/                  # HTTP clients with exponential retry
│   ├── config/                  # Viper configuration setup
│   ├── db/                      # Postgres and Redis connection logic
│   ├── domain/                  # Core data models (e.g., ViralItem)
│   ├── server/                  # HTTP handlers, Chi router, and Fx server hooks
│   └── service/                 # Business logic and database operations
├── config.yaml                  # Environment configuration
├── docker-compose.yaml          # Local infrastructure orchestration
└── Dockerfile                   # Multi-stage build for the API
```
