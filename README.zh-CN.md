# gitlink-cli

[![GitLink](https://img.shields.io/badge/GitLink-Gitlink%2Fgitlink--cli-green)](https://www.gitlink.org.cn/Gitlink/gitlink-cli)
[![License](https://img.shields.io/badge/License-MulanPSL--2.0-blue.svg)](https://license.coscl.org.cn/MulanPSL2)
[![Go Version](https://img.shields.io/badge/Go-1.26%2B-blue.svg)](https://golang.org)
[![npm version](https://img.shields.io/npm/v/@gitlink-ai/cli.svg)](https://www.npmjs.com/package/@gitlink-ai/cli)

[GitLink（确实开源）](https://www.gitlink.org.cn) 官方 CLI 工具 — 为人类和 AI Agent 双重设计。支持 **macOS、Linux、Windows**，覆盖仓库管理、Issue 追踪、Pull Request、CI/CD 和 AI 自动化工作流，包含 40+ 命令和 11 个 AI Agent [Skills](./skills/)。

**[English](./README.md)**

[安装](#安装与快速上手) · [AI Agent Skills](#ai-agent-skills) · [认证](#配置与使用) · [命令](#使用示例) · [贡献](#相关项目)

## 为什么选择 gitlink-cli？

- **Agent-Native 设计** — 开箱即用 11 个结构化 [Skills](./skills/)，兼容 Claude Code — Agent 零配置即可操作 GitLink
- **广泛覆盖** — 仓库、Issue、PR、分支、Release、CI、组织、搜索、用户 — 核心功能全覆盖
- **AI 友好 & 优化** — 每条命令都经过真实 Agent 测试，简洁参数、智能默认值、结构化输出
- **跨平台** — macOS、Linux、Windows (x64/arm64) 全支持，`npm` 一条命令安装
- **开源零门槛** — 木兰宽松许可证第2版（MulanPSL-2.0），`npm install` 即用
- **3 分钟上手** — 交互式登录或 `GITLINK_TOKEN` 环境变量，从安装到首次 API 调用仅需 3 步
- **安全可控** — OS 原生 keychain 凭证存储，`GITLINK_TOKEN` 环境变量支持 CI/CD 和非交互环境，自动 git remote 上下文解析
- **三层架构** — Shortcuts（人+AI友好）→ Raw API（全覆盖）→ Config（配置管理）

## 功能一览

| 分类 | 能力 |
|------|------|
| 📦 仓库 | 列出、创建、Fork、删除仓库，查看仓库信息 |
| 🐛 Issue | 创建、更新、关闭、评论 Issue |
| 🔀 PR | 创建、合并、Review Pull Request，查看变更文件 |
| 🌿 分支 | 创建、删除、保护分支 |
| 🏷️ 发布 | 创建、查看、删除 Release |
| 🏢 组织 | 管理组织、成员、团队 |
| 🔧 CI | 查看构建、日志、CI/CD 操作 |
| 🔍 搜索 | 搜索仓库、用户 |
| 👤 用户 | 查看用户资料和信息 |
| 📋 项目管理 | Sprint 管理、看板、周报 |
| 🤖 工作流 | AI 驱动的 Issue 分类、PR Review、Release Notes |

## 安装与快速上手

### 前置条件

- Node.js 14+（`npm`/`npx`）— 用于 npm 安装
- 支持平台：macOS、Linux、Windows（x64/arm64）
- Go 1.26+ — 仅从源码构建时需要

### 快速上手（人类用户）

> **AI 助手请注意：** 如果你是帮助用户安装的 AI Agent，请直接跳到 [快速上手（AI Agent）](#快速上手ai-agent)，其中包含你需要完成的所有步骤。

#### 安装

选择以下**任一**方式：

**方式 1 — 从 npm 安装（推荐）：**

```bash
# 安装 CLI
npm install -g @gitlink-ai/cli

# 安装 CLI Skill（必须，全平台通用）
gitlink-cli-install-skills

# 也可使用 npx 安装 Skill
npx skills add ccfos/gitlink-cli/skills -y -g
```

**方式 2 — 从源码构建：**

需要 Go 1.26+。

```bash
git clone https://www.gitlink.org.cn/Gitlink/gitlink-cli.git
cd gitlink-cli
make install

# 安装 CLI Skill（必须）
npx skills add ./skills -y -g
```

> **Windows 用户注意：** 请在 PowerShell 或 CMD 中运行 `npm install -g @gitlink-ai/cli`。从源码构建请使用 `go install .` 代替 `make install`。

#### 配置与使用

```bash
# 1. 配置（首次使用，交互式引导）
gitlink-cli config init

# 2. 登录（任选其一）
gitlink-cli auth login            # 用户名密码（推荐）
gitlink-cli auth login --token    # 或粘贴私人令牌
export GITLINK_TOKEN="your-token" # 或设置环境变量（适用于 CI/CD、非交互环境）

# 3. 开始使用
gitlink-cli repo +list
```

### 快速上手（AI Agent）

> 以下步骤面向 AI Agent。部分步骤需要用户在浏览器中完成操作。

**第 1 步 — 安装**

```bash
# 安装 CLI
npm install -g @gitlink-ai/cli

# 安装 CLI Skill（必须，全平台通用）
gitlink-cli-install-skills
```

**第 2 步 — 配置**

```bash
gitlink-cli config init
```

**第 3 步 — 登录**

交互环境：
```bash
gitlink-cli auth login
```

非交互环境（CI/CD、Trae 沙箱、MCP 等）：
```bash
export GITLINK_TOKEN="your-private-token"
```

> 获取私人令牌：GitLink 网页端 → 个人设置 → 私人令牌。

**第 4 步 — 验证**

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
│   ├── Design.md
│   ├── CODE_SYNC_STRATEGY_FINAL.md
│   └── ...
├── main.go
├── Makefile
├── go.mod
└── README.md
```

## 文档

- [Skills 使用指南](skills/README.md) — AI Agent Skills 详细说明
- [设计文档](doc/design.md) — 架构设计和开发计划

## 常见问题

### Q: 如何在脚本中使用 gitlink-cli？

使用 `GITLINK_TOKEN` 环境变量 + `--format json` 获取结构化输出：

```bash
export GITLINK_TOKEN="your-private-token"
gitlink-cli repo +list --format json | jq '.data.projects[] | .name'
```

### Q: 如何自动解析 owner/repo？

在 git 仓库目录下运行命令，CLI 会自动从 `git remote origin` 解析：

```bash
cd ~/my-gitlink-project
gitlink-cli issue +list  # 自动使用当前仓库
```

### Q: Token 过期了怎么办？

重新登录：

```bash
# 用户名密码登录
gitlink-cli auth login

# 或使用私人令牌（在 GitLink 网页端 个人设置 → 私人令牌 中生成）
gitlink-cli auth login --token
```

### Q: 如何在 CI/CD 或非交互环境（Trae 沙箱等）中使用？

设置 `GITLINK_TOKEN` 环境变量即可，无需 `auth login`：

```bash
export GITLINK_TOKEN="your-private-token"
gitlink-cli repo +list   # 直接可用
gitlink-cli auth status   # 显示 "✓ Logged in via GITLINK_TOKEN environment variable"
```

Token 优先级：`GITLINK_TOKEN` 环境变量 > keyring/文件存储的 token。不设置环境变量时完全兼容原有交互式登录。

### Q: Windows 上凭证存储在哪里？

gitlink-cli 使用 Windows Credential Manager 安全存储 Token。如果 Credential Manager 不可用，会自动降级到文件存储（`~/.config/gitlink-cli/credentials`）。

### Q: 如何查看完整的 API 参考？

查看 [skills/gitlink-shared/REFERENCE.md](skills/gitlink-shared/REFERENCE.md)

## 许可证

[MulanPSL-2.0](https://license.coscl.org.cn/MulanPSL2)
