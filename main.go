package main

import (
	"context"
	"embed"
	"fmt"
	"go-test/kit"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed frontend/dist
var assets embed.FS

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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// 主程序
func main() {
	test()

	// Create an instance of the app structure
	app := NewApp()

	if runtime.Environment(app.ctx).BuildType != "production" {
		runtime.BrowserOpenURL(app.ctx, "http://localhost:34115")
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "go",
		Width:     1024,
		Height:    768,
		Assets:    assets,
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		kit.LogRed(err.Error())
	}
}

func test() {
}
