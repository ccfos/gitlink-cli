# gitlink-cli

[![GitLink](https://img.shields.io/badge/GitLink-Gitlink%2Fgitlink--cli-green)](https://www.gitlink.org.cn/Gitlink/gitlink-cli)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](./LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.26%2B-blue.svg)](https://golang.org)
[![npm version](https://img.shields.io/npm/v/@gitlink-ai/cli.svg)](https://www.npmjs.com/package/@gitlink-ai/cli)

The official [GitLink（确实开源）](https://www.gitlink.org.cn) CLI tool — built for humans and AI Agents. Supports **macOS, Linux, and Windows**. Covers repository management, issue tracking, pull requests, CI/CD, and AI-powered workflows, with 40+ commands and 11 AI Agent [Skills](./skills/).

[Install](#installation--quick-start) · [AI Agent Skills](#ai-agent-skills) · [Auth](#首次使用) · [Commands](#使用示例) · [Contributing](#相关项目)

## Why gitlink-cli?

- **Agent-Native Design** — 11 structured [Skills](./skills/) out of the box, compatible with Claude Code — Agents can operate GitLink with zero extra setup
- **Wide Coverage** — Repository, Issue, PR, Branch, Release, CI, Org, Search, User — all core domains covered
- **AI-Friendly & Optimized** — Every command is tested with real Agents, featuring concise parameters, smart defaults, and structured output
- **Cross-Platform** — Runs on macOS, Linux, and Windows (x64/arm64), install via `npm` in one command
- **Open Source, Zero Barriers** — Apache 2.0 license, ready to use, just `npm install`
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

Before you start, make sure you have:

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

# Install CLI SKILL (required, works on all platforms)
gitlink-cli-install-skills
```

**Option 2 — From source:**

Requires Go 1.26+.

```bash
git clone https://www.gitlink.org.cn/Gitlink/gitlink-cli.git
cd gitlink-cli
make install

# Install CLI SKILL (required)
npx skills add ./skills -y -g
```

> **Windows 用户注意：** 请在 PowerShell 或 CMD 中运行 `npm install -g @gitlink-ai/cli`。从源码构建请使用 `go install .` 代替 `make install`。

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

# Install CLI SKILL (required, works on all platforms)
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
> 获取私人令牌：GitLink 网页端 → 个人设置 → 私人令牌。

**Step 4 — Verify**

```bash
gitlink-cli user +me
```

## 使用示例

### 仓库操作

```bash
# 列出仓库
gitlink-cli repo +list

# 查看仓库信息
gitlink-cli repo +info --owner Gitlink --repo forgeplus

# 创建仓库
gitlink-cli repo +create -n my-project -d "项目描述"

# Fork 仓库
gitlink-cli repo +fork --owner Gitlink --repo forgeplus
```

### Issue 管理

```bash
# 列出 Issue
gitlink-cli issue +list --owner Gitlink --repo forgeplus

# 创建 Issue
gitlink-cli issue +create --owner Gitlink --repo forgeplus -t "Bug: 登录失败" -b "复现步骤..."

# 查看 Issue
gitlink-cli issue +view --owner Gitlink --repo forgeplus -i 123

# 关闭 Issue
gitlink-cli issue +close --owner Gitlink --repo forgeplus -i 123

# 添加评论
gitlink-cli issue +comment --owner Gitlink --repo forgeplus -i 123 -b "已修复"
```

### Pull Request

```bash
# 列出 PR
gitlink-cli pr +list --owner Gitlink --repo forgeplus

# 创建 PR（同仓库分支）
gitlink-cli pr +create --owner Gitlink --repo forgeplus -t "feat: 搜索功能" --head feature/search --base master

# 创建 PR（从 Fork 仓库）
gitlink-cli pr +create --owner Gitlink --repo forgeplus -t "feat: 新功能" --head your_username/forgeplus:feature/my-feature --base master

# 查看 PR
gitlink-cli pr +view --owner Gitlink --repo forgeplus -i 42

# 合并 PR
gitlink-cli pr +merge --owner Gitlink --repo forgeplus -i 42

# 查看 PR 变更文件
gitlink-cli pr +files --owner Gitlink --repo forgeplus -i 42
```

### 发布管理

```bash
# 列出 Release
gitlink-cli release +list --owner Gitlink --repo forgeplus

# 创建 Release
gitlink-cli release +create --owner Gitlink --repo forgeplus -t v1.0.0 -n "v1.0.0 正式版" -b "更新内容..."

# 查看 Release
gitlink-cli release +view --owner Gitlink --repo forgeplus -i <version_id>
```

### 搜索

```bash
# 搜索仓库
gitlink-cli search +repos -k "machine learning"

# 搜索用户
gitlink-cli search +users -k "zhangsan"
```

### Raw API

Shortcuts 未覆盖的接口可通过 Raw API 直接调用：

```bash
# GET 请求
gitlink-cli api GET /users/me

# POST 请求
gitlink-cli api POST /Gitlink/forgeplus/issues --body '{"subject":"test","description":"..."}'

# 带查询参数
gitlink-cli api GET /Gitlink/forgeplus/commits --query 'page=1&limit=5'
```

## 全局参数

| 参数 | 说明 | 示例 |
|------|------|------|
| `--owner` | 仓库所有者 | `--owner Gitlink` |
| `--repo` | 仓库名称 | `--repo forgeplus` |
| `--format` | 输出格式（json/table/yaml） | `--format json` |
| `--debug` | 启用调试输出 | `--debug` |

**自动上下文解析**：在 git 仓库目录下，`--owner` 和 `--repo` 会自动从 `git remote origin` 解析。

## 分支约定

gitlink-cli 支持 GitHub 和 GitLink 的代码双向同步：

| 平台 | 主分支 |
|------|--------|
| GitHub | `main` |
| GitLink | `master` |

**本地 push 到 GitLink**：

```bash
# 方式 1：使用 git 命令
git push gitlink main:master

# 方式 2：配置 git remote
git config remote.gitlink.push refs/heads/main:refs/heads/master
git push gitlink
```

## AI Agent Skills

`skills/` 目录包含 11 个 Claude Code Agent Skill 文件，支持 AI 自动化操作 GitLink 平台。

详见 [skills/README.md](skills/README.md)

| Skill | 说明 |
|-------|------|
| `gitlink-shared` | 认证、全局参数、安全规则、API 注意事项 |
| `gitlink-repo` | 仓库操作（创建、查看、删除、Fork 等） |
| `gitlink-issue` | Issue 操作（创建、更新、关闭、评论等） |
| `gitlink-pr` | Pull Request 操作（创建、合并、Review 等） |
| `gitlink-release` | 发布管理（创建、查看、删除等） |
| `gitlink-org` | 组织管理（成员、团队等） |
| `gitlink-ci` | CI/CD 操作（构建、日志等） |
| `gitlink-search` | 搜索功能（仓库、用户等） |
| `gitlink-user` | 用户管理（个人信息等） |
| `gitlink-pm` | 项目管理（Sprint、看板、周报等） |
| `gitlink-workflow` | AI 自动化工作流（Issue 分类、PR Review、Release Notes 等） |

## 项目结构

```
gitlink-cli/
├── cmd/                      # Cobra 命令定义
│   ├── root.go               # 根命令 + 全局 flags
│   ├── auth/                 # 认证命令
│   ├── api/                  # Raw API 命令
│   ├── config/               # 配置命令
│   └── cmdutil/              # 全局工具
├── internal/                 # 内部包
│   ├── auth/                 # 登录、Token 存储、Transport
│   ├── client/               # HTTP 客户端 + 分页
│   ├── config/               # 配置文件管理
│   ├── context/              # git remote 解析
│   └── output/               # Envelope + Formatter
├── shortcuts/                # Shortcut 实现
│   ├── common/               # 框架（types, runner）
│   ├── repo/                 # 仓库 shortcuts
│   ├── issue/                # Issue shortcuts
│   ├── pr/                   # PR shortcuts
│   ├── branch/               # 分支 shortcuts
│   ├── release/              # Release shortcuts
│   ├── org/                  # 组织 shortcuts
│   ├── ci/                   # CI shortcuts
│   ├── search/               # 搜索 shortcuts
│   ├── user/                 # 用户 shortcuts
│   └── register.go           # 注册入口
├── skills/                   # AI Agent Skills
│   ├── README.md             # Skills 使用指南
│   ├── gitlink-shared/       # 共享规则
│   ├── gitlink-repo/         # 仓库 Skill
│   ├── gitlink-issue/        # Issue Skill
│   ├── gitlink-pr/           # PR Skill
│   ├── gitlink-pm/           # 项目管理 Skill
│   └── ...
├── doc/                      # 设计文档
│   ├── SKILLS_TEST_REPORT_2026-04-02.md
│   ├── CODE_SYNC_STRATEGY_FINAL.md
│   └── ...
├── main.go
├── Makefile
├── go.mod
└── README.md
```

## 文档

- [Skills 使用指南](skills/README.md) - AI Agent Skills 详细说明
- [设计文档](doc/design.md) - 架构设计和开发计划
- [测试报告](doc/SKILLS_TEST_REPORT_2026-04-02.md) - 功能测试报告
- [代码同步方案](doc/CODE_SYNC_STRATEGY_FINAL.md) - GitHub ↔ GitLink 同步设计

## 常见问题

### Q: 如何在脚本中使用 gitlink-cli？

A: 使用 `GITLINK_TOKEN` 环境变量 + `--format json` 获取结构化输出：

```bash
export GITLINK_TOKEN="your-private-token"
gitlink-cli repo +list --format json | jq '.data.projects[] | .name'
```

### Q: 如何自动解析 owner/repo？

A: 在 git 仓库目录下运行命令，CLI 会自动从 `git remote origin` 解析：

```bash
cd ~/my-gitlink-project
gitlink-cli issue +list  # 自动使用当前仓库
```

### Q: Token 过期了怎么办？

A: 重新登录：

```bash
# 用户名密码登录
gitlink-cli auth login

# 或使用私人令牌（在 GitLink 网页端 个人设置 → 私人令牌 中生成）
gitlink-cli auth login --token
```

### Q: 如何在 CI/CD 或非交互环境（Trae 沙箱等）中使用？

A: 设置 `GITLINK_TOKEN` 环境变量即可，无需 `auth login`：

```bash
export GITLINK_TOKEN="your-private-token"
gitlink-cli repo +list   # 直接可用
gitlink-cli auth status   # 显示 "✓ Logged in via GITLINK_TOKEN environment variable"
```

Token 优先级：`GITLINK_TOKEN` 环境变量 > keyring/文件存储的 token。不设置环境变量时完全兼容原有交互式登录。

Priority: `GITLINK_TOKEN` env var > keyring/file stored token. When env var is not set, the original interactive login flow works as before.

### Q: Windows 上凭证存储在哪里？

A: gitlink-cli 使用 Windows Credential Manager 安全存储 Token。如果 Credential Manager 不可用，会自动降级到文件存储 (`~/.config/gitlink-cli/credentials`)。

### Q: 如何查看完整的 API 参考？

A: 查看 [skills/gitlink-shared/REFERENCE.md](skills/gitlink-shared/REFERENCE.md)

## 相关项目

- [gitlink-bisync](https://www.gitlink.org.cn/wbtiger/gitlink-bisync) - GitHub ↔ GitLink 代码双向同步系统

## 许可证

[Apache License 2.0](LICENSE)
