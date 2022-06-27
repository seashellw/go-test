package util

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// 运行系统命令
func SystemCall(dir string, params ...string) {
	var build strings.Builder
	build.WriteString(">")
	build.WriteString(" ")
	build.WriteString(dir)
	for _, val := range params {
		build.WriteString(" ")
		build.WriteString(val)
	}
	log.Println(build.String())
	cmd := exec.Command("PowerShell", params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	cmd.Run()
}

// 在浏览器中打开链接
func BrowserOpen(url string) {
	SystemCall("./", "explorer", url)
}
