package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed client/dist
var clientFs embed.FS

func main() {
	r := gin.Default()
	r.Use(static)
	// util.BrowserOpen("http://localhost")
	r.Run(":80")
}

func static(c *gin.Context) {
	distFs, _ := fs.Sub(clientFs, "client/dist")
	fileServer := http.FileServer(http.FS(distFs))
	// fileServer := http.FileServer(http.FS(distFs))
	path := c.Request.URL.Path
	path = strings.TrimPrefix(path, "/")
	log.Println("请求路径", path)
	if path == "" {
		log.Println("path为空", path)
		c.FileFromFS("index.html", http.FS(distFs))
	}
	_, err := distFs.Open(path)
	if err != nil {
		log.Println("错误", err)
	} else {
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
