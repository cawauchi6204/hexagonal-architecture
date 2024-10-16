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
	port := core.MustGetEnv("DB_PORT")
	user := core.MustGetEnv("DB_ROOT_USER")
	password := core.MustGetEnv("DB_ROOT_PASSWORD")
	dbname := core.MustGetEnv("DB_DATABASE")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, dbname, port, dbname)
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
