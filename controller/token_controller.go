package controller

import (
	"github.com/fajarnugraha37/go-rest-api/util"
	"github.com/gofiber/fiber/v2"
)

func GetNewAccessToken(c *fiber.Ctx) error {
	token, err := util.GenerateToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"message":      nil,
		"access_token": token,
	})
}
