# release +create

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

创建一个新的发行版。

## 命令

```bash
# 创建发行版
gitlink-cli release +create --tag v1.0.0 --name "v1.0.0 Release"

# 创建带发布说明的发行版
gitlink-cli release +create --tag v1.0.0 --name "v1.0.0" --body "Bug fixes and improvements"

# 创建预发布版本，指定目标分支
gitlink-cli release +create --tag v2.0.0-beta.1 --name "v2.0.0 Beta" --target develop --prerelease true
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--tag, -t` | 是 | Tag 名称 |
| `--name, -n` | 是 | 发行版名称 |
| `--body, -b` | 否 | 发布说明 |
| `--target` | 否 | 目标分支（默认 `main`） |
| `--prerelease` | 否 | 标记为预发布（`true`/`false`，默认 `false`） |
| `--owner` | 是* | 仓库所有者（可从 git remote 自动推断） |
| `--repo` | 是* | 仓库名称（可从 git remote 自动推断） |
| `--format` | 否 | 输出格式：`json`/`table`/`yaml` |
| `--debug` | 否 | 启用调试输出 |

> *如果在 GitLink 仓库目录下执行，`--owner` 和 `--repo` 可自动推断。

## Workflow

> [!CAUTION]
> This is a **Write Operation** -- confirm user intent.

1. 确认用户希望创建的 tag 名称和发行版名称。
2. 执行 `release +create --tag <tag> --name <name>`。
3. 输出创建结果。

## References
- [gitlink-release](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
