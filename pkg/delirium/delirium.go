package delirium

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"math"
)

func RandomBase64String(numOfBytes int) (string, error) {
	b, err := RandomBytes(numOfBytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func RandomBytes(numOfBytes int) ([]byte, error) {
	b := make([]byte, numOfBytes)
	_, err := rand.Read(b)
	return b, err
}

func RandomInt16() (int16, error) {
	b, err := RandomBytes(2)
	if err != nil {
		return 0, nil
	}

	return int16(binary.NativeEndian.Uint16(b)), nil
}

func RandomInt32() (int32, error) {
	b, err := RandomBytes(4)
	if err != nil {
		return 0, err
	}

	return int32(binary.NativeEndian.Uint32(b)), nil
}

func RandomInt64() (int64, error) {
	b, err := RandomBytes(8)
	if err != nil {
		return 0, nil
	}

	return int64(binary.NativeEndian.Uint64(b)), nil
}

func RandomFloat32() (float32, error) {
	b, err := RandomBytes(4)
	if err != nil {
		return 0, err
	}

	return math.Float32frombits(binary.NativeEndian.Uint32(b)), nil
}

func RandomFloat64() (float64, error) {
	b, err := RandomBytes(8)
	if err != nil {
		return 0, nil
	}

	return math.Float64frombits(binary.NativeEndian.Uint64(b)), nil
}

func RandomUint16() (uint16, error) {
	b, err := RandomBytes(2)
	if err != nil {
		return 0, nil
	}

	return binary.NativeEndian.Uint16(b), nil
}

func RandomUint32() (uint32, error) {
	b, err := RandomBytes(4)
	if err != nil {
		return 0, err
	}

	return binary.NativeEndian.Uint32(b), nil
}

func RandomUint64() (uint64, error) {
	b, err := RandomBytes(8)
	if err != nil {
		return 0, nil
	}

	return binary.NativeEndian.Uint64(b), nil
}
