package helper

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func AutoCommitRollback(tx *sql.Tx) {
	if err0 := tx.Commit(); err0 != nil {
		if err1 := tx.Rollback(); err1 != nil {
			panic(err1)
		} else {
			panic(err0)
		}
	}
}

func AutoCommitRollbackX(tx *sqlx.Tx) {
	if err0 := tx.Commit(); err0 != nil {
		if err1 := tx.Rollback(); err1 != nil {
			panic(err1)
		} else {
			panic(err0)
		}
	}
}
