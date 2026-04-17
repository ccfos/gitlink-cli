package ci

import (
	"fmt"
	"net/url"

	"github.com/gitlink-org/gitlink-cli/shortcuts/common"
)

func Shortcuts() []*common.Shortcut {
	return []*common.Shortcut{
		{
			Name:        "builds",
			Description: "List CI builds",
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
				env, err := ctx.CallAPIWithQuery("GET", ctx.RepoPath()+"/builds", q)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "logs",
			Description: "View build logs",
			Flags: []common.Flag{
				{Name: "build", Short: "b", Usage: "Build number", Required: true},
				{Name: "stage", Short: "s", Usage: "Stage number", Default: "1"},
				{Name: "step", Usage: "Step number", Default: "1"},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				build, _ := ctx.RequireArg("build")
				stage := ctx.Arg("stage")
				step := ctx.Arg("step")
				if stage == "" {
					stage = "1"
				}
				if step == "" {
					step = "1"
				}
				env, err := ctx.CallAPI("GET", fmt.Sprintf("%s/builds/%s/logs/%s/%s", ctx.RepoPath(), build, stage, step), nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "restart",
			Description: "Restart a build",
			Flags: []common.Flag{
				{Name: "build", Short: "b", Usage: "Build number", Required: true},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				build, _ := ctx.RequireArg("build")
				env, err := ctx.CallAPI("POST", fmt.Sprintf("%s/builds/%s/restart", ctx.RepoPath(), build), nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "stop",
			Description: "Stop a build",
			Flags: []common.Flag{
				{Name: "build", Short: "b", Usage: "Build number", Required: true},
			},
			Run: func(ctx *common.RuntimeContext) error {
				if err := ctx.ResolveOwnerRepo(); err != nil {
					return err
				}
				build, _ := ctx.RequireArg("build")
				env, err := ctx.CallAPI("DELETE", fmt.Sprintf("%s/builds/%s/stop", ctx.RepoPath(), build), nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
	}
}
