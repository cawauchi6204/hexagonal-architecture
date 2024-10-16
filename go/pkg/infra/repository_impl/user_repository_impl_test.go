package repository_impl_test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
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
}
