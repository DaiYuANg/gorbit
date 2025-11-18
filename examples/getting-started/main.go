package main

import (
	"github.com/DaiYuANg/gorbit/framework"
	"github.com/DaiYuANg/gorbit/modules/config"
	"github.com/DaiYuANg/gorbit/modules/http"
	"github.com/gofiber/fiber/v3"
)

func main() {
	fw, err := framework.New()
	if err != nil {
		panic(err)
	}
	err = fw.Use(config.NewConfigModule(config.Options{
		Backend: config.BackendKoanf(),
		//Files:     []string{"config.yaml"},
		EnvFile:   ".env",
		EnvPrefix: "APP",
		OnlyEnv:   false,
		App: &config.AppConfig{
			Name: "my-app",
			Env:  "dev",
		},
	}))
	if err != nil {
		panic(err)
	}
	// 注入 Fiber 版本 HTTP 模块
	err = fw.Use(&http.Module{
		Server: &http.FiberAdapter{App: fiber.New()},
		Addr:   ":8080",
	})
	if err != nil {
		panic(err)
	}

	err = fw.Run()
	if err != nil {
		panic(err)
	}
}
