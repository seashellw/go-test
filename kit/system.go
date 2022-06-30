package kit

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

func ExecSync(dir string, params ...string) {
	cmd := getCmd(dir, params...)
	cmd.Run()
}

// 在浏览器中打开链接
func BrowserOpen(url string) {
	ExecSync("./", "explorer", url)
}

var configFileName = "go-config.json"

// 读配置文件，文件应为json格式
func ReadConfig() string {
	str := ""
	file, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return str
	}
	str = string(file)
	return str
}

// 写配置文件，字符串应为json格式
func WriteConfig(config string) {
	var out bytes.Buffer
	json.Indent(&out, []byte(config), "", "  ")
	ioutil.WriteFile(configFileName, out.Bytes(), 0644)
}
