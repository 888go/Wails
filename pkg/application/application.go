package application

import (
	"context"
	"sync"

	"github.com/wailsapp/wails/v2/internal/app"
	"github.com/wailsapp/wails/v2/internal/signal"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
)

// Application 是 Wails 主应用程序
type Application struct {
	application *app.App
	options     *options.App

	// running flag
	running bool

	shutdown sync.Once
}

// NewWithOptions 使用给定的选项创建一个新的Application

// ff:创建并按选项
// options:App选项
func NewWithOptions(options *options.App) *Application {
	if options == nil {
		return New()
	}
	return &Application{
		options: options,
	}
}

// New 创建一个使用默认选项的新 Application

// ff:创建并按默认选项
func New() *Application {
	return &Application{
		options: &options.App{},
	}
}

// 设置应用菜单 将设置应用程序的菜单

// ff:设置菜单
// appMenu:菜单
func (a *Application) SetApplicationMenu(appMenu *menu.Menu) {
	if a.running {
		a.application.SetApplicationMenu(appMenu)
		return
	}

	a.options.Menu = appMenu
}

// Run 启动应用程序

// ff:运行
func (a *Application) Run() error {
	err := applicationInit()
	if err != nil {
		return err
	}

	application, err := app.CreateApp(a.options)
	if err != nil {
		return err
	}

	a.application = application

	// Control-C handlers
	signal.OnShutdown(func() {
		a.application.Shutdown()
	})
	signal.Start()

	a.running = true

	err = a.application.Run()
	return err
}

// Quit 将关闭应用程序

// ff:退出
func (a *Application) Quit() {
	a.shutdown.Do(func() {
		a.application.Shutdown()
	})
}

// 将给定的结构体绑定到应用程序

// ff:
// boundStruct:
func (a *Application) Bind(boundStruct any) {
	a.options.Bind = append(a.options.Bind, boundStruct)
}


// ff:
// callback:
// eventType:
func (a *Application) On(eventType EventType, callback func()) {
	c := func(ctx context.Context) {
		callback()
	}

	switch eventType {
	case StartUp:
		a.options.OnStartup = c
	case ShutDown:
		a.options.OnShutdown = c
	case DomReady:
		a.options.OnDomReady = c
	}
}
