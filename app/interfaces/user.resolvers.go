package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-graph-demo/app/models"
	"net/http"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.DefaultResponse, error) {
	_, err := r.Repositories.Users.CreateUser(ctx, &input)
	if err != nil {
		return &models.DefaultResponse{
			Message: "failed to create user",
			Status:  http.StatusInternalServerError,
		}, err
	}

	return &models.DefaultResponse{
		Message: "successfully created user",
		Status:  http.StatusOK,
	}, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*models.UserResponse, error) {
	results, err := r.Repositories.Users.GetUserByID(ctx, id)
	if err != nil {
		return &models.UserResponse{
			Message: "failed to retrieve user",
			Status:  http.StatusInternalServerError,
		}, err
	}

	return &models.UserResponse{
		Message: "successfully retrieved user",
		Status:  http.StatusOK,
		Data:    []*models.User{results},
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) (*models.UserResponse, error) {
	results, err := r.Repositories.Users.GetUsers(ctx)
	if err != nil {
		return &models.UserResponse{
			Message: "failed to retrieve users",
			Status:  http.StatusInternalServerError,
		}, err
	}

	return &models.UserResponse{
		Message: "successfully retrieved users",
		Status:  http.StatusOK,
		Data:    results,
	}, nil
}
