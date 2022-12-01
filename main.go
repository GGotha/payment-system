package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Default()

	app := fiber.New(fiber.Config{
		AppName: "Payment System v1",
	})

	app.Route("/v1", func(api fiber.Router) {
		api.Get("/user", func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"status": "error", "message": "User doesn't exists", "data": nil})
		}).Name("foo") // /test/foo (name: test.foo)
	}, "test.")

	log.Fatal(app.Listen(":3333"))
}
