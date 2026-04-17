package user

import (
	"fmt"

	"github.com/gitlink-org/gitlink-cli/shortcuts/common"
)

func Shortcuts() []*common.Shortcut {
	return []*common.Shortcut{
		{
			Name:        "me",
			Description: "Show current authenticated user",
			Run: func(ctx *common.RuntimeContext) error {
				env, err := ctx.CallAPI("GET", "/users/me", nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
		{
			Name:        "info",
			Description: "Show user profile",
			Flags: []common.Flag{
				{Name: "login", Short: "l", Usage: "User login name", Required: true},
			},
			Run: func(ctx *common.RuntimeContext) error {
				login, err := ctx.RequireArg("login")
				if err != nil {
					return err
				}
				env, err := ctx.CallAPI("GET", fmt.Sprintf("/users/%s", login), nil)
				if err != nil {
					return err
				}
				return ctx.Output(env)
			},
		},
	}
}
