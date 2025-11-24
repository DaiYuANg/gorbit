package eventbus

import (
	"github.com/stanipetrosyan/go-eventbus"
	"go.uber.org/fx"
)

var Module = fx.Module("event_bus",
	fx.Provide(goeventbus.NewEventBus),
)
