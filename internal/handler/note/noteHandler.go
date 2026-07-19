package internal

import (
	"notes-api-fiber/database"
	internal "notes-api-fiber/internal/model"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Request struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Text     string `json:"text"`
}

func CreateNoteHandler(c fiber.Ctx) error {
	db := database.DB
	var req Request

	err := c.Bind().Body(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	note := new(internal.Note)
	note.ID = uuid.New()
	note.Title = req.Title
	note.Subtitle = req.Subtitle
	note.Text = req.Text

	err = db.Create(&note).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot create note",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"note": note,
	})
}
