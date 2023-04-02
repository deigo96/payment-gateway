package route

import (
	"account-service/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db gorm.DB, router *gin.RouterGroup) {
	protectedRoute := router.Group("")
	// publicRoute := router.Group("")
	// protectedRoute.Use(middleware.ValidateJwt(env.ServerUrl))
	// publicRoute.Use(middleware.RateLimit())

	NewPaymentRoute(env, timeout, db, protectedRoute)
}
