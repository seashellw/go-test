package main

import (
	"context"
	"embed"
	"go-test/app"
	"net/http"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	go test()

	this := &app.App{}

	go app.StartRouter(this)

	wails.Run(&options.App{
		Width:            1024,
		Height:           768,
		Assets:           assets,
		OnStartup:        func(ctx context.Context) { this.Ctx = ctx },
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		Bind:             []interface{}{this},
		AssetsHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Cross-Origin-Embedder-Policy", "require-corp")
			w.Header().Add("Cross-Origin-Opener-Policy", "same-origin")
		}),
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		},
	})

}

func test() {

}
