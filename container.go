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
func CreateContainer(withFxLogger bool, options ...fx.Option) (*fx.App, error) {
	var allOptions []fx.Option

	if withFxLogger {
		// 仅在需要时追加 commonOption
		allOptions = append(options, commonOption)
	} else {
		// 不追加 commonOption
		allOptions = options
	}
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

func CreateContainerWithFxLogger(options ...fx.Option) (*fx.App, error) {
	return CreateContainer(true, options...)
}
