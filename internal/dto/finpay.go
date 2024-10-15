package dto

type FinpayEwalletReq struct {
	Customer      *Customer      `json:"customer,omitempty"`
	Order         *Order         `json:"order,omitempty"`
	URL           *URL           `json:"url,omitempty"`
	SourceOfFunds *SourceOfFunds `json:"sourceOfFunds,omitempty"`
}

type FinpayEwalletRes struct {
	ResponseCode    string  `json:"responseCode,omitempty"`
	ResponseMessage string  `json:"responseMessage,omitempty"`
	AppUrl          string  `json:"appurl,omitempty"`
	RedirectURL     string  `json:"redirecturl,omitempty"`
	ExpiryLink      string  `json:"expiryLink,omitempty"`
	PaymentCode     any     `json:"paymentCode,omitempty"`
	ProcessingTime  float64 `json:"processingTime,omitempty"`
}

type FinpayEwalletNotify struct {
	Customer  *Customer `json:"customer,omitempty"`
	Order     *Order    `json:"order,omitempty"`
	Card      *Card     `json:"card,omitempty"`
	Meta      *Meta     `json:"meta,omitempty"`
	Result    *Result   `json:"result,omitempty"`
	Signature string    `json:"signature,omitempty"`
}

type Customer struct {
	Email       string `json:"email,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	MobilePhone string `json:"mobilePhone,omitempty"`
}

type Order struct {
	ID          string  `json:"id,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	Description string  `json:"description,omitempty"`
	Timeout     int     `json:"timeout,omitempty"`
	Items       []Item  `json:"item,omitempty"`
}

type Item struct {
	Name        string `json:"name,omitempty"`
	Quantity    string `json:"quantity,omitempty"`
	Category    string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	UnitPrice   string `json:"unitPrice,omitempty"`
}

type URL struct {
	CallbackURL string `json:"callbackUrl,omitempty"`
	SuccessURL  string `json:"successUrl,omitempty"`
	FailURL     string `json:"failUrl,omitempty"`
}

type SourceOfFunds struct {
	Type      string `json:"type,omitempty"`
	AccountId string `json:"accountId,omitempty"`
}

type Result struct {
	Payment Payment `json:"payment,omitempty"`
}

type Payment struct {
	Status     string `json:"status,omitempty"`
	StatusDesc string `json:"statusDesc,omitempty"`
	UserDesc   string `json:"userDesc,omitempty"`
	Datetime   string `json:"datetime,omitempty"`
	Reference  string `json:"reference,omitempty"`
	Channel    any    `json:"channel,omitempty"`
	Amount     int    `json:"amount,omitempty"`
}

type Meta struct {
	Data any `json:"data,omitempty"`
}

type CardInfo struct {
	Brand   string `json:"brand,omitempty"`
	Issuing string `json:"issuing,omitempty"`
	Type    string `json:"type,omitempty"`
	SubType string `json:"subType,omitempty"`
	Country string `json:"country,omitempty"`
}

type Card struct {
	Mask string   `json:"mask,omitempty"`
	Info CardInfo `json:"info,omitempty"`
}
