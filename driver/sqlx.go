package driver

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	dbX   *sqlx.DB
	muX   *sync.Mutex = &sync.Mutex{}
	onceX sync.Once   = sync.Once{}
)

func openConnX() {
	var err error
	fmt.Printf("Registered SQL Drivers: %s\n", sql.Drivers())
	dsn := buildDSN()
	dbX, err = sqlx.Open("mysql", *dsn)
	if err != nil {
		panic(err)
	}
	dbX.SetMaxIdleConns(10)
	dbX.SetMaxOpenConns(10)
	dbX.SetConnMaxIdleTime(time.Minute * 3)
	dbX.SetConnMaxLifetime(time.Hour * 1)
	if err = dbX.Ping(); err != nil {
		panic(err)
	}
}

func InitConnX() {
	if dbX == nil {
		onceX.Do(openConnX)
	}
}

func ExportDbHandleX() *sqlx.DB {
	muX.Lock()
	defer muX.Unlock()
	return dbX
}
