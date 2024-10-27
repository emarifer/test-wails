package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist components
var assets embed.FS

//go:embed build/appicon.png
var icon []byte
var version = "1.0.0"

func main() {
	// Create an instance of the app structure and custom Middleware
	app := NewApp()
	r := NewChiRouter()

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Wails Demo",
		Width:  1040,
		Height: 768,
		// MinWidth:          1040,
		// MinHeight:         768,
		// MaxWidth:          1280,
		// MaxHeight:         800,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer: &assetserver.Options{
			Assets: assets,
			Middleware: func(next http.Handler) http.Handler {
				r.NotFound(next.ServeHTTP)
				return r
			},
		},
		Menu:             AppMenu,
		Logger:           nil,
		LogLevel:         logger.DEBUG,
		OnStartup:        app.startup,
		OnDomReady:       app.domReady,
		OnBeforeClose:    app.beforeClose,
		OnShutdown:       app.shutdown,
		WindowStartState: options.Normal,
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
			ZoomFactor:          1.0,
		},
		// Linux platform specific options
		Linux: &linux.Options{
			Icon: icon,
			// WindowIsTranslucent: true,
			WebviewGpuPolicy: linux.WebviewGpuPolicyNever,
			// ProgramName:         "wails",
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "test-wails",
				Message: "",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

/*
https://github.com/PylotLight/wails-htmx-templ-template
https://github.com/wailsapp/wails/issues/2977#issuecomment-1791231550

COMMAND FOR BUILD:
wails dev -tags webkit2_40
wails build -tags webkit2_40

COMMAND FOR WINDOWS BUILD:
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc wails build -skipbindings -s -platform windows/amd64
// CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc wails build -skipbindings -s -platform windows/amd64

https://madin.dev/cross-wails/

https://www.google.com/search?q=wails+compile+sqlite+for+windows&sca_esv=ac88267094842ae4&sxsrf=ADLYWIIB60WncbSg7MnPBYV7977EHLjeSQ%3A1729878267256&ei=-9gbZ--gD5qQxc8Pw87I2AE&ved=0ahUKEwiviv_ciqqJAxUaSPEDHUMnEhsQ4dUDCA8&uact=5&oq=wails+compile+sqlite+for+windows&gs_lp=Egxnd3Mtd2l6LXNlcnAiIHdhaWxzIGNvbXBpbGUgc3FsaXRlIGZvciB3aW5kb3dzMggQABiABBiiBDIIEAAYgAQYogQyCBAAGIAEGKIESNZbUL0gWL5DcAJ4AZABAJgBrwGgAc4IqgEDMS44uAEDyAEA-AEBmAIKoAKBCMICChAAGLADGNYEGEfCAgQQIxgnwgIIECEYoAEYwwSYAwCIBgGQBgiSBwMyLjigB7wV&sclient=gws-wiz-serp
*/
