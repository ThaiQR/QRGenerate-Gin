package core

import "github.com/gin-gonic/gin"

func Extensions() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	return r
}
