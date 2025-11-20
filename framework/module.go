package framework

import "github.com/samber/do/v2"

type Module interface {
	Name() string
	Register(i do.Injector) error
	Init(ctx *AppContext) error  // 可选
	Start(ctx *AppContext) error // 可选
	Stop(ctx *AppContext) error  // 可选
}
