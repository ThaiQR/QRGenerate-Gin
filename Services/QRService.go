package services

import (
	helpers "github.com/ThaiQR/QRGenerate-Gin/Helpers"
	request "github.com/ThaiQR/QRGenerate-Gin/Model/Request"
	response "github.com/ThaiQR/QRGenerate-Gin/Model/Response"

	model "github.com/ThaiQR/QRGenerate-Gin/Model"

	"strconv"

	"github.com/spf13/viper"
)

type QRService struct{}

func (s *QRService) QRMerchantBillpaymentGenerate(value *request.MerchantBillpaymentRequest) *response.QRGenerateResponse {
	qrData := model.MerchantBillpaymentModel{}

	billerId := viper.GetString("merchant_settings.biller_id")
	merchantName := viper.GetString("merchant_settings.merchant_name")

	qrData.SetValues(billerId, value.Reference1, value.Reference2, strconv.FormatFloat(value.Amount, 'f', 2, 64), merchantName)

	response := &response.QRGenerateResponse{
		QrText:   helpers.QRMerchantBillpaymentGenerate(qrData),
		QrBase64: helpers.QRConvertToBase64String(helpers.QRMerchantBillpaymentGenerate(qrData)),
	}

	return response
}

func (s *QRService) QRMerchantPromptpayGenerate(value *request.MerchantPromptpayRequest) *response.QRGenerateResponse {
	qrData := model.MerchantPromptpayModel{}

	qrData.SetValues(value.RecieveID, strconv.FormatFloat(value.Amount, 'f', 2, 64), value.Onetime)

	response := &response.QRGenerateResponse{
		QrText:   helpers.QRMerchantPromptpayGenerate(qrData),
		QrBase64: helpers.QRConvertToBase64String(helpers.QRMerchantPromptpayGenerate(qrData)),
	}

	return response
}
