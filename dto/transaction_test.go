package dto

import (
	"reflect"
	"restfulapi/api"
	"restfulapi/helper"
	"restfulapi/model/domain"
	"testing"
)

// go test -v -cpu=1 -race -count=1 -failfast -run=TestWebReqUpdateTrxToDomainTrx ./
func TestWebReqUpdateTrxToDomainTrx(t *testing.T) {
	type args struct {
		from *api.ReqUpdateTransaction
	}
	tests := []struct {
		name string
		args args
		want *domain.Transaction
	}{
		{
			name: "Test TestWebReqUpdateTrxToDomainTrx 1",
			args: args{from: &api.ReqUpdateTransaction{
				Amount: helper.InlinePointer[float64](100000),
			}},
			want: &domain.Transaction{
				Amount: helper.InlinePointer[float64](100000),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WebReqUpdateTrxToDomainTrx(tt.args.from); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebReqUpdateTrxToDomainTrx() = %v, want %v", got, tt.want)
			} else {
				t.Logf("WebReqUpdateTrxToDomainTrx() = %v", got)
			}
		})
	}
}

// go test -v -cpu=1 -race -count=1 -failfast -run=TestDomainTrxToApiTrx ./
func TestDomainTrxToApiTrx(t *testing.T) {
	type args struct {
		from *domain.Transaction
	}
	tests := []struct {
		name string
		args args
		want *api.Transaction
	}{
		{
			name: "Test TestDomainTrxToApiTrx 1",
			args: args{from: &domain.Transaction{
				Amount: helper.InlinePointer[float64](100000),
			}},
			want: &api.Transaction{
				Amount: helper.InlinePointer[float64](100000),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DomainTrxToApiTrx(tt.args.from); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainTrxToApiTrx() = %v, want %v", got, tt.want)
			}
		})
	}
}
