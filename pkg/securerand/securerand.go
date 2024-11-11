package securerand

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

/*
Makes a cryptographically-secure random base64 string.

If the 'urlSafe' argument is true then it will use base64.URLEncoding.
If the 'urlSafe' argument is false then it will use base64.StdEncoding.
*/
func Base64String(numOfBytes int, urlSafe bool) (string, error) {
	b, err := Bytes(numOfBytes)
	if err != nil {
		return "", err
	}

	if urlSafe {
		return base64.URLEncoding.EncodeToString(b), nil
	} else {
		return base64.StdEncoding.EncodeToString(b), nil
	}
}

// Makes a cryptographically-secure random []byte.
func Bytes(numOfBytes int) ([]byte, error) {
	b := make([]byte, numOfBytes)
	_, err := rand.Read(b)
	return b, err
}

// Makes a cryptographically-secure random hexadecimal string.
func HexString(numOfBytes int) (string, error) {
	b, err := Bytes(numOfBytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
