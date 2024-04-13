package controllers

import (
	"encoding/json"
	db "github.com/andradecierdo/go-api/database"
	"github.com/andradecierdo/go-api/models"
	"github.com/andradecierdo/go-api/repositories"
	"github.com/andradecierdo/go-api/services"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(db *db.Queries) *UserController {
	return &UserController{
		UserService: services.NewUserService(repositories.NewUserRepository(db)),
	}
}

// CreateUser TODO add common function for error handling
func (controller UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userInput models.UserInput
	err := decoder.Decode(&userInput)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	user, err := controller.UserService.CreateUser(r.Context(), &userInput)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (controller UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := uuid.MustParse(mux.Vars(r)["id"])

	decoder := json.NewDecoder(r.Body)
	var userInput models.UserInput
	err := decoder.Decode(&userInput)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	resources, err := controller.UserService.UpdateUser(r.Context(), userId, &userInput)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(resources)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (controller UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := uuid.MustParse(mux.Vars(r)["id"])
	user, err := controller.UserService.GetUserById(r.Context(), userId)
	// TODO implement specific error handling
	if err != nil {
		NotFoundHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (controller UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	pagination := models.Pagination{
		Page:  page,
		Limit: limit,
	}
	resources, err := controller.UserService.GetUsers(r.Context(), pagination)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(resources)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (controller UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := uuid.MustParse(mux.Vars(r)["id"])
	err := controller.UserService.DeleteUser(r.Context(), userId)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal("User successfully deleted.")
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
