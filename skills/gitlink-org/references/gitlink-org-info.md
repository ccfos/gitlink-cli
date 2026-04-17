# org +info

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

查看组织详细信息。

## 命令

```bash
# 查看组织信息
gitlink-cli org +info --id Gitlink
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--id` / `-i` | 是 | 组织标识（login name） |

## 输出字段

| 字段 | 说明 |
|------|------|
| `name` | 组织名称 |
| `nickname` | 组织昵称 |
| `description` | 组织描述 |
| `num_projects` | 项目数量 |
| `num_users` | 成员数量 |
| `num_teams` | 团队数量 |
| `website` | 组织网站 |
| `location` | 所在地 |

## References

- [gitlink-org](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
