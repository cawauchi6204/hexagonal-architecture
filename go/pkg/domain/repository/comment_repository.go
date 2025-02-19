package repository

import (
	"context"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
)

type CommentRepository interface {
	// コメントを作成
	Create(ctx context.Context, comment *entity.Comment) error

	// IDでコメントを取得
	FindByID(ctx context.Context, id string) (*entity.Comment, error)

	// 投稿IDでコメントを取得
	FindByPostID(ctx context.Context, postID string) ([]*entity.Comment, error)

	// ユーザーIDでコメントを取得
	FindByUserID(ctx context.Context, userID string) ([]*entity.Comment, error)

	// コメントを更新
	Update(ctx context.Context, comment *entity.Comment) error

	// コメントを削除
	Delete(ctx context.Context, id string) error

	// コメントにいいねを追加
	AddLike(ctx context.Context, commentID, userID string) error

	// コメントのいいねを削除
	RemoveLike(ctx context.Context, commentID, userID string) error

	// コメントのいいね数を取得
	CountLikes(ctx context.Context, commentID string) (int, error)

	// ユーザーがいいねしたコメントを取得
	FindLikedByUser(ctx context.Context, userID string) ([]*entity.Comment, error)

	// 投稿のコメントをページネーション付きで取得
	FindByPostIDWithPagination(ctx context.Context, postID string, offset, limit int) ([]*entity.Comment, error)

	// 最新のコメントを取得
	FindLatest(ctx context.Context, limit int) ([]*entity.Comment, error)
}
