package internal

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID       uuid.UUID ` json:"id" Gorm:"primary_key:uuid;default:gen_random_uuid()"`
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Text     string    `json:"text"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}
