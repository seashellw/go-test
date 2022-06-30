package main

import (
	"embed"
	"flag"
	"go-test/api"
	"go-test/kit"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 启动参数中的url
var url = flag.String("u", "http://localhost", "启动URL")

//go:embed dist
var clientFs embed.FS

// 文件系统服务
var fileServer http.Handler

// 初始化
func init() {
	flag.Parse()
	distFs, _ := fs.Sub(clientFs, "dist")
	fileServer = http.FileServer(http.FS(distFs))
}

// 静态资源中间件
func handleStatic(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	_, err := fs.Stat(clientFs, "dist"+path)
	if err == nil || path == "/" {
		gin.WrapH(fileServer)(ctx)
		ctx.Abort()
	}
}

// 处理history模式，在无路由匹配时返回index
func handleHistory(ctx *gin.Context) {
	if ctx.Request.Method != "GET" || strings.HasPrefix(ctx.Request.URL.Path, "/api/") {
		return
	}
	req, _ := http.NewRequest("GET", "/", nil)
	fileServer.ServeHTTP(ctx.Writer, req)
	ctx.Abort()
}

// 在浏览器中打开url，若url不正确，则不打开
func handleOpenBrowser() {
	url := *url
	if !strings.HasPrefix(url, "http") {
		kit.LogBlue("The url used for the browser is incorrect. No browser open.")
		return
	}
	kit.BrowserOpen(url)
}

// 主程序
func main() {
	test()
	router := gin.Default()
	// 静态资源中间件
	router.Use(handleStatic)
	// 加载api路由
	api.Route(router)
	// 无路由匹配时，加载history模式中间件
	router.NoRoute(handleHistory)
	// 在浏览器中启动UI
	handleOpenBrowser()
	// 启动后端
	router.Run("localhost:80")
}

func test() {
}
