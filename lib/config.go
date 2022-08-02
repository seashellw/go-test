package lib

import (
	"bytes"
	"encoding/json"
	"path/filepath"
)

var configFileName = filepath.Join(FileGetCurrentPath(), "go-config.json")

// 读配置文件，文件应为json格式
func ConfigRead() string {
	return FileReadFile(configFileName)
}

// 写配置文件，字符串应为json格式
func ConfigWrite(config string) {
	var out bytes.Buffer
	json.Indent(&out, []byte(config), "", "  ")
	FileWriteFile(configFileName, out.String())
}
