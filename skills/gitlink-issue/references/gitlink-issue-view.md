# issue +view

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../../gitlink-shared/SKILL.md) 了解认证、全局参数和安全规则。

View details of a specific issue by its ID.

## 命令

```bash
# View issue #42
gitlink-cli issue +view -i 42

# View issue with JSON output
gitlink-cli issue +view -i 42 --format json
```

## 参数

| 参数 | 必填 | 说明 |
|------|------|------|
| `--id, -i` | **是** | Issue ID 或编号 |
| `--owner` | 否 | 仓库所有者（自动从 git remote 解析） |
| `--repo` | 否 | 仓库名称（自动从 git remote 解析） |
| `--format` | 否 | 输出格式: `json`/`table`/`yaml` |
| `--debug` | 否 | 开启调试输出 |

## API

```
GET /{owner}/{repo}/issues/{id}
```

## References

- [gitlink-issue](../SKILL.md)
- [gitlink-shared](../../gitlink-shared/SKILL.md)
