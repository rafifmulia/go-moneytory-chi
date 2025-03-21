package repository

import (
	"context"
	"database/sql"
	"restfulapi/api"
	"restfulapi/model/domain"
	"restfulapi/types"
)

type TransactionRepository interface {
	Save(ctx context.Context, tx *sql.Tx, trx *domain.Transaction)
	LookupById(ctx context.Context, tx types.DB, trxId *string) error
	FindById(ctx context.Context, tx types.DB, trxId *string) (*domain.Transaction, error)
	FindAll(ctx context.Context, tx types.DB, filter *api.GetTransactionParams) ([]*domain.Transaction, error)
	Update(ctx context.Context, tx *sql.Tx, trx *domain.Transaction)
	SoftDelete(ctx context.Context, tx *sql.Tx, trxId *string) error
	HardDelete(ctx context.Context, tx *sql.Tx, trxId *string) error
}
