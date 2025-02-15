package main

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		slog.Info("GET /hello")
		return c.JSON(map[string]string{"msg": "Hello 大阪!"})
	})
	_ = app.Listen(":" + os.Getenv("PORT"))
}
