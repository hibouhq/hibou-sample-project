// insecure_net.go adds a second batch of INTENTIONALLY insecure patterns, on top
// of the ones in insecure.go, so a pull request can introduce brand-new SAST
// findings rather than only shifting counts on existing ones. These are meant to
// be flagged, so there are no suppression annotations. Nothing here is reachable
// from the tested surface. Do not copy into production code.
// See ../SECURITY-FIXTURES.md.
package calc

import (
	"crypto/sha1"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
)

// InsecureClient returns an HTTP client that accepts any TLS certificate
// (gosec G402 — TLS InsecureSkipVerify set true).
func InsecureClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

// FetchURL performs a GET against a caller-supplied URL
// (gosec G107 — URL as taint input, server-side request forgery).
func FetchURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	return io.ReadAll(resp.Body)
}

// ReadUserFile reads whatever path the caller names, with no containment
// (gosec G304 — file inclusion via variable, path traversal).
func ReadUserFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// LegacyDigest hashes s with SHA-1 (gosec G401/G505 — weak, broken primitive).
func LegacyDigest(s string) string {
	h := sha1.Sum([]byte(s))
	return fmt.Sprintf("%x", h)
}
