package orm_converter

import (
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/model"
	"github.com/cawauchi6204/hexagonal-architecture-todo/schemas"
)

func ToModel(user *schemas.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Username,
	}
}
