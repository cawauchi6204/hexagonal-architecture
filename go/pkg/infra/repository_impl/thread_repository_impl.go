package repository_impl

import (
	"context"
	"database/sql"

	"github.com/cawauchi6204/hexagonal-architecture-todo/schemas"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ThreadRepositoryImpl struct {
	db *sql.DB
}

func NewThreadRepositoryImpl(db *sql.DB) *ThreadRepositoryImpl {
	return &ThreadRepositoryImpl{db: db}
}

// スレッドにあるコメントを取得する
func (impl *ThreadRepositoryImpl) FindAllCommentsInThread(ctx context.Context, threadID string) (*schemas.Thread, error) {
	mods := []qm.QueryMod{
		qm.Load(schemas.ThreadRels.Posts, qm.Where(schemas.ThreadTableColumns.ID+" = ?", threadID)),
	}
	row, err := schemas.Threads(mods...).One(ctx, impl.db)
	if err != nil {
		return nil, err
	}
	return row, nil
}
