package main

import (
	"github.com/DaiYuANg/gorbit"
	"github.com/DaiYuANg/gorbit/database/bun"
	"github.com/DaiYuANg/gorbit/eventbus"
	"github.com/DaiYuANg/gorbit/http/fiber"
	"github.com/DaiYuANg/gorbit/jwt"
	"github.com/DaiYuANg/gorbit/logger"
	"github.com/DaiYuANg/gorbit/modules/cli"
	"github.com/DaiYuANg/gorbit/modules/config"
	"github.com/DaiYuANg/gorbit/schedule"
	"github.com/DaiYuANg/gorbit/validator"
)

func main() {
	container, err := gorbit.CreateContainer(
		logger.NewModule(),
		cli.NewCLIModule(),
		config.NewConfigModule(UserConfig{}),
		schedule.Module,
		jwt.NewJwtModule(),
		bun.NewDatabaseModule(),
		validator.Module,
		eventbus.Module,
		fiber.NewFiberModule(),
	)
	if err != nil {
		panic(err)
	}
	container.Run()
}
