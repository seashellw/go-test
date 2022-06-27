package main

import (
	"embed"
	"go-test/api"
	"go-test/util"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed client/dist
var clientFs embed.FS

var httpFs http.FileSystem
var fileServer http.Handler

// 初始化静态资源文件系统
func initFile() {
	distFs, _ := fs.Sub(clientFs, "client/dist")
	httpFs = http.FS(distFs)
	fileServer = http.FileServer(httpFs)
}

// 发送index
func handleIndex(ctx *gin.Context) {
	req, _ := http.NewRequest("GET", "/", nil)
	fileServer.ServeHTTP(ctx.Writer, req)
	ctx.Abort()
}

// 静态资源中间件
func handleStatic(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	_, err := fs.Stat(clientFs, "client/dist"+path)
	if err == nil || path == "/" {
		gin.WrapH(fileServer)(ctx)
		ctx.Abort()
		return
	}
}

// 处理history模式，在无路由匹配时返回index
func handleHistory(ctx *gin.Context) {
	if ctx.Request.Method != "GET" || strings.HasPrefix(ctx.Request.URL.Path, "/api/") {
		return
	}
	handleIndex(ctx)
}

// 主程序
func main() {
	initFile()
	router := gin.Default()
	// 静态资源中间件
	router.Use(handleStatic)
	// 加载api路由
	api.UseApiRouter(router)
	// 无路由匹配时，加载history模式中间件
	router.NoRoute(handleHistory)
	// 在浏览器中启动UI
	util.BrowserOpen("http://localhost")
	// 启动后端
	router.Run("localhost:80")
}
