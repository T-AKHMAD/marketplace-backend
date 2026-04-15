package main

import (
	"context"
	"log"
	"github.com/T-AKHMAD/marketplace-backend/internal/cache"
	"github.com/T-AKHMAD/marketplace-backend/internal/db"
	"github.com/T-AKHMAD/marketplace-backend/internal/httpapi"
	"github.com/T-AKHMAD/marketplace-backend/internal/queue"
	"github.com/T-AKHMAD/marketplace-backend/internal/repository"
	"github.com/T-AKHMAD/marketplace-backend/internal/service"
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
