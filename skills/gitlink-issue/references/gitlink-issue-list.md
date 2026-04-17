# issue +list

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

List issues for the current repository, with optional state filter and pagination.

## 命令

```bash
# List open issues (default)
gitlink-cli issue +list

# List closed issues
gitlink-cli issue +list -s closed

# List all issues, page 2, 10 per page
gitlink-cli issue +list -s all -p 2 -l 10

# Output as JSON
gitlink-cli issue +list --format json
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--state, -s` | 否 | 按状态过滤: `open`、`closed`、`all`（默认 `open`） |
| `--page, -p` | 否 | 页码（默认 `1`） |
| `--limit, -l` | 否 | 每页数量（默认 `20`） |
| `--owner` | 否 | 仓库所有者（自动从 git remote 解析） |
| `--repo` | 否 | 仓库名称（自动从 git remote 解析） |
| `--format` | 否 | 输出格式: `json`/`table`/`yaml` |
| `--debug` | 否 | 开启调试输出 |

## API

```
GET /{owner}/{repo}/issues?state={state}&page={page}&limit={limit}
```

## References

- [gitlink-issue](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
