package dto

type FinpayEwalletReq struct {
	Customer      *Customer      `json:"customer"`
	Order         *Order         `json:"order"`
	URL           *URL           `json:"url"`
	SourceOfFunds *SourceOfFunds `json:"sourceOfFunds"`
}

type FinpayEwalletRes struct {
	ResponseCode    string  `json:"responseCode"`
	ResponseMessage string  `json:"responseMessage"`
	AppUrl          string  `json:"appurl"`
	RedirectURL     string  `json:"redirecturl"`
	ExpiryLink      string  `json:"expiryLink"`
	PaymentCode     any     `json:"paymentCode"`
	ProcessingTime  float64 `json:"processingTime"`
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
	Payment Payment `json:"payment"`
}

type Payment struct {
	Status     string `json:"status"`
	StatusDesc string `json:"statusDesc"`
	UserDesc   string `json:"userDesc"`
	Datetime   string `json:"datetime"`
	Reference  string `json:"reference"`
	Channel    any    `json:"channel"`
	Amount     int    `json:"amount"`
}
