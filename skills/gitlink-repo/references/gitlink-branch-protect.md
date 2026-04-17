# branch +protect

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

设置分支保护规则。

## 命令

```bash
# 保护指定分支
gitlink-cli branch +protect --name main

# 指定仓库
gitlink-cli branch +protect --name main --owner someone --repo myrepo
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--name, -n` | 是 | 要保护的分支名称 |
| `--owner` | 是* | 仓库所有者（可从 git remote 自动推断） |
| `--repo` | 是* | 仓库名称（可从 git remote 自动推断） |
| `--format` | 否 | 输出格式：`json`/`table`/`yaml` |
| `--debug` | 否 | 启用调试输出 |

> *如果在 GitLink 仓库目录下执行，`--owner` 和 `--repo` 可自动推断。

## Workflow

> [!CAUTION]
> This is a **Write Operation** -- confirm user intent.

1. 确认用户希望保护的分支名称。
2. 执行 `branch +protect --name <name>`。
3. 输出设置结果。

## References
- [gitlink-repo](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
