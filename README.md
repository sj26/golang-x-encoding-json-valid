# JSON IsValid

This is a single function, `func IsValid([]byte) error`, exposing some private
functionality in the go standard library [`encoding/json` package][json-package]. It does a
scan over encoded JSON data to make sure it is valid without too many
allocations.

The file `scanner.go` is copied directly from the json package. `valid.go` adds
the method, and `valid_test.go` contains a simple test.

  [json-package]: https://golang.org/pkg/encoding/json/

## License

`scanner.go` retains the BSD license from Go. The rest is also BSD licensed.
