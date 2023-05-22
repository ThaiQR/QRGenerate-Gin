package response

type QRGenerateResponse struct {
	QrText   string `json:"qr_text"`
	QrBase64 string `json:"qr_base64"`
}
