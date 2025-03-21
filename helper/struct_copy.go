package helper

import "reflect"

// Automate copying struct fields, with assumption each struct field names and types are identical.
// Not each field must identical, just make sure there are one that identical.
// And make sure you pass a pointer or interface (real interface, not any).
func StructCopy(from any, to any) {
	var (
		src, dst reflect.Value = reflect.ValueOf(from).Elem(), reflect.ValueOf(to).Elem()
	)
	for i := 0; i < src.NumField(); i++ {
		srcVal := src.Field(i)
		if IsEmptyValue(srcVal) {
			continue
		}
		srcField := src.Type().Field(i)
		dstVal := dst.FieldByName(srcField.Name)
		if dstVal.IsValid() && dstVal.CanSet() {
			if srcVal.Type() == dstVal.Type() {
				dstVal.Set(srcVal)
			}
		}
	}
}
