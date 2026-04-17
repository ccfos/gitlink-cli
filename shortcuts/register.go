package shortcuts

import (
	"github.com/spf13/cobra"

	"github.com/gitlink-org/gitlink-cli/shortcuts/branch"
	"github.com/gitlink-org/gitlink-cli/shortcuts/ci"
	"github.com/gitlink-org/gitlink-cli/shortcuts/common"
	"github.com/gitlink-org/gitlink-cli/shortcuts/issue"
	"github.com/gitlink-org/gitlink-cli/shortcuts/org"
	"github.com/gitlink-org/gitlink-cli/shortcuts/pr"
	"github.com/gitlink-org/gitlink-cli/shortcuts/release"
	"github.com/gitlink-org/gitlink-cli/shortcuts/repo"
	"github.com/gitlink-org/gitlink-cli/shortcuts/search"
	"github.com/gitlink-org/gitlink-cli/shortcuts/user"
)

// RegisterAll mounts all shortcut groups onto the root command.
func RegisterAll(root *cobra.Command) {
	groups := map[string][]*common.Shortcut{
		"repo":    repo.Shortcuts(),
		"issue":   issue.Shortcuts(),
		"pr":      pr.Shortcuts(),
		"release": release.Shortcuts(),
		"branch":  branch.Shortcuts(),
		"org":     org.Shortcuts(),
		"user":    user.Shortcuts(),
		"search":  search.Shortcuts(),
		"ci":      ci.Shortcuts(),
	}

	descriptions := map[string]string{
		"repo":    "Repository operations",
		"issue":   "Issue operations",
		"pr":      "Pull request operations",
		"release": "Release operations",
		"branch":  "Branch operations",
		"org":     "Organization operations",
		"user":    "User operations",
		"search":  "Search operations",
		"ci":      "CI/CD operations",
	}

	for name, shortcuts := range groups {
		groupCmd := &cobra.Command{
			Use:   name,
			Short: descriptions[name],
		}
		common.MountShortcuts(groupCmd, shortcuts)
		root.AddCommand(groupCmd)
	}
}
