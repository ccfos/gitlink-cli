package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/gitlink-org/gitlink-cli/internal/config"
)

type LoginResult struct {
	Username string `json:"username"`
	Login    string `json:"login"`
	UserID   int    `json:"user_id"`
	Token    string `json:"token"`
	// Error fields (only present on failure)
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Login authenticates with username/password and stores the session cookie.
func Login(username, password string) (*LoginResult, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	body := map[string]string{
		"login":    username,
		"password": password,
	}
	bodyJSON, _ := json.Marshal(body)

	loginURL := cfg.BaseURL + "/accounts/login.json"
	req, err := http.NewRequest("POST", loginURL, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")

	// Use a client that captures cookies across redirects
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result LoginResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Check for error response (has negative status)
	if result.Status < 0 {
		return nil, fmt.Errorf("%s", result.Message)
	}

	// Collect auth cookies from response (GitLink uses autologin_trustie for session persistence)
	var authCookies []string
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "autologin_trustie" || cookie.Name == "_educoder_session" {
			authCookies = append(authCookies, cookie.Name+"="+cookie.Value)
		}
	}
	// Also check cookie jar (captures cookies from redirect hops)
	if len(authCookies) == 0 {
		if u, err := url.Parse(loginURL); err == nil {
			for _, cookie := range jar.Cookies(u) {
				if cookie.Name == "autologin_trustie" || cookie.Name == "_educoder_session" {
					authCookies = append(authCookies, cookie.Name+"="+cookie.Value)
				}
			}
		}
	}

	if len(authCookies) == 0 {
		return nil, fmt.Errorf("login succeeded but no auth cookies received")
	}

	// Store as "cookie:<name>=<value>; <name>=<value>" format
	// Transport will send these as Cookie header
	tokenValue := "cookie:" + strings.Join(authCookies, "; ")

	if err := StoreToken(tokenValue); err != nil {
		return nil, fmt.Errorf("failed to store credentials: %w", err)
	}

	// Verify the stored credentials actually work
	if _, verifyErr := GetCurrentUser(); verifyErr != nil {
		// Clean up the bad token
		_ = DeleteToken()
		return nil, fmt.Errorf("login failed: credentials not accepted by API (%v)", verifyErr)
	}

	return &result, nil
}

// GetCurrentUser fetches the authenticated user info.
// Returns error if not authenticated or token is invalid.
func GetCurrentUser() (map[string]interface{}, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	client := NewHTTPClient()
	resp, err := client.Get(cfg.BaseURL + "/users/me.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(data))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	// GitLink returns {"status": -1, "message": "..."} for auth errors with HTTP 200
	if status, ok := result["status"].(float64); ok && status < 0 {
		msg, _ := result["message"].(string)
		return nil, fmt.Errorf("%s", msg)
	}

	// Verify we got actual user data
	if _, ok := result["login"]; !ok {
		return nil, fmt.Errorf("invalid response: missing login field")
	}

	return result, nil
}
