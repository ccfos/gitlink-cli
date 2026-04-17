# GitLink-GitHub 代码双向同步方案

**版本**: v1.0
**日期**: 2026-04-02
**状态**: 已确认

---

## 1 概述

实现 GitHub（主仓）和 GitLink（镜像仓）的代码双向同步，确保两个平台的代码始终一致。通过 Git hooks 在关键操作点进行同步检查，自动解决冲突或提示用户手动处理。

### 1.1 核心原则

- **GitHub 为主仓**：GitHub 是代码的唯一真实来源
- **GitLink 为镜像仓**：GitLink 作为备份和协作平台
- **Hook 控制在 GitLink**：所有同步逻辑通过 GitLink 本地 hooks 实现
- **GitHub 完全被动**：GitHub 不需要任何 hook 操作
- **冲突自动解决**：优先自动 rebase 解决，无法解决时提示用户

---

## 2 同步触发点

### 2.1 四个关键触发点

| 触发点 | Hook | 检查内容 | 动作 |
|--------|------|---------|------|
| **直接 commit** | `pre-commit` | GitLink HEAD vs GitHub HEAD | 不同步则阻止 commit |
| **直接 push** | `pre-push` | GitLink HEAD vs GitHub HEAD | 不同步则阻止 push |
| **创建 PR** | Webhook | 拉取最新 GitHub 代码 | 自动 rebase 到最新 GitHub 代码 |
| **Merge PR** | Webhook | 拉取最新 GitHub 代码 + 自动 rebase | 解决冲突后 merge，并推送到 GitHub |

### 2.2 触发点详解

#### 2.2.1 Pre-commit Hook（直接 commit）

**场景**：用户在 GitLink 本地修改代码后执行 `git commit`

**流程**：
```
1. 用户执行: git commit -m "..."
2. pre-commit hook 触发
3. 检查: GitLink HEAD == GitHub HEAD?
   - 是 → 允许 commit
   - 否 → 阻止 commit，提示用户执行 rebase
```

**实现**：
```bash
#!/bin/bash
# .git/hooks/pre-commit

GITHUB_HEAD=$(git ls-remote https://github.com/owner/repo HEAD | awk '{print $1}')
GITLINK_HEAD=$(git rev-parse HEAD)

if [ "$GITHUB_HEAD" != "$GITLINK_HEAD" ]; then
    echo "❌ Error: GitLink HEAD 不同步 GitHub"
    echo "请执行: git fetch github && git rebase github/main"
    exit 1
fi
exit 0
```

#### 2.2.2 Pre-push Hook（直接 push）

**场景**：用户在 GitLink 本地修改代码后执行 `git push`

**流程**：
```
1. 用户执行: git push origin main
2. pre-push hook 触发
3. 检查: GitLink HEAD == GitHub HEAD?
   - 是 → 允许 push
   - 否 → 阻止 push，提示用户执行 rebase
```

**实现**：
```bash
#!/bin/bash
# .git/hooks/pre-push

GITHUB_HEAD=$(git ls-remote https://github.com/owner/repo HEAD | awk '{print $1}')
GITLINK_HEAD=$(git rev-parse HEAD)

if [ "$GITHUB_HEAD" != "$GITLINK_HEAD" ]; then
    echo "❌ Error: GitLink HEAD 不同步 GitHub"
    echo "请执行: git fetch github && git rebase github/main"
    exit 1
fi
exit 0
```

#### 2.2.3 创建 PR 时同步（Webhook）

**场景**：用户在 GitLink 创建 PR

**流程**：
```
1. 用户在 GitLink 创建 PR: feature → main
2. GitLink webhook 触发
3. 拉取最新 GitHub 代码到 GitLink
4. 自动 rebase PR 分支到最新 GitHub main
5. 如果冲突可自动解决 → 直接解决
6. 如果冲突无法自动解决 → 提示用户手动解决
```

