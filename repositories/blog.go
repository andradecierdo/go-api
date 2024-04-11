package repositories

import (
	db "github.com/andradecierdo/go-api/database"
	"github.com/andradecierdo/go-api/models"
	"github.com/google/uuid"
)

type IBlogRepository interface {
	GetBlog(id uuid.UUID) (models.Blog, error)
	GetBlogs() ([]models.Blog, error)
}

type BlogRepository struct {
	db *db.Queries
}

func NewBlogRepository(db *db.Queries) *BlogRepository {
	return &BlogRepository{db}
}
