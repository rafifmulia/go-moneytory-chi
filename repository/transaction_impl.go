package repository

import (
	"context"
	"database/sql"
	"fmt"
	"restfulapi/api"
	"restfulapi/exception"
	"restfulapi/helper"
	"restfulapi/model/domain"
	"restfulapi/types"
	"strings"
	"time"
)

type TransactionRepositoryImpl struct{}

func NewTransactionRepositoryImpl() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (impl *TransactionRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, trx *domain.Transaction) {
	rs, err := tx.ExecContext(ctx, "insert into transaction (uuid, description, date, amount, ctm, mtm) values (?, ?, ?, ?, ?, ?);", trx.Uuid, trx.Description, trx.Date.Time.Format(time.RFC3339), trx.Amount, trx.Ctm, trx.Mtm)
	if err != nil {
		panic(err)
	}
	i, err := rs.RowsAffected()
	if err != nil {
		panic(err)
	}
	if i < 1 {
		panic("failed save transaction with")
	}
}

func (impl *TransactionRepositoryImpl) LookupById(ctx context.Context, tx types.DB, trxId *string) error {
	rs := tx.QueryRowContext(ctx, "select uuid, mtm from transaction where uuid = ? and dtm is null limit 1;", trxId)
	d := &domain.Transaction{}
	if err := rs.Scan(&d.Uuid, &d.Date); err != nil {
		return exception.NewNotFoundException("transaction is not found")
	}
	return nil
}

func (impl *TransactionRepositoryImpl) FindById(ctx context.Context, tx types.DB, trxId *string) (*domain.Transaction, error) {
	rs := tx.QueryRowContext(ctx, "select uuid, description, date, amount from transaction where uuid = ? and dtm is null limit 1;", trxId)
	d := &domain.Transaction{}
	if err := rs.Scan(&d.Uuid, &d.Description, &d.Date, &d.Amount); err != nil {
		return d, exception.NewNotFoundException("transaction is not found")
	}
	return d, nil
}

func (impl *TransactionRepositoryImpl) FindAll(ctx context.Context, tx types.DB, filter *api.GetTransactionParams) ([]*domain.Transaction, error) {
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = tx.QueryContext(ctx, "select uuid, description, date, amount from transaction where date between ? and ? and dtm is null order by date desc limit 100;", filter.RangeStart, filter.RangeEnd)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	items := make([]*domain.Transaction, 0, 100)
	for rows.Next() {
		d := &domain.Transaction{}
		if err := rows.Scan(&d.Uuid, &d.Description, &d.Date, &d.Amount); err != nil {
			panic(err)
		}
		items = append(items, d)
	}
	if len(items) == 0 {
		return items, exception.NewNotFoundException("transaction is empty")
	}
	return items, nil
}

func (impl *TransactionRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, trx *domain.Transaction) {
	var (
		w        strings.Builder = strings.Builder{}
		bindArgs []any           = make([]any, 0, 6)
		err      error
	)
	w.Grow(85)
	w.WriteString("update transaction set")
	helper.SqlColumnsUpdate(&w, trx, []string{"uuid"}, &bindArgs)
	w.WriteString(" where uuid = ?;")
	bindArgs = append(bindArgs, trx.Uuid)
	rs, err := tx.ExecContext(ctx, w.String(), bindArgs...)
	if err != nil {
		panic(err)
	}
	i, err := rs.RowsAffected()
	if err != nil {
		panic(err)
	}
	if i < 1 {
		panic(fmt.Sprintf("failed update transaction with id %s", *trx.Uuid))
	}
}

func (impl *TransactionRepositoryImpl) SoftDelete(ctx context.Context, tx *sql.Tx, trxId *string) error {
	rs, err := tx.ExecContext(ctx, "update transaction set dtm = ? where uuid = ?;", time.Now().Format(time.RFC3339), trxId)
	if err != nil {
		panic(err)
	}
	i, err := rs.RowsAffected()
	if err != nil {
		panic(err)
	}
	if i < 1 {
		return fmt.Errorf("failed soft delete transaction with id %s", *trxId)
	}
	return nil
}

func (impl *TransactionRepositoryImpl) HardDelete(ctx context.Context, tx *sql.Tx, trxId *string) error {
	rs, err := tx.ExecContext(ctx, "delete from transaction where uuid = ?;", trxId)
	if err != nil {
		panic(err)
	}
	i, err := rs.RowsAffected()
	if err != nil {
		panic(err)
	}
	if i < 1 {
		return fmt.Errorf("failed hard delete transaction with id %s", *trxId)
	}
	return nil
}
