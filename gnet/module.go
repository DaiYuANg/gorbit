package gnet

import (
	"context"
	"log/slog"

	"github.com/panjf2000/gnet/v2"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"go.uber.org/fx"
)

type ServerConstructor func(logger *slog.Logger) gnet.EventHandler

func NewModule(serverCtor ServerConstructor, opts ...Option) fx.Option {
	cfg := &Config{Addr: "tcp://:8080", Multicore: true}
	lo.ForEach(opts, func(item Option, _ int) { item(cfg) })

	return fx.Module(
		"gnet",
		fx.Provide(func(logger *slog.Logger) gnet.EventHandler {
			return serverCtor(logger)
		}),
		fx.Invoke(func(lc fx.Lifecycle, server gnet.EventHandler, logger *slog.Logger) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						if err := gnet.Run(
							server,
							cfg.Addr,
							gnet.WithMulticore(cfg.Multicore),
							gnet.WithLogger(&gnetSlogLogger{logger}),
						); err != nil {
							logger.Error("gnet start failed",
								slog.String("addr", cfg.Addr),
								slog.Any("err", oops.Wrap(err)),
							)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					if s, ok := server.(interface{ Stop() error }); ok {
						return s.Stop()
					}
					return nil
				},
			})
		}),
	)
}
