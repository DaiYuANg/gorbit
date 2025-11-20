package framework

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
)

func (f *Framework) Run() error {
	f.ctx.Publish(EventFrameworkRegisterStart, nil)

	// 1) Register 阶段（全部模块）
	for _, m := range f.modules {
		if err := m.Register(f.injector); err != nil {
			return fmt.Errorf("module %s register failed: %w", m.Name(), err)
		}
	}

	f.ctx.Publish(EventFrameworkRegisterDone, nil)

	// 2) Init 阶段
	f.ctx.Publish(EventFrameworkInitStart, nil)

	for _, m := range f.modules {
		if err := m.Init(f.ctx); err != nil {
			return fmt.Errorf("module %s init failed: %w", m.Name(), err)
		}
	}

	f.ctx.Publish(EventFrameworkInitDone, nil)

	// 3) Start 阶段
	f.ctx.Publish(EventFrameworkStart, nil)

	for _, m := range f.modules {
		if err := m.Start(f.ctx); err != nil {
			return fmt.Errorf("module %s start failed: %w", m.Name(), err)
		}
	}

	f.ctx.Publish(EventFrameworkStartDone, nil)
	return nil
}

func (f *Framework) Stop() {
	f.ctx.Publish(EventFrameworkStop, nil)

	for _, m := range f.modules {
		_ = m.Stop(f.ctx)
	}

	f.ctx.Publish(EventFrameworkStopDone, nil)
}
