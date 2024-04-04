package model

import (
	"github.com/google/uuid"
	"time"
)

type MediaType string

const (
	Video MediaType = "video"
	Photo MediaType = "photo"
)

type Media struct {
	ID       uuid.UUID `json:"id"`
	BlogID   uuid.UUID `json:"blogId"`
	Path     string    `json:"path"`
	Source   string    `json:"source"`
	Name     string    `json:"name"`
	Type     MediaType `json:"type"`
	Uploaded time.Time `json:"uploaded"`
}
