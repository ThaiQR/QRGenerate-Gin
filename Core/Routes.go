package core

import (
	controller "github.com/ThaiQR/QRGenerate-Gin/Controller"
	services "github.com/ThaiQR/QRGenerate-Gin/Services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	qrService := &services.QRService{}
	qrGenerateController := &controller.QRGenerateController{
		QRService: qrService,
	}

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", controller.Test)

	r.POST("/merchant-billpayment-qr", qrGenerateController.GetMerchantBillpaymentQR)
	r.POST("/merchant-promptpay-qr", qrGenerateController.GetMerchantPromptpayQR)

	r.Use(gin.Logger())

	return r
}
