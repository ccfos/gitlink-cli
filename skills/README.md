# GitLink CLI Skills 指南

[![GitLink](https://img.shields.io/badge/GitLink-wbtiger%2Fgitlink--cli-green)](https://www.gitlink.org.cn/wbtiger/gitlink-cli)

欢迎使用 GitLink CLI Skills！本指南帮助你快速上手和充分利用 gitlink-cli 的所有功能。

## 📚 Skills 是什么？

Skills 是为 Claude Code 和其他 AI 代理设计的结构化知识库，包含：

- **SKILL.md** — 命令参考和使用指南
- **REFERENCE.md** — API 详细参考和参数说明
- **TROUBLESHOOTING.md** — 常见问题和解决方案
- **examples/** — 真实工作流示例

AI 代理通过 Skills 可以自动化操作 GitLink 平台，无需手动查阅文档。

---

## 🚀 快速开始

### 1. 安装和认证

```bash
# 首次使用：登录
gitlink-cli auth login

# 验证登录状态
gitlink-cli auth status

# 查看当前用户
gitlink-cli user +me
```

详见: [gitlink-shared/examples/auth-workflow.md](gitlink-shared/examples/auth-workflow.md)

### 2. 查看可用命令

```bash
# 列出所有 Shortcuts
gitlink-cli --help

# 查看特定命令帮助
gitlink-cli repo --help
gitlink-cli issue --help
gitlink-cli pr --help
```

### 3. 使用 JSON 格式便于脚本处理

```bash
# 所有命令都支持 --format json
gitlink-cli user +me --format json
gitlink-cli repo +list --format json
gitlink-cli issue +list --format json
```

---

## 📁 Skills 目录结构

```
skills/
├── README.md                          # 本文件
├── gitlink-shared/                    # 共享基础规则
│   ├── SKILL.md                       # 认证、全局参数、安全规则、分支约定
│   ├── REFERENCE.md                   # API 详细参考、错误处理
│   ├── TROUBLESHOOTING.md             # 常见问题排查
│   └── examples/
│       └── auth-workflow.md           # 认证工作流示例
├── gitlink-repo/                      # 仓库管理
│   ├── SKILL.md                       # 仓库操作指南
│   ├── REFERENCE.md                   # 仓库 API 参考
│   └── examples/
│       └── repo-workflow.md           # 仓库管理工作流
├── gitlink-issue/                     # Issue 管理
│   ├── SKILL.md                       # Issue 操作指南
│   ├── REFERENCE.md                   # Issue API 参考
│   └── examples/
│       └── issue-workflow.md          # Issue 全流程工作流
├── gitlink-pr/                        # Pull Request
│   ├── SKILL.md                       # PR 操作指南
│   ├── REFERENCE.md                   # PR API 参考
│   └── examples/
│       └── pr-workflow.md             # PR 工作流
├── gitlink-branch/                    # 分支管理
│   ├── SKILL.md                       # 分支操作指南
│   └── examples/
│       └── branch-workflow.md         # 分支工作流
├── gitlink-release/                   # 版本发布
│   ├── SKILL.md                       # Release 操作指南
│   ├── REFERENCE.md                   # Release API 参考
│   └── examples/
│       └── release-workflow.md        # Release 工作流
├── gitlink-search/                    # 搜索功能
│   ├── SKILL.md                       # 搜索操作指南
│   └── examples/
│       └── search-workflow.md         # 搜索工作流
├── gitlink-user/                      # 用户管理
│   └── SKILL.md                       # 用户操作指南
├── gitlink-org/                       # 组织管理
│   ├── SKILL.md                       # 组织操作指南
│   └── examples/
│       └── org-workflow.md            # 组织工作流
├── gitlink-ci/                        # CI/CD
│   ├── SKILL.md                       # CI 操作指南
│   └── examples/
│       └── ci-workflow.md             # CI 工作流
├── gitlink-pm/                        # 项目管理
│   └── SKILL.md                       # PM 操作指南
└── gitlink-workflow/                  # AI 自动化工作流
    └── SKILL.md                       # 工作流模板（Issue 分类、PR Review、Release Notes）
```

---

## 📖 所有 Skills 概览

### 核心 Skills

| Skill | 说明 | 常用命令 |
|-------|------|----------|
| **gitlink-shared** | 认证、全局参数、API 参考、安全规则、分支约定 | `auth login`, `auth status` |
| **gitlink-repo** | 仓库管理 | `repo +list`, `repo +create`, `repo +info`, `repo +fork` |
| **gitlink-issue** | Issue 管理 | `issue +create`, `issue +list`, `issue +view`, `issue +close` |
| **gitlink-pr** | Pull Request | `pr +list`, `pr +create`, `pr +view`, `pr +merge`, `pr +review` |
| **gitlink-branch** | 分支管理 | `branch +list`, `branch +create`, `branch +protect` |
| **gitlink-release** | 版本发布 | `release +list`, `release +create`, `release +view` |

### 辅助 Skills

| Skill | 说明 | 常用命令 |
|-------|------|----------|
| **gitlink-search** | 搜索功能 | `search +repos`, `search +users` |
| **gitlink-user** | 用户管理 | `user +me`, `user +info` |
| **gitlink-org** | 组织管理 | `org +list`, `org +info`, `org +members` |
| **gitlink-ci** | CI/CD | `ci +builds`, `ci +logs` |
| **gitlink-pm** | 项目管理 | 通过 Raw API 访问 |
| **gitlink-workflow** | AI 工作流 | Issue 分类、PR Review、Release Notes |

---

## 🎯 使用场景

### 场景 1：查看仓库信息

```bash
# 在仓库目录下自动解析 owner/repo
cd ~/my-project
gitlink-cli repo +info

# 或显式指定
gitlink-cli repo +info --owner wbtiger --repo gitlink-cli
```

详见: [gitlink-repo/examples/repo-workflow.md](gitlink-repo/examples/repo-workflow.md)

### 场景 2：创建和管理 Issue

```bash
# 创建 Issue
gitlink-cli issue +create -t "Bug: 登录失败" -b "复现步骤..."

# 查看 Issue
gitlink-cli issue +view -i 123

# 添加评论
gitlink-cli issue +comment -i 123 -b "已修复"

# 关闭 Issue
gitlink-cli issue +close -i 123
```

详见: [gitlink-issue/examples/issue-workflow.md](gitlink-issue/examples/issue-workflow.md)

### 场景 3：管理分支和发布

```bash
# 创建分支
gitlink-cli branch +create -n develop

# 保护分支
gitlink-cli branch +protect -n master

# 创建 Release
gitlink-cli release +create -t v1.0.0 -n "v1.0.0 正式版" -b "更新内容..."

# 查看 Release
gitlink-cli release +view -i <version_id>
```

详见: [gitlink-release/examples/release-workflow.md](gitlink-release/examples/release-workflow.md)

### 场景 4：搜索和发现

```bash
# 搜索仓库
gitlink-cli search +repos -k "machine learning"

# 搜索用户
gitlink-cli search +users -k "zhangsan"

# 查看组织
gitlink-cli org +list
gitlink-cli org +info -i Gitlink
```

详见: [gitlink-search/examples/search-workflow.md](gitlink-search/examples/search-workflow.md)

---

## 📚 文档导航

### 快速查找

- **我想了解认证**: [gitlink-shared/SKILL.md](gitlink-shared/SKILL.md)
- **我想查看 API 细节**: [gitlink-shared/REFERENCE.md](gitlink-shared/REFERENCE.md)
- **我遇到了错误**: [gitlink-shared/TROUBLESHOOTING.md](gitlink-shared/TROUBLESHOOTING.md)
- **我想看工作流示例**: 查看各 Skill 下的 `examples/` 目录

### 按功能分类

**仓库操作**:
- [gitlink-repo/SKILL.md](gitlink-repo/SKILL.md) - 仓库命令
- [gitlink-branch/SKILL.md](gitlink-branch/SKILL.md) - 分支命令
- [gitlink-repo/examples/repo-workflow.md](gitlink-repo/examples/repo-workflow.md) - 完整工作流

**Issue 和 PR**:
- [gitlink-issue/SKILL.md](gitlink-issue/SKILL.md) - Issue 命令
- [gitlink-pr/SKILL.md](gitlink-pr/SKILL.md) - PR 命令
- [gitlink-issue/examples/issue-workflow.md](gitlink-issue/examples/issue-workflow.md) - Issue 工作流

**发布和搜索**:
- [gitlink-release/SKILL.md](gitlink-release/SKILL.md) - Release 命令
- [gitlink-search/SKILL.md](gitlink-search/SKILL.md) - 搜索命令

**组织和用户**:
- [gitlink-org/SKILL.md](gitlink-org/SKILL.md) - 组织命令
- [gitlink-user/SKILL.md](gitlink-user/SKILL.md) - 用户命令

---

## ❓ 常见问题

### Q: 如何在脚本中使用 gitlink-cli？

A: 使用 `--format json` 获取结构化输出：

```bash
gitlink-cli repo +list --format json | jq '.data.projects[] | .name'
```

### Q: 如何自动解析 owner/repo？

A: 在 git 仓库目录下运行命令，CLI 会自动从 `git remote origin` 解析：

```bash
cd ~/my-project
gitlink-cli repo +info  # 自动使用当前仓库
```

### Q: Token 过期了怎么办？

A: 重新登录：

```bash
gitlink-cli auth login
```

### Q: 如何查看完整的 API 参考？

A: 查看 [gitlink-shared/REFERENCE.md](gitlink-shared/REFERENCE.md)

### Q: 遇到错误怎么办？

A: 查看 [gitlink-shared/TROUBLESHOOTING.md](gitlink-shared/TROUBLESHOOTING.md)

---

## 🤖 AI Agent 使用

Claude Code 和其他 AI 代理可以直接使用这些 Skills 自动化操作 GitLink 平台：

```
用户: "帮我在 GitLink 上创建一个 Issue"
↓
AI 代理读取 gitlink-issue/SKILL.md
↓
AI 代理执行: gitlink-cli issue +create -t "..." -b "..."
↓
完成！
```

AI 代理可以：
- ✅ 自动创建和管理 Issue
- ✅ 自动创建和合并 PR
- ✅ 自动发布 Release
- ✅ 自动分类 Issue
- ✅ 自动生成 Release Notes
- ✅ 自动执行代码审查

---

## 📊 测试状态

✅ **生产就绪** (8.5/10)

- 8/8 常用场景通过
- 所有边界情况处理正确
- 完整的文档和示例

详见: [../doc/SKILLS_TEST_REPORT_2026-04-02.md](../doc/SKILLS_TEST_REPORT_2026-04-02.md)

---

## 🔗 相关资源

- [主项目 README](../README.md) - gitlink-cli 项目说明
- [设计文档](../doc/design.md) - 架构设计和开发计划
- [测试报告](../doc/SKILLS_TEST_REPORT_2026-04-02.md) - 功能测试报告
- [代码同步方案](../doc/CODE_SYNC_STRATEGY_FINAL.md) - GitHub ↔ GitLink 同步设计
- [gitlink-bisync](https://www.gitlink.org.cn/wbtiger/gitlink-bisync) - 代码双向同步系统

---

## 📞 获取帮助

- **命令帮助**: `gitlink-cli <command> --help`
- **故障排查**: [gitlink-shared/TROUBLESHOOTING.md](gitlink-shared/TROUBLESHOOTING.md)
- **API 参考**: [gitlink-shared/REFERENCE.md](gitlink-shared/REFERENCE.md)
- **工作流示例**: 查看各 Skill 下的 `examples/` 目录

---

## 🎓 下一步

1. 阅读 [gitlink-shared/SKILL.md](gitlink-shared/SKILL.md) 了解基础
2. 查看 [gitlink-shared/examples/auth-workflow.md](gitlink-shared/examples/auth-workflow.md) 完成认证
3. 根据需求选择相应的 Skill 文档
4. 参考 `examples/` 目录中的工作流示例
5. 使用 AI 代理自动化你的工作流

祝你使用愉快！🚀
