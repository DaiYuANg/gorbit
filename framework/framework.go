package framework

import (
	"github.com/samber/do/v2"
	"github.com/samber/oops"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

type Framework struct {
	injector do.Injector
	modules  []Module
	eventBus *goeventbus.EventBus
	ctx      *AppContext
}

func New() *Framework {
	inj := do.New()
	eb := goeventbus.New()

	ctx := newAppContext(inj, eb)

	return &Framework{
		injector: inj,
		eventBus: eb,
		modules:  []Module{},
		ctx:      ctx,
	}
}

func (f *Framework) RegisterModule(m Module) {
	f.modules = append(f.modules, m)
}

// --- lifecycle events topic ---
const (
	EventFrameworkRegisterStart = "framework.register.start"
	EventFrameworkRegisterDone  = "framework.register.done"
	EventFrameworkInitStart     = "framework.init.start"
	EventFrameworkInitDone      = "framework.init.done"
	EventFrameworkStart         = "framework.start.start"
	EventFrameworkStartDone     = "framework.start.done"
	EventFrameworkStop          = "framework.stop.start"
	EventFrameworkStopDone      = "framework.stop.done"
)
