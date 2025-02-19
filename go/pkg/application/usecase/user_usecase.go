package usecase

import (
	"context"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/model"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/repository"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/infra/orm_converter"
)

type UserUsecase struct {
	UserRepository repository.UserRepository
}

func (u *UserUsecase) List(ctx context.Context) ([]*model.User, error) {
	users, err := u.UserRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	users := orm_converter.ToModel(users)
	return users, nil
}
