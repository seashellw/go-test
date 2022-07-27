package lib

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
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

func FileGetAllFileList(path string) *[]*FileItem {
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
	return &list
}

func FileReadDir(path string) *[]*FileItem {
	list := []*FileItem{}
	dir, err := os.Stat(path)
	if err != nil || !dir.IsDir() {
		return &list
	}
	infoList, err := ioutil.ReadDir(path)
	if err != nil {
		return &list
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
	return &list
}

func FindDuplicateFile(path string) *[]*[]*FileItem {
	dir, err := os.Stat(path)
	if err != nil || !dir.IsDir() {
		return &[]*[]*FileItem{}
	}
	hashMap := map[string]*[]*FileItem{}
	wg := sync.WaitGroup{}
	for _, value := range *FileGetAllFileList(path) {
		wg.Add(1)
		item := *value
		go func() {
			defer wg.Done()
			hash := FileHash(item.Path)
			list, ok := hashMap[hash]
			if !ok {
				hashMap[hash] = &[]*FileItem{}
			}
			newList := append(*list, &item)
			hashMap[hash] = &newList
		}()
	}
	wg.Wait()
	list := []*[]*FileItem{}
	for _, value := range hashMap {
		if len(*value) > 1 {
			list = append(list, value)
		}
	}
	return &list
}
