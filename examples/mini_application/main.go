package main

import (
	"github.com/DaiYuANg/gorbit"
	"github.com/DaiYuANg/gorbit/cli"
	"github.com/DaiYuANg/gorbit/database/bun"
	"github.com/DaiYuANg/gorbit/eventbus"
	"github.com/DaiYuANg/gorbit/http/fiber"
	"github.com/DaiYuANg/gorbit/logger/zap_logger"
	"github.com/DaiYuANg/gorbit/modules/config"
	"github.com/DaiYuANg/gorbit/schedule"
	vm "github.com/DaiYuANg/gorbit/validator"
	"github.com/go-playground/validator/v10"
)

func main() {
	container, err := gorbit.CreateContainer(
		zap_logger.NewModule(),
		cli.NewCLIModule(),
		config.NewConfigModule(UserConfig{}),
		schedule.Module,
		//jwt.NewJwtModule(),
		bun.NewDatabaseModule(),
		vm.NewValidatorModule(validator.WithRequiredStructEnabled(), validator.WithPrivateFieldValidation()),
		eventbus.Module,
		fiber.NewFiberModule(),
	)
	if err != nil {
		panic(err)
	}
	container.Run()
}
