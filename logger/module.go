package logger

import (
	"log/slog"

	"go.uber.org/fx"
)

func NewModule(opts ...Option) fx.Option {
	return fx.Module("logger",
		fx.Provide(
			func(lc fx.Lifecycle) (*slog.Logger, error) {
				log, closer, err := NewLogger(opts...)
				if err != nil {
					return nil, err
				}

				RegisterLifecycle(lc, closer)
				return log, nil
			},
		),
	)
}
