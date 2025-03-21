package repository

import (
	"context"
	"restfulapi/api"
	"restfulapi/model/domain"
	"restfulapi/types"

	"github.com/jmoiron/sqlx"
)

type TransactionXRepository interface {
	Save(ctx context.Context, tx *sqlx.Tx, trx *domain.Transaction)
	LookupById(ctx context.Context, tx types.DBx, trxId *string) error
	FindById(ctx context.Context, tx types.DBx, trxId *string) (*domain.Transaction, error)
	FindAll(ctx context.Context, tx types.DBx, filter *api.GetTransactionParams) ([]*domain.Transaction, error)
	Update(ctx context.Context, tx *sqlx.Tx, trx *domain.Transaction)
	SoftDelete(ctx context.Context, tx *sqlx.Tx, trxId *string) error
	HardDelete(ctx context.Context, tx *sqlx.Tx, trxId *string) error
}
