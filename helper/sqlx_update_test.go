package helper

import (
	"database/sql"
	"restfulapi/model/domain"
	"strings"
	"testing"
	"time"
)

// go test -v -cpu=1 -race -count=1 -failfast -run=TestSqlXColumnsUpdate ./
func TestSqlXColumnsUpdate(t *testing.T) {
	var (
		w   strings.Builder
		trx *domain.Transaction
	)
	w = strings.Builder{}
	trx = &domain.Transaction{
		Uuid:   InlinePointer("01956b42-b552-7e8e-b1b8-d0a64eed5182"),
		Amount: InlinePointer[float64](100000),
		Date:   &sql.NullTime{Valid: true, Time: time.Now()},
		Mtm:    InlinePointer(time.Now()),
	}
	w.Grow(85)
	w.WriteString("update transaction set")
	SqlXColumnsUpdate(&w, trx, []string{"uuid"})
	w.WriteString(" where uuid = :uuid;")
	s := w.String()
	want := "update transaction set amount = :amount, mtm = :mtm, date = :date where uuid = :uuid;"
	if s != want {
		t.Fatalf("got %s\tshould %s\n", s, want)
	}
	t.Logf("%s\n", w.String())
}
