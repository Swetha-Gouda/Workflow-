package main

import "github.com/gofiber/fiber/v2"

func new() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello CMD Practice 🚀")
	})

	app.Listen(":3000")
}
