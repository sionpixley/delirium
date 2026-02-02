package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sionpixley/delirium/internal/constants"
	"github.com/sionpixley/delirium/pkg/pseudorand"
	"github.com/sionpixley/delirium/pkg/securerand"
	flag "github.com/spf13/pflag"
)

func execute(numOfBytes int, enc string, secure bool) {
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

func parseCLIOptions() (int, string, bool) {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		_, _ = fmt.Fprintln(w, constants.Help)
	}

	var numOfBytes int
	flag.IntVarP(&numOfBytes, "bytes", "B", 16, "number of bytes to use in the random algorithm")

	var enc string
	flag.StringVar(&enc, "encoding", "base64", "the encoding to use for the random algorithm")

	var secure bool
	flag.BoolVar(&secure, "secure", false, "if 'true', the random algorithm will be cryptographically-secure")

	var version bool
	flag.BoolVarP(&version, "version", "v", false, "print the version and exit")

	flag.Parse()

	if version {
		fmt.Println(constants.Version)
		os.Exit(0)
	}

	enc = strings.ToLower(enc)

	return numOfBytes, enc, secure
}
