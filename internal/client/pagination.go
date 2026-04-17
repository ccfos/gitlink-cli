package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// PaginateAll fetches all pages and returns combined results.
func (c *Client) PaginateAll(path string, params url.Values) ([]json.RawMessage, error) {
	if params == nil {
		params = url.Values{}
	}
	if params.Get("limit") == "" {
		params.Set("limit", "50")
	}

	var all []json.RawMessage
	page := 1

	for {
		params.Set("page", strconv.Itoa(page))
		env, err := c.Get(path, params)
		if err != nil {
			return nil, err
		}

		if !env.OK {
			return nil, fmt.Errorf("API error on page %d", page)
		}

		// Try to extract array from data
		var items []json.RawMessage
		switch data := env.Data.(type) {
		case []interface{}:
			for _, item := range data {
				raw, _ := json.Marshal(item)
				items = append(items, raw)
			}
		case map[string]interface{}:
			// Some endpoints wrap in {"data": [...], "total_count": N}
			if arr, ok := data["data"]; ok {
				if slice, ok := arr.([]interface{}); ok {
					for _, item := range slice {
						raw, _ := json.Marshal(item)
						items = append(items, raw)
					}
				}
			} else {
				// Single object, not paginated
				raw, _ := json.Marshal(data)
				return []json.RawMessage{raw}, nil
			}
		}

		if len(items) == 0 {
			break
		}

		all = append(all, items...)

		// Check if we got fewer items than limit
		limit, _ := strconv.Atoi(params.Get("limit"))
		if len(items) < limit {
			break
		}

		page++
	}

	return all, nil
}
