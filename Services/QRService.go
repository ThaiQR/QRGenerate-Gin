package services

import (
	helpers "github.com/ThaiQR/QRGenerate-Gin/Helpers"
	request "github.com/ThaiQR/QRGenerate-Gin/Model/Request"
	response "github.com/ThaiQR/QRGenerate-Gin/Model/Response"
	thaiqr "github.com/ThaiQR/ThaiQR-Go"

	"strconv"

	"github.com/spf13/viper"
)

type QRService struct{}

func (s *QRService) QRMerchantBillpaymentGenerate(value *request.MerchantBillpaymentRequest) *response.QRGenerateResponse {
	billerId := viper.GetString("merchant_settings.biller_id")
	merchantName := viper.GetString("merchant_settings.merchant_name")

	qrResult := thaiqr.MerchantBillpaymentQRGenerate(billerId, merchantName, value.Reference1, value.Reference2, strconv.FormatFloat(value.Amount, 'f', 2, 64), true)

	response := &response.QRGenerateResponse{
		QrText:   qrResult,
		QrBase64: helpers.QRConvertToBase64String(qrResult),
	}

	return response
}

func (s *QRService) QRMerchantPromptpayGenerate(value *request.MerchantPromptpayRequest) *response.QRGenerateResponse {
	qrResult := thaiqr.MerchantPromptpayQRGenerate(value.ReceiveID, strconv.FormatFloat(value.Amount, 'f', 2, 64), value.Onetime)

	response := &response.QRGenerateResponse{
		QrText:   qrResult,
		QrBase64: helpers.QRConvertToBase64String(qrResult),
	}

	return response
}
