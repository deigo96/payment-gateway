package domain

type PaymentUsecase interface {
	StoreTransaction(StoreTranscation) (interface{}, error)
	CancelTransaction(CancelParam) (*CancelTransaction, error)
}
