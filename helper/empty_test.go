package helper

import (
	"reflect"
	"restfulapi/model/domain"
	"testing"
	"time"
)

// go test -v -cpu=1 -race -count=1 -failfast -run=TestIsEmptyValue ./
func TestIsEmptyValue(t *testing.T) {
	println(time.Time{}.Format(time.RFC3339))
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestIsEmptyValue 1",
			args: args{v: reflect.ValueOf(domain.Transaction{})},
			want: true,
		},
		{
			name: "TestIsEmptyValue 2",
			args: args{v: reflect.ValueOf(domain.Transaction{Mtm: InlinePointer(time.Now())})},
			want: false,
		},
		{
			name: "TestIsEmptyValue 3",
			args: args{v: reflect.ValueOf(domain.Transaction{Mtm: InlinePointer(time.Time{})})},
			want: false,
		},
		{
			name: "TestIsEmptyValue 4",
			args: args{v: reflect.ValueOf(domain.Transaction{Mtm: nil})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmptyValue(tt.args.v); got != tt.want {
				t.Errorf("IsEmptyValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
