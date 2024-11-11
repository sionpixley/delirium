package delirium

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

// Makes a cryptographically-secure random base64 string.
func MakeRandomBase64String(numOfBytes int) (string, error) {
	b, err := MakeRandomBytes(numOfBytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

// Makes a cryptographically-secure random base64 URL-safe string.
func MakeRandomBase64UrlString(numOfBytes int) (string, error) {
	b, err := MakeRandomBytes(numOfBytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

// Makes a cryptographically-secure random []byte.
func MakeRandomBytes(numOfBytes int) ([]byte, error) {
	b := make([]byte, numOfBytes)
	_, err := rand.Read(b)
	return b, err
}

// Makes a cryptographically-secure random hexadecimal string.
func MakeRandomHexString(numOfBytes int) (string, error) {
	b, err := MakeRandomBytes(numOfBytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
