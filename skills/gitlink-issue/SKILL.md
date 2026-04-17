---
name: gitlink-issue
version: 1.0.0
description: "Issue 管理：创建、查看、更新、关闭 Issue，添加评论。当用户需要操作 GitLink Issue 时触发。"
metadata:
  requires:
    bins: ["gitlink-cli"]
  cliHelp: "gitlink-cli issue --help"
---

# gitlink-issue（Issue 操作）

**CRITICAL — 开始前必须先阅读 [`../gitlink-shared/SKILL.md`](../gitlink-shared/SKILL.md)，其中包含认证、权限处理和 API 注意事项。**
**CRITICAL — 所有 Shortcuts 在执行写入/删除操作前，务必先确认用户意图。**

> **前置条件：** 先阅读 [`../gitlink-shared/SKILL.md`](../gitlink-shared/SKILL.md) 了解认证和全局参数。

## Shortcuts

| Shortcut | 说明 | 需要认证 |
|----------|------|----------|
| `issue +list` | Issue 列表 | 否（公开项目） |
| `issue +create` | 创建 Issue | 是 |
| `issue +view` | Issue 详情 | 否（公开项目） |
| `issue +update` | 更新 Issue | 是 |
| `issue +close` | 关闭 Issue | 是 |
| `issue +comment` | 添加评论 | 是 |

## 使用示例

```bash
# 列出 Issue
gitlink-cli issue +list --owner Gitlink --repo forgeplus --state open

# 创建 Issue
gitlink-cli issue +create --owner myuser --repo myrepo --title "Bug: 登录失败" --body "复现步骤：..."

# 查看 Issue 详情
gitlink-cli issue +view --owner Gitlink --repo forgeplus --id 123

# 更新 Issue
gitlink-cli issue +update --id 123 --title "新标题" --body "更新描述"

# 关闭 Issue
gitlink-cli issue +close --id 123

# 添加评论
gitlink-cli issue +comment --id 123 --body "已修复，请验证"
```

## Raw API 补充

```bash
# 获取 Issue 评论列表
gitlink-cli api GET /issues/:issue_id/journals

# 批量更新 Issue
gitlink-cli api POST /:owner/:repo/issues/series_update --body '{"ids":[1,2,3],"status_id":"closed"}'

# Issue 认领
gitlink-cli api POST /issues/:issue_id/claims
```

## GitLink Issue 字段映射

| gitlink-cli 参数 | GitLink API 字段 | 说明 |
|------------------|-----------------|------|
| `--title` | `subject` | Issue 标题 |
| `--body` | `description` | Issue 描述 |
| `--assignee` | `assigned_to_id` | 指派人 ID |
| `--milestone` | `fixed_version_id` | 里程碑 ID |
| `--state` | `status_id` | 状态（5=关闭） |

## API 注意事项

- **创建 Issue 时必须包含 `done_ratio: 0`**，否则数据库报错（CLI 已自动处理）
- **更新/关闭 Issue 时必须包含 `subject` 字段**，即使只修改状态（CLI 已自动处理）
- 使用 Raw API 操作 Issue 时需手动添加这些字段
- Issue 评论路径为 `/issues/:id/journals`（不带 owner/repo 前缀）
