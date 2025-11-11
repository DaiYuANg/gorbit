package main

import (
	"github.com/daiyuang/gorbit/config"
	"github.com/daiyuang/gorbit/framework"
	"github.com/daiyuang/gorbit/http"
	"github.com/gofiber/fiber/v3"
	"github.com/samber/do/v2"
)

func main() {
	fw := framework.New(framework.Options{ConfigPath: "config.yaml"})

	fw.Use(config.NewKoanf(config.ConfigOptions{EnvPrefix: "TEST"}))
	fw.Use(http.NewFiber())

	// 注册路由
	app, _ := do.Invoke[*fiber.App](fw.Injector)
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, Framework!")
	})

	fw.Run()
}
