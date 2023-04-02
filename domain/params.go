package domain

type RegisterParams struct {
	AccountId  int  `json:"account_id"`
	UserId     int  `json:"user_id" validate:"required"`
	BranchId   int  `json:"branch_id"`
	IsApprover bool `json:"is_approver"`
	IsAdmin    bool `json:"is_admin"`
	IsActive   bool `json:"is_active"`
}

type EditParams struct {
	AccountId  int `json:"account_id"`
	UserId     int `json:"user_id"`
	BranchId   int `json:"branch_id"`
	IsApprover int `json:"is_approver"`
	IsAdmin    int `json:"is_admin"`
	IsActive   int `json:"is_active"`
}
