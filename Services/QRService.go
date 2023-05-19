package services

import (
	data "BankGW-Gin/Data"
	helpers "BankGW-Gin/Helpers"

	models "BankGW-Gin/Models"

	"strconv"

	"github.com/spf13/viper"
)

type QRService struct{}

func (s *QRService) QRGenerate(value *models.QRGenerateRequest) *models.QRGenerateResponse {
	qrData := data.QRGenerateData{}

	billerId := viper.GetString("merchant_settings.biller_id")
	merchantName := viper.GetString("merchant_settings.merchant_name")

	qrData.SetValues(billerId, value.Reference1, value.Reference2, strconv.FormatFloat(value.Amount, 'f', 2, 64), merchantName)

	response := &models.QRGenerateResponse{
		QrText:   helpers.QRMerchantGenerate(qrData),
		QrBase64: helpers.QRConvertToBase64(helpers.QRMerchantGenerate(qrData)),
	}

	return response
}
