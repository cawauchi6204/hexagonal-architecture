package usecase

import (
	"context"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/model"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/repository"
)

type UserUsecase struct {
	UserRepository repository.UserRepository
}

func (u *UserUsecase) List(ctx context.Context) ([]*model.User, error) {
	users, err := u.UserRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
