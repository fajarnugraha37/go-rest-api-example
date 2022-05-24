package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/fajarnugraha37/go-rest-api/config"
	"github.com/fajarnugraha37/go-rest-api/database"
	"github.com/fajarnugraha37/go-rest-api/middleware"
	"github.com/fajarnugraha37/go-rest-api/route"
	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	defer database.DB.Close()

	config := config.ServerConfig()
	app := fiber.New(config)
	middleware.ServerMiddleware(app)

	route.PublicGroup(app)
	route.PrivateGroup(app)
	route.NotFoundRoute(app)

	StartServer(app)
}

func StartServerWithGracefulShutdown(a *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

func StartServer(a *fiber.App) {
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
