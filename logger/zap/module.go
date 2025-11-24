package zap_logger

import (
	"go.uber.org/fx"
)

func NewModule(opts ...Option) fx.Option {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	return fx.Module(
		"logger_module",
		fx.Provide(
			func() *Config { return cfg },
			newLogger,
			sugaredLogger,
		),
		fx.Invoke(deferLogger),
	)
}
