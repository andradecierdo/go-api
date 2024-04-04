package data

import (
	"errors"
	"github.com/andradecierdo/go-api/models"
	"github.com/google/uuid"
	"time"
)

var (
	NotFoundErr = errors.New("not found")
)

type BlogMemStore struct {
	list map[uuid.UUID]models.Blog
}

func NewBlogMemStore() *BlogMemStore {
	// Initial Data
	var blogId = uuid.New()
	var userId = uuid.New()
	list := map[uuid.UUID]models.Blog{
		blogId: {
			ID:      blogId,
			UserId:  userId,
			Title:   "Blog One",
			Content: "This is my first blog",
			Date:    time.Now(),
			Media: []models.Media{
				{
					ID:       uuid.New(),
					BlogID:   blogId,
					Path:     "/media/photo/test.png",
					Source:   "s3",
					Name:     "Mt Ruapehu",
					Type:     models.Photo,
					Uploaded: time.Now(),
				},
				{
					ID:       uuid.New(),
					BlogID:   blogId,
					Path:     "/media/photo/vid.mp4",
					Source:   "s3",
					Name:     "Mt Cook",
					Type:     models.Video,
					Uploaded: time.Now(),
				},
			},
		},
	}
	return &BlogMemStore{
		list,
	}
}

func (m BlogMemStore) GetBlog(id uuid.UUID) (models.Blog, error) {
	if val, ok := m.list[id]; ok {
		return val, nil
	}

	return models.Blog{}, NotFoundErr
}

func (m BlogMemStore) GetBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	for _, value := range m.list {
		blogs = append(blogs, value)
	}

	return blogs, nil
}
