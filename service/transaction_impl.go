package service

import (
	"context"
	"restfulapi/api"
	"restfulapi/driver"
	"restfulapi/dto"
	"restfulapi/helper"
	"restfulapi/model/domain"
	"restfulapi/repository"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TransactionServiceImpl struct {
	repo repository.TransactionXRepository
	db   *sqlx.DB
}

func NewTransactionServiceImpl() TransactionService {
	var (
		db   *sqlx.DB                          = driver.ExportDbHandleX()
		repo repository.TransactionXRepository = repository.NewTransactionXRepositoryImpl()
	)
	return &TransactionServiceImpl{db: db, repo: repo}
}

func (impl *TransactionServiceImpl) Create(ctx context.Context, req *api.ReqCreateTransaction) *api.Transaction {
	var (
		now     time.Time = time.Now()
		tx      *sqlx.Tx
		err     error
		uuidv7  uuid.UUID
		trx     *domain.Transaction
		trxResp *api.Transaction
	)
	tx, err = impl.db.Beginx()
	if err != nil {
		panic(err)
	}
	defer helper.AutoCommitRollbackX(tx)
	trx = dto.WebReqCreateTrxToDomainTrx(req)
	uuidv7, err = uuid.NewV7()
	if err != nil {
		panic(err)
	}
	trx.Uuid = helper.InlinePointer(uuidv7.String())
	trx.Ctm = &now
	trx.Mtm = &now
	impl.repo.Save(ctx, tx, trx)
	trxResp = dto.DomainTrxToApiTrx(trx)
	return trxResp
}

func (impl *TransactionServiceImpl) FindById(ctx context.Context, trxId *string) (*api.Transaction, error) {
	var (
		err     error
		trx     *domain.Transaction
		trxResp *api.Transaction
	)
	trx, err = impl.repo.FindById(ctx, impl.db, trxId)
	if err != nil {
		return nil, err
	}
	trxResp = dto.DomainTrxToApiTrx(trx)
	return trxResp, nil
}

func (impl *TransactionServiceImpl) FindAll(ctx context.Context, params *api.GetTransactionParams) ([]*api.Transaction, error) {
	var (
		err  error
		trxs []*domain.Transaction
		trxl []*api.Transaction
	)
	trxs, err = impl.repo.FindAll(ctx, impl.db, params)
	if err != nil {
		return nil, err
	}
	trxl = dto.DomainTrxsToApiTrxs(trxs)
	return trxl, nil
}

func (impl *TransactionServiceImpl) Update(ctx context.Context, req *api.ReqUpdateTransaction, trxId *string) (*api.Transaction, error) {
	var (
		now     time.Time = time.Now()
		tx      *sqlx.Tx
		err     error
		trx     *domain.Transaction
		trxResp *api.Transaction
	)
	tx, err = impl.db.Beginx()
	if err != nil {
		panic(err)
	}
	defer helper.AutoCommitRollbackX(tx)
	err = impl.repo.LookupById(ctx, tx, trxId)
	if err != nil {
		return nil, err
	}
	trx = dto.WebReqUpdateTrxToDomainTrx(req)
	trx.Uuid = trxId
	trx.Mtm = &now
	impl.repo.Update(ctx, tx, trx)
	trx, err = impl.repo.FindById(ctx, tx, trxId)
	if err != nil {
		return nil, err
	}
	trxResp = dto.DomainTrxToApiTrx(trx)
	return trxResp, nil
}

func (impl *TransactionServiceImpl) Delete(ctx context.Context, trxId *string) error {
	var (
		tx  *sqlx.Tx
		err error
	)
	tx, err = impl.db.Beginx()
	if err != nil {
		panic(err)
	}
	defer helper.AutoCommitRollbackX(tx)
	err = impl.repo.SoftDelete(ctx, tx, trxId)
	if err != nil {
		return err
	}
	return nil
}
