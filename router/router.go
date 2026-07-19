package router

import (
	"notes-api-fiber/internal/routes/note"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("", logger.New())
	// Notes Routes
	note.SetupNoteRoutes(api)
}
