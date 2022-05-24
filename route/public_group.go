package route

import (
	"github.com/fajarnugraha37/go-rest-api/controller"
	"github.com/gofiber/fiber/v2"
)

func PublicGroup(app *fiber.App) {
	route := app.Group("/api")

	route.Get("/books", controller.GetBooks)

	route.Get("/book/:id", controller.GetBook)

	route.Get("/token/new", controller.GetNewAccessToken)
}
