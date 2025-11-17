package config

import (
	"sync"

	"github.com/DaiYuANg/gorbit/framework"
	"github.com/samber/do/v2"
)

var (
	mu                sync.Mutex
	registeredTargets []any
)

// ProvideConfig 用于用户注册自己的 config 结构体（例如 &DatabaseConfig{}）
// 返回一个轻量 Module：在 Register 阶段把 target append 到全局列表。
// 这样不依赖模块注册顺序。
func ProvideConfig(target any) framework.Module {
	return &targetModule{target: target}
}

type targetModule struct{ target any }

func (t *targetModule) Name() string { return "config:target" }
func (t *targetModule) Register(_ do.Injector) error {
	mu.Lock()
	registeredTargets = append(registeredTargets, t.target)
	mu.Unlock()
	return nil
}
func (*targetModule) Init(_ do.Injector) error  { return nil }
func (*targetModule) Start(_ do.Injector) error { return nil }
func (*targetModule) Stop(_ do.Injector) error  { return nil }
