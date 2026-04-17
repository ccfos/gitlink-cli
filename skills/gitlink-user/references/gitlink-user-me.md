# user +me

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

查看当前登录用户的基本信息。

## 命令

```bash
# 查看当前用户
gitlink-cli user +me

# JSON 格式
gitlink-cli user +me --format json
```

## 输出字段

| 字段 | 说明 |
|------|------|
| `login` | 用户名 |
| `username` | 显示名称 |
| `email` | 邮箱 |
| `user_id` | 用户 ID |
| `admin` | 是否管理员 |

## References

- [gitlink-user](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
