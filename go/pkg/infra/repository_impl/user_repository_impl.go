package repository_impl

import (
	"context"
	"database/sql"

	"github.com/cawauchi6204/hexagonal-architecture-todo/schemas"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) FindAll(ctx context.Context) (schemas.UserSlice, error) {
	rows, err := schemas.Users().All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
