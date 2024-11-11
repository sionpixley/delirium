# Delirium

Delirium is a Go library that produces cryptographically-secure random values. Delirium is a wrapper around the standard library [crypto/rand](https://pkg.go.dev/crypto/rand), but it makes it easier to produce values of smaller size than crypto/rand is usually used for.

## How to install

`go get github.com/sionpixley/delirium`

## How to use

Make sure to include the import:

`import "github.com/sionpixley/delirium/pkg/delirium"`

### Producing random bytes

`randomBytes := delirium.RandomBytes(<insert-num-of-bytes>)`

### Producing random base64 strings

`randomBase64 := delirium.RandomBase64String(<insert-num-of-bytes>)`

### Producing random integers

Delirium can produce random signed and unsigned integers of sizes: 16-bit, 32-bit, and 64-bit.

```
randomInt16 := delirium.RandomInt16()
randomUint16 := delirium.RandomUint16()

randomInt32 := delirium.RandomInt32()
randomUint32 := delirium.RandomUint32()

randomInt64 := delirium.RandomInt64()
randomUint64 := delirium.RandomUint64()
```

### Producing random floats

Delirium can produce random float32 or float64 values.

```
randomFloat32 := delirium.RandomFloat32()
randomFloat64 := delirium.RandomFloat64()
```
