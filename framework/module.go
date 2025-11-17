package framework

import "github.com/samber/do/v2"

type Module interface {
	Name() string
	Register(i do.Injector) error
	Init(i do.Injector) error  // 可选
	Start(i do.Injector) error // 可选
	Stop(i do.Injector) error  // 可选
}
