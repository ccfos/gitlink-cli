# GitLink API 参考

## 认证端点

| 端点 | 方法 | 说明 |
|------|------|------|
| `/accounts/login` | POST | 用户名密码登录 |
| `/users/me` | GET | 获取当前用户信息 |

## 全局参数

| 参数 | 位置 | 说明 |
|------|------|------|
| `access_token` | Query | OAuth2 Token（自动注入） |
| `page` | Query | 分页页码（默认 1） |
| `limit` | Query | 每页条数（默认 20） |

## 响应格式

**成功响应**:
```json
{
  "ok": true,
  "data": { ... },
  "meta": { "page": 1, "limit": 20, "total_count": 100 }
}
```

**错误响应**:
```json
{
  "ok": false,
  "error": {
    "code": 401,
    "message": "无效token",
    "suggestion": "请先运行 gitlink-cli auth login 登录"
  }
}
```

## 常见错误码

| 错误码 | 含义 | 解决方案 |
|--------|------|----------|
| 401 | 未认证 | 运行 `gitlink-cli auth login` |
| 403 | 权限不足 | 确认账户权限或联系项目管理员 |
| 404 | 资源不存在 | 检查 owner/repo/id 是否正确 |
| 422 | 参数校验失败 | 检查请求参数 |
| -1 | GitLink 业务错误 | 查看 message 字段获取详情 |

## API 特殊性

### 必需字段

| 操作 | 必需字段 | 说明 |
|------|----------|------|
| Issue 创建 | `done_ratio: 0` | 数据库约束 |
| Issue 更新 | `subject` | 即使只改状态也需要 |
| Release 查看 | `version_id` | 不能用 tag_name |

### 端点前缀

| 操作 | 前缀 | 示例 |
|------|------|------|
| 分支操作 | `/v1/` | `/v1/:owner/:repo/branches` |
| Issue 评论 | 无 | `/issues/:id/journals` |
| 仓库操作 | 无 | `/:owner/:repo/info` |

### 已知 Bug

| Bug | 影响 | 状态 |
|-----|------|------|
| Branch 删除返回"不存在" | 无法删除分支 | 待 GitLink 修复 |
| Release 删除返回"不存在" | 无法删除发布 | 待 GitLink 修复 |
| Create File 返回"已存在" | 无法通过 API 创建文件 | 待 GitLink 修复 |