**实现逻辑**：
```
POST /webhook/pr-created
├─ 获取 PR 信息 (source_branch, target_branch)
├─ git fetch github main
├─ git checkout source_branch
├─ git rebase github/main
│  ├─ 冲突可解决 → 自动解决 + git rebase --continue
│  └─ 冲突无法解决 → 提示用户，PR 标记为 "需要手动 rebase"
└─ 更新 PR 状态
```

#### 2.2.4 Merge PR 时同步（Webhook）

**场景**：用户在 GitLink 合并 PR

**流程**：
```
1. 用户在 GitLink 点击 "Merge PR"
2. GitLink webhook 触发
3. 拉取最新 GitHub 代码到 GitLink
4. 自动 rebase PR 分支到最新 GitHub main
5. 如果冲突可自动解决 → 直接解决 + merge
6. 如果冲突无法自动解决 → 停止 merge，提示用户
7. Merge 成功后，自动 push 到 GitHub
```

**实现逻辑**：
```
POST /webhook/pr-merge
├─ 获取 PR 信息 (source_branch, target_branch)
├─ git fetch github main
├─ git checkout source_branch
├─ git rebase github/main
│  ├─ 冲突可解决 → 自动解决 + git rebase --continue
│  └─ 冲突无法解决 → 停止 merge，返回错误
├─ git checkout target_branch
├─ git merge source_branch
├─ git push github target_branch
└─ 更新 PR 状态为 "已合并"
```

---

## 3 冲突处理策略

### 3.1 冲突类型与解决方案

| 冲突类型 | 原因 | 解决方案 |
|---------|------|--------|
| **文件内容冲突** | 同一文件同一行被修改 | 自动 rebase（Git 尝试自动合并） |
| **文件删除冲突** | 一边删除，一边修改 | 提示用户手动选择 |
| **文件重命名冲突** | 同一文件被重命名为不同名称 | 提示用户手动选择 |
| **二进制文件冲突** | 二进制文件被修改 | 提示用户手动选择 |

### 3.2 自动解决策略

**可自动解决的冲突**：
- 不同文件的修改
- 同一文件不同行的修改
- 简单的文本冲突（Git 能自动合并）

**实现**：
```bash
git rebase github/main --no-edit

# 如果有冲突，尝试自动解决
if [ $? -ne 0 ]; then
    # 尝试使用 ours 或 theirs 策略
    git rebase --continue --strategy=recursive -X ours
fi
```

### 3.3 无法自动解决时的处理

**流程**：
```
1. Rebase 失败，存在冲突
2. 返回错误信息给用户
3. PR 标记为 "冲突" 状态
4. 用户本地手动解决冲突
5. 用户执行: git rebase --continue
6. 用户 push 到 GitLink
7. Pre-push hook 检查 → 通过
8. 用户重新点击 "Merge PR"
```

---

## 4 实现架构

### 4.1 组件清单

| 组件 | 位置 | 职责 |
|------|------|------|
| **Pre-commit Hook** | `.git/hooks/pre-commit` | 检查 commit 前的同步状态 |
| **Pre-push Hook** | `.git/hooks/pre-push` | 检查 push 前的同步状态 |
| **PR Created Webhook** | GitLink 服务端 | 创建 PR 时自动 rebase |
| **PR Merged Webhook** | GitLink 服务端 | Merge PR 时自动 rebase + push GitHub |
| **Sync CLI Command** | `gitlink-cli sync` | 手动触发同步（可选） |

### 4.2 数据流

```
GitHub (主仓)
    ↓ (git fetch)
GitLink 本地仓库
    ├─ pre-commit hook (检查同步)
    ├─ pre-push hook (检查同步)
    └─ webhook (创建/合并 PR 时自动 rebase)
    ↓ (git push)
GitLink 远程仓库
    ↓ (webhook)
GitHub (推送更新)
```

---

## 5 用户工作流

### 5.1 场景 1：直接 commit 到 GitLink

