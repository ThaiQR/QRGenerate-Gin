package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler godoc
// @summary Health Check
// @description Health checking for the service
// @id HealthCheckHandler
// @accept json
// @produce json
// @Success 200 {string} string "OK"
// @router / [get]
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
