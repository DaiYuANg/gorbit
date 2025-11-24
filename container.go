package gorbit

import (
	"log/slog"

	"github.com/samber/oops"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var commonOption = fx.Options(
	fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
		fxLogger := &fxevent.SlogLogger{Logger: logger}
		fxLogger.UseLogLevel(slog.LevelInfo)
		return fxLogger
	}),
)

// 创建 Container
func CreateContainer(options ...fx.Option) (*fx.App, error) {
	allOptions := append(options, commonOption)
	// 先验证配置
	if err := fx.ValidateApp(allOptions...); err != nil {
		return nil, oops.
			With("when", "creating fx container").
			Wrap(err)
	}

	// 创建 fx.App
	app := fx.New(allOptions...)
	return app, nil
}
