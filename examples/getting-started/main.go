package main

import (
	"github.com/daiyuang/gorbit/framework"
	"github.com/daiyuang/gorbit/http"
	"github.com/gofiber/fiber/v3"
)

func main() {
	fw := framework.New()

	// 注入 Fiber 版本 HTTP 模块
	fw.Use(&http.Module{
		Server: &http.FiberAdapter{App: fiber.New()},
		Addr:   ":8080",
	})

	fw.Run()
}
