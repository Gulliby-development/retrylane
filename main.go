package main

import (
	"io"
	"os"
)

// main reads standard input only when readsStdin says the command line is a
// valid use of a command; every other reply needs no input.
func main() {
	args := os.Args[1:]
	text := ""
	if readsStdin(args) {
		data, _ := io.ReadAll(os.Stdin)
		text = string(data)
	}
	status, out, errText := run(args, text)
	_, _ = os.Stdout.WriteString(out)
	_, _ = os.Stderr.WriteString(errText)
	os.Exit(status)
}
