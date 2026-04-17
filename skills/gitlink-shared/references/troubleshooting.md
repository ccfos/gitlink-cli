# 常见问题排查

## 认证问题

### Q: 遇到 401 错误

**症状**: `[-1] 无效token` 或 `[-1] 请登录后再操作`

**原因**:
- Token 过期（7 天有效期）
- Token 存储损坏
- 网络问题导致 Token 未正确注入

**解决**:
```bash
# 重新登录
gitlink-cli auth login

# 或使用新 Token
gitlink-cli auth login --token
```

---

## 权限问题

### Q: 遇到 403 错误

**症状**: `[-1] 您没有权限进行该操作`

**原因**:
- 无仓库管理权限
- 无 Issue 编辑权限
- 无分支删除权限

**解决**:
```bash
# 检查当前用户权限
gitlink-cli user +me

# 确认仓库所有者
gitlink-cli repo +info --owner xxx --repo yyy

# 联系项目管理员申请权限
```

---

## Issue 操作问题

### Q: Issue 创建失败 - "Mysql2::Error: done_ratio cannot be null"

**原因**: 使用 Raw API 创建 Issue 时缺少 `done_ratio` 字段

**解决**:
```bash
# 使用 issue +create shortcut（已自动处理）
gitlink-cli issue +create -t "标题" -b "描述"

# 或使用 Raw API 时添加 done_ratio
gitlink-cli api POST /:owner/:repo/issues --body '{
  "subject": "标题",
  "description": "描述",
  "done_ratio": 0
}'
```

### Q: Issue 关闭失败 - "验证失败: 标题不能为空"

**原因**: 更新 Issue 时缺少 `subject` 字段

**解决**:
```bash
# 使用 issue +close shortcut（已自动处理）
gitlink-cli issue +close -i 123

# 或使用 Raw API 时添加 subject
gitlink-cli api PUT /:owner/:repo/issues/123 --body '{
  "subject": "当前标题",
  "status_id": 5
}'
```

---

## Release 操作问题

### Q: Release 查看返回 HTML 页面

**原因**: 使用了 tag_name 而非 version_id

**解决**:
```bash
# 先获取 version_id
gitlink-cli release +list --format json
# 从返回结果中找到 version_id 字段

# 使用 version_id 查看
gitlink-cli release +view -i <version_id>
```

---

## 分支操作问题

### Q: Branch 删除失败 - "分支不存在"

**原因**: GitLink API Bug - DELETE 端点实现有问题

**解决**:
```bash
# 暂无 CLI 解决方案
# 建议通过 Web UI 删除分支
# 或联系 GitLink 团队修复 API
```

---

## 网络问题

### Q: 请求超时或连接失败

**症状**: `request failed: context deadline exceeded`

**原因**:
- 网络连接不稳定
- GitLink 服务器响应慢
- 防火墙阻止

**解决**:
```bash
# 检查网络连接
ping www.gitlink.org.cn

# 启用调试模式查看详细信息
gitlink-cli user +me --debug

# 重试操作
```

---

## 调试技巧

### 启用调试输出

```bash
gitlink-cli <command> --debug
```

### 查看完整 API 请求

```bash
gitlink-cli api GET /users/me --debug
```

### 检查认证状态

```bash
gitlink-cli auth status
```

### 验证 API 连接

```bash
gitlink-cli user +me
```
