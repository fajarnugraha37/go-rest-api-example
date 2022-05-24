package database

import (
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	var err error
	if DB != nil {
		err = DB.Ping()
		if err == nil {
			return
		}
	}

	DB, err = sqlx.Open("mysql", "root:root@/app")
	if err != nil {
		panic(err)
	}

	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	DB.SetMaxOpenConns(maxConn)
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	DB.SetMaxIdleConns(maxIdleConn)
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
	DB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))
}
