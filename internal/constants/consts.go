package constants

const (
	HELP string = `Description:

    Delirium is a CLI tool that produces random values. Supports cryptographically-secure random values.

Usage:

    delirium [option...]

Options:

    -B <int>
        Number of bytes to use in the random algorithm. Default is '16'.
    -encoding=<value>
        Chooses the encoding used on the random algorithm. Valid values are 'base64', 'base64url', or 'hex'.
        Default is 'base64'.
    -h, -help
        Prints help and usage information.
    -secure
        If this flag is added, the random algorithm will produce cryptographically-secure random values.
    -v, -version
        Prints the current version of Delirium.`

	INVALID_ENCODING_VALUE string = "delirium error: invalid value for the encoding flag"

	VERSION string = "v0.3.5"
)
