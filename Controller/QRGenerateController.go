package controller

import (
	models "BankGW-Gin/Models"
	services "BankGW-Gin/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QRGenerateController struct {
	QRService *services.QRService
}

// @summary TQR generate
// @description For Generate TQR
// @id TQRGen
// @accept json
// @produce json
// @Param models.QRGenerateRequest body models.QRGenerateRequest true "Body for Generate TQR"
// @Success 200 {string} string "OK"
// @router /qrgen [post]
func (qr *QRGenerateController) GetQR(c *gin.Context) {
	var value models.QRGenerateRequest

	if err := c.ShouldBindJSON(&value); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := qr.QRService.QRGenerate(&value)
	// if err != "" {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR"})
	// 	return
	// }

	// valueStr, err := json.Marshal(value)
	// resultStr, err := json.Marshal(result)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Fatal(string(valueStr))
	// }

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"response": result,
	})
}
