package dto

type PaymentRequest struct {
	PaymentType string `json:"paymentMethod" validate:"required"`
	PromoCode   string `json:"promoCode"`
}

type TransactionResponse struct {
	PaymentMethod          string                 `json:"paymentMethod"`
	LimitTransactionDate   string                 `json:"limitTransactionDate"`
	TotalPrice             float64                `json:"totalPrice"`
	VirtualAccountResponse VirtualAccountResponse `json:"virtualAccount,omitempty"`
	EWalletResponse        EWalletResponse        `json:"eWallet,omitempty"`
}

type VirtualAccountResponse struct {
	VirtualAccountNumber string `json:"virtualAccountNumber,omitempty"`
	BankName             string `json:"bankName,omitempty"`
}

type EWalletResponse struct {
	Actions []EWalletAction `json:"actions,omitempty"`
}

type EWalletAction struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
