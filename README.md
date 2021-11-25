# pter

Go Property Tester base on testing/quick.

## usage

```
$ go get github.com/a3d21/pter
```

## example
goto [example_test.go](./example_test.go).
```go
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
```