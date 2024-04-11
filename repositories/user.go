package repositories

import (
	"context"
	db "github.com/andradecierdo/go-api/database"
	"github.com/andradecierdo/go-api/models"
)

type UserRepository struct {
	db *db.Queries
}

func NewUserRepository(db *db.Queries) *UserRepository {
	return &UserRepository{db}
}

func toUserModel(dbUser db.User) models.User {
	return models.User{
		ID:        dbUser.ID,
		FirstName: dbUser.Firstname,
		LastName:  dbUser.Lastname,
		Email:     dbUser.Email,
		CreatedAt: dbUser.Createdat,
	}
}

func (repo *UserRepository) CreateUser(input *models.UserInput, ctx context.Context) (models.User, error) {
	args := &db.CreateUserParams{
		Firstname: input.FirstName,
		Lastname:  input.LastName,
		Email:     input.Email,
	}

	user, err := repo.db.CreateUser(ctx, *args)
	if err != nil {
		return models.User{}, err
	}

	return toUserModel(user), nil
}
