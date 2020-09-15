package main

import (
	"github.com/alakiani/HelloFiber/book"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", helloworld)

	app.Listen(3000)
}

func helloworld(c *fiber.Ctx) {
	c.Send("Hello, World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
}
