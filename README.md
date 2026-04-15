# Marketplace Backend (Go)

Production-style backend service for a classifieds platform.

## Tech Stack
	•	Go (Gin, gRPC)
	•	PostgreSQL (pgxpool)
	•	Redis (caching)
	•	RabbitMQ (event-driven architecture + DLQ)



## 🧠 Architecture

## The project follows a clean layered architecture:

```text
HTTP (Gin) / gRPC
        ↓
     Service
        ↓
   Repository
        ↓
   PostgreSQL
```

## Key Design Decisions
	•	Single source of truth → PostgreSQL
	•	Read-through cache → Redis with TTL
	•	Event-driven communication → RabbitMQ
	•	Dead Letter Queue (DLQ) for failed messages
	•	Context propagation across all layers
	•	Error abstraction via domain-level errors



## 📦 Features
	•	Create listing (REST + gRPC)
	•	Get all listings (with Redis cache)
	•	Get listing by ID (with Redis cache)
	•	Cache invalidation on write
	•	Asynchronous event publishing (RabbitMQ)
	•	Worker for processing events
	•	Dead Letter Queue for failed processing


## 🏗️ Project Structure

```text
cmd/api        - REST API server
cmd/grpc       - gRPC server
cmd/worker     - background worker

internal/
  httpapi      - handlers & middleware
  service      - business logic
  repository   - data layer
  db           - postgres connection
  cache        - redis client
  queue        - rabbitmq
  grpc         - gRPC server
```

## ⚡ How to Run

### 1. Start dependencies

```bash
docker compose up -d
```


### 2. Run API server

```bash
go run ./cmd/api
```

### 3. Run gRPC server

```bash
go run ./cmd/grpc
```

### 4. Run worker

```bash
go run ./cmd/worker
```

## 🔧 API Examples

### Create listing

```bash
curl -X POST localhost:8080/listings \
  -H "Content-Type: application/json" \
  -d '{"title":"test","description":"desc","price":100}'
```

### Get listings

```bash
curl localhost:8080/listings
```

### Get by ID

```bash
curl localhost:8080/listings/1
```


## 🔌 gRPC Example

### List available services:

```bash
grpcurl -plaintext localhost:50051 list
```
### Call method:

```bash
grpcurl -plaintext \
  -d '{"id":1}' \
  localhost:50051 listing.ListingService/GetListing
```

## 📊 What I Focused On

This project demonstrates:
	•	Building a production-style backend
	•	Designing scalable architecture
	•	Implementing caching strategies
	•	Working with message queues
	•	Handling failures (DLQ)
	•	Using context propagation
	•	Supporting both REST and gRPC interfaces

💡 Future Improvements
	•	Outbox pattern for reliable event delivery
	•	Authentication (JWT)
	•	Pagination & filtering
	•	Integration tests
	•	gRPC streaming



👨‍💻 Author

Akhmad
Backend Developer (Go)
