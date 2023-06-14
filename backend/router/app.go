package router

import (
	"github.com/gin-gonic/gin"
	"hakaryzhang.com/amis-test/config"
	"hakaryzhang.com/amis-test/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(config.Cors())

	r.GET("/test/01", service.Test01)

	return r
}
