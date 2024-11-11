package pseudorand

import (
	"encoding/base64"
	"encoding/hex"
	"math/rand/v2"
)

/*
Makes a random base64 string.
It is not cryptographically-secure. Use the 'securerand' package if you need cryptographically-secure values.

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
// It is not cryptographically-secure. Use the 'securerand' package if you need cryptographically-secure values.
func Bytes(numOfBytes int) []byte {
	b := make([]byte, numOfBytes)
	for i := 0; i < numOfBytes; i += 1 {
		b[i] = byte(rand.Uint32())
	}
	return b
}

// Makes a random hexadecimal string.
// It is not cryptographically-secure. Use the 'securerand' package if you need cryptographically-secure values.
func Hex(numOfBytes int) string {
	b := Bytes(numOfBytes)
	return hex.EncodeToString(b)
}
