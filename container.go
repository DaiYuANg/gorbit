package gorbit

import (
	"log/slog"

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

func CreateContainer(option ...fx.Option) (*fx.App, error) {
	options := append(option, commonOption)
	err := fx.ValidateApp(options...)
	if err != nil {
		return nil, err
	}
	app := fx.New(
		commonOption,
		fx.Options(option...),
	)

	return app, nil
}
