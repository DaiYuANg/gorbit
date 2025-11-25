package eventbus

import (
	"github.com/stanipetrosyan/go-eventbus"
	"go.uber.org/fx"
)

func NewEventBusModule() fx.Option {
	return fx.Module("event_bus",
		fx.Provide(goeventbus.NewEventBus),
	)
}
