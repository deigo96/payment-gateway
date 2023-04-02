package route

import (
	"account-service/api/controller"
	"account-service/bootstrap"
	"account-service/repository"
	"account-service/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewPaymentRoute(env *bootstrap.Env, timeout time.Duration, db gorm.DB, router *gin.RouterGroup) {
	pr := repository.NewPaymentRepository(&db, *env)
	pc := &controller.PaymentController{
		Payment: usecase.NewPaymentUsecase(pr, timeout, *env),
	}

	router.POST("/store-transaction", pc.StoreTransactions)
	router.PUT("/cancel-transaction", pc.CancelTransaction)

}