package model

import (
	"github.com/google/uuid"
	"time"
)

type Blog struct {
	ID      uuid.UUID `json:"id"`
	UserId  uuid.UUID `json:"userId"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
	Media   []Media   `json:"media"`
}
