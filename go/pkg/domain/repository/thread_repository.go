package repository

import (
	"context"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
)

type ThreadRepository interface {
	// スレッドを作成
	Create(ctx context.Context, thread *entity.Thread) error

	// IDでスレッドを取得
	FindByID(ctx context.Context, id string) (*entity.Thread, error)

	// タグでスレッドを取得
	FindByTag(ctx context.Context, tagID string) ([]*entity.Thread, error)

	// ユーザーIDでスレッドを取得
	FindByUserID(ctx context.Context, userID string) ([]*entity.Thread, error)

	// 全スレッドを取得（ページネーション対応）
	FindAll(ctx context.Context, offset, limit int) ([]*entity.Thread, error)

	// スレッド情報を更新
	Update(ctx context.Context, thread *entity.Thread) error

	// スレッドを削除
	Delete(ctx context.Context, id string) error

	// スレッドにタグを追加
	AddTag(ctx context.Context, threadID, tagID string) error

	// スレッドからタグを削除
	RemoveTag(ctx context.Context, threadID, tagID string) error

	// スレッドの投稿を取得
	FindPosts(ctx context.Context, threadID string) ([]*entity.Post, error)

	// スレッドの投稿数を取得
	CountPosts(ctx context.Context, threadID string) (int, error)

	// スレッドのタグを取得
	FindTags(ctx context.Context, threadID string) ([]*entity.Tag, error)

	// 最新のスレッドを取得
	FindLatest(ctx context.Context, limit int) ([]*entity.Thread, error)

	// 人気のスレッドを取得（投稿数やコメント数などに基づく）
	FindPopular(ctx context.Context, limit int) ([]*entity.Thread, error)
}
