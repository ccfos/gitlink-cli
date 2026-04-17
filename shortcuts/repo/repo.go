package repo

import (
	"fmt"
	"net/url"

	"github.com/gitlink-org/gitlink-cli/shortcuts/common"
)

func Shortcuts() []*common.Shortcut {
	return []*common.Shortcut{
		{
			Name:        "list",
			Description: "List repositories for a user or organization",
			Flags: []common.Flag{
				{Name: "user", Short: "u", Usage: "User login (default: current user)"},
				{Name: "category", Short: "c", Usage: "Filter: manage/mirror/sync/fork/all (default: manage)", Default: "manage"},
				{Name: "page", Short: "p", Usage: "Page number", Default: "1"},
				{Name: "limit", Short: "l", Usage: "Items per page", Default: "20"},
			},
			Run: func(ctx *common.RuntimeContext) error {
				user := ctx.Arg("user")
				q := url.Values{}
				q.Set("page", ctx.Arg("page"))
				q.Set("limit", ctx.Arg("limit"))
				if cat := ctx.Arg("category"); cat != "" && cat != "all" {
					q.Set("category", cat)
				}

				path := "/projects"
				if user != "" {
					path = fmt.Sprintf("/users/%s/projects", user)
				}
				env, err := ctx.CallAPIWithQuery("GET", path, q)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "info",
			Description: "Show repository details",
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				env, err := ctx.CallAPI("GET", ctx.RepoPath(), nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "create",
			Description: "Create a new repository",
			Flags: []common.Flag{
				{Name: "name", Short: "n", Usage: "Repository name", Required: true},
				{Name: "description", Short: "d", Usage: "Repository description"},
				{Name: "private", Usage: "Make repository private (true/false)", Default: "false"},
			},
			Run: func(ctx *common.RuntimeContext) error {
				name, err := ctx.RequireArg("name")
				if err != nil {
					return err
				}
				// Get current user login for the create path
				userEnv, err := ctx.CallAPI("GET", "/users/me", nil)
				if err != nil {
					return fmt.Errorf("failed to get current user: %w", err)
				}
				userData, _ := userEnv.Data.(map[string]interface{})
				login, _ := userData["login"].(string)
				if login == "" {
					return fmt.Errorf("cannot determine current user login")
				}
				userID, _ := userData["user_id"].(float64)
				body := map[string]interface{}{
					"name":            name,
					"repository_name": name,
					"user_id":         int(userID),
				}
				if desc := ctx.Arg("description"); desc != "" {
					body["description"] = desc
				}
				if ctx.Arg("private") == "true" {
					body["private"] = true
				}
				env, err := ctx.CallAPI("POST", fmt.Sprintf("/%s/%s", login, name), body)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "fork",
			Description: "Fork a repository",
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				env, err := ctx.CallAPI("POST", ctx.RepoPath()+"/forks", nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "delete",
			Description: "Delete a repository",
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				env, err := ctx.CallAPI("DELETE", ctx.RepoPath(), nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
	}
}
