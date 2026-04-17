# pr +list

> **前置条件：** 先阅读 [`../../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

列出仓库的 Pull Request 列表。

## 命令

```bash
# 列出 PR（默认 state=open）
gitlink-cli pr +list

# 指定仓库和状态
gitlink-cli pr +list --owner Gitlink --repo forgeplus --state open

# 分页
gitlink-cli pr +list --page 2 --limit 10

# 查看已合并的 PR（注意：state 仅影响统计计数）
gitlink-cli pr +list --state merged --format json
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--state` / `-s` | 否 | 过滤状态：`open`、`merged`、`closed`（默认 `open`） |
| `--page` / `-p` | 否 | 页码（默认 `1`） |
| `--limit` / `-l` | 否 | 每页条数（默认 `20`） |

## API

```
GET /{owner}/{repo}/pulls?state={state}&page={page}&limit={limit}
```

## 注意事项

- `--state` 参数**仅影响响应中的统计计数**（open_count / merged_count / closed_count），返回的 PR 列表可能包含所有状态的 PR
- 如需精确过滤，请在客户端通过 `pull_request_status` 字段二次过滤：
  - `0` = open
  - `1` = merged
  - `2` = closed
- 返回的每条 PR 包含 `pull_request_id` 字段，用于 `pr +view`、`pr +merge`、`pr +close` 等操作

## References

- [gitlink-shared SKILL.md](../../gitlink-shared/SKILL.md) -- 认证与全局参数
- [gitlink-pr SKILL.md](../SKILL.md) -- PR 操作总览
