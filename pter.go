package pter

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	fuzz "github.com/google/gofuzz"
)

// FuzzArgs fuzz non-nil args generator for testing/quick
func FuzzArgs(fuzzer *fuzz.Fuzzer, assertion interface{}) func(args []reflect.Value, rand *rand.Rand) {
	ft := reflect.TypeOf(assertion)
	return func(args []reflect.Value, rand *rand.Rand) {
		for i := 0; i < ft.NumIn(); i++ {
			args[i] = newValue(fuzzer, ft.In(i))
		}
	}
}

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

func QuickCheck(t *testing.T, assertion interface{}, maxCount int) {
	f := fuzz.New()
	QuickCheckWithFuzzer(t, f, assertion, maxCount)
}

func QuickCheckWithFuzzer(t *testing.T, fuzzer *fuzz.Fuzzer, assertion interface{}, maxCount int) {
	if err := quick.Check(assertion, &quick.Config{
		MaxCount: maxCount,
		Values:   FuzzArgs(fuzzer, assertion),
	}); err != nil {
		t.Error(err)
	}
}
