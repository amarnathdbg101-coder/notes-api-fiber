package note

import (
	internal "notes-api-fiber/internal/handler/note"

	"github.com/gofiber/fiber/v3"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	note.Post("", internal.CreateNoteHandler)
	note.Delete("/", internal.DeleteNoteHandler)
	note.Patch("/:title", internal.UpdateNoteHandler)

}
