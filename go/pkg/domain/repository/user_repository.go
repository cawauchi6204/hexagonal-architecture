package repository

import (
	"context"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/model"
)

type UserRepository interface {
	Find(ctx context.Context, id int) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
}
