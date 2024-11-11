package delirium

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"math"
)

// Makes a cryptographically-secure random base64 string.
func RandomBase64String(numOfBytes int) (string, error) {
	b, err := RandomBytes(numOfBytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

// Makes a cryptographically-secure random []byte.
func RandomBytes(numOfBytes int) ([]byte, error) {
	b := make([]byte, numOfBytes)
	_, err := rand.Read(b)
	return b, err
}

// Makes a cryptographically-secure random int16.
func RandomInt16() (int16, error) {
	b, err := RandomBytes(2)
	if err != nil {
		return 0, nil
	}

	return int16(binary.NativeEndian.Uint16(b)), nil
}

// Makes a cryptographically-secure random int32.
func RandomInt32() (int32, error) {
	b, err := RandomBytes(4)
	if err != nil {
		return 0, err
	}

	return int32(binary.NativeEndian.Uint32(b)), nil
}

// Makes a cryptographically-secure random int64.
func RandomInt64() (int64, error) {
	b, err := RandomBytes(8)
	if err != nil {
		return 0, nil
	}

	return int64(binary.NativeEndian.Uint64(b)), nil
}

// Makes a cryptographically-secure random float32.
func RandomFloat32() (float32, error) {
	b, err := RandomBytes(4)
	if err != nil {
		return 0, err
	}

	return math.Float32frombits(binary.NativeEndian.Uint32(b)), nil
}

// Makes a cryptographically-secure random float64.
func RandomFloat64() (float64, error) {
	b, err := RandomBytes(8)
	if err != nil {
		return 0, nil
	}

	return math.Float64frombits(binary.NativeEndian.Uint64(b)), nil
}

// Makes a cryptographically-secure random uint16.
func RandomUint16() (uint16, error) {
	b, err := RandomBytes(2)
	if err != nil {
		return 0, nil
	}

	return binary.NativeEndian.Uint16(b), nil
}

// Makes a cryptographically-secure random uint32.
func RandomUint32() (uint32, error) {
	b, err := RandomBytes(4)
	if err != nil {
		return 0, err
	}

	return binary.NativeEndian.Uint32(b), nil
}

// Makes a cryptographically-secure random uint64.
func RandomUint64() (uint64, error) {
	b, err := RandomBytes(8)
	if err != nil {
		return 0, nil
	}

	return binary.NativeEndian.Uint64(b), nil
}
