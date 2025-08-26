# Go 模块管理 vs Node.js 包管理对比

## 命令对比表

| 功能         | Node.js                           | Go                                     |
| ------------ | --------------------------------- | -------------------------------------- |
| 初始化项目   | `npm init`                        | `go mod init <module-name>`            |
| 安装依赖     | `npm install <package>`           | `go get <package>`                     |
| 安装特定版本 | `npm install <package>@<version>` | `go get <package>@<version>`           |
| 更新所有依赖 | `npm update`                      | `go get -u ./...`                      |
| 更新特定依赖 | `npm update <package>`            | `go get -u <package>`                  |
| 查看过期依赖 | `npm outdated`                    | `go list -u -m all`                    |
| 删除依赖     | `npm uninstall <package>`         | 删除代码中的 import，然后`go mod tidy` |
| 清理无用依赖 | `npm prune`                       | `go mod tidy`                          |
| 下载依赖     | `npm install`                     | `go mod download`                      |
| 查看依赖树   | `npm list`                        | `go mod graph`                         |
| 验证依赖     | `npm audit`                       | `go mod verify`                        |

## 具体使用示例

### 1. 查看可更新的依赖

```bash
# 查看所有依赖的最新版本信息
go list -u -m all

# 只查看有更新的依赖
go list -u -m all | grep -E '\[.*\]'
```

### 2. 更新依赖

```bash
# 更新所有依赖到最新版本
go get -u ./...

# 更新特定依赖
go get -u github.com/gin-gonic/gin

# 更新到特定版本
go get github.com/gin-gonic/gin@v1.9.2
```

### 3. 添加新依赖

```bash
# 添加新依赖
go get github.com/some/package

# 添加特定版本
go get github.com/some/package@v1.2.3

# 添加最新的预发布版本
go get github.com/some/package@latest
```

### 4. 清理和整理

```bash
# 清理未使用的依赖，添加缺失的依赖
go mod tidy

# 下载依赖到本地缓存
go mod download

# 将依赖复制到vendor目录
go mod vendor
```

### 5. 查看依赖信息

```bash
# 查看模块信息
go list -m all

# 查看依赖关系图
go mod graph

# 查看为什么需要某个依赖
go mod why github.com/gin-gonic/gin
```

## Go 模块的优势

1. **版本锁定**：go.sum 确保依赖版本一致性
2. **语义版本**：遵循 semver 规范
3. **最小版本选择**：选择满足要求的最小版本
4. **去中心化**：不依赖中央包管理器
5. **安全性**：校验和验证包完整性

## 实际工作流程

### 开发时添加新依赖：

1. 在代码中添加 import
2. 运行 `go mod tidy` 自动添加依赖
3. 或者直接运行 `go get <package>`

### 定期更新依赖：

1. 运行 `go list -u -m all` 查看可更新的包
2. 运行 `go get -u ./...` 更新所有依赖
3. 运行 `go mod tidy` 清理
4. 测试确保更新没有破坏功能
