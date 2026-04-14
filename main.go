package main

import (
	"fmt"
	"os"

	"github.com/xdevplatform/xurl/cmd"
)

// main is the entry point for xurl, a command-line tool for interacting
// with the X (Twitter) API using OAuth credentials.
// Personal fork: added non-zero exit code printing for easier debugging.
// Note: also printing exit code to stderr to make scripting easier.
// TODO: consider supporting XURL_DEBUG env var to toggle verbose error output.
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		if os.Getenv("XURL_VERBOSE") != "" {
			fmt.Fprintf(os.Stderr, "Exit code: 1\n")
		}
		os.Exit(1)
	}
}
