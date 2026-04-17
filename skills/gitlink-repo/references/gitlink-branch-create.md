# branch +create

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

创建一个新分支。

## 命令

```bash
# 从 master 创建分支
gitlink-cli branch +create --name feature/new-feature

# 从指定分支创建
gitlink-cli branch +create --name hotfix/bug-123 --from develop

# 指定仓库
gitlink-cli branch +create --name feature/x --owner someone --repo myrepo
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--name, -n` | 是 | 新分支名称 |
| `--from, -f` | 否 | 源分支或 commit（默认 `master`） |
| `--owner` | 是* | 仓库所有者（可从 git remote 自动推断） |
| `--repo` | 是* | 仓库名称（可从 git remote 自动推断） |
| `--format` | 否 | 输出格式：`json`/`table`/`yaml` |
| `--debug` | 否 | 启用调试输出 |

> *如果在 GitLink 仓库目录下执行，`--owner` 和 `--repo` 可自动推断。

## Workflow

> [!CAUTION]
> This is a **Write Operation** -- confirm user intent.

1. 确认用户希望创建的分支名称和源分支。
2. 执行 `branch +create --name <name> --from <source>`。
3. 输出创建结果。

## References
- [gitlink-repo](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