```bash
# 1. 用户在 GitLink 本地修改代码
cd gitlink-cli
echo "new code" >> file.txt

# 2. 执行 commit
git commit -m "feat: add new feature"

# 3. Pre-commit hook 检查
# ✅ 如果 HEAD 同步 → commit 成功
# ❌ 如果 HEAD 不同步 → commit 失败，提示 rebase

# 4. 如果失败，用户手动 rebase
git fetch github
git rebase github/main

# 5. 重新 commit
git commit -m "feat: add new feature"
```

### 5.2 场景 2：创建 PR

```bash
# 1. 用户创建 feature 分支
git checkout -b feature/new-feature

# 2. 修改代码并 commit
git commit -m "feat: implement feature"

# 3. Push 到 GitLink
git push origin feature/new-feature

# 4. 在 GitLink 创建 PR: feature/new-feature → main
# GitLink webhook 自动触发：
# - 拉取最新 GitHub 代码
# - 自动 rebase feature 分支到最新 GitHub main
# - 如果冲突可解决 → 自动解决
# - 如果冲突无法解决 → PR 标记为 "需要手动 rebase"
```

### 5.3 场景 3：Merge PR

```bash
# 1. 用户在 GitLink 点击 "Merge PR"
# GitLink webhook 自动触发：
# - 拉取最新 GitHub 代码
# - 自动 rebase PR 分支到最新 GitHub main
# - 如果冲突可解决 → 自动解决 + merge
# - 如果冲突无法解决 → merge 失败，提示用户

# 2. Merge 成功后，自动 push 到 GitHub
# GitHub 代码自动更新
```

### 5.4 场景 4：直接 push 到 GitLink

```bash
# 1. 用户本地修改代码并 commit
git commit -m "fix: bug fix"

# 2. 执行 push
git push origin main

# 3. Pre-push hook 检查
# ✅ 如果 HEAD 同步 → push 成功
# ❌ 如果 HEAD 不同步 → push 失败，提示 rebase

# 4. 如果失败，用户手动 rebase
git fetch github
git rebase github/main
git push origin main
```

---

## 6 配置与部署

### 6.1 Hook 安装

在 gitlink-cli 项目中创建 hooks：

```bash
# 创建 hooks 目录
mkdir -p .githooks

# 创建 pre-commit hook
cat > .githooks/pre-commit << 'EOF'
#!/bin/bash
GITHUB_HEAD=$(git ls-remote https://github.com/owner/repo HEAD | awk '{print $1}')
GITLINK_HEAD=$(git rev-parse HEAD)
if [ "$GITHUB_HEAD" != "$GITLINK_HEAD" ]; then
    echo "❌ Error: GitLink HEAD 不同步 GitHub"
    echo "请执行: git fetch github && git rebase github/main"
    exit 1
fi
exit 0
EOF

# 创建 pre-push hook
cat > .githooks/pre-push << 'EOF'
#!/bin/bash
GITHUB_HEAD=$(git ls-remote https://github.com/owner/repo HEAD | awk '{print $1}')
GITLINK_HEAD=$(git rev-parse HEAD)
if [ "$GITHUB_HEAD" != "$GITLINK_HEAD" ]; then
    echo "❌ Error: GitLink HEAD 不同步 GitHub"
    echo "请执行: git fetch github && git rebase github/main"
    exit 1
fi
exit 0
EOF

# 设置权限
chmod +x .githooks/pre-commit .githooks/pre-push

# 配置 Git 使用这些 hooks
git config core.hooksPath .githooks
```

### 6.2 Webhook 配置

在 GitLink 项目设置中配置 webhooks：

**PR Created Webhook**：
- URL: `https://your-server/webhook/pr-created`
- 事件: Pull Request Created
- 负载: PR 信息（source_branch, target_branch, pr_id）

**PR Merged Webhook**：
- URL: `https://your-server/webhook/pr-merged`
- 事件: Pull Request Merged
- 负载: PR 信息（source_branch, target_branch, pr_id）

