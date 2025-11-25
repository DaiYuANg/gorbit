package main

import (
	"log/slog"

	"github.com/DaiYuANg/gorbit"
	"github.com/DaiYuANg/gorbit/cli"
	"github.com/DaiYuANg/gorbit/database/bun"
	"github.com/DaiYuANg/gorbit/eventbus"
	"github.com/DaiYuANg/gorbit/logger/zap_logger"
	"github.com/DaiYuANg/gorbit/modules/config"
	"github.com/DaiYuANg/gorbit/scheduler"
	vm "github.com/DaiYuANg/gorbit/validator"
	"github.com/go-playground/validator/v10"
	"github.com/samber/oops"
	"go.uber.org/fx"
)

func main() {
	container, err := gorbit.CreateContainerWithFxLogger(
		zap_logger.NewModule(),
		cli.NewCLIModule(),
		config.NewConfigModule[ApplicationConfig](ApplicationConfig{}, config.WithEnvPrefix[ApplicationConfig]("MINI_")),
		schedule.NewSchedulerModule(),
		//jwt.NewJwtModule(),
		bun.NewDatabaseModule(),
		vm.NewValidatorModule(validator.WithRequiredStructEnabled(), validator.WithPrivateFieldValidation()),
		eventbus.NewEventBusModule(),
		fx.Module("internal", fx.Invoke(
			func(applicationConfig *ApplicationConfig, logger *slog.Logger) {
				logger.Info("ApplicationConfig", "config", slog.AnyValue(applicationConfig))
			},
		)),
	)
	if err != nil {
		panic(oops.Wrap(err))
	}
	container.Run()
}
