package lib

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

func FileHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", errors.New("文件打开错误")
	}
	defer file.Close()
	md5h := md5.New()
	io.Copy(md5h, file)
	return hex.EncodeToString(md5h.Sum(nil)), nil
}
