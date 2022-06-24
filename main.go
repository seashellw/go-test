package main

import (
	"embed"
	"go-test/util"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed client/dist
var clientFs embed.FS

func main() {
	r := gin.Default()
	distFs, _ := fs.Sub(clientFs, "client/dist")
	r.StaticFS("/", http.FS(distFs))
	util.BrowserOpen("http://localhost")
	r.Run(":80")
}
