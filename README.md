# Delirium

Delirium is a Go library that produces random values. Delirium comes with two Go packages: `securerand` and `pseudorand`. The `securerand` package is a wrapper around the Go standard library [crypto/rand](https://pkg.go.dev/crypto/rand) and produces cryptographically-secure random values. The `pseudorand` package is a wrapper around the Go standard library [math/rand/v2](https://pkg.go.dev/math/rand/v2) and produces random values that are not cryptographically-secure. 

## Table of contents

1. [Project directory structure](#project-directory-structure)
2. [How to install](#how-to-install)
3. [How to use securerand](#how-to-use-securerand)
    1. [Making random bytes](#making-random-bytes)
    2. [Making random base64 strings](#making-random-base64-strings)
    3. [Making random URL-safe base64 strings](#making-random-url-safe-base64-strings)
    4. [Making random hexadecimal strings](#making-random-hexadecimal-strings)
4. [How to use pseudorand](#how-to-use-pseudorand)
    1. [Making random bytes](#making-random-bytes-1)
    2. [Making random base64 strings](#making-random-base64-strings-1)
    3. [Making random URL-safe base64 strings](#making-random-url-safe-base64-strings-1)
    4. [Making random hexadecimal strings](#making-random-hexadecimal-strings-1)

## Project directory structure

```
delirium
├── CODE_OF_CONDUCT.md
├── LICENSE
├── README.md
├── SECURITY.md
├── go.mod
└── pkg
    ├── pseudorand
    │   └── pseudorand.go
    └── securerand
        └── securerand.go
```

## How to install

`go get github.com/sionpixley/delirium`

## How to use securerand

Make sure to include the import:

`import "github.com/sionpixley/delirium/pkg/securerand"`

### Making random bytes

`secureRandomBytes, err := securerand.Bytes(<insert-num-of-bytes>)`

### Making random base64 strings

`secureRandomBase64, err := securerand.Base64String(<insert-num-of-bytes>, false)`

### Making random URL-safe base64 strings

`secureRandomBase64, err := securerand.Base64String(<insert-num-of-bytes>, true)`

### Making random hexadecimal strings

`secureRandomHex, err := securerand.HexString(<insert-num-of-bytes>)`

## How to use pseudorand

Make sure to include the import:

`import "github.com/sionpixley/delirium/pkg/pseudorand"`

### Making random bytes

`randomBytes := pseudorand.Bytes(<insert-num-of-bytes>)`

### Making random base64 strings

`randomBase64, err := pseudorand.Base64String(<insert-num-of-bytes>, false)`

### Making random URL-safe base64 strings

`randomBase64, err := pseudorand.Base64String(<insert-num-of-bytes>, true)`

### Making random hexadecimal strings

`randomHex, err := pseudorand.HexString(<insert-num-of-bytes>)`
