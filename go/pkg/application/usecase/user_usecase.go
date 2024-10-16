package usecase

import (
	"context"
	"fmt"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/model"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/repository"
)

type UserUsecase struct {
	UserRepository repository.UserRepository
}

func (u *UserUsecase) List(ctx context.Context) ([]*model.User, error) {
	fmt.Println("ここまできた1")
	users, err := u.UserRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("ここまできた2")
	fmt.Printf("usersは: %v\n", users)
	return users, nil
}
