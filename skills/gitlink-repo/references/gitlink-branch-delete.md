# branch +delete

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

删除一个分支。

## 命令

```bash
# 删除指定分支
gitlink-cli branch +delete --name feature/old-feature

# 指定仓库
gitlink-cli branch +delete --name feature/old-feature --owner someone --repo myrepo
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--name, -n` | 是 | 要删除的分支名称 |
| `--owner` | 是* | 仓库所有者（可从 git remote 自动推断） |
| `--repo` | 是* | 仓库名称（可从 git remote 自动推断） |
| `--format` | 否 | 输出格式：`json`/`table`/`yaml` |
| `--debug` | 否 | 启用调试输出 |

> *如果在 GitLink 仓库目录下执行，`--owner` 和 `--repo` 可自动推断。

## Workflow

> [!CAUTION]
> This is a **Destructive Operation** -- confirm user intent.

1. 确认用户确实希望删除该分支（此操作不可逆）。
2. 执行 `branch +delete --name <name>`。
3. 输出删除结果。

## References
- [gitlink-repo](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
