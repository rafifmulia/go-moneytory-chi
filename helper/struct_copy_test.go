package helper

import (
	"reflect"
	"restfulapi/api"
	"restfulapi/model/domain"
	"testing"
)

// go test -v -cpu=1 -race -count=1 -failfast -run=TestStructCopy ./
func TestStructCopy(t *testing.T) {
	type args struct {
		from any
		to   any
	}
	from := &api.ReqUpdateTransaction{
		Amount: InlinePointer[float64](100000),
	}
	to := &domain.Transaction{}
	tests := []struct {
		name string
		args args
	}{
		{name: "TestStructCopy 1", args: args{from: from, to: to}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StructCopy(tt.args.from, tt.args.to)
			if !reflect.DeepEqual(from.Amount, to.Amount) {
				t.Errorf("WebReqUpdateTrxToDomainTrx() = %v, want %v", from, to)
			}
		})
	}
}
