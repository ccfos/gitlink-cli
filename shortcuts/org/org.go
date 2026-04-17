package org

import (
	"fmt"
	"net/url"

	"github.com/gitlink-org/gitlink-cli/shortcuts/common"
)

func Shortcuts() []*common.Shortcut {
	return []*common.Shortcut{
		{
			Name:        "list",
			Description: "List organizations",
			Flags: []common.Flag{
				{Name: "page", Short: "p", Usage: "Page number", Default: "1"},
				{Name: "limit", Short: "l", Usage: "Items per page", Default: "20"},
			},
			Run: func(ctx *common.RuntimeContext) error {
				q := url.Values{}
				q.Set("page", ctx.Arg("page"))
				q.Set("limit", ctx.Arg("limit"))
				env, err := ctx.CallAPIWithQuery("GET", "/organizations", q)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "info",
			Description: "Show organization details",
			Flags: []common.Flag{
				{Name: "id", Short: "i", Usage: "Organization ID or login", Required: true},
			},
			Run: func(ctx *common.RuntimeContext) error {
				id, _ := ctx.RequireArg("id")
				env, err := ctx.CallAPI("GET", fmt.Sprintf("/organizations/%s", id), nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "members",
			Description: "List organization members",
			Flags: []common.Flag{
				{Name: "id", Short: "i", Usage: "Organization ID", Required: true},
				{Name: "page", Short: "p", Usage: "Page number", Default: "1"},
				{Name: "limit", Short: "l", Usage: "Items per page", Default: "20"},
			},
			Run: func(ctx *common.RuntimeContext) error {
				id, _ := ctx.RequireArg("id")
				q := url.Values{}
				q.Set("page", ctx.Arg("page"))
				q.Set("limit", ctx.Arg("limit"))
				env, err := ctx.CallAPIWithQuery("GET", fmt.Sprintf("/organizations/%s/organization_users", id), q)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "create",
			Description: "Create an organization",
			Flags: []common.Flag{
				{Name: "name", Short: "n", Usage: "Organization name", Required: true},
				{Name: "description", Short: "d", Usage: "Description"},
			},
			Run: func(ctx *common.RuntimeContext) error {
				name, _ := ctx.RequireArg("name")
				payload := map[string]interface{}{
					"name": name,
				}
				if d := ctx.Arg("description"); d != "" {
					payload["description"] = d
				}
				env, err := ctx.CallAPI("POST", "/organizations", payload)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
	}
}
