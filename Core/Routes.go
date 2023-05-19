package core

import (
	controller "BankGW-Gin/Controller"
	services "BankGW-Gin/Services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", controller.Test)

	qrService := &services.QRService{}
	qrGenerateController := &controller.QRGenerateController{
		QRService: qrService,
	}

	r.POST("/qrgen", qrGenerateController.GetQR)

	r.Use(gin.Logger())

	return r
}
