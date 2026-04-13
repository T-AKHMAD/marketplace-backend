package main

import (
	"context"
	"log"
	"marketplace/internal/cache"
	"marketplace/internal/db"
	"marketplace/internal/httpapi"
	"marketplace/internal/queue"
	"marketplace/internal/repository"
	"marketplace/internal/service"
)

func main() {
	ctx := context.Background()
	db := db.NewPostgres()

	repo := repository.NewListingPostgresRepository(db)
	svc := service.NewListingService(repo)
	rdb := cache.NewRedisClient()
	_, ch := queue.NewRabbitMQ()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	}

	r := httpapi.NewRouter(svc, rdb, ch)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
