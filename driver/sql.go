package driver

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	mu   *sync.Mutex = &sync.Mutex{}
	once sync.Once   = sync.Once{}
)

func buildDSN() *string {
	var dsn *string = new(string)
	s := "%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s"
	*dsn = fmt.Sprintf(s, os.Getenv("SQL_USER"), os.Getenv("SQL_PASW"), os.Getenv("SQL_HOST"), os.Getenv("SQL_PORT"), os.Getenv("SQL_DBNAME"), url.QueryEscape(os.Getenv("TZ")))
	return dsn
}

func openConn() {
	var err error
	fmt.Printf("Registered SQL Drivers: %s\n", sql.Drivers())
	dsn := buildDSN()
	db, err = sql.Open("mysql", *dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetConnMaxLifetime(time.Hour * 1)
	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func InitConn() {
	if db == nil {
		once.Do(openConn)
	}
}

func ExportDbHandle() *sql.DB {
	mu.Lock()
	defer mu.Unlock()
	return db
}
