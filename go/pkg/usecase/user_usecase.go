package usecase

import (
	"context"
	"time"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(ctx context.Context, username, email, password string) (*entity.User, error)
	Login(ctx context.Context, email, password string) (*entity.User, error)
	UpdateProfile(ctx context.Context, userID, username, email string) (*entity.User, error)
	Follow(ctx context.Context, followerID, followedID string) error
	Unfollow(ctx context.Context, followerID, followedID string) error
	GetProfile(ctx context.Context, userID string) (*entity.User, error)
	SearchUsers(ctx context.Context, query string) ([]*entity.User, error)
	GetFollowers(ctx context.Context, userID string) ([]*entity.User, error)
	GetFollowing(ctx context.Context, userID string) ([]*entity.User, error)
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) Register(ctx context.Context, username, email, password string) (*entity.User, error) {
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) Login(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) UpdateProfile(ctx context.Context, userID, username, email string) (*entity.User, error) {
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	user.Username = username
	user.Email = email

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) Follow(ctx context.Context, followerID, followedID string) error {
	return uc.userRepo.CreateFollow(ctx, followerID, followedID)
}

func (uc *userUseCase) Unfollow(ctx context.Context, followerID, followedID string) error {
	return uc.userRepo.DeleteFollow(ctx, followerID, followedID)
}

func (uc *userUseCase) GetProfile(ctx context.Context, userID string) (*entity.User, error) {
	return uc.userRepo.FindByID(ctx, userID)
}

func (uc *userUseCase) SearchUsers(ctx context.Context, query string) ([]*entity.User, error) {
	// TODO: 実装する検索ロジック
	return nil, nil
}

func (uc *userUseCase) GetFollowers(ctx context.Context, userID string) ([]*entity.User, error) {
	return uc.userRepo.FindFollowers(ctx, userID)
}

func (uc *userUseCase) GetFollowing(ctx context.Context, userID string) ([]*entity.User, error) {
	return uc.userRepo.FindFollowing(ctx, userID)
}

func (uc *userUseCase) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	return uc.userRepo.FindAll(ctx)
}
