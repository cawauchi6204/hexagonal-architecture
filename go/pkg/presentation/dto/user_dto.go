package dto

import (
	"time"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
)

// リクエスト構造体
type RegisterUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateProfileRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

// レスポンス構造体
type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
}

// エンティティからDTOへの変換
func NewUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func NewUserListResponse(users []*entity.User) *UserListResponse {
	response := &UserListResponse{
		Users: make([]UserResponse, len(users)),
	}
	for i, user := range users {
		response.Users[i] = *NewUserResponse(user)
	}
	return response
}

// フォロー関連
type FollowRequest struct {
	FollowedID string `json:"followed_id" validate:"required"`
}

type FollowResponse struct {
	Success bool `json:"success"`
}

type FollowListResponse struct {
	Users []UserResponse `json:"users"`
}
