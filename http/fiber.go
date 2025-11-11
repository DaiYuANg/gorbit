package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/samber/do/v2"
)

func NewFiber(config ...fiber.Config) func(i do.Injector) error {
	return func(i do.Injector) error {
		do.Provide(i, func(i do.Injector) (*fiber.App, error) {
			app := fiber.New(config...)
			return app, nil
		})
		return nil
	}
}
