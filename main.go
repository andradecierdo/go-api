package main

import (
	"github.com/andradecierdo/go-api/data"
	"github.com/andradecierdo/go-api/services"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	health := healthCheckHandler{}
	router.HandleFunc("/health", health.HealthCheck)

	store := data.NewBlogMemStore()
	blogService := services.NewBlogService(store)

	router.HandleFunc("/blogs", blogService.GetBlogs).Methods("GET")
	router.HandleFunc("/blogs/{id}", blogService.GetBlog).Methods("GET")

	// Run the server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}

type healthCheckHandler struct{}

func (h *healthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Health Check"))
	if err != nil {
		return
	}
}
