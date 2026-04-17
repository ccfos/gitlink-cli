# repo +info

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

查看仓库详细信息。

## 命令

```bash
# 查看当前仓库信息（从 git remote 推断 owner/repo）
gitlink-cli repo +info

# 指定仓库
gitlink-cli repo +info --owner someone --repo myrepo

# 输出为 JSON
gitlink-cli repo +info --format json
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--owner` | 是* | 仓库所有者（可从 git remote 自动推断） |
| `--repo` | 是* | 仓库名称（可从 git remote 自动推断） |
| `--format` | 否 | 输出格式：`json`/`table`/`yaml` |
| `--debug` | 否 | 启用调试输出 |

> *如果在 GitLink 仓库目录下执行，`--owner` 和 `--repo` 可自动推断。

## References
- [gitlink-repo](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
