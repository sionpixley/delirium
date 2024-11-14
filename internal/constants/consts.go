package constants

const (
	HELP string = `
Description:

  Delirium is a CLI tool that can produce cryptographically-secure and non-cryptographically-secure random values.

Usage:

  delirium [option...]

Options:

  -B <int>             : Number of bytes to use in the random algorithm. Default is '16'.
  -encoding <value>    : Chooses the encoding used on the random algorithm. Valid values are 'base64', 'base64url', or 'hex'.
                         Default is 'base64'.
  -h                   : Prints help and usage information.
  -secure=[<boolean>]  : If 'true', the random algorithm will produce cryptographically-secure random values. Default is 'false'.`

	INVALID_ENCODING_VALUE string = "delirium: invalid value for the encoding flag"
)
