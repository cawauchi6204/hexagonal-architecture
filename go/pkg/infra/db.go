package infra

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/application/core"
	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func ReadDSN() string {
	// Cloud SQL接続情報
	instanceConnectionName := os.Getenv("INSTANCE_CONNECTION_NAME") // project:region:instance
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	var dsn string
	if os.Getenv("ENV") == "local" {
		// ローカル開発環境用DSN
		port := core.MustGetEnv("DB_PORT")
		host := core.MustGetEnv("DB_HOST")
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			dbUser, dbPass, host, port, dbName)
	} else {
		// Cloud SQL用DSN
		dsn = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=true",
			dbUser, dbPass, instanceConnectionName, dbName)
	}

	return dsn
}

// InitDB db接続を初期化
func InitDB() *sql.DB {
	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	boil.DebugMode = true

	db := sqldblogger.OpenDriver(
		ReadDSN(),
		&mysql.MySQLDriver{},
		loggerAdapter,
	)

	return db
}
