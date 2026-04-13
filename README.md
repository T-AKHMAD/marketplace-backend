Marketplace Backend (Go)

Production-style backend сервис для размещения объявлений.

🚀 Стек
	•	Go (Gin)
	•	PostgreSQL (pgxpool)
	•	Redis (cache)
	•	RabbitMQ (events + worker + DLQ)

🧠 Архитектура
	•	Clean Architecture:
	•	Handler → Service → Repository
	•	Context propagation (request-scoped)
	•	Error abstraction (domain errors)
	•	Read-through cache (Redis)
	•	Event-driven (RabbitMQ)

📦 Возможности
	•	Создание объявления (POST /listings)
	•	Получение списка (GET /listings)
	•	Получение по ID (GET /listings/:id)
	•	Кэширование через Redis
	•	Асинхронные события через RabbitMQ
	•	Dead Letter Queue для обработки ошибок

⚡ Как запустить
```bash
docker compose up -d
go run ./cmd/api
go run ./cmd/worker
```
🔧 Примеры

Создание
```bash
curl -X POST localhost:8080/listings \
  -H "Content-Type: application/json" \
  -d '{"title":"test","description":"desc","price":100}'
```

Получение

```bash
curl localhost:8080/listings
```

📊 Что реализовано
	•	PostgreSQL как источник истины
	•	Redis cache с TTL и invalidation
	•	RabbitMQ producer/consumer
	•	DLQ для невалидных сообщений
	•	Context propagation
	•	Обработка ошибок (domain layer)

