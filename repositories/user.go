package repositories

import (
	"context"
	db "github.com/andradecierdo/go-api/database"
	"github.com/andradecierdo/go-api/models"
	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, input models.UserInput) (*models.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, input models.UserInput) (*models.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*models.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetUsers(ctx context.Context, pagination models.Pagination) ([]*models.User, error)
}

type UserRepository struct {
	db *db.Queries
	IUserRepository
}

func NewUserRepository(db *db.Queries) *UserRepository {
	return &UserRepository{db: db}
}

func toUserModel(dbUser db.User) *models.User {
	return &models.User{
		ID:        dbUser.ID,
		FirstName: dbUser.Firstname,
		LastName:  dbUser.Lastname,
		Email:     dbUser.Email,
		IsAdmin:   dbUser.Isadmin,
		CreatedAt: dbUser.Createdat,
	}
}

func (repo *UserRepository) CreateUser(ctx context.Context, input *models.UserInput) (*models.User, error) {
	args := &db.CreateUserParams{
		Firstname: input.FirstName,
		Lastname:  input.LastName,
		Email:     input.Email,
	}

	user, err := repo.db.CreateUser(ctx, *args)
	if err != nil {
		return nil, err
	}

	return toUserModel(user), nil
}

func (repo *UserRepository) UpdateUser(ctx context.Context, id uuid.UUID, input *models.UserInput) (*models.User, error) {
	args := &db.UpdateUserParams{
		ID:        id,
		Firstname: input.FirstName,
		Lastname:  input.LastName,
		Email:     input.Email,
	}

	user, err := repo.db.UpdateUser(ctx, *args)
	if err != nil {
		return nil, err
	}

	return toUserModel(user), nil
}

func (repo *UserRepository) GetUserById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := repo.db.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return toUserModel(user), nil
}

func (repo *UserRepository) GetUsers(ctx context.Context, pagination models.Pagination) ([]*models.User, error) {
	offset := (pagination.Page - 1) * pagination.Limit

	args := &db.ListUsersParams{
		Limit:  int64(pagination.Limit),
		Offset: int64(offset),
	}

	users, err := repo.db.ListUsers(ctx, *args)
	if err != nil {
		return []*models.User{}, err
	}

	userList := make([]*models.User, 0)
	for _, user := range users {
		userList = append(userList, toUserModel(user))
	}

	return userList, nil
}

func (repo *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return repo.db.DeleteUser(ctx, id)
}
