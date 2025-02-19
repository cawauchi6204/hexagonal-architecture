package repository

import (
	"context"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
)

type PostRepository interface {
	// 投稿を作成
	Create(ctx context.Context, post *entity.Post) error

	// IDで投稿を取得
	FindByID(ctx context.Context, id string) (*entity.Post, error)

	// スレッドIDで投稿を取得
	FindByThreadID(ctx context.Context, threadID string) ([]*entity.Post, error)

	// ユーザーIDで投稿を取得
	FindByUserID(ctx context.Context, userID string) ([]*entity.Post, error)

	// 投稿を更新
	Update(ctx context.Context, post *entity.Post) error

	// 投稿を削除
	Delete(ctx context.Context, id string) error

	// 投稿のコメントを取得
	FindComments(ctx context.Context, postID string) ([]*entity.Comment, error)

	// 投稿のコメント数を取得
	CountComments(ctx context.Context, postID string) (int, error)

	// スレッド内の投稿をページネーション付きで取得
	FindByThreadIDWithPagination(ctx context.Context, threadID string, offset, limit int) ([]*entity.Post, error)

	// 最新の投稿を取得
	FindLatest(ctx context.Context, limit int) ([]*entity.Post, error)

	// ユーザーの投稿をページネーション付きで取得
	FindByUserIDWithPagination(ctx context.Context, userID string, offset, limit int) ([]*entity.Post, error)

	// 特定のスレッド内で投稿を検索
	SearchInThread(ctx context.Context, threadID string, query string) ([]*entity.Post, error)
}
