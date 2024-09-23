package infra

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

func ReadDSN() string {
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_ROOT_USER")
	password := os.Getenv("DB_ROOT_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, dbname, port, dbname)
}

// func ReadDSN() string {
// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	dbname := os.Getenv("DB_NAME")

// 	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// }

// InitDB db接続を初期化
func InitDB() *sql.DB {
	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))

	db := sqldblogger.OpenDriver(
		ReadDSN(),
		&mysql.MySQLDriver{},
		loggerAdapter,
	)

	return db
}
