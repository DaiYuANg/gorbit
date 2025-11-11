package framework

import (
	"github.com/gofiber/fiber/v3"
	"github.com/samber/do/v2"
	"github.com/samber/oops"
)

func New(opts Options) *Framework {
	i := do.New()
	fw := &Framework{
		Injector: i,
		Options:  opts,
	}
	return fw
}

// Use 注册模块
func (fw *Framework) Use(register func(i do.Injector) error) error {
	return oops.Wrap(register(fw.Injector))
}

// Run 启动服务（例如启动 HTTP server）
func (fw *Framework) Run() error {
	if fiberApp, err := do.Invoke[*fiber.App](fw.Injector); err == nil {
		return oops.Wrap(fiberApp.Listen(":8080"))
	}
	return nil
}
