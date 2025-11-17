package http

import (
	"github.com/gofiber/fiber/v3"
)

type FiberAdapter struct {
	App *fiber.App
}

func (fa *FiberAdapter) Listen(addr string) error {
	return fa.App.Listen(addr)
}

func (fa *FiberAdapter) Shutdown() error {
	return fa.App.Shutdown()
}
