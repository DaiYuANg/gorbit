package gorbit

import (
	"log/slog"

	"github.com/samber/oops"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var commonOption = fx.Options(
	fx.WithLogger(func(log *slog.Logger) fxevent.Logger {
		fxLogger := &fxevent.SlogLogger{Logger: log}
		fxLogger.UseLogLevel(slog.LevelDebug)
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
			Wrap(err).(*oops.OopsError)
	}

	// 创建 fx.App
	app := fx.New(options...)

	return app, nil
}
