package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sionpixley/delirium/internal/constants"
	"github.com/sionpixley/delirium/pkg/pseudorand"
	"github.com/sionpixley/delirium/pkg/securerand"
)

func main() {
	defer fmt.Println()

	flag.Usage = func() {
		fmt.Println(constants.HELP)
		fmt.Println()
	}

	var numOfBytes int
	flag.IntVar(&numOfBytes, "B", 16, "number of bytes to use in the random algorithm")

	var enc string
	flag.StringVar(&enc, "encoding", "base64", "the encoding to use for the random algorithm")

	var useSecure bool
	flag.BoolVar(&useSecure, "secure", false, "if 'true', the random algorithm will be cryptographically-secure")

	flag.Parse()

	if useSecure {
		switch enc {
		case "base64":
			output, err := securerand.Base64String(numOfBytes, false)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Print(output)
		case "base64url":
			output, err := securerand.Base64String(numOfBytes, true)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Print(output)
		case "hex":
			output, err := securerand.HexString(numOfBytes)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Print(output)
		default:
			log.Fatal(constants.INVALID_ENCODING_VALUE)
		}
	} else {
		switch enc {
		case "base64":
			fmt.Print(pseudorand.Base64String(numOfBytes, false))
		case "base64url":
			fmt.Print(pseudorand.Base64String(numOfBytes, true))
		case "hex":
			fmt.Print(pseudorand.Hex(numOfBytes))
		default:
			log.Fatal(constants.INVALID_ENCODING_VALUE)
		}
	}
}
