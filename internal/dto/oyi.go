package dto

type OYIStatus struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type OYICreateEwalletReq struct {
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

type OYIEwalletRes struct {
	Status           OYIStatus `json:"status"`
	EwalletTrxStatus string    `json:"ewallet_trx_status"`
	Amount           float64   `json:"amount"`
	TrxId            string    `json:"trx_id"`
	CustomerId       string    `json:"customer_id"`
	PartnerTrxId     string    `json:"partner_trx_id"`
	EwalletCode      string    `json:"ewallet_code"`
	EwalletURL       string    `json:"ewallet_url"`
}

type OYIEwalletNotify struct {
	Success            bool    `json:"success"`
	PartnerTrxId       string  `json:"partner_trx_id,omitempty"`
	TrxId              string  `json:"trx_id"`
	RefNumber          string  `json:"ref_number,omitempty"`
	CustomerId         string  `json:"customer_id"`
	Amount             float64 `json:"amount"`
	EwalletCode        string  `json:"ewallet_code"`
	MobileNumber       string  `json:"mobile_number"`
	SuccessRedirectURL string  `json:"success_redirect_url"`
	SettlementTime     string  `json:"settlement_time"`
	SettlementStatus   string  `json:"settlement_status"`
}

type OYIStatusEwalletReq struct {
	PartnerTrxId string `json:"partner_trx_id"`
}
