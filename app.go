package main

import (
	"context"
	"go-test/lib"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// 读取配置文件
func (a *App) ConfigGet() string {
	return lib.ConfigRead()
}

// 写入配置文件
func (a *App) ConfigSet(config string) {
	lib.ConfigWrite(config)
}

// 发送HTTP请求
func (a *App) HttpFetch(request *lib.HttpRequest) *lib.HttpResponse {
	return lib.HttpFetch(request)
}

func (a *App) SysGetCpuPercent() float64 {
	return lib.SysGetCpuPercent()
}

func (a *App) SysGetMemPercent() float64 {
	return lib.SysGetMemPercent()
}

func (a *App) FileGetAllFileList(path string) []*lib.FileItem {
	return lib.FileGetAllFileList(path)
}

func (a *App) FileHash(path string) string {
	return lib.FileHash(path)
}

func (a *App) FileReadDir(path string) []*lib.FileItem {
	return lib.FileReadDir(path)
}

func (a *App) DialogDirSelect() string {
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return ""
	}
	return dir
}
