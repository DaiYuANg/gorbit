package fiber

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/samber/oops"
	"go.uber.org/fx"
)

// NewFiberModule 返回 fx.Module
func NewFiberModule(opts ...FiberOption) fx.Option {
	options := defaultFiberOptions()
	for _, o := range opts {
		o(options)
	}

	return fx.Module("fiber_http",
		fx.Provide(func() *fiber.App {
			app := fiber.New(options.Config)

			if options.EnableRecover {
				app.Use(recover.New())
			}
			if options.EnableLogger {
				app.Use(logger.New())
			}

			if options.Custom != nil {
				options.Custom(app)
			}

			return app
		}),
		fx.Invoke(func(lc fx.Lifecycle, app *fiber.App, logger *slog.Logger) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					// 不启动监听，由用户自行调用 Listen 或自己在 invoke 中启动
					logger.Info("Fiber module ready")
					go func() {
						err := app.Listen(fmt.Sprintf(":%d", options.Port))
						if err != nil {
							panic(oops.Wrap(err))
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return app.Shutdown()
				},
			})
		}),
	)
}
