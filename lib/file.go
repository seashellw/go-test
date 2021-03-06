package lib

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"sync"
)

type FileItem struct {
	// 文件全名
	Path string
	// 文件名或目录名
	Name string
	// 是否是目录
	IsDir bool
	// 文件大小，以字节记
	Size int64
}

func FileHash(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()
	md5h := md5.New()
	io.Copy(md5h, file)
	return hex.EncodeToString(md5h.Sum(nil))
}

func FileGetAllFileList(path string) []*FileItem {
	list := []*FileItem{}
	filepath.Walk(path, func(path string, info fs.FileInfo, _ error) error {
		if !info.IsDir() {
			list = append(list, &FileItem{
				Name:  info.Name(),
				Path:  path,
				IsDir: false,
				Size:  info.Size(),
			})
		}
		return nil
	})
	return list
}

func FileReadDir(path string) []*FileItem {
	if strings.HasSuffix(path, `:`) {
		path = path + `\`
	}
	list := []*FileItem{}
	dir, err := os.Stat(path)
	if err != nil || !dir.IsDir() {
		return list
	}
	infoList, err := ioutil.ReadDir(path)
	if err != nil {
		return list
	}
	for _, value := range infoList {
		item := &FileItem{
			Name:  value.Name(),
			Path:  filepath.Join(path, value.Name()),
			IsDir: value.IsDir(),
		}
		if !value.IsDir() {
			item.Size = value.Size()
		}
		list = append(list, item)
	}
	return list
}

func FileGetCurrentPath() string {
	path, _ := os.Getwd()
	return path
}

func FileGetDesktopPath() string {
	u, _ := user.Current()
	return filepath.Join(u.HomeDir, "Desktop")
}

var fileLock = map[string]*sync.RWMutex{}

func checkLock(path string) *sync.RWMutex {
	val, ok := fileLock[path]
	if !ok {
		val = &sync.RWMutex{}
		fileLock[path] = val
	}
	return val
}

func FileWriteFile(path string, value string) {
	lock := checkLock(path)
	lock.Lock()
	defer lock.Unlock()
	ioutil.WriteFile(path, []byte(value), 0644)
}

func FileReadFile(path string) string {
	lock := checkLock(path)
	lock.RLock()
	defer lock.RUnlock()
	file, _ := ioutil.ReadFile(configFileName)
	return string(file)
}
