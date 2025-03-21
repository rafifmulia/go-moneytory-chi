package helper

import (
	"fmt"
	"reflect"
	"strings"
)

// Set column update if struct fields exist and non empty.
// Variable data and skipCol should never be edited.
func SqlXColumnsUpdate(w *strings.Builder, data any, skipCol []string) {
	var (
		v reflect.Value = reflect.ValueOf(data).Elem()
	)
	typeOf := v.Type()
	if typeOf.Kind() != reflect.Struct {
		panic(fmt.Sprintf("SqlXColumnsUpdate data must struct, got %v", typeOf.Kind()))
	}
	hasWritten := false // Check if there are recently written "column update" before adding the next "column update", and in between, insert the delimiter.
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := typeOf.Field(i).Tag.Get("db")
		if skipColSet(&fieldName, &skipCol) {
			continue
		}
		if !IsEmptyValue(field) {
			addDelimiterNotAtEnd2(w, &hasWritten, fmt.Sprintf(" %s = :%s", fieldName, fieldName))
		}
	}
}

func addDelimiterNotAtEnd2(w *strings.Builder, hasWritten *bool, colSet string) {
	if *hasWritten {
		w.WriteString(",")
	} else {
		*hasWritten = true
	}
	w.WriteString(colSet)
}
