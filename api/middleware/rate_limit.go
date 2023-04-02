package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimit() gin.HandlerFunc {
	// Create a new limiter that allows 10 requests per minute
	limiter := rate.NewLimiter(3, 6)

	return func(ctx *gin.Context) {
		if !limiter.Allow() {

			ctx.AbortWithError(http.StatusTooManyRequests, errors.New("too many requests"))
		}

		// If the limiter doesn't allow the request, return an error
	}
}
