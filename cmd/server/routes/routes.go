package routes

import (
	"toughleaf/internal/templates"

	"github.com/gofiber/fiber/v3"
)

func Routes(r *fiber.App) {
	r.Get("/", func(c fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return templates.Home().Render(c.Context(), c.Response().BodyWriter())
	})
	r.Post("/register", func(c fiber.Ctx) error {
		return c.SendString("Register functionality not implemented yet")
	})
	r.Get("/users", func(c fiber.Ctx) error {
		return c.SendString("Users functionality not implemented yet")
	})
}
