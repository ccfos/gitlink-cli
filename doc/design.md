# gitlink-cli 整体设计文档

## Context

GitLink（确实开源）是 CCF 官方开源协作平台，Forgeplus 后端提供 490+ API 端点，但缺少官方 CLI 工具。本项目采用业界主流的分层 CLI 架构（Shortcuts → API Commands → Raw API），开发 gitlink-cli，覆盖高频 Git Forge 场景，并内置 AI Agent Skills 支持 Claude Code 自动化操作。

**设计原则**：单实例（gitlink.org.cn）、Go + Cobra、三层命令体系、Claude Code Skills 优先。

---

## 1 项目结构

```
gitlink-cli/
├── cmd/
│   ├── root.go              # 根命令、全局 flags
│   ├── auth/                 # auth login / logout / status
│   ├── api/                  # api GET/POST/PUT/DELETE（原始层）
│   ├── config/               # config init / set / get / list
│   └── service/              # 元数据驱动的 API 命令层
├── internal/
│   ├── auth/
│   │   ├── login.go          # 用户名密码登录 + Token 粘贴
│   │   ├── token_store.go    # OS Keychain 存储
│   │   └── transport.go      # http.RoundTripper 自动注入 Token
│   ├── client/
│   │   ├── client.go         # HTTP 客户端 + 错误解包
│   │   └── pagination.go     # Kaminari 分页迭代器
│   ├── config/
│   │   └── config.go         # ~/.config/gitlink-cli/config.yaml
│   ├── output/
│   │   ├── envelope.go       # {ok, data, error, meta} 标准输出
│   │   └── formatter.go      # --format json/table/yaml
│   ├── context/
│   │   └── repo.go           # git remote → owner/repo 自动解析
│   └── registry/
│       ├── loader.go         # 元数据加载（内嵌 JSON）
│       └── meta_data.json    # API 元数据（路径、参数、说明）
├── shortcuts/
│   ├── common/
│   │   ├── types.go          # Shortcut / Flag / RuntimeContext 定义
│   │   └── runner.go         # CallAPI / PaginateAll / ResolveOwnerRepo
│   ├── repo/                 # repo +create / +clone / +fork / +list / +info
│   ├── issue/                # issue +list / +create / +view / +close / +comment
│   ├── pr/                   # pr +list / +create / +view / +merge / +review
│   ├── release/              # release +list / +create / +download
│   ├── branch/               # branch +list / +protect / +unprotect
│   ├── org/                  # org +list / +info / +members
│   ├── user/                 # user +me / +info
│   ├── search/               # search +repos / +issues / +users
│   ├── ci/                   # ci +builds / +logs / +restart
│   └── register.go           # 注册所有 shortcuts 到 cobra
├── skills/
│   ├── gitlink-shared/       # SKILL.md — 认证、全局参数、安全规则
│   ├── gitlink-repo/         # SKILL.md + references/ — 仓库操作
│   ├── gitlink-issue/        # SKILL.md + references/ — Issue 操作
│   ├── gitlink-pr/           # SKILL.md + references/ — PR 操作
│   ├── gitlink-ci/           # SKILL.md + references/ — CI/CD 操作
│   ├── gitlink-org/          # SKILL.md + references/ — 组织管理
│   ├── gitlink-release/      # SKILL.md + references/ — 发布管理
│   ├── gitlink-search/       # SKILL.md + references/ — 搜索
│   ├── gitlink-user/         # SKILL.md + references/ — 用户管理
│   ├── gitlink-pm/           # SKILL.md + references/ — 项目管理
│   └── gitlink-workflow/     # SKILL.md — AI 自动化工作流（Issue 分类、PR Review 等）
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

---

## 2 三层命令体系

### 2.1 Layer 1: Shortcuts（快捷命令，`+` 前缀）

面向高频场景的语义化封装，MVP 覆盖 ~43 个：

| 领域 | Shortcuts | 数量 |
|------|-----------|------|
| repo | `+create` `+clone` `+fork` `+list` `+info` `+delete` `+settings` | 7 |
| issue | `+list` `+create` `+view` `+update` `+close` `+comment` `+assign` `+label` | 8 |
| pr | `+list` `+create` `+view` `+merge` `+close` `+review` `+files` `+diff` | 8 |
| release | `+list` `+create` `+view` `+delete` `+download` | 5 |
| branch | `+list` `+create` `+delete` `+protect` `+unprotect` | 5 |
| org | `+list` `+info` `+members` `+create` | 4 |
| ci | `+builds` `+logs` `+restart` `+stop` | 4 |
| user | `+me` `+info` | 2 |

**Shortcut 声明式定义**：

```go
type Shortcut struct {
    Name        string
    Description string
    Flags       []Flag
    Run         func(ctx *RuntimeContext) error
}

