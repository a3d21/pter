package pter

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	fuzz "github.com/google/gofuzz"
)

var (
	defaultFuzzer = fuzz.New()
	defaultConfig = &Config{MaxCount: 10000}
)

// ArgsFn generate test args
type ArgsFn func() []interface{}

// ToQuickValueFn transform to testing/quick value fn
func (f ArgsFn) ToQuickValueFn() func(values []reflect.Value, rand *rand.Rand) {
	return func(values []reflect.Value, rand *rand.Rand) {
		args := f()
		for i, a := range args {
			values[i] = reflect.ValueOf(a)
		}
	}
}

// FuzzArgs fuzz non-nil args generator for testing/quick
func FuzzArgs(fuzzer *fuzz.Fuzzer, assertion interface{}) ArgsFn {
	ft := reflect.TypeOf(assertion)
	return func() []interface{} {
		alen := ft.NumIn()
		args := make([]interface{}, alen, alen)
		for i := 0; i < alen; i++ {
			args[i] = newValue(fuzzer, ft.In(i)).Interface()
		}
		return args
	}
}

type Config struct {
	// MaxCount sets the maximum test count
	MaxCount int
	// Args generator
	Args ArgsFn
	// Fuzzer generate args, use `defaultFuzzer` when Fuzzer==nil
	Fuzzer *fuzz.Fuzzer
}

func QuickCheck(t *testing.T, assertion interface{}, c *Config) {
	if c == nil {
		cc := *defaultConfig
		c = &cc
	}
	if c.MaxCount == 0 {
		c.MaxCount = defaultConfig.MaxCount
	}
	if c.Args == nil {
		fuzzer := defaultFuzzer
		if c.Fuzzer != nil {
			fuzzer = c.Fuzzer
		}
		c.Args = FuzzArgs(fuzzer, assertion)
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: c.MaxCount,
		Values:   c.Args.ToQuickValueFn(),
	}); err != nil {
		t.Helper()
		t.Error(err)
	}
}
