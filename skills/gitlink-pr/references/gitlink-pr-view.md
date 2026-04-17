# pr +view

> **前置条件：** 先阅读 [`../../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

查看 Pull Request 详情。

## 命令

```bash
# 查看 PR 详情
gitlink-cli pr +view --id 14200

# 简写
gitlink-cli pr +view -i 14200

# JSON 格式输出
gitlink-cli pr +view -i 14200 --format json
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--id` / `-i` | 是 | PR 编号（`pull_request_id`，从 `pr +list` 或 `pr +create` 获取） |

## API

```
GET /{owner}/{repo}/pulls/{id}
```

## 注意事项

- `--id` 使用的是 `pull_request_id`，可从 `pr +list` 或 `pr +create` 的返回结果中获取
- 返回内容包含 PR 标题、描述、状态、源/目标分支、作者等详细信息
- PR 状态字段 `pull_request_status`：`0` = open, `1` = merged, `2` = closed
- 建议在执行 `pr +merge` 或 `pr +close` 前先用 `pr +view` 确认 PR 当前状态

## References

- [gitlink-shared SKILL.md](../../gitlink-shared/SKILL.md) -- 认证与全局参数
- [gitlink-pr SKILL.md](../SKILL.md) -- PR 操作总览
