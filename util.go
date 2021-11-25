package pter

import (
	"reflect"

	fuzz "github.com/google/gofuzz"
)

func newValue(fuzzer *fuzz.Fuzzer, typ reflect.Type) reflect.Value {
	if typ.Kind() == reflect.Ptr {
		etyp := typ.Elem()
		v := reflect.New(etyp)
		fuzzer.Fuzz(v.Interface())
		return v
	}

	v := reflect.New(typ)
	fuzzer.Fuzz(v.Interface())
	return v.Elem()
}
