package repositories

import (
	"github.com/andradecierdo/go-api/models"
	"github.com/google/uuid"
)

type BlogRepository interface {
	GetBlog(id uuid.UUID) (models.Blog, error)
	GetBlogs() ([]models.Blog, error)
}
