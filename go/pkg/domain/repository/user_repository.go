package repository

import (
	"context"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
)

type UserRepository interface {
	// ユーザーを作成
	Create(ctx context.Context, user *entity.User) error

	// IDでユーザーを取得
	FindByID(ctx context.Context, id string) (*entity.User, error)

	// メールアドレスでユーザーを取得
	FindByEmail(ctx context.Context, email string) (*entity.User, error)

	// ユーザー名でユーザーを取得
	FindByUsername(ctx context.Context, username string) (*entity.User, error)

	// 全ユーザーを取得
	FindAll(ctx context.Context) ([]*entity.User, error)

	// ユーザー情報を更新
	Update(ctx context.Context, user *entity.User) error

	// ユーザーを削除
	Delete(ctx context.Context, id string) error

	// フォロワーを取得
	FindFollowers(ctx context.Context, userID string) ([]*entity.User, error)

	// フォロー中のユーザーを取得
	FindFollowing(ctx context.Context, userID string) ([]*entity.User, error)

	// フォロー関係を作成
	CreateFollow(ctx context.Context, followerID, followedID string) error

	// フォロー関係を削除
	DeleteFollow(ctx context.Context, followerID, followedID string) error
}
