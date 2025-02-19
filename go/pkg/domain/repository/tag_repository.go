package repository

import (
	"context"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/entity"
)

type TagRepository interface {
	// タグを作成
	Create(ctx context.Context, tag *entity.Tag) error

	// IDでタグを取得
	FindByID(ctx context.Context, id string) (*entity.Tag, error)

	// 名前でタグを取得
	FindByName(ctx context.Context, name string) (*entity.Tag, error)

	// 全タグを取得
	FindAll(ctx context.Context) ([]*entity.Tag, error)

	// タグを更新
	Update(ctx context.Context, tag *entity.Tag) error

	// タグを削除
	Delete(ctx context.Context, id string) error

	// タグに関連するスレッドを取得
	FindThreads(ctx context.Context, tagID string) ([]*entity.Thread, error)

	// タグのスレッド数を取得
	CountThreads(ctx context.Context, tagID string) (int, error)

	// 人気のタグを取得（使用頻度順）
	FindPopular(ctx context.Context, limit int) ([]*entity.Tag, error)

	// タグ名で部分一致検索
	SearchByName(ctx context.Context, query string) ([]*entity.Tag, error)

	// ユーザーが使用したタグを取得
	FindByUserID(ctx context.Context, userID string) ([]*entity.Tag, error)

	// スレッドに関連するタグを取得
	FindByThreadID(ctx context.Context, threadID string) ([]*entity.Tag, error)
}
