package pter

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing"

	fuzz "github.com/google/gofuzz"
)

func TestAddSpec(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}

	assertion := func(a, b int) bool {
		return add(a, b) == a+b
	}

	QuickCheck(t, assertion, &Config{
		MaxCount: 10000,
	})
}

type AType struct {
	Name   string
	Typ    int
	I64Ptr *int64
	Map    map[string]string
	Slice  []string
}

func TestJsonSpecWithCustomFuzzer(t *testing.T) {
	assertion := func(a *AType) bool {
		bs, err := json.Marshal(a)
		if err != nil {
			return false
		}

		got := &AType{}
		err = json.Unmarshal(bs, got)
		if err != nil {
			return false
		}

		return reflect.DeepEqual(a, got)
	}

	f := fuzz.New().NilChance(0)
	QuickCheck(t, assertion, &Config{
		MaxCount: 2000,
		Fuzzer:   f,
	})
}

func TestMultiplySpecWithCustomArgsFn(t *testing.T) {
	multiply := func(a, b int) int {
		return a * b
	}

	assertion := func(a, b int) bool {
		return multiply(a, b) == a*b
	}

	QuickCheck(t, assertion, &Config{
		MaxCount: 10000,
		Args: func() []interface{} {
			return []interface{}{rand.Int(), rand.Int()}
		},
	})
}
