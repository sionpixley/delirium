package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/sionpixley/delirium/internal/constants"
	"github.com/sionpixley/delirium/pkg/pseudorand"
	"github.com/sionpixley/delirium/pkg/securerand"
)

func main() {
	flag.Usage = func() {
		fmt.Println(constants.Help)
	}

	var numOfBytes int
	flag.IntVar(&numOfBytes, "B", 16, "number of bytes to use in the random algorithm")

	var enc string
	flag.StringVar(&enc, "encoding", "base64", "the encoding to use for the random algorithm")

	var secure bool
	flag.BoolVar(&secure, "secure", false, "if 'true', the random algorithm will be cryptographically-secure")

	var version bool
	flag.BoolVar(&version, "v", false, "print the version and exit")
	flag.BoolVar(&version, "version", false, "print the version and exit")

	if version {
		fmt.Println(constants.Version)
		return
	}

	flag.Parse()

	enc = strings.ToLower(enc)

	if secure {
		switch enc {
		case "base64":
			output, err := securerand.Base64String(numOfBytes, false)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(output)
		case "base64url":
			output, err := securerand.Base64String(numOfBytes, true)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(output)
		case "hex":
			output, err := securerand.HexString(numOfBytes)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(output)
		default:
			log.Fatalln(constants.InvalidEncodingError)
		}
	} else {
		switch enc {
		case "base64":
			fmt.Println(pseudorand.Base64String(numOfBytes, false))
		case "base64url":
			fmt.Println(pseudorand.Base64String(numOfBytes, true))
		case "hex":
			fmt.Println(pseudorand.HexString(numOfBytes))
		default:
			log.Fatalln(constants.InvalidEncodingError)
		}
	}
}
