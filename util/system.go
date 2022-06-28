package util

import (
	"log"
	"os"
	"os/exec"
)

type System struct{}

func (sys System) getCmd(dir string, params ...string) *exec.Cmd {
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
	return cmd
}

func (sys System) Sync(dir string, params ...string) System {
	cmd := sys.getCmd(dir, params...)
	cmd.Run()
	return sys
}

func (sys System) Async(dir string, params ...string) System {
	cmd := sys.getCmd(dir, params...)
	cmd.Start()
	return sys
}

// 在浏览器中打开链接
func BrowserOpen(url string) {
	sys := System{}
	sys.Async("./", "explorer", url)
}
