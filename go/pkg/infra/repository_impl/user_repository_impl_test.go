package repository_impl_test

import (
	"database/sql"
	"testing"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/infra/repository_impl"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	db, err := sql.Open("mysql", "root:rootpassword@tcp(localhost:33333)/test")
	if err != nil {
		t.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("./user_repository_impl_test"),
	)
	if err != nil {
		t.Fatalf("Could not create fixtures: %v", err)
	}

	err = fixtures.Load()
	if err != nil {
		t.Fatalf("Could not load fixtures: %v", err)
	}

	// テストケースを追加
	t.Run("ユーザーの取得テスト", func(t *testing.T) {
		repo := repository_impl.NewUserRepository(db)
		user, err := repo.GetByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "テストユーザー", user.Name)
	})

	t.Run("存在しないユーザーの取得テスト", func(t *testing.T) {
		repo := repository_impl.NewUserRepository(db)
		user, err := repo.GetByID(999)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
