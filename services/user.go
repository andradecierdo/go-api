package services

import (
	"encoding/json"
	"github.com/andradecierdo/go-api/models"
	"github.com/andradecierdo/go-api/repositories"
	"net/http"
)

type UserService struct {
	Repository *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		Repository: repo,
	}
}

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
