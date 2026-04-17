# gitlink-cli

[![GitLink](https://img.shields.io/badge/GitLink-Gitlink%2Fgitlink--cli-green)](https://www.gitlink.org.cn/Gitlink/gitlink-cli)
[![License](https://img.shields.io/badge/License-MulanPSL--2.0-blue.svg)](https://license.coscl.org.cn/MulanPSL2)
[![Go Version](https://img.shields.io/badge/Go-1.26%2B-blue.svg)](https://golang.org)
[![npm version](https://img.shields.io/npm/v/@gitlink-ai/cli.svg)](https://www.npmjs.com/package/@gitlink-ai/cli)

The official [GitLink](https://www.gitlink.org.cn) CLI tool — built for humans and AI Agents. Supports **macOS, Linux, and Windows**. Covers repository management, issue tracking, pull requests, CI/CD, and AI-powered workflows, with 40+ commands and 11 AI Agent [Skills](./skills/).

**[中文文档](./README.zh-CN.md)**

[Install](#installation--quick-start) · [AI Agent Skills](#ai-agent-skills) · [Auth](#configure--use) · [Commands](#usage-examples) · [Contributing](#related-projects)

## Why gitlink-cli?

- **Agent-Native Design** — 11 structured [Skills](./skills/) out of the box, compatible with Claude Code — Agents can operate GitLink with zero extra setup
- **Wide Coverage** — Repository, Issue, PR, Branch, Release, CI, Org, Search, User — all core domains covered
- **AI-Friendly & Optimized** — Every command is tested with real Agents, featuring concise parameters, smart defaults, and structured output
- **Cross-Platform** — Runs on macOS, Linux, and Windows (x64/arm64), install via `npm` in one command
- **Open Source, Zero Barriers** — MulanPSL-2.0 license, ready to use, just `npm install`
- **Up and Running in 3 Minutes** — Interactive login or `GITLINK_TOKEN` env var, from install to first API call in just 3 steps
- **Secure & Controllable** — OS-native keychain credential storage, `GITLINK_TOKEN` env var for CI/CD & non-interactive environments, auto git remote context resolution
- **Three-Layer Architecture** — Shortcuts (human & AI friendly) → Raw API (full coverage) → Config (configuration management)

## Features

| Category | Capabilities |
|----------|-------------|
| 📦 Repo | List, create, fork, delete repositories, view repo info |
| 🐛 Issue | Create, update, close, comment on issues |
| 🔀 PR | Create, merge, review pull requests, view changed files |
| 🌿 Branch | Create, delete, protect branches |
| 🏷️ Release | Create, view, delete releases |
| 🏢 Org | Manage organizations, members, teams |
| 🔧 CI | View builds, logs, CI/CD operations |
| 🔍 Search | Search repositories, users |
| 👤 User | View user profiles and info |
| 📋 PM | Sprint management, kanban boards, weekly reports |
| 🤖 Workflow | AI-powered issue triage, PR review, release notes |

## Installation & Quick Start

### Requirements

- Node.js 14+ (`npm`/`npx`) — for npm installation
- Supported platforms: macOS, Linux, Windows (x64/arm64)
- Go 1.26+ — only required for building from source

### Quick Start (Human Users)

> **Note for AI assistants:** If you are an AI Agent helping the user with installation, jump directly to [Quick Start (AI Agent)](#quick-start-ai-agent), which contains all the steps you need to complete.

#### Install

Choose **one** of the following methods:

**Option 1 — From npm (recommended):**

```bash
# Install CLI
npm install -g @gitlink-ai/cli

# Install CLI Skills (required, works on all platforms)
gitlink-cli-install-skills

# Or install Skills with npx
npx skills add ccfos/gitlink-cli/skills -y -g
```

**Option 2 — From source:**

Requires Go 1.26+.

```bash
git clone https://www.gitlink.org.cn/Gitlink/gitlink-cli.git
cd gitlink-cli
make install

# Install CLI Skills (required)
npx skills add ./skills -y -g
```

> **Windows users:** Run `npm install -g @gitlink-ai/cli` in PowerShell or CMD. For building from source, use `go install .` instead of `make install`.

#### Configure & Use

```bash
# 1. Configure (one-time, interactive guided setup)
gitlink-cli config init

# 2. Log in (choose one)
gitlink-cli auth login            # Username/password (recommended)
gitlink-cli auth login --token    # Or paste a private token
export GITLINK_TOKEN="your-token" # Or set env var (for CI/CD, non-interactive environments)

# 3. Start using
gitlink-cli repo +list
```

### Quick Start (AI Agent)

> The following steps are for AI Agents. Some steps require the user to complete actions in a browser.

**Step 1 — Install**

```bash
# Install CLI
npm install -g @gitlink-ai/cli

# Install CLI Skills (required, works on all platforms)
gitlink-cli-install-skills
```

**Step 2 — Configure**

```bash
gitlink-cli config init
```

**Step 3 — Login**

For interactive environments:
```bash
gitlink-cli auth login
```

For non-interactive environments (CI/CD, Trae sandbox, MCP, etc.):
```bash
export GITLINK_TOKEN="your-private-token"
```

> To get a private token, go to GitLink web → Settings → Private Tokens.

**Step 4 — Verify**

```bash
gitlink-cli user +me
```

## Usage Examples

### Repository Operations

```bash
# List repositories
gitlink-cli repo +list

# View repository info
gitlink-cli repo +info --owner Gitlink --repo forgeplus

# Create a repository
gitlink-cli repo +create -n my-project -d "Project description"

# Fork a repository
gitlink-cli repo +fork --owner Gitlink --repo forgeplus
```

### Issue Management

```bash
# List issues
gitlink-cli issue +list --owner Gitlink --repo forgeplus

# Create an issue
gitlink-cli issue +create --owner Gitlink --repo forgeplus -t "Bug: Login failed" -b "Steps to reproduce..."

# View an issue
gitlink-cli issue +view --owner Gitlink --repo forgeplus -i 123

# Close an issue
gitlink-cli issue +close --owner Gitlink --repo forgeplus -i 123

# Add a comment
gitlink-cli issue +comment --owner Gitlink --repo forgeplus -i 123 -b "Fixed"
```

### Pull Requests

```bash
# List PRs
gitlink-cli pr +list --owner Gitlink --repo forgeplus

# Create a PR (same-repo branch)
gitlink-cli pr +create --owner Gitlink --repo forgeplus -t "feat: Search feature" --head feature/search --base master

# Create a PR (from a fork)
gitlink-cli pr +create --owner Gitlink --repo forgeplus -t "feat: New feature" --head your_username/forgeplus:feature/my-feature --base master

# View a PR
gitlink-cli pr +view --owner Gitlink --repo forgeplus -i 42

# Merge a PR
gitlink-cli pr +merge --owner Gitlink --repo forgeplus -i 42

# View changed files
gitlink-cli pr +files --owner Gitlink --repo forgeplus -i 42
```

### Release Management

```bash
# List releases
gitlink-cli release +list --owner Gitlink --repo forgeplus

# Create a release
gitlink-cli release +create --owner Gitlink --repo forgeplus -t v1.0.0 -n "v1.0.0 Stable" -b "Changelog..."

# View a release
gitlink-cli release +view --owner Gitlink --repo forgeplus -i <version_id>
```

### Search

```bash
# Search repositories
gitlink-cli search +repos -k "machine learning"

# Search users
gitlink-cli search +users -k "zhangsan"
```

### Raw API

For endpoints not covered by shortcuts, use the Raw API directly:

```bash
# GET request
gitlink-cli api GET /users/me

# POST request
gitlink-cli api POST /Gitlink/forgeplus/issues --body '{"subject":"test","description":"..."}'

# With query parameters
gitlink-cli api GET /Gitlink/forgeplus/commits --query 'page=1&limit=5'
```

## Global Parameters

| Parameter | Description | Example |
|-----------|-------------|---------|
| `--owner` | Repository owner | `--owner Gitlink` |
| `--repo` | Repository name | `--repo forgeplus` |
| `--format` | Output format (json/table/yaml) | `--format json` |
| `--debug` | Enable debug output | `--debug` |

**Automatic context resolution:** When running inside a git repository, `--owner` and `--repo` are automatically resolved from `git remote origin`.

## Branch Conventions

gitlink-cli supports bidirectional code sync between GitHub and GitLink:

| Platform | Default Branch |
|----------|---------------|
| GitHub | `main` |
| GitLink | `master` |

**Push to GitLink from local:**

```bash
# Method 1: Use git command directly
git push gitlink main:master

# Method 2: Configure git remote
git config remote.gitlink.push refs/heads/main:refs/heads/master
git push gitlink
```

## AI Agent Skills

The `skills/` directory contains 11 Claude Code Agent Skill files for AI-automated GitLink operations.

See [skills/README.md](skills/README.md) for details.

| Skill | Description |
|-------|-------------|
| `gitlink-shared` | Authentication, global parameters, safety rules, API notes |
| `gitlink-repo` | Repository operations (create, view, delete, fork, etc.) |
| `gitlink-issue` | Issue operations (create, update, close, comment, etc.) |
| `gitlink-pr` | Pull request operations (create, merge, review, etc.) |
| `gitlink-release` | Release management (create, view, delete, etc.) |
| `gitlink-org` | Organization management (members, teams, etc.) |
| `gitlink-ci` | CI/CD operations (builds, logs, etc.) |
| `gitlink-search` | Search (repositories, users, etc.) |
| `gitlink-user` | User management (profile info, etc.) |
| `gitlink-pm` | Project management (sprints, kanban, weekly reports, etc.) |
| `gitlink-workflow` | AI-powered workflows (issue triage, PR review, release notes, etc.) |

## Project Structure

```
gitlink-cli/
├── cmd/                      # Cobra command definitions
│   ├── root.go               # Root command + global flags
│   ├── auth/                 # Authentication commands
│   ├── api/                  # Raw API commands
│   ├── config/               # Configuration commands
│   └── cmdutil/              # Global utilities
├── internal/                 # Internal packages
│   ├── auth/                 # Login, token storage, transport
│   ├── client/               # HTTP client + pagination
│   ├── config/               # Config file management
│   ├── context/              # Git remote resolution
│   └── output/               # Envelope + formatter
├── shortcuts/                # Shortcut implementations
│   ├── common/               # Framework (types, runner)
│   ├── repo/                 # Repository shortcuts
│   ├── issue/                # Issue shortcuts
│   ├── pr/                   # PR shortcuts
│   ├── branch/               # Branch shortcuts
│   ├── release/              # Release shortcuts
│   ├── org/                  # Organization shortcuts
│   ├── ci/                   # CI shortcuts
│   ├── search/               # Search shortcuts
│   ├── user/                 # User shortcuts
│   └── register.go           # Registration entry point
├── skills/                   # AI Agent Skills
│   ├── README.md             # Skills guide
│   ├── gitlink-shared/       # Shared rules
│   ├── gitlink-repo/         # Repository skill
│   ├── gitlink-issue/        # Issue skill
│   ├── gitlink-pr/           # PR skill
│   ├── gitlink-pm/           # Project management skill
│   └── ...
├── doc/                      # Design documents
│   ├── Design.md
│   ├── CODE_SYNC_STRATEGY_FINAL.md
│   └── ...
├── main.go
├── Makefile
├── go.mod
└── README.md
```

## Documentation

- [Skills Guide](skills/README.md) — AI Agent Skills detailed documentation
- [Design Document](doc/design.md) — Architecture design and development plan

## FAQ

### Q: How do I use gitlink-cli in scripts?

Use the `GITLINK_TOKEN` environment variable + `--format json` for structured output:

```bash
export GITLINK_TOKEN="your-private-token"
gitlink-cli repo +list --format json | jq '.data.projects[] | .name'
```

### Q: How does automatic owner/repo resolution work?

When running inside a git repository, the CLI automatically resolves `--owner` and `--repo` from `git remote origin`:

```bash
cd ~/my-gitlink-project
gitlink-cli issue +list  # Automatically uses the current repository
```

### Q: What if my token expires?

Re-authenticate:

```bash
# Username/password login
gitlink-cli auth login

# Or use a private token (generate at GitLink web → Settings → Private Tokens)
gitlink-cli auth login --token
```

### Q: How do I use gitlink-cli in CI/CD or non-interactive environments (e.g. Trae sandbox)?

Set the `GITLINK_TOKEN` environment variable — no `auth login` needed:

```bash
export GITLINK_TOKEN="your-private-token"
gitlink-cli repo +list   # Ready to use
gitlink-cli auth status   # Shows "✓ Logged in via GITLINK_TOKEN environment variable"
```

Priority: `GITLINK_TOKEN` env var > keyring/file stored token. When the env var is not set, the original interactive login flow works as before.

### Q: Where are credentials stored on Windows?

gitlink-cli uses Windows Credential Manager for secure token storage. If Credential Manager is unavailable, it automatically falls back to file storage (`~/.config/gitlink-cli/credentials`).

### Q: Where can I find the full API reference?

See [skills/gitlink-shared/REFERENCE.md](skills/gitlink-shared/REFERENCE.md).

## License

[MulanPSL-2.0](https://license.coscl.org.cn/MulanPSL2)
