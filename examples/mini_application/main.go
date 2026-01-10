package main

import (
	"log/slog"

	"github.com/DaiYuANg/gorbit"
	"github.com/DaiYuANg/gorbit/cli"
	"github.com/DaiYuANg/gorbit/config"
	"github.com/DaiYuANg/gorbit/logger"
	"github.com/DaiYuANg/gorbit/scheduler"
	"github.com/samber/oops"
	"go.uber.org/fx"
)

func main() {
	container, err := gorbit.CreateContainerWithFxLogger(
		logger.NewModule(),
		cli.NewCLIModule(),
		config.NewConfigModule[ApplicationConfig](ApplicationConfig{}, config.WithEnvPrefix[ApplicationConfig]("MINI_")),
		schedule.NewSchedulerModule(),
		fx.Module("internal",
			fx.Invoke(
				func(applicationConfig *ApplicationConfig, logger *slog.Logger) {
					logger.Info("ApplicationConfig", "config", slog.AnyValue(applicationConfig))
				},
			),
		),
	)
	if err != nil {
		panic(oops.Wrap(err))
	}
	container.Run()
}
