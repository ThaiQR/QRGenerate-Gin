package controller

import (
	"net/http"

	request "github.com/ThaiQR/QRGenerate-Gin/Model/Request"
	services "github.com/ThaiQR/QRGenerate-Gin/Services"

	"github.com/gin-gonic/gin"
)

type QRGenerateController struct {
	QRService *services.QRService
}

// @summary TQR Billpayment generate
// @description For Generate TQR
// @id TQR Merchant Billpayment
// @accept json
// @produce json
// @Param request.MerchantBillpaymentRequest body request.MerchantBillpaymentRequest true "Body for Generate TQR"
// @Success 200 {string} string "OK"
// @router /merchant-billpayment-qr [post]
func (qr *QRGenerateController) GetMerchantBillpaymentQR(c *gin.Context) {
	var value request.MerchantBillpaymentRequest

	if err := c.ShouldBindJSON(&value); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := qr.QRService.QRMerchantBillpaymentGenerate(&value)

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"response": result,
	})
}

// @summary TQR Promptpay generate
// @description For Generate TQR
// @id TQR Promptpay Promptpay
// @accept json
// @produce json
// @Param request.MerchantPromptpayRequest body request.MerchantPromptpayRequest true "Body for Generate TQR"
// @Success 200 {string} string "OK"
// @router /merchant-promptpay-qr [post]
func (qr *QRGenerateController) GetMerchantPromptpayQR(c *gin.Context) {
	var value request.MerchantPromptpayRequest

	if err := c.ShouldBindJSON(&value); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := qr.QRService.QRMerchantPromptpayGenerate(&value)

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"response": result,
	})
}
