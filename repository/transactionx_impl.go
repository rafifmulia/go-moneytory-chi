package repository

import (
	"context"
	"fmt"
	"restfulapi/api"
	"restfulapi/exception"
	"restfulapi/helper"
	"restfulapi/model/domain"
	"restfulapi/types"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type TransactionXRepositoryImpl struct{}

func NewTransactionXRepositoryImpl() TransactionXRepository {
	return &TransactionXRepositoryImpl{}
}

func (impl *TransactionXRepositoryImpl) Save(ctx context.Context, tx *sqlx.Tx, trx *domain.Transaction) {
	rs, err := tx.NamedExecContext(ctx, "insert into transaction (uuid, description, date, amount, ctm, mtm) values (:uuid, :description, :date, :amount, :ctm, :mtm);", trx)
	if err != nil {
		panic(err)
	}
	i, err := rs.RowsAffected()
	if err != nil {
		panic(err)
	}
	if i < 1 {
		panic("failed save transaction")
	}
}

func (impl *TransactionXRepositoryImpl) LookupById(ctx context.Context, tx types.DBx, trxId *string) error {
	rs := tx.QueryRowxContext(ctx, "select uuid, mtm from transaction where uuid = ? and dtm is null limit 1;", trxId)
	d := &domain.Transaction{}
	if err := rs.StructScan(d); err != nil {
		return exception.NewNotFoundException("transaction is not found")
	}
	return nil
}

func (impl *TransactionXRepositoryImpl) FindById(ctx context.Context, tx types.DBx, trxId *string) (*domain.Transaction, error) {
	rs := tx.QueryRowxContext(ctx, "select uuid, description, date, amount from transaction where uuid = ? and dtm is null limit 1;", trxId)
	d := &domain.Transaction{}
	if err := rs.StructScan(d); err != nil {
		return d, exception.NewNotFoundException("transaction is not found")
	}
	return d, nil
}

func (impl *TransactionXRepositoryImpl) FindAll(ctx context.Context, tx types.DBx, filter *api.GetTransactionParams) ([]*domain.Transaction, error) {
	var (
		rows *sqlx.Rows
		err  error
	)
	rows, err = tx.QueryxContext(ctx, "select uuid, description, date, amount from transaction where date between ? and ? and dtm is null order by date desc limit 100;", filter.RangeStart, filter.RangeEnd)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	items := make([]*domain.Transaction, 0, 100)
	for rows.Next() {
		d := &domain.Transaction{}
		if err := rows.StructScan(d); err != nil {
			panic(err)
		}
		items = append(items, d)
	}
	if len(items) == 0 {
		return items, exception.NewNotFoundException("Transaction is empty")
	}
	return items, nil
}

func (impl *TransactionXRepositoryImpl) Update(ctx context.Context, tx *sqlx.Tx, trx *domain.Transaction) {
	var (
		w   strings.Builder = strings.Builder{}
		err error
	)
	w.Grow(111)
	w.WriteString("update transaction set")
	helper.SqlXColumnsUpdate(&w, trx, []string{"uuid"})
	w.WriteString(" where uuid = :uuid;")
	rs, err := tx.NamedExecContext(ctx, w.String(), trx)
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

func (impl *TransactionXRepositoryImpl) SoftDelete(ctx context.Context, tx *sqlx.Tx, trxId *string) error {
	rs, err := tx.ExecContext(ctx, "update transaction set dtm = ? where uuid = ?;", time.Now().Format(time.RFC3339), trxId)
	if err != nil {
		panic(err)
	}
	i, err := rs.RowsAffected()
	if err != nil {
		panic(err)
	}
	if i < 1 {
		return exception.NewNotFoundException(fmt.Sprintf("failed soft delete transaction, trxid %s not found", *trxId))
	}
	return nil
}

func (impl *TransactionXRepositoryImpl) HardDelete(ctx context.Context, tx *sqlx.Tx, trxId *string) error {
	rs, err := tx.ExecContext(ctx, "delete from transaction where uuid = ?;", trxId)
	if err != nil {
		panic(err)
	}
	i, err := rs.RowsAffected()
	if err != nil {
		panic(err)
	}
	if i < 1 {
		return exception.NewNotFoundException(fmt.Sprintf("failed hard delete transaction, trxid %s not found", *trxId))
	}
	return nil
}
