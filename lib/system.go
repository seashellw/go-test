package lib

import (
	"os"
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
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

func SysGetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func SysGetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}
