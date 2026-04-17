package branch

import (
	"fmt"
	"net/url"

	"github.com/gitlink-org/gitlink-cli/shortcuts/common"
)

func Shortcuts() []*common.Shortcut {
	return []*common.Shortcut{
		{
			Name:        "list",
			Description: "List branches",
			Flags: []common.Flag{
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
				env, err := ctx.CallAPIWithQuery("GET", "/v1"+ctx.RepoPath()+"/branches", q)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "create",
			Description: "Create a branch",
			Flags: []common.Flag{
				{Name: "name", Short: "n", Usage: "Branch name", Required: true},
				{Name: "from", Short: "f", Usage: "Source branch or commit", Default: "master"},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				name, _ := ctx.RequireArg("name")
				from := ctx.Arg("from")
				if from == "" {
					from = "master"
				}
				payload := map[string]interface{}{
					"new_branch_name": name,
					"old_branch_name": from,
				}
				env, err := ctx.CallAPI("POST", "/v1"+ctx.RepoPath()+"/branches", payload)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "delete",
			Description: "Delete a branch",
			Flags: []common.Flag{
				{Name: "name", Short: "n", Usage: "Branch name", Required: true},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				name, _ := ctx.RequireArg("name")
				payload := map[string]interface{}{
					"branch_name": name,
				}
				env, err := ctx.CallAPI("POST", "/v1"+ctx.RepoPath()+"/branches/delete", payload)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "protect",
			Description: "Set branch protection",
			Flags: []common.Flag{
				{Name: "name", Short: "n", Usage: "Branch name", Required: true},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				name, _ := ctx.RequireArg("name")
				payload := map[string]interface{}{
					"branch_name": name,
				}
				env, err := ctx.CallAPI("POST", ctx.RepoPath()+"/protected_branches", payload)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "unprotect",
			Description: "Remove branch protection",
			Flags: []common.Flag{
				{Name: "name", Short: "n", Usage: "Branch name", Required: true},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				name, _ := ctx.RequireArg("name")
				env, err := ctx.CallAPI("DELETE", fmt.Sprintf("%s/protected_branches/%s", ctx.RepoPath(), name), nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
	}
}
