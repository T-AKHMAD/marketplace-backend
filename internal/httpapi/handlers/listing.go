package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"marketplace/internal/domain"
	"marketplace/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type ListingHandler struct {
	svc *service.ListingService
	rdb *redis.Client
	ch  *amqp091.Channel
}

func NewListingHandler(svc *service.ListingService, rdb *redis.Client, ch *amqp091.Channel) *ListingHandler {
	return &ListingHandler{
		svc: svc,
		rdb: rdb,
		ch:  ch,
	}
}

func (h *ListingHandler) Create(ctx *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Price       int64  `json:"price"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if req.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	if req.Price <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid price"})
		return
	}

	l, err := h.svc.Create(ctx, req.Title, req.Description, req.Price)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	h.rdb.Del(ctx, "listings:all")
	h.rdb.Del(ctx, "listing:"+strconv.FormatInt(l.ID, 10))

	body, err := json.Marshal(l)
	if err != nil {
		log.Printf("failed to marshal listing: %v", err)
	} else {
		err = h.ch.Publish(
			"",
			"listing_created",
			false,
			false,
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		if err != nil {
			log.Printf("Failed to publish listing_created: %v", err)
		}
	}

	ctx.JSON(http.StatusCreated, l)

}

func (h *ListingHandler) List(ctx *gin.Context) {

	data, err := h.rdb.Get(ctx, "listings:all").Result()
	if err == nil {
		ctx.Data(http.StatusOK, "application/json", []byte(data))
		return
	}
	if err != redis.Nil {
		log.Printf("Redis error: %v", err)
	}

	list, err := h.svc.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	jsonData, err := json.Marshal(list)
	if err != nil {
		log.Printf("failed to marshal listings: %v", err)
	} else {
		if err := h.rdb.Set(ctx, "listings:all", jsonData, time.Minute).Err(); err != nil {
			log.Printf("failed to set cache: %v", err)
		}
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *ListingHandler) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	key := "listing:" + idStr

	data, err := h.rdb.Get(ctx, key).Result()
	if err == nil {
		ctx.Data(http.StatusOK, "application/json", []byte(data))
		return
	}

	if err != redis.Nil {
		log.Printf("Redis error: %v", err)
	}

	listing, err := h.svc.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	jsonData, err := json.Marshal(listing)
	if err != nil {
		log.Printf("failed to marshal listing: %v", err)
	} else {
		if err := h.rdb.Set(ctx, key, jsonData, 30*time.Second).Err(); err != nil {
			log.Printf("failed to set cache: %v", err)
		}
	}
	ctx.JSON(http.StatusOK, listing)
}
