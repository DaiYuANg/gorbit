package logger

import (
	"context"

	"go.uber.org/fx"
)

func RegisterLifecycle(lc fx.Lifecycle, closer Closer) {
	if closer == nil {
		return
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return closer()
		},
	})
}
