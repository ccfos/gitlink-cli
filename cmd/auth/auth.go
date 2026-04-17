package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	internalAuth "github.com/gitlink-org/gitlink-cli/internal/auth"
)

const envTokenVar = "GITLINK_TOKEN"

func NewAuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authentication commands",
	}
	cmd.AddCommand(newLoginCmd())
	cmd.AddCommand(newLogoutCmd())
	cmd.AddCommand(newStatusCmd())
	return cmd
}

func newLoginCmd() *cobra.Command {
	var tokenMode bool

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to GitLink",
		RunE: func(cmd *cobra.Command, args []string) error {
			if tokenMode {
				return loginWithToken()
			}
			return loginWithPassword()
		},
	}
	cmd.Flags().BoolVar(&tokenMode, "token", false, "Login by pasting an existing token")
	return cmd
}

func loginWithPassword() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username/Email/Phone: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("failed to read password: %w", err)
	}
	fmt.Println()
	password := string(passwordBytes)

	result, err := internalAuth.Login(username, password)
	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	fmt.Printf("✓ Logged in as %s\n", result.Login)
	return nil
}

func loginWithToken() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Paste your token: ")
	token, _ := reader.ReadString('\n')
	token = strings.TrimSpace(token)

	if token == "" {
		return fmt.Errorf("token cannot be empty")
	}

	if err := internalAuth.StoreToken(token); err != nil {
		return fmt.Errorf("failed to store token: %w", err)
	}

	fmt.Println("✓ Token saved")
	return nil
}

func newLogoutCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Logout from GitLink",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := internalAuth.DeleteToken(); err != nil {
				return fmt.Errorf("failed to delete token: %w", err)
			}
			fmt.Println("✓ Logged out")
			return nil
		},
	}
}

func newStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show authentication status",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Check env var token first
			if envToken := os.Getenv(envTokenVar); envToken != "" {
				fmt.Printf("✓ Logged in via %s environment variable\n", envTokenVar)
			}

			token, err := internalAuth.LoadToken()
			if err != nil || token == "" {
				if os.Getenv(envTokenVar) == "" {
					fmt.Println("✗ Not logged in")
					fmt.Println("  Run: gitlink-cli auth login")
					fmt.Printf("  Or set %s environment variable\n", envTokenVar)
				}
				return nil
			}

			user, err := internalAuth.GetCurrentUser()
			if err != nil {
				fmt.Printf("✓ Token stored (but cannot verify: %v)\n", err)
				return nil
			}

			login, _ := user["login"].(string)
			name, _ := user["name"].(string)
			if login != "" {
				fmt.Printf("✓ Logged in as %s", login)
				if name != "" {
					fmt.Printf(" (%s)", name)
				}
				fmt.Println()
			} else {
				fmt.Println("✓ Token stored (user info unavailable)")
			}
			return nil
		},
	}
}
