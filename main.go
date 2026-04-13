package main

import (
	"os"

	"github.com/xdevplatform/xurl/cmd"
)

// main is the entry point for xurl, a command-line tool for interacting
// with the X (Twitter) API using OAuth credentials.
func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
