package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gitlink-org/gitlink-cli/cmd/cmdutil"
	authCmd "github.com/gitlink-org/gitlink-cli/cmd/auth"
	apiCmd "github.com/gitlink-org/gitlink-cli/cmd/api"
	configCmd "github.com/gitlink-org/gitlink-cli/cmd/config"
	"github.com/gitlink-org/gitlink-cli/shortcuts"
)

var Version = "dev"

var rootCmd = &cobra.Command{
	Use:   "gitlink-cli",
	Short: "GitLink CLI — command-line tool for gitlink.org.cn",
	Long:  `gitlink-cli is a command-line interface for the GitLink (确实开源) platform, providing repository management, issue tracking, pull requests, CI/CD, and AI-powered workflows.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cmdutil.Owner, "owner", "", "Repository owner (auto-detected from git remote)")
	rootCmd.PersistentFlags().StringVar(&cmdutil.Repo, "repo", "", "Repository name (auto-detected from git remote)")
	rootCmd.PersistentFlags().StringVar(&cmdutil.Format, "format", "", "Output format: json, table, yaml (default: table)")
	rootCmd.PersistentFlags().BoolVar(&cmdutil.Debug, "debug", false, "Enable debug output")

	rootCmd.AddCommand(authCmd.NewAuthCmd())
	rootCmd.AddCommand(apiCmd.NewAPICmd())
	rootCmd.AddCommand(configCmd.NewConfigCmd())
	rootCmd.AddCommand(versionCmd)

	shortcuts.RegisterAll(rootCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gitlink-cli %s\n", Version)
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}
