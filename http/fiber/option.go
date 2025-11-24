package fiber

import "github.com/gofiber/fiber/v2"

// FiberOption 功能型 Option
type FiberOption func(*fiberOptions)

type fiberOptions struct {
	Config        fiber.Config
	EnableLogger  bool
	EnableRecover bool
	Port          int
	Custom        func(*fiber.App)
}

// 默认值
func defaultFiberOptions() *fiberOptions {
	return &fiberOptions{
		Config:        fiber.Config{},
		EnableLogger:  true,
		EnableRecover: true,
		Custom:        nil,
		Port:          8080,
	}
}

// Option helper
func WithLogger(enabled bool) FiberOption {
	return func(o *fiberOptions) { o.EnableLogger = enabled }
}

func WithRecover(enabled bool) FiberOption {
	return func(o *fiberOptions) { o.EnableRecover = enabled }
}

func WithCustomHandler(f func(app *fiber.App)) FiberOption {
	return func(o *fiberOptions) { o.Custom = f }
}

func WithConfig(cfg fiber.Config) FiberOption {
	return func(o *fiberOptions) { o.Config = cfg }
}

func WithPort() {

}
