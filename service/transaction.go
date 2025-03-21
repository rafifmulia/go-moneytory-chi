package service

import (
	"context"
	"restfulapi/api"
)

type TransactionService interface {
	Create(ctx context.Context, req *api.ReqCreateTransaction) *api.Transaction
	FindById(ctx context.Context, trxId *string) (*api.Transaction, error)
	FindAll(ctx context.Context, params *api.GetTransactionParams) ([]*api.Transaction, error)
	Update(ctx context.Context, req *api.ReqUpdateTransaction, trxId *string) (*api.Transaction, error)
	Delete(ctx context.Context, trxId *string) error
}
