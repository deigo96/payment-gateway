package controller

import (
	"account-service/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	Payment domain.PaymentUsecase
}

func (payment *PaymentController) StoreTransactions(c *gin.Context) {
	var req domain.StoreTranscation
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest,domain.BuildResponse(http.StatusBadRequest, "Bad request, invalid request body", domain.EmptyObj{}))
		return
	}
	
	_, err := payment.Payment.StoreTransaction(req)
	if err != nil {
		code, err := domain.ErrorResponse(err.Error())
		c.JSON(code, domain.BuildResponse(code, err, domain.EmptyObj{}))
		return
	}

	c.JSON(http.StatusCreated, domain.BuildResponse(http.StatusCreated, "Success store transaction", domain.EmptyObj{}))

}


func (payment *PaymentController) CancelTransaction(c *gin.Context) {
	var req domain.CancelParam
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest,domain.BuildResponse(http.StatusBadRequest, "Bad request, invalid request body", domain.EmptyObj{}))
		return
	}

	cancel, err := payment.Payment.CancelTransaction(req)
	if err != nil {
		code, err := domain.ErrorResponse(err.Error())
		c.JSON(code, domain.BuildResponse(code, err, domain.EmptyObj{}))
		return
	}

	c.JSON(http.StatusOK, domain.BuildResponse(http.StatusOK, "Transaction successfully canceled", cancel))
}

