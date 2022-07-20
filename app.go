package main

import (
	"context"
	"go-test/lib"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
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
