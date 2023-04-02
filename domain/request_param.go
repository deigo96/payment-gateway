package domain

type StoreTranscation struct {
	OrderId   string   `json:"_,omitempty"`
	CodeBank  string   `json:"code_bank"`
	Customers Customer `json:"customers"`
	Items     []Items  `json:"items"`
}

type CancelParam struct {
	OrderId       string `json:"order_id"`
	TransactionId string `json:"transaction_id"`
}

type Customer struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Items struct {
	Id       string `json:"item_id"`
	ItemName string `json:"item_name"`
	Price    int64  `json:"price"`
	Quantity int32  `json:"qty"`
}