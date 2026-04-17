# repo +fork

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

Fork 一个仓库到当前认证用户下。

## 命令

```bash
# Fork 当前仓库（从 git remote 推断 owner/repo）
gitlink-cli repo +fork

# Fork 指定仓库
gitlink-cli repo +fork --owner someone --repo their-repo
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--owner` | 是* | 仓库所有者（可从 git remote 自动推断） |
| `--repo` | 是* | 仓库名称（可从 git remote 自动推断） |
| `--format` | 否 | 输出格式：`json`/`table`/`yaml` |
| `--debug` | 否 | 启用调试输出 |

> *如果在 GitLink 仓库目录下执行，`--owner` 和 `--repo` 可自动推断。

## Workflow

> [!CAUTION]
> This is a **Write Operation** -- confirm user intent.

1. 确认用户希望 fork 的目标仓库。
2. 执行 `repo +fork --owner <owner> --repo <repo>`。
3. 输出 fork 结果。

## References
- [gitlink-repo](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
