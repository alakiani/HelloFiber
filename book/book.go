package book

import (
	"github.com/gofiber/fiber"
)

// GetBooks sends all books
func GetBooks(c *fiber.Ctx) {
	c.Send("All Books")
}

// GetBook sends all books
func GetBook(c *fiber.Ctx) {
	c.Send("All Books")
}

// NewBook sends all books
func NewBook(c *fiber.Ctx) {
	c.Send("All Books")
}

// DeleteBook sends all books
func DeleteBook(c *fiber.Ctx) {
	c.Send("All Books")
}
