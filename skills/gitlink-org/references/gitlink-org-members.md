# org +members

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

列出组织成员。

## 命令

```bash
# 列出组织成员
gitlink-cli org +members --id Gitlink

# 分页
gitlink-cli org +members --id Gitlink --page 1 --limit 50
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--id` / `-i` | 是 | 组织标识（login name） |
| `--page` / `-p` | 否 | 页码（默认 1） |
| `--limit` / `-l` | 否 | 每页数量（默认 20） |

## 输出字段

返回 `organization_users` 数组，每项包含 `user` 对象：

| 字段 | 说明 |
|------|------|
| `user.login` | 用户名 |
| `user.name` | 显示名称 |
| `user.mail` | 邮箱 |

## References

- [gitlink-org](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
