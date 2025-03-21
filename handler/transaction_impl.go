package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"restfulapi/api"
	"restfulapi/exception"
	"restfulapi/helper"
	"restfulapi/libs"
	"restfulapi/service"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
)

type TransactionHandlerImpl struct {
	svc      service.TransactionService
	validate *validator.Validate
}

func NewTransactionHandler() TransactionHandler {
	var (
		validate *validator.Validate        = libs.ExportValidator()
		svc      service.TransactionService = service.NewTransactionServiceImpl()
	)
	return &TransactionHandlerImpl{svc: svc, validate: validate}
}

func (impl *TransactionHandlerImpl) ListTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		now         time.Time       = time.Now()
		rctx        context.Context = r.Context()
		err         error
		queryParams url.Values                = r.URL.Query()
		trxParams   *api.GetTransactionParams = &api.GetTransactionParams{}
		trxResp     []*api.Transaction
	)
	w.Header().Set("Content-Type", "application/json")
	if err := form.NewDecoder().Decode(trxParams, queryParams); err != nil {
		panic(err)
	}
	switch trxParams.Filter {
	case "today":
		trxParams.RangeStart = helper.StartOfDay(&now).Format(time.RFC3339)
		trxParams.RangeEnd = helper.EndOfDay(&now).Format(time.RFC3339)
	case "week":
		trxParams.RangeStart = helper.StartOfWeek(&now).Format(time.RFC3339)
		trxParams.RangeEnd = helper.EndOfWeek(&now).Format(time.RFC3339)
	case "year":
		trxParams.RangeStart = helper.StartOfYear(&now).Format(time.RFC3339)
		trxParams.RangeEnd = helper.EndOfYear(&now).Format(time.RFC3339)
	case "custom":
		trxParams.RangeStart = helper.StartOfDay(helper.StrDateToTime(trxParams.RangeStart)).Format(time.RFC3339)
		trxParams.RangeEnd = helper.EndOfDay(helper.StrDateToTime(trxParams.RangeEnd)).Format(time.RFC3339)
	default: // case "month":
		trxParams.RangeStart = helper.StartOfMonth(&now).Format(time.RFC3339)
		trxParams.RangeEnd = helper.EndOfMonth(&now).Format(time.RFC3339)
	}
	ctx, cancel := context.WithCancel(rctx)
	defer cancel()
	trxResp, err = impl.svc.FindAll(ctx, trxParams)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(api.RespListTransactions{
		Meta: &api.Meta{
			Code:    200,
			Message: "Success list transactions",
		},
		Data: trxResp,
	})
	if err != nil {
		panic(err)
	}
}

func (impl *TransactionHandlerImpl) GetTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		rctx    context.Context = r.Context()
		trxId   string          = chi.URLParam(r, "trxId")
		trxResp *api.Transaction
	)
	w.Header().Set("Content-Type", "application/json")
	if len(trxId) != 36 {
		panic(exception.NewUnprocessableEntityException("Transaction id is not valid"))
	}
	ctx, cancel := context.WithCancel(rctx)
	defer cancel()
	trxResp, err = impl.svc.FindById(ctx, &trxId)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(api.RespDetailTransaction{
		Meta: &api.Meta{
			Code:    200,
			Message: "Success get detail transaction",
		},
		Data: trxResp,
	})
	if err != nil {
		panic(err)
	}
}

func (impl *TransactionHandlerImpl) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		rctx    context.Context           = r.Context()
		trxReq  *api.ReqCreateTransaction = &api.ReqCreateTransaction{}
		trxResp *api.Transaction
	)
	w.Header().Set("Content-Type", "application/json")
	err = r.ParseForm()
	if err != nil {
		panic(exception.NewBadRequestException(err.Error()))
	}
	err = form.NewDecoder().Decode(trxReq, r.PostForm)
	if err != nil {
		panic(err)
	}
	err = impl.validate.Struct(trxReq)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(rctx)
	defer cancel()
	trxResp = impl.svc.Create(ctx, trxReq)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(api.RespDetailTransaction{
		Meta: &api.Meta{
			Code:    201,
			Message: "Success create transaction",
		},
		Data: trxResp,
	})
	if err != nil {
		panic(err)
	}
}

func (impl *TransactionHandlerImpl) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		rctx    context.Context           = r.Context()
		trxId   string                    = chi.URLParam(r, "trxId")
		trxReq  *api.ReqUpdateTransaction = &api.ReqUpdateTransaction{}
		trxResp *api.Transaction
	)
	w.Header().Set("Content-Type", "application/json")
	if len(trxId) != 36 {
		panic(exception.NewUnprocessableEntityException("Transaction id is not valid"))
	}
	err = r.ParseForm()
	if err != nil {
		panic(exception.NewBadRequestException(err.Error()))
	}
	err = form.NewDecoder().Decode(trxReq, r.PostForm)
	if err != nil {
		panic(err)
	}
	err = impl.validate.Struct(trxReq)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(rctx)
	defer cancel()
	trxResp, err = impl.svc.Update(ctx, trxReq, &trxId)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(api.RespDetailTransaction{
		Meta: &api.Meta{
			Code:    200,
			Message: "Success update transaction",
		},
		Data: trxResp,
	})
	if err != nil {
		panic(err)
	}
}

func (impl *TransactionHandlerImpl) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		rctx  context.Context = r.Context()
		trxId string          = chi.URLParam(r, "trxId")
	)
	w.Header().Set("Content-Type", "application/json")
	if len(trxId) != 36 {
		panic(exception.NewUnprocessableEntityException("Transaction id is not valid"))
	}
	ctx, cancel := context.WithCancel(rctx)
	defer cancel()
	err = impl.svc.Delete(ctx, &trxId)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(api.RespDetailTransaction{
		Meta: &api.Meta{
			Code:    200,
			Message: "Success delete transaction",
		},
	})
	if err != nil {
		panic(err)
	}
}
