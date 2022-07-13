package main

import (
	"context"
	"go-test/kit"
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
	return kit.ReadConfig()
}

func (a *App) ConfigSet(config string) {
	kit.WriteConfig(config)
}
