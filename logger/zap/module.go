package zap_logger

import (
	"log/slog"

	"github.com/samber/lo"
	slogzap "github.com/samber/slog-zap/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule(opts ...Option) fx.Option {
	cfg := defaultConfig()
	lo.ForEach(opts, func(item Option, _ int) {
		item(cfg)
	})
	zapLogger := newLogger(cfg)
	logger := slog.New(slogzap.Option{Level: slog.LevelInfo, Logger: zapLogger}.NewZapHandler())
	return fx.Module(
		"logger",
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
