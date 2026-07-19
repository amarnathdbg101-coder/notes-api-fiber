package internal

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Text     string    `json:"text"`
}
