package http

import "github.com/gofiber/fiber/v3"

func newRouter() *fiber.App {
	router := fiber.New()
	return router
}
