package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ServerMiddleware(a *fiber.App) {
	// docs ==>: https://docs.gofiber.io/api/middleware
	a.Use(
		cors.New(),
		logger.New(),
	)
}
