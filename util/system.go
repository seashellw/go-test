package util

import (
	"log"
	"os"
	"os/exec"
)

// 运行系统命令
func SystemCall(dir string, params ...string) {
	logStr := ">"
	logStr += " " + dir
	for _, val := range params {
		logStr += " " + val
	}
	log.Println(logStr)
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
