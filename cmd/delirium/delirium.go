package main

func main() {
	numOfBytes, enc, secure := parseCLIOptions()
	execute(numOfBytes, enc, secure)
}
