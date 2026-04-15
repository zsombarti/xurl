// Package cmd provides the CLI commands for xurl.
// xurl is a command-line HTTP client for the X (Twitter) API
// with built-in authentication support.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"

	// flags
	verbose    bool
	headers    []string
	method     string
	body       string
	prettyJSON bool
)

// rootCmd is the base command for xurl.
var rootCmd = &cobra.Command{
	Use:   "xurl [flags] <url>",
	Short: "A curl-like CLI for the X (Twitter) API",
	Long: `xurl is a command-line tool for making authenticated requests
to the X (Twitter) API. It handles OAuth authentication automatically
using credentials stored in the environment or a config file.

Examples:
  xurl https://api.twitter.com/2/users/me
  xurl -X POST -d '{"text":"Hello world"}' https://api.twitter.com/2/tweets
  xurl -H "Content-Type: application/json" https://api.twitter.com/2/users/me`,
	Args:          cobra.ExactArgs(1),
	RunE:          runRequest,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs the root command and handles top-level errors.
func Execute(version string) {
	Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output (show request/response headers)")
	rootCmd.Flags().StringArrayVarP(&headers, "header", "H", nil, "HTTP headers to include (can be specified multiple times)")
	rootCmd.Flags().StringVarP(&method, "request", "X", "", "HTTP method to use (default: GET, or POST if body is provided)")
	rootCmd.Flags().StringVarP(&body, "data", "d", "", "Request body data")
	// Default pretty-print to false so raw output can be piped to jq without interference
	rootCmd.Flags().BoolVarP(&prettyJSON, "pretty", "p", false, "Pretty-print JSON responses")
}

// runRequest is the main handler for the root command.
// It builds and executes an authenticated HTTP request to the X API.
func runRequest(cmd *cobra.Command, args []string) error {
	targetURL := args[0]

	// Determine HTTP method
	httpMethod := method
	if httpMethod == "" {
		if body != "" {
			httpMethod = "POST"
		} else {
			httpMethod = "GET"
		}
	}

	client, err := NewClient()
	if err != nil {
		return fmt.Errorf("failed to initialize client: %w", err)
	}

	resp, err := client.Do(httpMethod, targetURL, headers, body, verbose)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if prettyJSON {
		return resp.PrintPretty(os.Stdout)
	}
	return resp.Print(os.Stdout)
}
