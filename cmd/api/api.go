package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"

	"github.com/gitlink-org/gitlink-cli/cmd/cmdutil"
	"github.com/gitlink-org/gitlink-cli/internal/client"
	"github.com/gitlink-org/gitlink-cli/internal/output"
)

func NewAPICmd() *cobra.Command {
	apiCmd := &cobra.Command{
		Use:   "api <METHOD> <PATH>",
		Short: "Make raw API requests to GitLink",
		Long:  `Send arbitrary HTTP requests to the GitLink API. Authentication is injected automatically.`,
		Example: `  gitlink-cli api GET /users/me
  gitlink-cli api GET /projects --query 'page=1&limit=10'
  gitlink-cli api POST /:owner/:repo/issues --body '{"subject":"Bug","description":"..."}'`,
		Args: cobra.ExactArgs(2),
		RunE: runAPI,
	}

	apiCmd.Flags().String("body", "", "Request body (JSON string)")
	apiCmd.Flags().String("query", "", "Query parameters (key=val&key2=val2)")
	apiCmd.Flags().StringSlice("header", nil, "Additional headers (key:value)")

	return apiCmd
}

func runAPI(c *cobra.Command, args []string) error {
	method := strings.ToUpper(args[0])
	path := args[1]

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	cli, err := client.New()
	if err != nil {
		return err
	}
	cli.Debug = cmdutil.Debug

	var body interface{}
	bodyStr, _ := c.Flags().GetString("body")
	if bodyStr != "" {
		if err := json.Unmarshal([]byte(bodyStr), &body); err != nil {
			return fmt.Errorf("invalid JSON body: %w", err)
		}
	}

	var query url.Values
	queryStr, _ := c.Flags().GetString("query")
	if queryStr != "" {
		var err error
		query, err = url.ParseQuery(queryStr)
		if err != nil {
			return fmt.Errorf("invalid query string: %w", err)
		}
	}

	env, err := cli.Do(method, path, body, query)
	if err != nil {
		if apiErr, ok := err.(*client.APIError); ok {
			errEnv := output.ErrorEnvelope(apiErr.Code, apiErr.Message, "")
			return output.Print(errEnv, resolveFormat())
		}
		return err
	}

	return output.Print(env, resolveFormat())
}

func resolveFormat() string {
	f := cmdutil.Format
	if f == "" {
		return "json"
	}
	return f
}
