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

func (a *App) ConfigGet() string {
	return lib.ConfigRead()
}

func (a *App) ConfigSet(config string) {
	lib.ConfigWrite(config)
}
