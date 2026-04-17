# pr +files

> **前置条件：** 先阅读 [`../../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

查看 Pull Request 的变更文件列表及 diff 内容。`pr +diff` 是此命令的别名，行为完全相同。

## 命令

```bash
# 查看变更文件
gitlink-cli pr +files --id 14200

# 简写
gitlink-cli pr +files -i 14200

# JSON 格式
gitlink-cli pr +files -i 14200 --format json

# pr +diff 是等效别名
gitlink-cli pr +diff -i 14200
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--id` / `-i` | 是 | PR 编号（`pull_request_id`） |

## API

```
GET /{owner}/{repo}/pulls/{id}/files
```

## 注意事项

- 返回内容包含变更文件列表和 diff 内容
- `pr +diff` 与 `pr +files` 调用相同的 API 端点，返回结果一致
- 适合在 code review 场景中使用，先 `pr +view` 查看概览，再 `pr +files` 查看具体变更

## References

- [gitlink-shared SKILL.md](../../gitlink-shared/SKILL.md) -- 认证与全局参数
- [gitlink-pr SKILL.md](../SKILL.md) -- PR 操作总览
