package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"sync"
)

var configFileName = "go-config.json"

var configLock = sync.RWMutex{}

// 读配置文件，文件应为json格式
func ConfigRead() string {
	configLock.RLock()
	defer configLock.RUnlock()
	str := ""
	file, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return str
	}
	str = string(file)
	return str
}

// 写配置文件，字符串应为json格式
func ConfigWrite(config string) {
	configLock.Lock()
	defer configLock.Unlock()
	var out bytes.Buffer
	json.Indent(&out, []byte(config), "", "  ")
	ioutil.WriteFile(configFileName, out.Bytes(), 0644)
}
