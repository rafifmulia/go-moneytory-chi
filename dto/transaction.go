package dto

import (
	"database/sql"
	"restfulapi/api"
	"restfulapi/helper"
	"restfulapi/model/domain"
)

func WebReqCreateTrxToDomainTrx(from *api.ReqCreateTransaction) *domain.Transaction {
	var (
		d *domain.Transaction = &domain.Transaction{} // Uuid, Ctm, Mtm, Dtm is missing.
	)
	helper.StructCopy(from, d)
	if from.Date != nil && !from.Date.IsZero() {
		d.Date = &sql.NullTime{
			Time:  *from.Date,
			Valid: true,
		}
	}
	return d
}

func WebReqUpdateTrxToDomainTrx(from *api.ReqUpdateTransaction) *domain.Transaction {
	var (
		d *domain.Transaction = &domain.Transaction{} // Uuid, Ctm, Mtm, Dtm is missing.
	)
	helper.StructCopy(from, d)
	if from.Date != nil && !from.Date.IsZero() {
		d.Date = &sql.NullTime{
			Time:  *from.Date,
			Valid: true,
		}
	}
	return d
}

func DomainTrxToApiTrx(from *domain.Transaction) *api.Transaction {
	var (
		d *api.Transaction = &api.Transaction{}
	)
	helper.StructCopy(from, d)
	if from.Date != nil && from.Date.Valid {
		d.Date = &from.Date.Time
	}
	return d
}

func DomainTrxsToApiTrxs(from []*domain.Transaction) []*api.Transaction {
	l := make([]*api.Transaction, 0, len(from))
	for i := range from {
		d := DomainTrxToApiTrx(from[i])
		l = append(l, d)
	}
	return l
}
