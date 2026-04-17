# release +view

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

查看发行版详情。注意：必须使用 version_id（数字 ID），不支持 tag_name。

## 命令

```bash
# 查看指定发行版（使用 version_id）
gitlink-cli release +view --id 12345

# 指定仓库
gitlink-cli release +view --id 12345 --owner someone --repo myrepo

# 输出为 JSON
gitlink-cli release +view --id 12345 --format json
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--id, -i` | 是 | 发行版 ID（version_id，数字 ID，**不是** tag_name） |
| `--owner` | 是* | 仓库所有者（可从 git remote 自动推断） |
| `--repo` | 是* | 仓库名称（可从 git remote 自动推断） |
| `--format` | 否 | 输出格式：`json`/`table`/`yaml` |
| `--debug` | 否 | 启用调试输出 |

> *如果在 GitLink 仓库目录下执行，`--owner` 和 `--repo` 可自动推断。

> **重要：** `--id` 参数必须传 version_id（从 `release +list` 返回结果中获取），不能传 tag name。

## References
- [gitlink-release](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
