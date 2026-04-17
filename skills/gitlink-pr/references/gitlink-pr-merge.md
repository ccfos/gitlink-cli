# pr +merge

> **前置条件：** 先阅读 [`../../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

合并一个 Pull Request。

## 命令

```bash
# 默认 merge 方式合并
gitlink-cli pr +merge --id 14200

# 使用 squash 方式
gitlink-cli pr +merge -i 14200 --method squash

# 使用 rebase 方式
gitlink-cli pr +merge -i 14200 -m rebase
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--id` / `-i` | 是 | PR 编号（`pull_request_id`） |
| `--method` / `-m` | 否 | 合并方式：`merge`、`rebase`、`squash`（默认 `merge`） |

## API

```
POST /{owner}/{repo}/pulls/{id}/pr_merge
Body: { "do": "merge" }
```

## Workflow

> [!CAUTION]
> This is a **Write Operation** -- confirm user intent before executing.

1. 使用 `pr +view -i <id>` 确认 PR 状态为 open（`pull_request_status: 0`）
2. 确认用户意图和合并方式
3. 执行 `pr +merge -i <id> -m <method>`

## 注意事项

- 合并方式说明：
  - `merge` -- 创建合并提交（默认）
  - `rebase` -- 变基到目标分支
  - `squash` -- 压缩为单个提交
- API 请求体使用 `"do"` 字段指定合并方式
- 合并前建议先用 `pr +view` 确认 PR 处于 open 状态
- 需要对目标分支有写入权限

## References

- [gitlink-shared SKILL.md](../../gitlink-shared/SKILL.md) -- 认证与全局参数
- [gitlink-pr SKILL.md](../SKILL.md) -- PR 操作总览
