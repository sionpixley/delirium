# Delirium

Delirium is a Go library that produces cryptographically-secure random string values. Delirium is a wrapper around the standard library [crypto/rand](https://pkg.go.dev/crypto/rand).

## How to install

`go get github.com/sionpixley/delirium`

## How to use

Make sure to include the import:

`import "github.com/sionpixley/delirium/pkg/delirium"`

### Making random bytes

`randomBytes := delirium.MakeRandomBytes(<insert-num-of-bytes>)`

### Making random base64 strings

`randomBase64 := delirium.MakeRandomBase64String(<insert-num-of-bytes>)`

### Making random URL-safe base64 strings

`randomBase64 := delirium.MakeRandomBase64UrlString(<insert-num-of-bytes>)`

### Making random hexadecimal strings

`randomHex := delirium.MakeRandomHexString(<insert-num-of-bytes>)`
