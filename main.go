package main

import (
	"context"
	"embed"
	"go-test/lib"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

// 主程序
func main() {
	test()

	// Create an instance of the app structure
	app := &App{}

	// Create application with options
	err := wails.Run(&options.App{
		Width:            1024,
		Height:           768,
		Assets:           assets,
		OnStartup:        func(ctx context.Context) { app.ctx = ctx },
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		},
	})

	if err != nil {
		lib.LogRed(err.Error())
	}
}

func test() {
}
