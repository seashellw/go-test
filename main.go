package main

import (
	"embed"
	"go-test/kit"

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
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:      "Go",
		Width:      1024,
		Height:     768,
		Assets:     assets,
		OnStartup:  app.startup,
		OnDomReady: app.domready,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			DisableWindowIcon: true,
		},
	})

	if err != nil {
		kit.LogRed(err.Error())
	}
}

func test() {
}
