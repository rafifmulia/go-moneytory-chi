package domain

import (
	"database/sql"
	"time"
)

type Transaction struct {
	Uuid        *string       `db:"uuid"`
	Description *string       `db:"description"`
	Amount      *float64      `db:"amount"`
	Ctm         *time.Time    `db:"ctm"`
	Mtm         *time.Time    `db:"mtm"`
	Date        *sql.NullTime `db:"date"`
	Dtm         *sql.NullTime `db:"dtm"`
}
