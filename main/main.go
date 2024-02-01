package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Server Listen By fiber\n")
	})

	log.Fatal(app.Listen(":3000"))
}
