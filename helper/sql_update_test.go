package helper

import (
	"database/sql"
	"reflect"
	"restfulapi/model/domain"
	"strings"
	"testing"
	"time"
)

// go test -v -cpu=1 -race -count=1 -failfast -run=TestSqlColumnsUpdate ./
func TestSqlColumnsUpdate(t *testing.T) {
	var (
		w        strings.Builder
		bindArgs []interface{}
		trx      *domain.Transaction
	)
	w = strings.Builder{}
	trx = &domain.Transaction{
		Uuid:   InlinePointer("01956b42-b552-7e8e-b1b8-d0a64eed5182"),
		Amount: InlinePointer[float64](100000),
		Date:   &sql.NullTime{Valid: true, Time: time.Now()},
		Mtm:    InlinePointer(time.Now()),
	}
	bindArgs = make([]interface{}, 0, 4)
	w.Grow(85)
	w.WriteString("update transaction set")
	SqlColumnsUpdate(&w, trx, []string{"uuid"}, &bindArgs)
	w.WriteString(" where uuid = ?;")
	bindArgs = append(bindArgs, trx.Uuid)
	s := w.String()
	want := "update transaction set amount = ?, mtm = ?, date = ? where uuid = ?;"
	if s != want {
		t.Fatalf("got %s\tshould %s\n", s, want)
	}
	t.Logf("%s\n", w.String())
	if len(bindArgs) != cap(bindArgs) {
		t.Errorf("bindArgs -> len:%d cap:%d\n", len(bindArgs), cap(bindArgs))
		for i := range bindArgs {
			v := reflect.ValueOf(bindArgs[i])
			if v.Kind() == reflect.Ptr {
				t.Errorf("%v\t", reflect.Indirect(v))
			} else {
				t.Errorf("%v\t", bindArgs[i])
			}
		}
	}
}
