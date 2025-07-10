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
	cfg         *koanf.Koanf
	coreCfg     *AppConfig
	userCfg     T
}
