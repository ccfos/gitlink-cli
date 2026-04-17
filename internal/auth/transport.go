package auth

import (
	"net/http"
	"os"
	"strings"
)

// Transport wraps an http.RoundTripper and injects authentication.
type Transport struct {
	Base http.RoundTripper
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Priority: GITLINK_TOKEN env var > stored token (keyring/file)
	token := os.Getenv("GITLINK_TOKEN")
	if token == "" {
		token, _ = LoadToken()
	}
	if token != "" {
		if strings.HasPrefix(token, "cookie:") {
			// Cookie-based auth: token stored as "cookie:<name>=<value>"
			cookiePart := strings.TrimPrefix(token, "cookie:")
			// Append to existing cookies
			existing := req.Header.Get("Cookie")
			if existing != "" {
				req.Header.Set("Cookie", existing+"; "+cookiePart)
			} else {
				req.Header.Set("Cookie", cookiePart)
			}
		} else {
			// Private token: use access_token query parameter
			q := req.URL.Query()
			q.Set("access_token", token)
			req.URL.RawQuery = q.Encode()
		}
	}
	if req.Body != nil && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	base := t.Base
	if base == nil {
		base = http.DefaultTransport
	}
	return base.RoundTrip(req)
}

func NewHTTPClient() *http.Client {
	return &http.Client{
		Transport: &Transport{},
	}
}
