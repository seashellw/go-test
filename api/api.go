package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用api路由
func UseApiRouter(engine *gin.Engine) {
	api := engine.Group("/api")

	api.POST("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})
}
