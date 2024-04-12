package services

import (
	"encoding/json"
	"github.com/andradecierdo/go-api/models"
	"github.com/andradecierdo/go-api/repositories"
	"github.com/andradecierdo/go-api/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserService struct {
	Repository *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		Repository: repo,
	}
}

// TODO enhance error handling
func (service UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userInput models.UserInput
	err := decoder.Decode(&userInput)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	resources, err := service.Repository.CreateUser(&userInput, r.Context())
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

func (service UserService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := uuid.MustParse(mux.Vars(r)["id"])

	decoder := json.NewDecoder(r.Body)
	var userInput models.UserInput
	err := decoder.Decode(&userInput)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	resources, err := service.Repository.UpdateUser(userId, &userInput, r.Context())
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

func (service UserService) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := uuid.MustParse(mux.Vars(r)["id"])
	resources, err := service.Repository.GetUserById(userId, r.Context())
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

func (service UserService) GetUsers(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	pagination := models.Pagination{
		Page:  utils.Page(page, 1),
		Limit: utils.Page(limit, 10),
	}
	resources, err := service.Repository.GetUsers(pagination, r.Context())
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

func (service UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := uuid.MustParse(mux.Vars(r)["id"])
	err := service.Repository.DeleteUser(userId, r.Context())
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
