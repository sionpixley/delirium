/*
Package pseudorand produces non-cryptographically-secure random values.
For cryptographically-secure random values, use package [securerand].

This package is a wrapper around [math/rand/v2].

[math/rand/v2]: https://pkg.go.dev/math/rand/v2
[securerand]: https://pkg.go.dev/github.com/sionpixley/delirium/pkg/securerand
*/
package pseudorand

import (
	"encoding/base64"
	"encoding/hex"
	"math/rand/v2"
)

/*
Makes a random base64 string.
It is not cryptographically-secure. Use package securerand if you need cryptographically-secure values.

If the 'urlSafe' argument is true then it will use base64.URLEncoding.
If the 'urlSafe' argument is false then it will use base64.StdEncoding.
*/
func Base64String(numOfBytes int, urlSafe bool) string {
	b := Bytes(numOfBytes)

	if urlSafe {
		return base64.URLEncoding.EncodeToString(b)
	} else {
		return base64.StdEncoding.EncodeToString(b)
	}
}

// Makes a random []byte.
// It is not cryptographically-secure. Use package securerand if you need cryptographically-secure values.
func Bytes(numOfBytes int) []byte {
	b := make([]byte, numOfBytes)
	for i := range numOfBytes {
		b[i] = byte(rand.Uint32())
	}
	return b
}

// Makes a random hexadecimal string.
// It is not cryptographically-secure. Use package securerand if you need cryptographically-secure values.
func Hex(numOfBytes int) string {
	b := Bytes(numOfBytes)
	return hex.EncodeToString(b)
}
