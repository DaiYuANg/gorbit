package logger

import (
	"log/slog"
	"os"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module("logger", fx.Provide(newLogger))
}

func newLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}
