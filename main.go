package main

import (
	"fmt"
	"os"

	"github.com/xdevplatform/xurl/cmd"
)

// main is the entry point for xurl, a command-line tool for interacting
// with the X (Twitter) API using OAuth credentials.
// Personal fork: added non-zero exit code printing for easier debugging.
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
