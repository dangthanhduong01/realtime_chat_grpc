package db

import (
	"context"
	"snowApp/internal/model"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (model.User, error)
	GetUserByID(ctx context.Context, id string) (model.User, error)
}

var _ Querier = (*Queries)(nil)
