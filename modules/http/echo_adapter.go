package http

import (
	"context"

	"github.com/labstack/echo/v4"
)

type EchoAdapter struct {
	E *echo.Echo
}

func (ea *EchoAdapter) Listen(addr string) error {
	return ea.E.Start(addr)
}

func (ea *EchoAdapter) Shutdown() error {
	return ea.E.Shutdown(context.Background())
}