type Flag struct {
    Name     string
    Short    string
    Usage    string
    Required bool
    Default  interface{}
}
```

**RuntimeContext 核心方法**：

```go
type RuntimeContext struct {
    Client     *client.Client
    Owner      string  // 自动从 git remote 解析或 --owner 指定
    Repo       string  // 自动从 git remote 解析或 --repo 指定
    Format     string  // json / table / yaml
}

func (ctx *RuntimeContext) CallAPI(method, path string, body interface{}) (*Envelope, error)
func (ctx *RuntimeContext) PaginateAll(path string, params url.Values) ([]json.RawMessage, error)
func (ctx *RuntimeContext) ResolveOwnerRepo() error  // git remote 解析
func (ctx *RuntimeContext) Output(data interface{}) error
```

### 2.2 Layer 2: API Commands（元数据驱动）

自动从 `meta_data.json` 生成，覆盖 ~84 个常用端点：

```
gitlink-cli repos list
gitlink-cli repos get --owner foo --repo bar
gitlink-cli issues list --owner foo --repo bar --state open
gitlink-cli pulls create --owner foo --repo bar --title "..." --head dev --base main
```

元数据格式：

```json
{
  "repos": {
    "list": {
      "method": "GET",
      "path": "/api/:owner/:repo",
      "params": {
        "owner": {"type": "string", "required": true, "from": "path"},
        "repo": {"type": "string", "required": true, "from": "path"}
      },
      "description": "获取仓库信息"
    }
  }
}
```

### 2.3 Layer 3: Raw API（原始调用）

```bash
gitlink-cli api GET /api/users/me
gitlink-cli api POST /api/:owner/:repo/issues --body '{"subject":"bug","description":"..."}'
gitlink-cli api GET /api/projects --query 'page=1&limit=10'
```

覆盖全部 490+ 端点，自动注入认证 Header。

---

## 3 认证（Auth）

### GitLink 认证特点

- OAuth2 (Doorkeeper) 签发 Bearer Token
- Token 有效期 7 天，到期需重新登录
- 不支持 OAuth Device Flow（无 `/device/code` 端点）

### 认证流程

**方式 1：用户名密码登录**

```bash
gitlink-cli auth login
# 交互式输入用户名和密码
# 调用 POST /api/accounts/login 获取 Token
# Token 存入 OS Keychain
```

**方式 2：Token 粘贴**

```bash
gitlink-cli auth login --token
# 交互式粘贴已有 Token
# 存入 OS Keychain
```

**其他命令**：

```bash
gitlink-cli auth status    # 查看登录状态和 Token 过期时间
gitlink-cli auth logout    # 清除 Token
```

### Token 存储

使用 `zalando/go-keyring` 库，跨平台：
- macOS: Keychain
- Linux: Secret Service (GNOME Keyring / KDE Wallet)
- Windows: Credential Manager

Fallback: `~/.config/gitlink-cli/credentials`（文件权限 0600）

---

## 4 配置（Config）

配置文件：`~/.config/gitlink-cli/config.yaml`

```yaml
# 单实例，无需 instance 管理
base_url: https://www.gitlink.org.cn/api
default_format: table    # json | table | yaml
editor: vim              # Issue/PR 正文编辑器
pager: less              # 长输出分页器
```

```bash
gitlink-cli config init          # 首次初始化
gitlink-cli config set key val   # 设置配置项
gitlink-cli config get key       # 读取配置项
gitlink-cli config list          # 列出所有配置
```

---

## 5 API 客户端

### 5.1 GitLink 响应特殊处理

GitLink API 常规错误返回 **HTTP 200 + JSON body 中 status 非 200**：

```json
{"status": 403, "message": "You are not authorized"}
{"status": -1, "message": "参数错误"}
```

客户端必须在 HTTP 层和 JSON body 层双重检查：

```go
func (c *Client) Do(req *http.Request) (*Envelope, error) {
    resp, err := c.http.Do(req)
    // 1. 检查 HTTP 状态码
    if resp.StatusCode >= 400 { ... }
    // 2. 解析 JSON body
    var raw map[string]interface{}
    json.Decode(resp.Body, &raw)
    // 3. 检查 body 中的 status 字段
    if status, ok := raw["status"]; ok && status != 200 {
        return nil, &APIError{Code: status, Message: raw["message"]}
    }
    // 4. 封装为标准 Envelope
    return &Envelope{OK: true, Data: raw}, nil
}
```

### 5.2 分页

GitLink 使用 Kaminari 风格分页：

```
GET /api/projects?page=1&limit=20
→ Response headers: X-Total / X-Page / X-Limit
→ 或 body 中: total_count / page / limit
```

分页迭代器：

```go
func (c *Client) PaginateAll(path string, params url.Values) ([]json.RawMessage, error) {
    var all []json.RawMessage
    page := 1
    for {
        params.Set("page", strconv.Itoa(page))
        resp, err := c.Get(path, params)
        items := resp.Data.([]interface{})
        if len(items) == 0 { break }
        all = append(all, items...)
        page++
    }
    return all, nil
}
```

### 5.3 输出 Envelope

统一输出格式：

```json
{
  "ok": true,
  "data": { ... },
  "meta": {
    "page": 1,
    "limit": 20,
    "total_count": 156,
    "identity": "user:zhangsan"
  }
}
```

错误格式：

```json
{
  "ok": false,
  "error": {
    "code": 403,
    "message": "You are not authorized",
    "suggestion": "请先运行 gitlink-cli auth login 登录"
  }
}
```

---

## 6 Git Remote 上下文解析

gitlink-cli 可从当前目录的 git remote 自动推断 `owner` 和 `repo`：

```go
func ResolveOwnerRepo() (owner, repo string, err error) {
    // 1. 检查 --owner/--repo flags
    // 2. 解析 .git/config 中的 remote "origin" URL
    //    支持: https://www.gitlink.org.cn/owner/repo.git
    //          git@www.gitlink.org.cn:owner/repo.git
    // 3. 提取 owner 和 repo
}
```

使用示例：

```bash
cd ~/projects/my-gitlink-repo
gitlink-cli issue +list              # 自动解析 owner/repo
gitlink-cli issue +list --owner foo --repo bar  # 显式指定
```

---

## 7 Schema 自省

```bash
gitlink-cli schema list                    # 列出所有 API 域
gitlink-cli schema show repos              # 查看 repos 域下的接口
gitlink-cli schema show repos.list         # 查看具体接口参数详情
```

从内嵌的 `meta_data.json` 读取，帮助用户和 AI Agent 发现可用 API。

---

## 8 AI Agent Skills 设计

### 8.1 Skills 目录结构

```
skills/
├── gitlink-shared/
│   └── SKILL.md             # 认证、全局参数、安全规则、错误处理
├── gitlink-repo/
│   ├── SKILL.md             # Shortcuts 总览 + 快速决策
│   └── references/
│       ├── repo-create.md
│       ├── repo-fork.md
│       └── repo-settings.md
├── gitlink-issue/
│   ├── SKILL.md
│   └── references/
│       ├── issue-create.md
│       ├── issue-update.md
│       └── issue-comment.md
├── gitlink-pr/
│   ├── SKILL.md
│   └── references/
│       ├── pr-create.md
│       ├── pr-merge.md
│       └── pr-review.md
├── gitlink-ci/
│   ├── SKILL.md
│   └── references/
│       └── ci-builds.md
├── gitlink-org/
│   └── SKILL.md
├── gitlink-release/
│   └── SKILL.md
├── gitlink-search/
│   └── SKILL.md
├── gitlink-user/
│   └── SKILL.md
├── gitlink-pm/
│   └── SKILL.md
└── gitlink-workflow/
    └── SKILL.md             # AI 自动化工作流 recipes
