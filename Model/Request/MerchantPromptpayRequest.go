package request

type MerchantPromptpayRequest struct {
	RecieveID string  `json:"recieveId"`
	Amount    float64 `json:"amount"`
	Onetime   bool    `json:"onetime"`
}
