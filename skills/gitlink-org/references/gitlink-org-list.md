# org +list

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

列出当前用户所属的组织。

## 命令

```bash
# 列出我的组织
gitlink-cli org +list

# JSON 格式输出
gitlink-cli org +list --format json
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--format` | 否 | 输出格式：json / table / yaml |

## 输出字段

| 字段 | 说明 |
|------|------|
| `login` | 组织标识 |
| `name` | 组织名称 |
| `description` | 组织描述 |

## References

- [gitlink-org](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
