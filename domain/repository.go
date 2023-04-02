package domain

import (
	"encoding/json"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

type PreTranscations struct {
	UserId int64
	Items json.RawMessage 
	Status int32
	RequestJson json.RawMessage
	CreatedAt time.Time
	Quantity int32
	OrderId string
	ExpiryAt *time.Time
	PaymentVia *string
	PaymentCode *string
	TransactionId *string
	GrossAmount *int64
	PaymentType *string
	StatusTransaction *string
}

type PaymentRepository interface {
	GetListBank() ([]ListBank, error)
	GetBankByCode(string) (*ListBank, error)
	StoreTranscation(coreapi.ChargeReq, PreTranscations) error
	CancelTransaction(CancelParam) (*PreTranscations, error)
}
