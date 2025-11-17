package config

import (
	"github.com/samber/do/v2"
	"github.com/samber/oops"
)

type Module struct {
	opts    Options
	backend Backend
}

func NewConfigModule(opts Options) *Module {
	if opts.App == nil {
		opts.App = &AppConfig{}
	}
	return &Module{
		opts:    opts,
		backend: opts.Backend,
	}
}

func (m *Module) Name() string { return "config" }

func (m *Module) Register(i do.Injector) error {
	// 将 backend 本身注入，让业务模块也能直接拿到 raw backend（可选）
	do.Provide[Backend](i, func(injector do.Injector) (Backend, error) {
		return m.backend, nil
	})
	// 注入 AppConfig 占位（值将在 Init 阶段被填充）
	do.Provide[AppConfig](i, func(injector do.Injector) (AppConfig, error) {
		return *m.opts.App, nil
	})
	return nil
}

func (m *Module) Init(i do.Injector) error {
	// 1) 加载所有来源（文件/dotenv/env）由 backend 负责
	if err := m.backend.Load(m.opts); err != nil {
		return oops.Wrap(err)
	}

	// 2) 将 AppConfig 从 backend Unmarshal（path "" 表示根）
	if err := m.backend.Unmarshal("", m.opts.App); err != nil {
		return oops.Wrapf(err, "unmarshal app config")
	}

	// 3) 为所有已注册的 target 执行 Unmarshal 并注入到 do.Injector
	mu.Lock()
	targets := make([]any, len(registeredTargets))
	copy(targets, registeredTargets)
	mu.Unlock()

	for _, t := range targets {
		if err := m.backend.Unmarshal("", t); err != nil {
			return oops.Wrapf(err, "unmarshal target %T", t)
		}
		// 把解析后的 target 注入 DI（按类型）
		do.Provide[any](i, func(injector do.Injector) (any, error) {
			return t, nil
		})
	}

	return nil
}

func (m *Module) Start(i do.Injector) error {
	// no-op
	return nil
}

func (m *Module) Stop(i do.Injector) error {
	// no-op
	return nil
}
