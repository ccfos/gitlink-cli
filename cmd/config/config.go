package config

import (
	"fmt"

	"github.com/spf13/cobra"

	internalConfig "github.com/gitlink-org/gitlink-cli/internal/config"
)

func NewConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage gitlink-cli configuration",
	}
	cmd.AddCommand(newInitCmd())
	cmd.AddCommand(newSetCmd())
	cmd.AddCommand(newGetCmd())
	cmd.AddCommand(newListCmd())
	return cmd
}

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := internalConfig.DefaultConfig()
			if err := internalConfig.Save(cfg); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}
			fmt.Printf("✓ Config initialized at %s\n", internalConfig.ConfigPath())
			return nil
		},
	}
}

func newSetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a configuration value",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := internalConfig.Set(args[0], args[1]); err != nil {
				return err
			}
			fmt.Printf("✓ %s = %s\n", args[0], args[1])
			return nil
		},
	}
}

func newGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <key>",
		Short: "Get a configuration value",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			val, err := internalConfig.Get(args[0])
			if err != nil {
				return err
			}
			if val == "" {
				fmt.Printf("%s: (not set)\n", args[0])
			} else {
				fmt.Printf("%s: %s\n", args[0], val)
			}
			return nil
		},
	}
}

func newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all configuration values",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := internalConfig.Load()
			if err != nil {
				return err
			}
			fmt.Printf("base_url:       %s\n", cfg.BaseURL)
			fmt.Printf("default_format: %s\n", cfg.Format)
			fmt.Printf("editor:         %s\n", cfg.Editor)
			fmt.Printf("pager:          %s\n", cfg.Pager)
			fmt.Printf("\nConfig file: %s\n", internalConfig.ConfigPath())
			return nil
		},
	}
}
