package pr

import (
	"fmt"
	"net/url"

	"github.com/gitlink-org/gitlink-cli/shortcuts/common"
)

func Shortcuts() []*common.Shortcut {
	return []*common.Shortcut{
		{
			Name:        "list",
			Description: "List pull requests",
			Flags: []common.Flag{
				{Name: "state", Short: "s", Usage: "Filter: open, merged, closed", Default: "open"},
				{Name: "page", Short: "p", Usage: "Page number", Default: "1"},
				{Name: "limit", Short: "l", Usage: "Items per page", Default: "20"},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				q := url.Values{}
				q.Set("page", ctx.Arg("page"))
				q.Set("limit", ctx.Arg("limit"))
				if s := ctx.Arg("state"); s != "" {
					q.Set("state", s)
				}
				env, err := ctx.CallAPIWithQuery("GET", ctx.RepoPath()+"/pulls", q)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "create",
			Description: "Create a pull request",
			Flags: []common.Flag{
				{Name: "title", Short: "t", Usage: "PR title", Required: true},
				{Name: "body", Short: "b", Usage: "PR description"},
				{Name: "head", Usage: "Source branch", Required: true},
				{Name: "base", Usage: "Target branch", Default: "master"},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				title, _ := ctx.RequireArg("title")
				head, _ := ctx.RequireArg("head")
				base := ctx.Arg("base")
				if base == "" {
					base = "master"
				}
				payload := map[string]interface{}{
					"title": title,
					"head":  head,
					"base":  base,
				}
				if b := ctx.Arg("body"); b != "" {
					payload["body"] = b
				}
				env, err := ctx.CallAPI("POST", ctx.RepoPath()+"/pulls", payload)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "view",
			Description: "View pull request details",
			Flags: []common.Flag{
				{Name: "id", Short: "i", Usage: "PR number", Required: true},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				id, _ := ctx.RequireArg("id")
				env, err := ctx.CallAPI("GET", fmt.Sprintf("%s/pulls/%s", ctx.RepoPath(), id), nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "merge",
			Description: "Merge a pull request",
			Flags: []common.Flag{
				{Name: "id", Short: "i", Usage: "PR number", Required: true},
				{Name: "method", Short: "m", Usage: "Merge method: merge, rebase, squash", Default: "merge"},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				id, _ := ctx.RequireArg("id")
				method := ctx.Arg("method")
				if method == "" {
					method = "merge"
				}
				payload := map[string]interface{}{
					"do": method,
				}
				env, err := ctx.CallAPI("POST", fmt.Sprintf("%s/pulls/%s/pr_merge", ctx.RepoPath(), id), payload)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "close",
			Description: "Close a pull request",
			Flags: []common.Flag{
				{Name: "id", Short: "i", Usage: "PR number", Required: true},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				id, _ := ctx.RequireArg("id")
				env, err := ctx.CallAPI("POST", fmt.Sprintf("%s/pulls/%s/refuse_merge", ctx.RepoPath(), id), nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "files",
			Description: "List changed files in a pull request",
			Flags: []common.Flag{
				{Name: "id", Short: "i", Usage: "PR number", Required: true},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				id, _ := ctx.RequireArg("id")
				env, err := ctx.CallAPI("GET", fmt.Sprintf("%s/pulls/%s/files", ctx.RepoPath(), id), nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "diff",
			Description: "Show diff for a pull request",
			Flags: []common.Flag{
				{Name: "id", Short: "i", Usage: "PR number", Required: true},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				id, _ := ctx.RequireArg("id")
				env, err := ctx.CallAPI("GET", fmt.Sprintf("%s/pulls/%s/files", ctx.RepoPath(), id), nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
	}
}
