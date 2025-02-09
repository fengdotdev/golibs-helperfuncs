package web

import (
	"fmt"
	"io"
	"net/http"
)

// GetRemoteResource makes an HTTP GET request to the provided URL,
// checks for a successful response, reads the body, and returns its content.
func GetRemoteResource(url string) ([]byte, error) {
	// Perform the HTTP GET request.
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching URL %q: %w", url, err)
	}
	// Ensure that the response body is closed when the function returns.
	defer resp.Body.Close()

	// Check if the server returned a successful status code.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d %s", resp.StatusCode, resp.Status)
	}

	// Read the response body.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return data, nil
}
