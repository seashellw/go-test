package main

import (
	"embed"
	"flag"
	"go-test/api"
	"go-test/util"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 初始化启动参数
func initFlag() string {
	url := flag.String("u", "http://localhost", "启动URL")
	flag.Parse()
	return *url
}

//go:embed dist
var clientFs embed.FS

var httpFs http.FileSystem
var fileServer http.Handler

// 初始化静态资源文件系统
func initFile() {
	distFs, _ := fs.Sub(clientFs, "dist")
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
	_, err := fs.Stat(clientFs, "dist"+path)
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
	url := initFlag()
	initFile()
	router := gin.Default()
	// 静态资源中间件
	router.Use(handleStatic)
	// 加载api路由
	api.UseApiRouter(router)
	// 无路由匹配时，加载history模式中间件
	router.NoRoute(handleHistory)
	// 在浏览器中启动UI
	util.BrowserOpen(url)
	// 启动后端
	router.Run("localhost:80")
}
