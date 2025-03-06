# Delirium

Delirium is a CLI tool and Go library that produces random values. Supports cryptographically-secure random values. 

## Table of contents

1. [Project structure](#project-structure)
2. [How to install the CLI](#how-to-install-delirium-cli-tool)
3. [How to use the CLI](#how-to-use-delirium-cli)
    1. [Printing help and usage information](#printing-help-and-usage-information)
    2. [Making random base64 strings](#making-random-base64-strings)
    3. [Making random URL-safe base64 strings](#making-random-url-safe-base64-strings)
    4. [Making random hexadecimal strings](#making-random-hexadecimal-strings)
4. [How to install the library](#how-to-install-delirium-library)
5. [How to use the library](#how-to-use-delirium-library)
    1. [securerand](#securerand)
        1. [Making random bytes](#making-random-bytes)
        2. [Making random base64 strings](#making-random-base64-strings-1)
        3. [Making random URL-safe base64 strings](#making-random-url-safe-base64-strings-1)
        4. [Making random hexadecimal strings](#making-random-hexadecimal-strings-1)
    2. [pseudorand](#pseudorand)
        1. [Making random bytes](#making-random-bytes-1)
        2. [Making random base64 strings](#making-random-base64-strings-2)
        3. [Making random URL-safe base64 strings](#making-random-url-safe-base64-strings-2)
        4. [Making random hexadecimal strings](#making-random-hexadecimal-strings-2)
6. [Building from source](#building-from-source)
    1. [Building directly](#building-directly)
        1. [Required technologies](#required-technologies)
        2. [Building on Unix-like systems](#building-on-unix-like-systems)
        3. [Building on Windows](#building-on-windows)
    2. [Dockerfile](#dockerfile)
        1. [Required technologies](#required-technologies-1)
        2. [Building an image](#building-an-image)
7. [Contributing](#contributing)

## Project structure

```
.
├── CODE_OF_CONDUCT.md
├── Dockerfile
├── LICENSE
├── README.md
├── SECURITY.md
├── cmd
│   └── delirium
│       └── delirium.go
├── go.mod
├── internal
│   └── constants
│       └── consts.go
└── pkg
    ├── pseudorand
    │   └── pseudorand.go
    └── securerand
        └── securerand.go
```

## How to install Delirium CLI tool

I currently do not provide any prebuilt binaries for this tool. In order to install the CLI, you will need to install Go first and then run this command after Go has been installed:

`go install github.com/sionpixley/delirium/cmd/delirium@latest`

Make sure the GOPATH bin directory is in your PATH environment variable or the `delirium` command won't be found.

## How to use Delirium CLI

The `delirium` command does not require sudo/admin to use.

### Printing help and usage information

`delirium -h`

### Making random base64 strings

By default, the `delirium` command will produce non-cryptographically-secure base64 strings with 16 bytes. If these defaults are ok with you, then run:

`delirium`

To change the number of bytes to use in the algorithm, specify the number of bytes with the `-B` flag:

`delirium -B 20`

To change the algorithm to be cryptographically-secure, add the `-secure` flag:

`delirium -B 20 -secure`

### Making random URL-safe base64 strings

To change the encoding to use base64's URL-safe encoding, specify the `-encoding` flag with the value `base64url`:

`delirium -encoding=base64url`

Regardless of the encoding used, you can still specify the number of bytes to use and whether to make the algorithm cryptographically-secure:

`delirium -encoding=base64url -B 20 -secure`

### Making random hexadecimal strings

To change the encoding to use hexadecimal, specify the `-encoding` flag with the value `hex`:

`delirium -encoding=hex`

Regardless of the encoding used, you can still specify the number of bytes to use and whether to make the algorithm cryptographically-secure:

`delirium -encoding=hex -B 20 -secure`

### Printing your current version of Delirium

`delirium -v`

## How to install Delirium library

`go get github.com/sionpixley/delirium`

## How to use Delirium library

Delirium is a Go module that comes with two packages: `securerand` and `pseudorand`. The `securerand` package is a wrapper around the Go standard library [crypto/rand](https://pkg.go.dev/crypto/rand) and produces cryptographically-secure random values. The `pseudorand` package is a wrapper around the Go standard library [math/rand/v2](https://pkg.go.dev/math/rand/v2) and produces random values that are not cryptographically-secure.

### securerand

Make sure to include the import:

`import "github.com/sionpixley/delirium/pkg/securerand"`

#### Making random bytes

`secureRandomBytes, err := securerand.Bytes(<insert-num-of-bytes>)`

#### Making random base64 strings

`secureRandomBase64, err := securerand.Base64String(<insert-num-of-bytes>, false)`

#### Making random URL-safe base64 strings

`secureRandomBase64, err := securerand.Base64String(<insert-num-of-bytes>, true)`

#### Making random hexadecimal strings

`secureRandomHex, err := securerand.HexString(<insert-num-of-bytes>)`

### pseudorand

Make sure to include the import:

`import "github.com/sionpixley/delirium/pkg/pseudorand"`

#### Making random bytes

`randomBytes := pseudorand.Bytes(<insert-num-of-bytes>)`

#### Making random base64 strings

`randomBase64 := pseudorand.Base64String(<insert-num-of-bytes>, false)`

#### Making random URL-safe base64 strings

`randomBase64 := pseudorand.Base64String(<insert-num-of-bytes>, true)`

#### Making random hexadecimal strings

`randomHex := pseudorand.HexString(<insert-num-of-bytes>)`

## Building from source

There are two main ways to build Delirium from source: [Building with Go directly](#building-directly) or [building the Dockerfile](#dockerfile).

Building with the Dockerfile is good for quick local testing.

### Building directly

#### Required technologies

- Go 1.24.1

#### Building on Unix-like systems

`go build -o delirium ./cmd/delirium`

#### Building on Windows

`go build -o delirium.exe ./cmd/delirium`

### Dockerfile

#### Required technologies

- Docker

#### Building an image

`docker build -t delirium .`

## Contributing

All contributions are welcome! If you wish to contribute to the project, the best way would be forking this repo and making a pull request from your fork with all of your suggested changes.
