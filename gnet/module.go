package gnet

import (
	"context"
	"log/slog"

	"github.com/panjf2000/gnet/v2"
	"github.com/samber/oops"
	"go.uber.org/fx"
)

func NewModule(userServer gnet.EventHandler, logger *slog.Logger, opts ...Option) fx.Option {
	cfg := &Config{
		Addr:      ":8880",
		Multicore: true,
	}
	for _, o := range opts {
		o(cfg)
	}

	return fx.Module(
		"v_server",
		fx.Invoke(func(lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						if err := gnet.Run(userServer, cfg.Addr, gnet.WithMulticore(cfg.Multicore)); err != nil {
							logger.Error("gnet start failed",
								slog.String("addr", cfg.Addr),
								slog.Any("err", oops.Wrap(err)),
							)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					// 如果 userServer 支持 Stop，可以用断言
					if s, ok := userServer.(interface{ Stop() error }); ok {
						return s.Stop()
					}
					return nil
				},
			})
		}),
	)
}
