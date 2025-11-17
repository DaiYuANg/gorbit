package framework

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/samber/do/v2"
	"github.com/samber/oops"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

func New(opts ...Options) (*Framework, error) {
	i := do.New()
	do.NewWithOpts(&do.InjectorOpts{
		HookBeforeRegistration: []func(scope *do.Scope, serviceName string){
			func(scope *do.Scope, serviceName string) {

			},
		},
		HookAfterRegistration:    nil,
		HookBeforeInvocation:     nil,
		HookAfterInvocation:      nil,
		HookBeforeShutdown:       nil,
		HookAfterShutdown:        nil,
		Logf:                     nil,
		HealthCheckParallelism:   0,
		HealthCheckGlobalTimeout: 0,
		HealthCheckTimeout:       0,
		StructTagKey:             "",
	})
	scheduler, err := gocron.NewScheduler(
		gocron.WithLogger(
			gocron.NewLogger(gocron.LogLevelInfo),
		),
	)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	eventbus := goeventbus.NewEventBus()
	fw := &Framework{
		injector: i,
		//Options:  opts,
		eventBus:  &eventbus,
		scheduler: &scheduler,
	}
	return fw, err
}

// Use 注册模块
func (fw *Framework) Use(m Module) error {
	fw.modules = append(fw.modules, m)
	return oops.Wrap(m.Register(fw.injector))
}

// Run 启动服务（例如启动 HTTP server）
func (fw *Framework) Run() error {
	// Init 阶段
	for _, m := range fw.modules {
		if err := m.Init(fw.injector); err != nil {
			return oops.Wrapf(err, "init module %s", m.Name())
		}
	}

	// Start 阶段
	for _, m := range fw.modules {
		if err := m.Start(fw.injector); err != nil {
			return oops.Wrapf(err, "start module %s", m.Name())
		}
	}

	// 阻塞等待（例如 HTTP 模块）
	return nil
}
