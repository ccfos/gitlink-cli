package common

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gitlink-org/gitlink-cli/cmd/cmdutil"
	"github.com/gitlink-org/gitlink-cli/internal/client"
	"github.com/gitlink-org/gitlink-cli/internal/context"
	"github.com/gitlink-org/gitlink-cli/internal/output"
)

// Shortcut defines a high-level CLI command (e.g., repo +create, issue +list).
type Shortcut struct {
	Name        string
	Description string
	Flags       []Flag
	Run         func(ctx *RuntimeContext) error
}

// Flag defines a command-line flag for a shortcut.
type Flag struct {
	Name     string
	Short    string
	Usage    string
	Required bool
	Default  string
}

// RuntimeContext provides helpers for shortcut implementations.
type RuntimeContext struct {
	Client *client.Client
	Owner  string
	Repo   string
	Format string
	Args   map[string]string
}

// NewRuntimeContext creates a RuntimeContext with auto-resolved owner/repo.
func NewRuntimeContext(args map[string]string) (*RuntimeContext, error) {
	cli, err := client.New()
	if err != nil {
		return nil, err
	}
	cli.Debug = cmdutil.Debug

	format := cmdutil.Format
	if format == "" {
		format = "json"
	}

	return &RuntimeContext{
		Client: cli,
		Owner:  cmdutil.Owner,
		Repo:   cmdutil.Repo,
		Format: format,
		Args:   args,
	}, nil
}

// ResolveOwnerRepo resolves owner and repo from flags or git remote.
func (ctx *RuntimeContext) ResolveOwnerRepo() error {
	owner, repo, err := context.ResolveOwnerRepo(ctx.Owner, ctx.Repo)
	if err != nil {
		return err
	}
	ctx.Owner = owner
	ctx.Repo = repo
	return nil
}

// CallAPI makes an API call and returns the envelope.
func (ctx *RuntimeContext) CallAPI(method, path string, body interface{}) (*output.Envelope, error) {
	return ctx.Client.Do(method, path, body, nil)
}

// CallAPIWithQuery makes an API call with query parameters.
func (ctx *RuntimeContext) CallAPIWithQuery(method, path string, query url.Values) (*output.Envelope, error) {
	return ctx.Client.Do(method, path, nil, query)
}

// PaginateAll fetches all pages.
func (ctx *RuntimeContext) PaginateAll(path string, params url.Values) ([]json.RawMessage, error) {
	return ctx.Client.PaginateAll(path, params)
}

// Output prints the envelope in the configured format.
func (ctx *RuntimeContext) Output(env *output.Envelope) error {
	return output.Print(env, ctx.Format)
}

// OutputData wraps data in a success envelope and prints it.
func (ctx *RuntimeContext) OutputData(data interface{}) error {
	return output.Print(output.SuccessEnvelope(data, nil), ctx.Format)
}

// RepoPath returns the API path prefix for the current owner/repo.
func (ctx *RuntimeContext) RepoPath() string {
	return fmt.Sprintf("/%s/%s", ctx.Owner, ctx.Repo)
}

// Arg returns a flag value, or the default if not set.
func (ctx *RuntimeContext) Arg(name string) string {
	if v, ok := ctx.Args[name]; ok {
		return v
	}
	return ""
}

// RequireArg returns a flag value or an error if not set.
func (ctx *RuntimeContext) RequireArg(name string) (string, error) {
	v := ctx.Arg(name)
	if v == "" {
		return "", fmt.Errorf("required flag --%s is missing", name)
	}
	return v, nil
}
