# issue +create

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

Create a new issue in the current repository. The `done_ratio` field is automatically set to `0`.

## 命令

```bash
# Create issue with title only
gitlink-cli issue +create -t "Bug: login page crashes"

# Create issue with title and description
gitlink-cli issue +create -t "Feature request" -b "Add dark mode support"

# Create issue with assignee and milestone
gitlink-cli issue +create -t "Fix CI pipeline" -b "Flaky tests" -a user123 -m 5
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--title, -t` | **是** | Issue 标题（映射为 API 字段 `subject`） |
| `--body, -b` | 否 | Issue 描述（映射为 API 字段 `description`） |
| `--assignee, -a` | 否 | 指派人登录名（映射为 API 字段 `assigned_to_id`） |
| `--milestone, -m` | 否 | 里程碑 ID（映射为 API 字段 `fixed_version_id`） |
| `--label` | 否 | 标签 ID |
| `--owner` | 否 | 仓库所有者（自动从 git remote 解析） |
| `--repo` | 否 | 仓库名称（自动从 git remote 解析） |
| `--format` | 否 | 输出格式: `json`/`table`/`yaml` |
| `--debug` | 否 | 开启调试输出 |

## API

```
POST /{owner}/{repo}/issues
Body: { "subject": title, "done_ratio": 0, "description": body, ... }
```

## Workflow

1. **Confirm** the issue title (and optional body) with the user before creating.
2. **Execute** `gitlink-cli issue +create -t "..." -b "..."`.
3. **Report** the created issue ID and URL to the user.

> [!CAUTION]
> This is a **Write Operation** -- confirm user intent before executing.

## References

- [gitlink-issue](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
