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

func DeleteNoteHandler(c fiber.Ctx) error {
	db := database.DB
	var note internal.Note
	err := db.Where("title = ?", c.Params("title")).First(&note).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot Find",
		})
	}
	err = db.Delete(&note).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot delete note",
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"error": "Note has been deleted",
	})
}
func UpdateNoteHandler(c fiber.Ctx) error {
	db := database.DB
	var note internal.Note
	err := db.Where("title = ?", c.Params("title")).First(&note).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot update note",
		})
	}
	note.Title = c.Params("title")
	note.Subtitle = c.Params("subtitle")
	note.Text = c.Params("text")
	err = db.Save(&note).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot update note",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"note": note,
	})
}