---

## 7 风险与缓解

### 7.1 潜在风险

| 风险 | 影响 | 缓解措施 |
|------|------|--------|
| **Rebase 失败** | PR 无法合并 | 提示用户手动解决，PR 标记为冲突 |
| **GitHub 网络不可达** | Hook 超时 | 设置超时时间，失败时提示用户 |
| **Webhook 失败** | 同步不及时 | 重试机制 + 手动同步命令 |
| **并发 merge** | 数据不一致 | 使用分布式锁或队列 |

### 7.2 缓解方案

**Hook 超时处理**：
```bash
timeout 5 git ls-remote https://github.com/owner/repo HEAD
if [ $? -eq 124 ]; then
    echo "⚠️ Warning: GitHub 网络超时，跳过同步检查"
    exit 0  # 允许操作继续
fi
```

**Webhook 重试**：
```
失败 → 等待 5 秒 → 重试
失败 → 等待 10 秒 → 重试
失败 → 等待 30 秒 → 重试
失败 → 记录日志，通知管理员
```

---

## 8 手动同步命令（可选）

提供 CLI 命令供用户手动触发同步：

```bash
# 手动同步 GitLink 到最新 GitHub 代码
gitlink-cli sync --from github --to gitlink

# 手动同步 GitHub 到最新 GitLink 代码（不推荐）
gitlink-cli sync --from gitlink --to github

# 查看同步状态
gitlink-cli sync status
```

---

## 9 验证与测试

### 9.1 测试场景

| 场景 | 预期结果 | 验证方法 |
|------|---------|--------|
| GitHub 有新代码，GitLink 直接 commit | Commit 失败，提示 rebase | 执行 commit，检查错误信息 |
| GitHub 有新代码，GitLink 直接 push | Push 失败，提示 rebase | 执行 push，检查错误信息 |
| 创建 PR 时 GitHub 有新代码 | 自动 rebase，PR 创建成功 | 创建 PR，检查 PR 状态 |
| Merge PR 时 GitHub 有新代码 | 自动 rebase + merge，推送 GitHub | Merge PR，检查 GitHub 代码 |
| 冲突无法自动解决 | PR 标记为��突，提示用户 | 创建冲突 PR，检查状态 |

### 9.2 测试命令

```bash
# 1. 测试 pre-commit hook
git commit -m "test"

# 2. 测试 pre-push hook
git push origin main

# 3. 测试 PR 创建同步
# 在 GitLink 创建 PR，检查是否自动 rebase

# 4. 测试 PR 合并同步
# 在 GitLink 合并 PR，检查 GitHub 是否更新
```

---

## 10 后续优化

### 10.1 短期（1-2 周）

- [ ] 实现 pre-commit 和 pre-push hooks
- [ ] 实现 PR Created 和 PR Merged webhooks
- [ ] 编写测试用例
- [ ] 文档完善

### 10.2 中期（2-4 周）

- [ ] 实现手动同步 CLI 命令
- [ ] 添加同步状态监控
- [ ] 优化冲突自动解决策略
- [ ] 添加日志和告警

### 10.3 长期（1 个月+）

- [ ] Issue 和 PR 元数据同步
- [ ] 评论和 Review 同步
- [ ] 自动化测试和 CI/CD 集成
- [ ] 性能优化和扩展性改进

---

## 11 总结

本方案通过 **Git hooks + Webhooks** 的组合，实现了 GitHub 和 GitLink 的代码双向同步。核心特点：

✅ **GitHub 为主仓**：确保代码唯一真实来源
✅ **自动冲突解决**：优先自动 rebase，无法解决时提示用户
✅ **多触发点**：commit、push、PR 创建、PR 合并都有同步检查
✅ **用户友好**：清晰的错误提示和恢复指导
✅ **低风险**：失败时提示用户，不会自动破坏代码

---

**审批人**: [待确认]
**实施日期**: [待定]
**联系人**: [待定]
