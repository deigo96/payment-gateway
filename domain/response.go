package domain

type ListBank struct {
	CodeBank string
	Bank string
}

type CancelTransaction struct {
	OrderId string `json:"order_id"`
	PaymentVia string `json:"payment_via"`
	StatusTransaction string `json:"status_transaction"`
}
