package route

import "github.com/gofiber/fiber/v2"

func NotFoundRoute(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "sorry, endpoint is not found",
		})
	})
}
