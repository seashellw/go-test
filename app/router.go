package app

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func StartRouter(this *App) {
	r := gin.Default()
	r.Use(cors.Default())
	handleApp(this, r.Group("/api/app"))

	r.Run("localhost:9001")
}

func handleApp(this *App, r *gin.RouterGroup) {
	r.POST("/reload", func(c *gin.Context) {
		runtime.WindowReloadApp(this.Ctx)
		c.AbortWithStatus(http.StatusOK)
	})
}
