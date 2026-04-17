# pr +create

> **前置条件：** 先阅读 [`../../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

创建一个 Pull Request。源分支必须与目标分支存在实际代码差异。

## 命令

```bash
# 基本创建
gitlink-cli pr +create --title "feat: 新增搜索功能" --head feature/search --base master

# 附带描述
gitlink-cli pr +create -t "fix: 修复登录问题" --head bugfix/login --base master -b "修复了 Token 过期后的重定向问题"

# 指定仓库
gitlink-cli pr +create --owner myorg --repo myrepo --title "docs: 更新 README" --head docs-update --base master
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--title` / `-t` | 是 | PR 标题 |
| `--head` | 是 | 源分支（包含变更的分支） |
| `--base` | 否 | 目标分支（默认 `master`） |
| `--body` / `-b` | 否 | PR 描述 |

## API

```
POST /{owner}/{repo}/pulls
Body: { "title": "...", "head": "...", "base": "...", "body": "..." }
```

## Workflow

> [!CAUTION]
> This is a **Write Operation** -- confirm user intent before executing.

创建 PR 的完整流程（从零开始）：

### Step 1: 创建分支

```bash
gitlink-cli branch +create --name feature-branch --from master
```

### Step 2: 在分支上创建文件（content 必须 base64 编码）

```bash
# 生成 base64 内容
CONTENT=$(echo -n "文件内容" | base64)

# 通过 Raw API 创建文件
gitlink-cli api POST /:owner/:repo/create_file --body '{
  "filepath": "path/to/new-file.md",
  "content": "'$CONTENT'",
  "branch": "feature-branch",
  "message": "add new file"
}'
```

### Step 3: 创建 PR

```bash
gitlink-cli pr +create --title "feat: 新功能" --head feature-branch --base master --body "添加了新文件"
```

### 更新已有文件的流程

如需修改已有文件而非创建新文件：

```bash
# Step 2a: 获取文件 SHA
gitlink-cli api GET /:owner/:repo/sub_entries --query 'filepath=path/to/file.md&ref=feature-branch'
# 从返回的 entries.sha 获取 SHA 值

# Step 2b: 更新文件（content 必须 base64 编码）
CONTENT=$(echo -n "更新后的内容" | base64)
gitlink-cli api PUT /:owner/:repo/update_file --body '{
  "filepath": "path/to/file.md",
  "content": "'$CONTENT'",
  "sha": "<从 sub_entries 获取的 sha>",
  "branch": "feature-branch",
  "message": "update file"
}'
```

## 注意事项

- **源分支与目标分支必须有实际代码差异**，否则 API 返回 "分支内容相同，无需创建合并请求"
- GitLink 默认主分支为 `master`（非 `main`），`--base` 默认值为 `master`
- `create_file` 的 `content` 字段**必须 base64 编码**，不编码会返回 "文件已存在" 错误
- 创建成功后返回的 `pull_request_id` 用于后续 view/merge/close 操作

## References

- [gitlink-shared SKILL.md](../../gitlink-shared/SKILL.md) -- 认证与全局参数、文件操作 API
- [gitlink-pr SKILL.md](../SKILL.md) -- PR 操作总览
