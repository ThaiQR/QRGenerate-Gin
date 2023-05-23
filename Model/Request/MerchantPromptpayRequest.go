package request

type MerchantPromptpayRequest struct {
	ReceiveID string  `json:"receiveId"`
	Amount    float64 `json:"amount"`
	Onetime   bool    `json:"onetime"`
}
