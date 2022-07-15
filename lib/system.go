package lib

import (
	"os"
	"os/exec"
)

func getCmd(dir string, params ...string) *exec.Cmd {
	LogBlue(append([]string{">", dir}, params...)...)
	cmd := exec.Command("PowerShell", params...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func SysExecSync(dir string, params ...string) {
	cmd := getCmd(dir, params...)
	cmd.Run()
}

// 在浏览器中打开链接
func SysBrowserOpen(url string) {
	SysExecSync("./", "explorer", url)
}
