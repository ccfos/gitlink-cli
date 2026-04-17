# user +info

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

查看指定用户的详细信息。

## 命令

```bash
# 查看用户信息
gitlink-cli user +info --login zhangsan
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--login` | 是 | 用户名（login name） |

## 输出字段

| 字段 | 说明 |
|------|------|
| `login` | 用户名 |
| `name` | 显示名称 |
| `user_identity` | 身份（教授、学生等） |
| `user_projects_count` | 项目数量 |
| `user_org_count` | 组织数量 |
| `description` | 个人描述 |

## References

- [gitlink-user](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
