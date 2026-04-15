package httpapi

import (
	"github.com/T-AKHMAD/marketplace-backend/internal/httpapi/handlers"
	"github.com/T-AKHMAD/marketplace-backend/internal/httpapi/middleware"
	"github.com/T-AKHMAD/marketplace-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

func NewRouter(svc *service.ListingService, rdb *redis.Client, ch *amqp091.Channel) *gin.Engine {
	r := gin.New()
	h := handlers.NewListingHandler(svc, rdb, ch)

	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())

	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
	r.POST("/listings", h.Create)
	r.GET("/listings", h.List)
	r.GET("/listings/:id", h.GetByID)

	return r
}
