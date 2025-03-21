package helper

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
)

var (
	sqlBindVar string = "?"
)

// Set column update if struct fields exist and non empty.
// Variable data and skipCol should never be edited.
func SqlColumnsUpdate(w *strings.Builder, data any, skipCol []string, bindArgs *[]interface{}) {
	var (
		v reflect.Value = reflect.ValueOf(data).Elem()
	)
	typeOf := v.Type()
	if typeOf.Kind() != reflect.Struct {
		panic(fmt.Sprintf("SqlColumnsUpdate data must struct, got %v", typeOf.Kind()))
	}
	hasWritten := false // Check if there are recently written "column update" before adding the next "column update", and in between, insert the delimiter.
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := typeOf.Field(i).Tag.Get("db")
		if skipColSet(&fieldName, &skipCol) {
			continue
		}
		if !IsEmptyValue(field) {
			addDelimiterNotAtEnd(w, &field, &hasWritten, fmt.Sprintf(" %s = %s", fieldName, sqlBindVar), bindArgs)
		}
	}
}

func addDelimiterNotAtEnd(w *strings.Builder, field *reflect.Value, hasWritten *bool, colSet string, bindArgs *[]interface{}) {
	if *hasWritten {
		w.WriteString(",")
	} else {
		*hasWritten = true
	}
	var val any
	switch field.Type() {
	case reflect.TypeOf(time.Time{}):
		val = field.Interface().(time.Time).Format(time.RFC3339)
	case reflect.TypeOf(sql.NullString{}):
		val = field.Interface().(sql.NullString).String
	case reflect.TypeOf(sql.NullTime{}):
		val = field.Interface().(sql.NullTime).Time.Format(time.RFC3339)
	default:
		val = field.Interface()
	}
	*bindArgs = append(*bindArgs, val)
	w.WriteString(colSet)
}

func skipColSet(fieldName *string, skipCol *[]string) bool {
	for i := range *skipCol {
		if *fieldName == "" || *fieldName == (*skipCol)[i] {
			return true
		}
	}
	return false
}
