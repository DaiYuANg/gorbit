package framework

import "github.com/samber/do/v2"

type Module interface {
	Name() string
	Register(i do.Injector) error
	Init(ctx *AppContext) error  // 可选
	Start(ctx *AppContext) error // 可选
	Stop(ctx *AppContext) error  // 可选
}

// 可独立使用的优先级接口
type Prioritizer interface {
	Priority() int
}

// 默认优先级实现，可嵌入模块
type Priority struct {
	Value int
}

func (p Priority) Priority() int {
	return p.Value
}
