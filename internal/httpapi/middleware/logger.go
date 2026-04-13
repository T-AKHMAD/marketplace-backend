package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		latency := time.Since(start)

		log.Printf("%s %s %d %s", ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status(), latency)
	}
}