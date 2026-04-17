# release +delete

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

删除一个发行版。注意：必须使用 version_id（数字 ID）。

## 命令

```bash
# 删除指定发行版（使用 version_id）
gitlink-cli release +delete --id 12345

# 指定仓库
gitlink-cli release +delete --id 12345 --owner someone --repo myrepo
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--id, -i` | 是 | 发行版 ID（version_id，数字 ID） |
| `--owner` | 是* | 仓库所有者（可从 git remote 自动推断） |
| `--repo` | 是* | 仓库名称（可从 git remote 自动推断） |
| `--format` | 否 | 输出格式：`json`/`table`/`yaml` |
| `--debug` | 否 | 启用调试输出 |

> *如果在 GitLink 仓库目录下执行，`--owner` 和 `--repo` 可自动推断。

## Workflow

> [!CAUTION]
> This is a **Destructive Operation** -- confirm user intent.

1. 确认用户确实希望删除该发行版（此操作不可逆）。
2. 如果用户只知道 tag name，先执行 `release +list` 获取 version_id。
3. 执行 `release +delete --id <version_id>`。
4. 输出删除结果。

## References
- [gitlink-release](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
