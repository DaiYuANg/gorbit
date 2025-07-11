package framework

import (
	"github.com/gofiber/fiber/v3"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/v2"
	"go.uber.org/fx"
)

type Framework[T any] struct {
	fxOpts      []fx.Option
	fiberRoutes []func(router fiber.Router)
	cfgParser   *koanf.Koanf
	cfg         T
	app         *fx.App
}

func (t Framework[T]) Run() {
	t.app.Run()
}
