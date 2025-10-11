package framework

import (
	"context"

	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
)

type App struct {
	container *do.RootScope
	opts      Options
	rootCmd   *cobra.Command
}

func New(opts ...Option) *App {
	o := defaultOptions()
	for _, opt := range opts {
		opt(&o)
	}

	injector := do.New()

	app := &App{
		container: injector,
		opts:      o,
		rootCmd:   &cobra.Command{Use: o.Name, Short: o.Description},
	}

	// 注册配置模块
	if o.ConfigLoader != nil {
		cfg, err := o.ConfigLoader()
		if err != nil {
			panic(err)
		}
		do.ProvideValue(injector, cfg)
	}

	// 注册命令
	if o.RootRun != nil {
		app.rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			return o.RootRun(ctx, injector)
		}
	}

	// 注册子命令
	for _, sub := range o.SubCommands {
		app.rootCmd.AddCommand(sub)
	}

	return app
}

func (a *App) Run() error {
	return a.rootCmd.Execute()
}
