# GitLink-GitHub 代码双向同步方案（最终版）

**版本**: v2.0（最终确认）
**日期**: 2026-04-02
**状态**: 已确认，可实施

---

## 1 核心原则

- **GitHub 为主仓**：GitHub 是代码的唯一真实来源
- **GitLink 为镜像仓**：GitLink 作为备份和协作平台
- **所有操作在 GitLink 服务端**：无需用户本地配置
- **Main 分支保护**：只能通过 PR merge，禁止直接 push
- **关键点同步**：PR create/patch/merge 时自动同步 GitHub

---

## 2 同步触发点

### 2.1 三个关键触发点

| 触发点 | 事件 | 操作 |
|--------|------|------|
| **PR Create** | 用户创建 PR | fetch GitHub + rebase |
| **PR Patch** | 用户修改 PR（push 新 commit） | fetch GitHub + rebase |
| **PR Merge** | 用户点击 merge | fetch GitHub + rebase + merge + push GitHub |

### 2.2 详细流程

#### 2.2.1 PR Create 时同步

**触发**：用户在 GitLink 创建 PR（feature → main）

**流程**：
```
1. GitLink webhook 接收 PR created 事件
2. 执行：
   ├─ git fetch github main
   ├─ git checkout feature_branch
   ├─ git rebase github/main
   ├─ 检查是否有冲突
   │  ├─ 有冲突 → PR 标记为 "需要 rebase"
   │  │           返回错误信息给用户
   │  └─ 无冲突 → PR 状态正常，可以 merge
   └─ 完成
```

**用户体验**：
- 如果无冲突：PR 创建成功，可以 merge
- 如果有冲突：PR 创建成功，但标记为冲突，提示用户本地解决冲突后重新 push

#### 2.2.2 PR Patch 时同步

**触发**：用户修改 PR（在 feature 分支上新增 commit 并 push）

**流程**：
```
1. GitLink webhook 接收 PR updated 事件
2. 执行：
   ├─ git fetch github main
   ├─ git checkout feature_branch
   ├─ git rebase github/main
   ├─ 检查是否有冲突
   │  ├─ 有冲突 → PR 标记为 "需要 rebase"
   │  │           返回错误信息给用户
   │  └─ 无冲突 → PR 状态正常，可以 merge
   └─ 完成
```

**用户体验**：
- 每次 push 新 commit 时，自动检查是否与 GitHub 最新代码冲突
- 如果有冲突，立即提示用户

#### 2.2.3 PR Merge 时同步

**触发**：用户在 GitLink 点击 "Merge PR"

**流程**：
```
1. GitLink webhook 接收 PR merged 事件
2. 执行（事务性操作）：
   ├─ git fetch github main
   ├─ git checkout feature_branch
   ├─ git rebase github/main
   ├─ 检查是否有冲突
   │  ├─ 有冲突 → merge 失败
   │  │           返回错误信息给用户
   │  │           PR 状态回滚到 "open"
   │  └─ 无冲突 → 继续
   ├─ git checkout main
   ├─ git merge feature_branch
   ├─ git push github main
   ├─ 检查 push 是否成功
   │  ├─ 失败 → merge 失败，回滚
   │  └─ 成功 → merge 成功，PR 标记为 "merged"
   └─ 完成
```

**用户体验**：
- 点击 merge 后，自动完成所有同步操作
- 如果有冲突或 push 失败，立即反馈给用户
- 成功后，GitHub 和 GitLink 代码自动同步

---

## 3 Main 分支保护

### 3.1 保护规则

| 规则 | 说明 |
|------|------|
| 禁止直接 push | 用户不能直接 push 到 main 分支 |
| 只能 PR merge | main 分支只能通过 PR merge 更新 |
| 禁止 force push | GitLink 禁止所有 force push 操作 |

### 3.2 实现

在 GitLink 服务端配置分支保护：

```
项目设置 → 分支保护
├─ 分支名称: main
├─ 禁止直接 push: ✅
├─ 禁止 force push: ✅
└─ 只能通过 PR merge: ✅
```

---

## 4 用户工作流

### 4.1 场景 1：创建 PR

```bash
# 1. 用户在本地创建 feature 分支
git checkout -b feature/new-feature

# 2. 修改代码并 commit
git commit -m "feat: implement feature"

# 3. Push 到 GitLink
git push origin feature/new-feature

# 4. 在 GitLink 创建 PR: feature/new-feature → main
# GitLink webhook 自动触发：
# - fetch github main
# - rebase feature 到 github/main
# - 如果无冲突 → PR 创建成功
# - 如果有冲突 → PR 标记为 "需要 rebase"，提示用户
```

### 4.2 场景 2：修改 PR（Patch）

```bash
# 1. 用户在 feature 分支继续修改
git add .
git commit -m "fix: address review comments"

# 2. Push 到 GitLink
git push origin feature/new-feature

# 3. GitLink webhook 自动触发：
# - fetch github main
# - rebase feature 到 github/main
# - 如果无冲突 → PR 更新成功
# - 如果有冲突 → PR 标记为 "需要 rebase"，提示用户
```

### 4.3 场景 3：Merge PR

