package config

import (
	"go-test/kit"
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	Config string
}

type response struct {
	Config string
}

func Use(router *gin.RouterGroup) {
	router.GET("/config", func(ctx *gin.Context) {
		res := response{}
		res.Config = kit.ReadConfig()
		ctx.JSON(http.StatusOK, res)
	})

	router.POST("/configSet", func(ctx *gin.Context) {
		var req request
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			kit.HandleError(err, ctx)
			return
		}
		kit.WriteConfig(req.Config)
		ctx.JSON(http.StatusOK, nil)
	})
}
