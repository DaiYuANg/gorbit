package main

import (
	"log/slog"

	"github.com/DaiYuANg/gorbit"
	"github.com/DaiYuANg/gorbit/eventbus"
	gorbit_gnet "github.com/DaiYuANg/gorbit/gnet"
	"github.com/DaiYuANg/gorbit/logger"
	"github.com/panjf2000/gnet/v2"
)

func main() {
	container, err := gorbit.CreateContainerWithFxLogger(
		logger.NewModule(),
		eventbus.NewEventBusModule(),
		gorbit_gnet.NewModule(func(logger *slog.Logger) gnet.EventHandler {
			return &testServer{
				logger: logger,
			}
		}),
	)
	if err != nil {
		panic(err)
	}
	container.Run()
}
