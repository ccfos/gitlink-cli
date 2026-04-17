package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gitlink-org/gitlink-cli/internal/auth"
	"github.com/gitlink-org/gitlink-cli/internal/config"
	"github.com/gitlink-org/gitlink-cli/internal/output"
)

type Client struct {
	HTTP    *http.Client
	BaseURL string
	Debug   bool
}

type APIError struct {
	StatusCode int
	Code       interface{}
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("[%v] %s", e.Code, e.Message)
}

func New() (*Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	return &Client{
		HTTP:    auth.NewHTTPClient(),
		BaseURL: cfg.BaseURL,
	}, nil
}

func (c *Client) Do(method, path string, body interface{}, query url.Values) (*output.Envelope, error) {
	// Append .json suffix if not already present (GitLink API convention)
	// Handle paths that may already contain query strings (e.g., /path?key=val)
	if idx := strings.Index(path, "?"); idx != -1 {
		basePath := path[:idx]
		queryStr := path[idx:]
		if !strings.HasSuffix(basePath, ".json") {
			path = basePath + ".json" + queryStr
		}
	} else if !strings.HasSuffix(path, ".json") {
		path += ".json"
	}
	fullURL := c.BaseURL + path
	if query != nil && len(query) > 0 {
		sep := "?"
		if strings.Contains(fullURL, "?") {
			sep = "&"
		}
		fullURL += sep + query.Encode()
	}

	// Replace path params
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, fullURL, bodyReader)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		fmt.Printf("→ %s %s\n", method, fullURL)
	}

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if c.Debug {
		fmt.Printf("← %d %s\n", resp.StatusCode, string(respData[:min(len(respData), 200)]))
	}

	// Check HTTP-level errors
	if resp.StatusCode >= 400 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Code:       resp.StatusCode,
			Message:    fmt.Sprintf("HTTP %d: %s", resp.StatusCode, strings.TrimSpace(string(respData))),
		}
	}

	// Parse JSON
	var raw map[string]interface{}
	if err := json.Unmarshal(respData, &raw); err != nil {
		// Not JSON, return as-is
		return output.SuccessEnvelope(string(respData), nil), nil
	}

	// Check GitLink error-in-body pattern
	if status, ok := raw["status"]; ok {
		var statusCode float64
		switch v := status.(type) {
		case float64:
			statusCode = v
		case int:
			statusCode = float64(v)
		}
		if statusCode != 0 && statusCode != 200 && statusCode != 1 {
			msg, _ := raw["message"].(string)
			suggestion := suggestFix(int(statusCode))
			return output.ErrorEnvelope(int(statusCode), msg, suggestion), &APIError{
				StatusCode: int(statusCode),
				Code:       int(statusCode),
				Message:    msg,
			}
		}
	}

	// Auto-parse JSON string data (GitLink API quirk: some endpoints return data as JSON string)
	if dataStr, ok := raw["data"].(string); ok {
		var parsedData interface{}
		if err := json.Unmarshal([]byte(dataStr), &parsedData); err == nil {
			raw["data"] = json.RawMessage(dataStr)
		}
	}

	// Build meta from pagination info
	var meta *output.Meta
	if tc, ok := raw["total_count"]; ok {
		meta = &output.Meta{}
		if v, ok := tc.(float64); ok {
			meta.TotalCount = int(v)
		}
		if v, ok := raw["page"].(float64); ok {
			meta.Page = int(v)
		}
		if v, ok := raw["limit"].(float64); ok {
			meta.Limit = int(v)
		}
	}

	return output.SuccessEnvelope(raw, meta), nil
}

func (c *Client) Get(path string, query url.Values) (*output.Envelope, error) {
	return c.Do("GET", path, nil, query)
}

func (c *Client) Post(path string, body interface{}) (*output.Envelope, error) {
	return c.Do("POST", path, body, nil)
}

func (c *Client) Put(path string, body interface{}) (*output.Envelope, error) {
	return c.Do("PUT", path, body, nil)
}

func (c *Client) Delete(path string, query url.Values) (*output.Envelope, error) {
	return c.Do("DELETE", path, nil, query)
}

func suggestFix(code int) string {
	switch code {
	case 401:
		return "请先运行 gitlink-cli auth login 登录"
	case 403:
		return "权限不足，请确认账户权限或联系项目管理员"
	case 404:
		return "资源不存在，请检查 owner/repo/id 是否正确"
	case 422:
		return "参数校验失败，请检查请求参数"
	default:
		return ""
	}
}
