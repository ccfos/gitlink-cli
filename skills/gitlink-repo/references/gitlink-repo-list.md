# repo +list

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

列出当前用户或指定用户/组织的仓库列表。

## 命令

```bash
# 列出当前用户的仓库
gitlink-cli repo +list

# 列出指定用户的仓库
gitlink-cli repo +list --user someone

# 按类别过滤（manage/mirror/sync/fork/all）
gitlink-cli repo +list --category fork

# 分页
gitlink-cli repo +list --page 2 --limit 10

# 输出为 JSON
gitlink-cli repo +list --format json
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--user, -u` | 否 | 用户登录名（默认：当前认证用户） |
| `--category, -c` | 否 | 过滤类别：`manage`/`mirror`/`sync`/`fork`/`all`（默认 `manage`） |
| `--page, -p` | 否 | 页码（默认 `1`） |
| `--limit, -l` | 否 | 每页条数（默认 `20`） |
| `--owner` | 否 | 全局参数 - 仓库所有者 |
| `--repo` | 否 | 全局参数 - 仓库名称 |
| `--format` | 否 | 输出格式：`json`/`table`/`yaml` |
| `--debug` | 否 | 启用调试输出 |

## References
- [gitlink-repo](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
