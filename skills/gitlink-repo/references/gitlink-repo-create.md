# repo +create

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

创建一个新仓库（在当前认证用户下）。

## 命令

```bash
# 创建公开仓库
gitlink-cli repo +create --name my-new-repo

# 创建带描述的私有仓库
gitlink-cli repo +create --name my-new-repo --description "A great project" --private true
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--name, -n` | 是 | 仓库名称 |
| `--description, -d` | 否 | 仓库描述 |
| `--private` | 否 | 是否私有（`true`/`false`，默认 `false`） |
| `--format` | 否 | 输出格式：`json`/`table`/`yaml` |
| `--debug` | 否 | 启用调试输出 |

## Workflow

> [!CAUTION]
> This is a **Write Operation** -- confirm user intent.

1. 确认用户希望创建的仓库名称。
2. 执行 `repo +create --name <name>`。
3. 输出创建结果。

## References
- [gitlink-repo](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
