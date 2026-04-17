# search +users

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

按关键词搜索 GitLink 平台上的用户。

## 命令

```bash
# 搜索用户
gitlink-cli search +users --keyword "zhangsan"

# 限制结果数量
gitlink-cli search +users --keyword "tiger" --limit 10
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--keyword` / `-k` | 是 | 搜索关键词（用户名、姓名等） |
| `--page` / `-p` | 否 | 页码（默认 1） |
| `--limit` / `-l` | 否 | 每页数量（默认 20） |

## References

- [gitlink-search](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
