package framework

import (
	"github.com/samber/do/v2"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

type Framework struct {
	injector do.Injector
	modules  []Module
	eventBus goeventbus.EventBus
	ctx      *AppContext
	appID    string
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
