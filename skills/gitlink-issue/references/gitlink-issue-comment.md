# issue +comment

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

Add a comment to an existing issue.

## 命令

```bash
# Add a comment to issue #42
gitlink-cli issue +comment -i 42 -b "This has been fixed in commit abc123"

# Add a multi-line comment
gitlink-cli issue +comment -i 42 -b "Investigation results:
- Root cause: null pointer in auth module
- Fix: add nil check before dereference"
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--id, -i` | **是** | Issue ID |
| `--body, -b` | **是** | 评论内容（映射为 API 字段 `content`） |
| `--owner` | 否 | 仓库所有者（自动从 git remote 解析） |
| `--repo` | 否 | 仓库名称（自动从 git remote 解析） |
| `--format` | 否 | 输出格式: `json`/`table`/`yaml` |
| `--debug` | 否 | 开启调试输出 |

## API

```
POST /issues/{id}/journals
Body: { "content": body }
```

## Workflow

1. **Confirm** the comment content with the user before posting.
2. **Execute** `gitlink-cli issue +comment -i {id} -b "..."`.
3. **Report** that the comment was added successfully.

> [!CAUTION]
> This is a **Write Operation** -- confirm user intent before executing.

## References

- [gitlink-issue](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
