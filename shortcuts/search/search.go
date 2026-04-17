package search

import (
	"net/url"

	"github.com/gitlink-org/gitlink-cli/shortcuts/common"
)

func Shortcuts() []*common.Shortcut {
	return []*common.Shortcut{
		{
			Name:        "repos",
			Description: "Search repositories",
			Flags: []common.Flag{
				{Name: "keyword", Short: "k", Usage: "Search keyword", Required: true},
				{Name: "page", Short: "p", Usage: "Page number", Default: "1"},
				{Name: "limit", Short: "l", Usage: "Items per page", Default: "20"},
			},
			Run: func(ctx *common.RuntimeContext) error {
				keyword, _ := ctx.RequireArg("keyword")
				q := url.Values{}
				q.Set("search", keyword)
				q.Set("page", ctx.Arg("page"))
				q.Set("limit", ctx.Arg("limit"))
				env, err := ctx.CallAPIWithQuery("GET", "/projects", q)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "users",
			Description: "Search users",
			Flags: []common.Flag{
				{Name: "keyword", Short: "k", Usage: "Search keyword", Required: true},
				{Name: "page", Short: "p", Usage: "Page number", Default: "1"},
				{Name: "limit", Short: "l", Usage: "Items per page", Default: "20"},
			},
			Run: func(ctx *common.RuntimeContext) error {
				keyword, _ := ctx.RequireArg("keyword")
				q := url.Values{}
				q.Set("search", keyword)
				q.Set("page", ctx.Arg("page"))
				q.Set("limit", ctx.Arg("limit"))
				env, err := ctx.CallAPIWithQuery("GET", "/users/list", q)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
	}
}
