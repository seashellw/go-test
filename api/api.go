package api

import (
	"go-test/api/fetch"

	"github.com/gin-gonic/gin"
)

// 使用api路由
func Route(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	fetch.Use(apiGroup)
}
