package main

import (
	"notes-api-fiber/database"
	"notes-api-fiber/router"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	database.Connect()

	router.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
