package framework

import (
	"github.com/samber/do/v2"
	"github.com/samber/oops"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

func New(opts ...Options) *Framework {
	i := do.New()
	eventbus := goeventbus.NewEventBus()
	fw := &Framework{
		Injector: i,
		//Options:  opts,
		EventBus: &eventbus,
	}
	return fw
}

// Use 注册模块
func (fw *Framework) Use(m Module) error {
	fw.Modules = append(fw.Modules, m)
	return oops.Wrap(m.Register(fw.Injector))
}

// Run 启动服务（例如启动 HTTP server）
func (fw *Framework) Run() error {
	// Init 阶段
	for _, m := range fw.Modules {
		if err := m.Init(fw.Injector); err != nil {
			return oops.Wrapf(err, "init module %s", m.Name())
		}
	}

	// Start 阶段
	for _, m := range fw.Modules {
		if err := m.Start(fw.Injector); err != nil {
			return oops.Wrapf(err, "start module %s", m.Name())
		}
	}

	// 阻塞等待（例如 HTTP 模块）
	return nil
}
