package controller

import (
	"time"

	"github.com/fajarnugraha37/go-rest-api/database"
	"github.com/fajarnugraha37/go-rest-api/database/model"
	"github.com/fajarnugraha37/go-rest-api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetBooks(c *fiber.Ctx) error {
	db, err := database.OpenConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	books, err := db.GetBooks()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "books were not found",
			"count":   0,
			"books":   nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"count":   len(books),
		"books":   books,
	})
}

func GetBook(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	db, err := database.OpenConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	book, err := db.GetBook(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "book with the given ID is not found",
			"book":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"book":    book,
	})
}

func CreateBook(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := util.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "unauthorized, check expiration time of your token",
		})
	}

	book := &model.Book{}
	if err := c.BodyParser(book); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	db, err := database.OpenConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	validate := util.NewValidator()

	book.ID = uuid.New()
	book.CreatedAt = time.Now()
	book.BookStatus = 1 // 0 == draft, 1 == active

	if err := validate.Struct(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": util.ValidatorErrors(err),
		})
	}

	if err := db.CreateBook(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"book":    book,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := util.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "unauthorized, check expiration time of your token",
		})
	}

	book := &model.Book{}
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	db, err := database.OpenConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	foundedBook, err := db.GetBook(book.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "book with this ID not found",
		})
	}

	book.UpdatedAt = time.Now()

	validate := util.NewValidator()

	if err := validate.Struct(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": util.ValidatorErrors(err),
		})
	}

	if err := db.UpdateBook(foundedBook.ID, book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DeleteBook(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := util.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "unauthorized, check expiration time of your token",
		})
	}

	book := &model.Book{}
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	validate := util.NewValidator()
	if err := validate.StructPartial(book, "id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": util.ValidatorErrors(err),
		})
	}

	db, err := database.OpenConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	foundedBook, err := db.GetBook(book.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "book with this ID not found",
		})
	}

	if err := db.DeleteBook(foundedBook.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
