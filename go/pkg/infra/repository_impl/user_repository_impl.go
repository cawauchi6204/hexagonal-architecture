package repository_impl

import (
	"context"
	"database/sql"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/model"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/repository"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/infra/orm_converter"
	"github.com/cawauchi6204/hexagonal-architecture-todo/schemas"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Find(ctx context.Context, id int) (*model.User, error) {
	user, err := schemas.Users().One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return orm_converter.ToModel(user), nil
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context) ([]*model.User, error) {
	rows, err := schemas.Users().All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, 0, len(rows))
	for _, row := range rows {
		users = append(users, orm_converter.ToModel(row))
	}
	return users, nil
}
