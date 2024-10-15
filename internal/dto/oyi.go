package dto

type OYIStatus struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type OYITransactionReq struct {
	CustomerId         string  `json:"customer_id,omitempty"`
	SubMerchantId      string  `json:"sub_merchant_id,omitempty"`
	PartnerTrxId       string  `json:"partner_trx_id,omitempty"`
	Amount             float64 `json:"amount,omitempty"`
	EwalletCode        string  `json:"ewallet_code,omitempty"`
	MobileNumber       string  `json:"mobile_number,omitempty"`
	Email              string  `json:"email,omitempty"`
	SuccessRedirectURL string  `json:"success_redirect_url,omitempty"`
	ExpirationTime     int     `json:"expiration_time,omitempty"`
}

type OYITransactionRes struct {
	Status           OYIStatus `json:"status"`
	EwalletTrxStatus string    `json:"ewallet_trx_status"`
	Amount           float64   `json:"amount"`
	TrxID            string    `json:"trx_id"`
	CustomerID       string    `json:"customer_id"`
	PartnerTrxID     string    `json:"partner_trx_id"`
	EwalletCode      string    `json:"ewallet_code"`
	EwalletURL       string    `json:"ewallet_url"`
}
