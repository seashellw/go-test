package util

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

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

func BrowserOpen(url string) {
	SystemCall("./", "explorer", url)
}
