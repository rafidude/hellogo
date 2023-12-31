package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// create fiber instance
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("World")
	})

	app.Get("/env", func(c *fiber.Ctx) error {
		return c.SendString("Hello, ENV! " + os.Getenv("TEST_ENV"))
	})

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("port: " + port)
	app.Listen(":" + port)
}
