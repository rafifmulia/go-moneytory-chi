package repository

import (
	"context"
	"restfulapi/api"
	"restfulapi/exception"
	"restfulapi/model/domain"
	"restfulapi/types"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

type TransactionXRepositoryMock struct {
	Mock *mock.Mock
}

func NewTransactionXRepositoryMock(mock *mock.Mock) *TransactionXRepositoryMock {
	return &TransactionXRepositoryMock{Mock: mock}
}

func (impl *TransactionXRepositoryMock) Save(ctx context.Context, tx *sqlx.Tx, trx *domain.Transaction) {
	args := impl.Mock.Called(ctx, tx, trx)
	d := args.Get(0)
	if d != nil {
		panic(d)
	}
}

func (impl *TransactionXRepositoryMock) LookupById(ctx context.Context, tx types.DBx, trxId *string) error {
	args := impl.Mock.Called(trxId)
	d := args.Get(0)
	if d != nil {
		return d.(error)
	}
	return nil
}

func (impl *TransactionXRepositoryMock) FindById(ctx context.Context, tx types.DBx, trxId *string) (*domain.Transaction, error) {
	args := impl.Mock.Called(trxId)
	err := args.Get(1)
	if err != nil {
		return nil, err.(error)
	}
	d := args.Get(0)
	if d == nil {
		return nil, exception.NewNotFoundException("transaction is not found")
	}
	return d.(*domain.Transaction), nil
}

func (impl *TransactionXRepositoryMock) FindAll(ctx context.Context, tx types.DBx, filter *api.GetTransactionParams) ([]*domain.Transaction, error) {
	args := impl.Mock.Called(filter)
	err := args.Get(1)
	if err != nil {
		return nil, err.(error)
	}
	d := args.Get(0)
	if d == nil {
		return nil, exception.NewNotFoundException("transaction is empty")
	}
	return d.([]*domain.Transaction), nil
}

func (impl *TransactionXRepositoryMock) Update(ctx context.Context, tx *sqlx.Tx, trx *domain.Transaction) {
	args := impl.Mock.Called(trx)
	d := args.Get(0)
	if d != nil {
		panic(d)
	}
}

func (impl *TransactionXRepositoryMock) SoftDelete(ctx context.Context, tx *sqlx.Tx, trxId *string) error {
	args := impl.Mock.Called(trxId)
	d := args.Get(0)
	if d != nil {
		return d.(error)
	}
	return nil
}

func (impl *TransactionXRepositoryMock) HardDelete(ctx context.Context, tx *sqlx.Tx, trxId *string) error {
	args := impl.Mock.Called(trxId)
	d := args.Get(0)
	if d != nil {
		return d.(error)
	}
	return nil
}