```

### 8.2 gitlink-shared/SKILL.md 核心内容

```markdown
# gitlink-cli 共享规则

## 认证
- 首次使用：`gitlink-cli auth login`
- Token 有效期 7 天，过期需重新登录
- 遇到 401/403 错误时，引导用户重新登录

## 上下文解析
- 在 git 仓库目录下自动解析 owner/repo
- 可通过 --owner/--repo 显式指定

## 输出格式
- 默认 table 格式，AI 场景建议 --format json
- 所有输出遵循 {ok, data, error, meta} Envelope

## 安全规则
- 禁止输出 Token 到终端明文
- 写入/删除操作前必须确认用户意图
- 使用 --dry-run 预览危险请求
```

### 8.3 gitlink-workflow/SKILL.md — AI 工作流

提供 Claude Code 可直接调用的高级工作流模板：

| 工作流 | 描述 |
|--------|------|
| Issue Triage | 自动分类新 Issue（bug/feature/question），添加标签 |
| PR Review | 获取 PR diff，分析代码质量，添加 review 评论 |
| Release Notes | 从 commits 自动生成版本发布说明 |
| Repo Setup | 初始化仓库（README、License、.gitignore、分支保护） |
| Sprint Report | 汇总 Issue/PR 统计，生成周报 |

---

## 9 完整命令参考

```
gitlink-cli
├── auth
│   ├── login              # 登录
│   ├── logout             # 登出
│   └── status             # 认证状态
├── config
│   ├── init               # 初始化配置
│   ├── set                # 设置配置项
│   ├── get                # 读取配置项
│   └── list               # 列出所有配置
├── repo
│   ├── +create            # 创建仓库
│   ├── +clone             # 克隆仓库
│   ├── +fork              # Fork 仓库
│   ├── +list              # 仓库列表
│   ├── +info              # 仓库详情
│   ├── +delete            # 删除仓库
│   └── +settings          # 仓库设置
├── issue
│   ├── +list              # Issue 列表
│   ├── +create            # 创建 Issue
│   ├── +view              # Issue 详情
│   ├── +update            # 更新 Issue
│   ├── +close             # 关闭 Issue
│   ├── +comment           # 添加评论
│   ├── +assign            # 指派
│   └── +label             # 标签管理
├── pr
│   ├── +list              # PR 列表
│   ├── +create            # 创建 PR
│   ├── +view              # PR 详情
│   ├── +merge             # 合并 PR
│   ├── +close             # 关闭 PR
│   ├── +review            # 代码审查
│   ├── +files             # 变更文件
│   └── +diff              # 查看 Diff
├── release
│   ├── +list              # 发布列表
│   ├── +create            # 创建发布
│   ├── +view              # 发布详情
│   ├── +delete            # 删除发布
│   └── +download          # 下载附件
├── branch
│   ├── +list              # 分支列表
│   ├── +create            # 创建分支
│   ├── +delete            # 删除分支
│   ├── +protect           # 设置保护
│   └── +unprotect         # 取消保护
├── org
│   ├── +list              # 组织列表
│   ├── +info              # 组织详情
│   ├── +members           # 成员列表
│   └── +create            # 创建组织
├── ci
│   ├── +builds            # 构建列表
│   ├── +logs              # 构建日志
│   ├── +restart           # 重启构建
│   └── +stop              # 停止构建
├── user
│   ├── +me                # 当前用户
│   └── +info              # 用户详情
├── search
│   ├── +repos             # 搜索仓库
│   ├── +issues            # 搜索 Issue
│   └── +users             # 搜索用户
├── schema
│   ├── list               # API 域列表
│   └── show               # 接口详情
├── api
│   ├── GET                # 原始 GET
│   ├── POST               # 原始 POST
│   ├── PUT                # 原始 PUT
│   └── DELETE             # 原始 DELETE
└── version                # 版本信息
```

---

## 10 关键文件清单

实现时需要修改/创建的核心文件：

| 文件 | 说明 |
|------|------|
| `cmd/root.go` | 根命令、全局 flags（--owner, --repo, --format, --debug） |
| `cmd/auth/*.go` | login / logout / status |
| `cmd/api/api.go` | Raw API 层 |
| `cmd/service/service.go` | 元数据驱动命令生成 |
| `internal/auth/login.go` | 用户名密码登录逻辑 |
| `internal/auth/token_store.go` | Keychain 存储 |
| `internal/auth/transport.go` | Bearer Token 注入 |
| `internal/client/client.go` | HTTP 客户端 + 错误解包 |
| `internal/client/pagination.go` | 分页迭代器 |
| `internal/config/config.go` | 配置文件读写 |
| `internal/context/repo.go` | git remote 解析 |
| `internal/output/envelope.go` | Envelope 输出 |
| `internal/output/formatter.go` | table / json / yaml 格式化 |
| `internal/registry/loader.go` | 元数据加载 |
| `internal/registry/meta_data.json` | API 元数据 |
| `shortcuts/common/types.go` | Shortcut 核心类型 |
| `shortcuts/common/runner.go` | RuntimeContext |
| `shortcuts/repo/*.go` | 仓库 shortcuts |
| `shortcuts/issue/*.go` | Issue shortcuts |
| `shortcuts/pr/*.go` | PR shortcuts |
| `shortcuts/register.go` | Shortcut 注册 |
| `skills/gitlink-shared/SKILL.md` | 共享 Skill |
| `skills/gitlink-*/SKILL.md` | 各领域 Skill |

---

## 11 开发计划

### Phase 1: Foundation（第 1-2 周）

- 项目骨架（go mod init, Makefile, cobra root）
- `internal/config` — 配置管理
- `internal/auth` — 登录 + Keychain + Transport
- `internal/client` — HTTP 客户端 + 错误解包
- `internal/output` — Envelope + Formatter
- `cmd/auth` — login / logout / status
- `cmd/config` — init / set / get / list

**验证**：`gitlink-cli auth login` → `gitlink-cli api GET /api/users/me` 返回当前用户

### Phase 2: Framework（第 3-4 周）

- `internal/context/repo.go` — git remote 解析
- `internal/client/pagination.go` — 分页迭代器
- `shortcuts/common/` — Shortcut 框架 + RuntimeContext
- `cmd/api/` — Raw API 层
- `internal/registry/` — 元数据加载 + `cmd/service/`
- `cmd/schema/` — Schema 自省

**验证**：`gitlink-cli api GET /api/projects` + `gitlink-cli schema list`

### Phase 3: Core Shortcuts（第 5-7 周）

按优先级实现 shortcuts：
1. `user +me` / `user +info`
2. `repo +list` / `+info` / `+create` / `+fork` / `+clone`
3. `issue +list` / `+create` / `+view` / `+close` / `+comment`
4. `pr +list` / `+create` / `+view` / `+merge` / `+review`
5. `branch +list` / `+protect`
6. `release +list` / `+create`
7. `org +list` / `+info` / `+members`
8. `ci +builds` / `+logs`
9. `search +repos` / `+issues`

**验证**：完整 CRUD 工作流测试

### Phase 4: AI Skills（第 8-9 周）

- 编写 11 个 SKILL.md 文件
- 每个 Skill 包含：命令参考、参数说明、返回值、使用示例、错误处理
- 编写 workflow recipes（Issue Triage, PR Review, Release Notes）
- Claude Code 集成测试

**验证**：Claude Code 使用 Skills 自动完成 Issue 创建 → PR 创建 → 合并全流程

### Phase 5: Polish & Release（第 10-11 周）

- 补充单元测试和集成测试
- 完善 README 和用户文档
- goreleaser 配置，多平台构建
- Homebrew / APT / Scoop 包分发
- 发布 v0.1.0

---

## 12 验证方案

| 阶段 | 验证方式 |
|------|----------|
| Phase 1 | `gitlink-cli auth login` + `gitlink-cli api GET /api/users/me` |
| Phase 2 | `gitlink-cli api GET /api/projects` + `gitlink-cli schema list` |
| Phase 3 | 端到端：创建仓库 → 创建 Issue → 创建 PR → 合并 → 发布 Release |
| Phase 4 | Claude Code Skills 集成测试：AI 自动完成 Issue/PR 工作流 |
| Phase 5 | `goreleaser --snapshot` 多平台构建 + 安装脚本测试 |
