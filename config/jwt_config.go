package config

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
)

// Create config for JWT authentication middleware.
func JwtConfig(ErrorHandler func(*fiber.Ctx, error) error) jwtMiddleware.Config {
	return jwtMiddleware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: ErrorHandler,
	}
}
