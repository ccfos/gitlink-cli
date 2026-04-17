package context

import (
	"fmt"
	"net/url"
	"os/exec"
	"strings"
)

// ResolveOwnerRepo auto-detects owner and repo from git remote origin.
// Returns the explicitly provided values if non-empty.
func ResolveOwnerRepo(flagOwner, flagRepo string) (string, string, error) {
	if flagOwner != "" && flagRepo != "" {
		return flagOwner, flagRepo, nil
	}

	owner, repo, err := fromGitRemote()
	if err != nil {
		if flagOwner == "" || flagRepo == "" {
			return "", "", fmt.Errorf("cannot detect owner/repo from git remote: %w\nUse --owner and --repo flags to specify explicitly", err)
		}
	}

	if flagOwner != "" {
		owner = flagOwner
	}
	if flagRepo != "" {
		repo = flagRepo
	}

	return owner, repo, nil
}

func fromGitRemote() (string, string, error) {
	out, err := exec.Command("git", "remote", "get-url", "origin").Output()
	if err != nil {
		return "", "", fmt.Errorf("not a git repository or no remote 'origin'")
	}
	remote := strings.TrimSpace(string(out))
	return parseRemoteURL(remote)
}

func parseRemoteURL(remote string) (string, string, error) {
	// SSH format: git@www.gitlink.org.cn:owner/repo.git
	if strings.HasPrefix(remote, "git@") {
		parts := strings.SplitN(remote, ":", 2)
		if len(parts) != 2 {
			return "", "", fmt.Errorf("cannot parse SSH remote: %s", remote)
		}
		return parsePathSegments(parts[1])
	}

	// HTTPS format: https://www.gitlink.org.cn/owner/repo.git
	u, err := url.Parse(remote)
	if err != nil {
		return "", "", fmt.Errorf("cannot parse remote URL: %s", remote)
	}
	return parsePathSegments(u.Path)
}

func parsePathSegments(path string) (string, string, error) {
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, ".git")
	parts := strings.SplitN(path, "/", 3)
	if len(parts) < 2 {
		return "", "", fmt.Errorf("cannot extract owner/repo from path: %s", path)
	}
	return parts[0], parts[1], nil
}
