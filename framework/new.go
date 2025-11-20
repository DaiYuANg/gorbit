package framework

import (
	"log/slog"

	"github.com/google/uuid"
	"github.com/samber/do/v2"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

func New(modules ...Module) *Framework {
	inj := do.New()
	slog.Info("Framework init")
	eb := goeventbus.NewEventBus()

	ctx := newAppContext(inj, eb)
	slog.Info("Framework Context Wire Finish")
	return &Framework{
		injector: inj,
		eventBus: eb,
		modules:  modules,
		ctx:      ctx,
		appID:    uuid.New().String(),
	}
}

func (f *Framework) RegisterModule(m Module) {
	f.modules = append(f.modules, m)
}
