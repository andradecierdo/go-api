package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andradecierdo/go-api/data"
	"github.com/andradecierdo/go-api/repositories"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type BlogService struct {
	Repository repositories.BlogRepository
}

func NewBlogService(repo repositories.BlogRepository) *BlogService {
	return &BlogService{
		Repository: repo,
	}
}

func (service BlogService) GetBlogs(w http.ResponseWriter, r *http.Request) {
	resources, err := service.Repository.GetBlogs()

	jsonBytes, err := json.Marshal(resources)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (service BlogService) GetBlog(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// Retrieve recipe from the store
	blog, err := service.Repository.GetBlog(uuid.MustParse(id))
	if err != nil {
		fmt.Println(err)
		// Not found
		if errors.Is(err, data.NotFoundErr) {
			NotFoundHandler(w, r)
			return
		}

		// Every other error
		InternalServerErrorHandler(w, r)
		return
	}

	// Convert the struct into JSON payload
	jsonBytes, err := json.Marshal(blog)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	// Write the results
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("500 Internal Server Error"))
	if err != nil {
		return
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("404 Not Found"))
	if err != nil {
		return
	}
}
