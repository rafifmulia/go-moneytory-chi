package types

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Supaya di interface repository dapat menerima *sqlx.Tx atau *sqlx.DB.
// Hal ini dikarenakan select statement bisa saja dilakukan tanpa perlu membuat *sqlx.DB.BeginTx,
// dan select statement bisa saja dilakukan bersamaan dengan proses alter data (DML) yang membutuhkan TCL di SQL.

// DBx
// Either database handler (connection) or database transaction.
type DBx interface {
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	// NamedStmt(stmt *sqlx.NamedStmt) *sqlx.NamedStmt
	// NamedStmtContext(ctx context.Context, stmt *sqlx.NamedStmt) *sqlx.NamedStmt
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	// Stmt(stmt *sql.Stmt) *sql.Stmt
	// StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
	// Stmtx(stmt interface{}) *sqlx.Stmt
	// StmtxContext(ctx context.Context, stmt interface{}) *sqlx.Stmt
	// *sqlx.DB | *sqlx.Tx // Type Union. *Cocok jika menggunakan struct generic.
}
