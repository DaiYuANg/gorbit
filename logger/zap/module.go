package zap_logger

import (
	"log/slog"

	slogzap "github.com/samber/slog-zap/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule(opts ...Option) fx.Option {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	zapLogger := newLogger(cfg)
	logger := slog.New(slogzap.Option{Level: slog.LevelDebug, Logger: zapLogger}.NewZapHandler())
	return fx.Module(
		"logger_module",
		fx.Provide(
			func() *Config { return cfg },
			func() *zap.Logger {
				return zapLogger
			},
			func() *slog.Logger {
				return logger
			},
			sugaredLogger,
		),
		fx.Invoke(deferLogger),
	)
}
