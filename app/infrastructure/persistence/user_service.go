package persistence

import (
	"context"

	"go-graph-demo/app/models"
)

type UserService interface {
	CreateUser(ctx context.Context, input *models.NewUser) (model *models.User, err error)
	GetUserByID(ctx context.Context, input int) (model *models.User, err error)
	GetUsers(ctx context.Context) (model []*models.User, err error)
}
