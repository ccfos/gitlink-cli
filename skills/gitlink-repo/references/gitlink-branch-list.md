# branch +list

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

列出仓库的所有分支。

## 命令

```bash
# 列出当前仓库的分支
gitlink-cli branch +list

# 指定仓库并分页
gitlink-cli branch +list --owner someone --repo myrepo --page 1 --limit 10

# 输出为 JSON
gitlink-cli branch +list --format json
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--owner` | 是* | 仓库所有者（可从 git remote 自动推断） |
| `--repo` | 是* | 仓库名称（可从 git remote 自动推断） |
| `--page, -p` | 否 | 页码（默认 `1`） |
| `--limit, -l` | 否 | 每页条数（默认 `20`） |
| `--format` | 否 | 输出格式：`json`/`table`/`yaml` |
| `--debug` | 否 | 启用调试输出 |

> *如果在 GitLink 仓库目录下执行，`--owner` 和 `--repo` 可自动推断。

## References
- [gitlink-repo](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
