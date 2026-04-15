package cmd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// RequestOptions holds configuration for an HTTP request.
type RequestOptions struct {
	Method  string
	Headers []string
	Body    string
	Timeout int
	Verbose bool
}

// BuildRequest constructs an *http.Request from the given URL and options.
func BuildRequest(rawURL string, opts RequestOptions) (*http.Request, error) {
	// Ensure the URL has a scheme
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "https://" + rawURL
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL %q: %w", rawURL, err)
	}

	var bodyReader io.Reader
	if opts.Body != "-" && opts.Body != "" {
		bodyReader = strings.NewReader(opts.Body)
	} else if opts.Body == "-" {
		bodyReader = os.Stdin
	}

	req, err := http.NewRequest(opts.Method, parsedURL.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	// Parse and apply custom headers
	for _, h := range opts.Headers {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid header format %q (expected Key: Value)", h)
		}
		req.Header.Set(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
	}

	// Set a default Content-Type for requests with a body
	if bodyReader != nil && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// ExecuteRequest sends the HTTP request and returns the response.
func ExecuteRequest(req *http.Request, opts RequestOptions) (*http.Response, error) {
	timeout := time.Duration(opts.Timeout) * time.Second
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	client := &http.Client{
		Timeout: timeout,
	}

	if opts.Verbose {
		fmt.Fprintf(os.Stderr, "> %s %s\n", req.Method, req.URL.String())
		for key, values := range req.Header {
			for _, v := range values {
				fmt.Fprintf(os.Stderr, "> %s: %s\n", key, v)
			}
		}
		fmt.Fprintln(os.Stderr)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if opts.Verbose {
		fmt.Fprintf(os.Stderr, "< %s\n", resp.Status)
		for key, values := range resp.Header {
			for _, v := range values {
				fmt.Fprintf(os.Stderr, "< %s: %s\n", key, v)
			}
		}
		fmt.Fprintln(os.Stderr)
	}

	return resp, nil
}
