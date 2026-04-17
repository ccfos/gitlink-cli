# pr +close

> **前置条件：** 先阅读 [`../../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

关闭（拒绝合并）一个 Pull Request。

## 命令

```bash
# 关闭 PR
gitlink-cli pr +close --id 14200

# 简写
gitlink-cli pr +close -i 14200
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--id` / `-i` | 是 | PR 编号（`pull_request_id`） |

## API

```
POST /{owner}/{repo}/pulls/{id}/refuse_merge
```

## Workflow

> [!CAUTION]
> This is a **Write Operation** -- confirm user intent before executing.

1. 使用 `pr +view -i <id>` 确认 PR 状态为 open（`pull_request_status: 0`）
2. 确认用户确实要关闭此 PR（此操作会拒绝合并）
3. 执行 `pr +close -i <id>`

## 注意事项

- 此操作调用 `refuse_merge` 端点，即**拒绝合并**
- 关闭后 PR 状态变为 closed（`pull_request_status: 2`）
- 需要对仓库有相应权限

## References

- [gitlink-shared SKILL.md](../../gitlink-shared/SKILL.md) -- 认证与全局参数
- [gitlink-pr SKILL.md](../SKILL.md) -- PR 操作总览
