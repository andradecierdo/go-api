package main

import (
	"database/sql"
	"fmt"
	"github.com/andradecierdo/go-api/data"
	dbConn "github.com/andradecierdo/go-api/database"
	"github.com/andradecierdo/go-api/database/controllers"
	"github.com/andradecierdo/go-api/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// TODO add separate file for db config
	env, _ := godotenv.Read(".env")
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env["DB_HOST"],
		env["DB_PORT"],
		env["DB_USER"],
		env["DB_PASSWORD"],
		env["DB_NAME"],
	)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	db := dbConn.New(conn)

	health := healthCheckHandler{}
	router.HandleFunc("/health", health.HealthCheck)

	// Users
	userController := controllers.NewUserController(db)
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userController.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	// Blogs
	// TODO refactor blog service
	store := data.NewBlogMemStore()
	blogService := services.NewBlogService(store)
	router.HandleFunc("/blogs", blogService.GetBlogs).Methods("GET")
	router.HandleFunc("/blogs/{id}", blogService.GetBlog).Methods("GET")

	// Run the server
	serverErr := http.ListenAndServe(":8080", router)
	if serverErr != nil {
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
