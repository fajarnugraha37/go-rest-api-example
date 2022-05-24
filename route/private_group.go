package route

import (
	"github.com/fajarnugraha37/go-rest-api/controller"
	"github.com/fajarnugraha37/go-rest-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateGroup(app *fiber.App) {
	route := app.Group("/api")

	route.Post("/book", middleware.JWTProtected(), controller.CreateBook)

	route.Put("/book", middleware.JWTProtected(), controller.UpdateBook)

	route.Delete("/book", middleware.JWTProtected(), controller.DeleteBook)
}
