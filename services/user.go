package services

import (
	"context"
	"github.com/andradecierdo/go-api/models"
	"github.com/andradecierdo/go-api/repositories"
	"github.com/andradecierdo/go-api/utils"
	"github.com/google/uuid"
)

type UserService struct {
	Repository *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		Repository: repo,
	}
}

func (service UserService) CreateUser(ctx context.Context, userInput *models.UserInput) (*models.User, error) {
	// TODO add input validation
	return service.Repository.CreateUser(ctx, userInput)
}

func (service UserService) GetUserById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return service.Repository.GetUserById(ctx, id)
}

func (service UserService) UpdateUser(ctx context.Context, id uuid.UUID, input *models.UserInput) (*models.User, error) {
	// TODO add input validation
	return service.Repository.UpdateUser(ctx, id, input)
}

func (service UserService) GetUsers(ctx context.Context, pagination models.Pagination) ([]*models.User, error) {
	pagination.Page = utils.Page(pagination.Page, 1)
	pagination.Limit = utils.Page(pagination.Limit, 10)

	return service.Repository.GetUsers(ctx, pagination)
}

func (service UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	// TODO add validation
	return service.Repository.DeleteUser(ctx, id)
}
