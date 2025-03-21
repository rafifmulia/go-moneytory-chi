package types

import (
	"context"
	"database/sql"
)

// ref: https://stackoverflow.com/questions/75780240/golang-struct-method-to-take-in-multiple-argument-types
// Supaya di interface repository dapat menerima *sql.Tx atau *sql.DB.
// Hal ini dikarenakan select statement bisa saja dilakukan tanpa perlu membuat *sql.DB.BeginTx,
// dan select statement bisa saja dilakukan bersamaan dengan proses alter data (DML) yang membutuhkan TCL di SQL.

// DB
// Either database handler (connection) or database transaction.
type DB interface {
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	// *sql.DB | *sql.Tx // Type Union. *Cocok jika menggunakan struct generic.
}
