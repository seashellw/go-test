package kit

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(err error, ctx *gin.Context) {
	e := err.Error()
	if e == "EOF" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求体格式错误"})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"error": e})
}
