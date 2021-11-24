package pter

import (
	"encoding/json"
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

	QuickCheck(t, assertion, 2000)
}

type AType struct {
	Name   string
	Typ    int
	I64Ptr *int64
	Map    map[string]string
	Slice  []string
}

func TestJsonSpec(t *testing.T) {
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
	QuickCheckWithFuzzer(t, f, assertion, 2000)
}
