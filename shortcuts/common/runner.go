package common

import (
	"github.com/spf13/cobra"
)

// MountShortcut converts a Shortcut into a cobra.Command and adds it as a subcommand.
func MountShortcut(parent *cobra.Command, s *Shortcut) {
	cmd := &cobra.Command{
		Use:   "+" + s.Name,
		Short: s.Description,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Collect flag values
			flagValues := make(map[string]string)
			for _, f := range s.Flags {
				val, _ := cmd.Flags().GetString(f.Name)
				if val != "" {
					flagValues[f.Name] = val
				}
			}

			ctx, err := NewRuntimeContext(flagValues)
			if err != nil {
				return err
			}

			return s.Run(ctx)
		},
	}

	for _, f := range s.Flags {
		if f.Short != "" {
			cmd.Flags().StringP(f.Name, f.Short, f.Default, f.Usage)
		} else {
			cmd.Flags().String(f.Name, f.Default, f.Usage)
		}
		if f.Required {
			cmd.MarkFlagRequired(f.Name)
		}
	}

	parent.AddCommand(cmd)
}

// MountShortcuts mounts multiple shortcuts under a parent command.
func MountShortcuts(parent *cobra.Command, shortcuts []*Shortcut) {
	for _, s := range shortcuts {
		MountShortcut(parent, s)
	}
}
