package models

type QRGenerateRequest struct {
	Reference1 string  `json:"reference1"`
	Reference2 string  `json:"reference2"`
	Amount     float64 `json:"amount"`
}