```bash
# 1. 用户在 GitLink 点击 "Merge PR"
# GitLink webhook 自动触发：
# - fetch github main
# - rebase feature 到 github/main
# - merge feature 到 main
# - push 到 github main
# - 如果无冲突 → merge 成功，GitHub 自动更新
# - 如果有冲突 → merge 失败，提示用户
```

### 4.4 场景 4：直接 Push（不允许）

```bash
# 用户尝试直接 push 到 main
git push origin main

# GitLink 拒绝：
# ❌ Error: 禁止直接 push 到 main 分支
# 请通过 PR merge 提交代码
```

---

## 5 冲突处理

### 5.1 冲突检测

在 PR create/patch/merge 时，GitLink 自动检查是否有冲突：

```bash
git rebase github/main

# 如果有冲突，rebase 会失败
if [ $? -ne 0 ]; then
    # 有冲突
    return error "冲突检测"
fi
```

### 5.2 冲突提示

当检测到冲突时，返回给用户：

```json
{
  "ok": false,
  "error": {
    "code": "CONFLICT",
    "message": "PR 与 GitHub 最新代码有冲突",
    "details": {
      "conflicted_files": ["file1.js", "file2.js"],
      "suggestion": "请在本地解决冲突后重新 push"
    }
  }
}
```

### 5.3 用户解决冲突

```bash
# 1. 用户本地拉取最新代码
git fetch origin
git fetch github

# 2. 本地 rebase 到 github/main
git rebase github/main

# 3. 手动解决冲突
# 编辑冲突文件，解决冲突

# 4. 继续 rebase
git add .
git rebase --continue

# 5. 强制推送到 GitLink（覆盖之前的 commit）
git push origin feature/new-feature --force-with-lease

# 6. GitLink 再次检查冲突
# 如果无冲突 → PR 更新成功
```

---

## 6 错误处理

### 6.1 常见错误

| 错误 | 原因 | 解决方案 |
|------|------|--------|
| 冲突 | PR 与 GitHub 最新代码冲突 | 本地解决冲突后重新 push |
| Push 失败 | GitHub 网络问题 | 重试或联系管理员 |
| Merge 失败 | 冲突或权限问题 | 检查冲突或权限 |

### 6.2 错误恢复

**如果 PR merge 失败**：

```
1. GitLink 自动回滚 merge 操作
2. PR 状态回到 "open"
3. 用户收到错误提示
4. 用户解决问题后重新 merge
```

---

## 7 实现清单

### 7.1 GitLink 服务端

- [ ] 配置 main 分支保护（禁止直接 push、禁止 force push）
- [ ] 实现 PR created webhook
  - [ ] fetch github main
  - [ ] rebase feature 到 github/main
  - [ ] 检查冲突，标记 PR 状态
- [ ] 实现 PR updated webhook
  - [ ] fetch github main
  - [ ] rebase feature 到 github/main
  - [ ] 检查冲突，标记 PR 状态
- [ ] 实现 PR merged webhook（事务性操作）
  - [ ] fetch github main
  - [ ] rebase feature 到 github/main
  - [ ] merge feature 到 main
  - [ ] push 到 github main
  - [ ] 失败时回滚

### 7.2 错误提示

- [ ] 冲突时返回详细错误信息
- [ ] 包含冲突文件列表
- [ ] 包含解决方案建议

### 7.3 文档

- [ ] 更新 README，说明 main 分支保护规则
- [ ] 编写用户指南，说明 PR 工作流
- [ ] 编写故障排查指南

---

## 8 验证与测试

### 8.1 测试场景

| 场景 | 预期结果 |
|------|---------|
| 创建无冲突 PR | PR 创建成功 |
| 创建有冲突 PR | PR 标记为冲突，提示用户 |
| Patch 无冲突 | PR 更新成功 |
| Patch 有冲突 | PR 标记为冲突，提示用户 |
| Merge 无冲突 | Merge 成功，GitHub 自动更新 |
| Merge 有冲突 | Merge 失败，PR 回滚到 open |
| 直接 push main | 拒绝，提示只能 PR merge |
| Force push | 拒绝，提示禁止 force push |

### 8.2 测试命令

```bash
# 1. 创建 PR
git checkout -b feature/test
echo "test" >> file.txt
git commit -m "test"
git push origin feature/test
# 在 GitLink 创建 PR

# 2. 修改 PR
echo "test2" >> file.txt
git commit -m "test2"
git push origin feature/test

# 3. Merge PR
# 在 GitLink 点击 merge

# 4. 验证 GitHub 是否更新
git log --oneline  # 检查 GitHub 是否有新 commit
```

---

## 9 总结

**方案特点**：

✅ **简洁**：只在 PR create/patch/merge 时同步
✅ **安全**：main 分支保护，禁止直接 push
✅ **可靠**：事务性 merge，失败自动回滚
✅ **用户友好**：清晰的错误提示和恢复指导
✅ **无需本地配置**：所有操作在 GitLink 服务端

**预期效果**：

- GitHub 和 GitLink 代码始终一致
- 用户只需正常 git 操作
- 冲突自动检测，提示用户解决
- 代码质量有保证（PR review + merge）

---

**审批**：已确认
**实施日期**：待定
**联系人**：待定
