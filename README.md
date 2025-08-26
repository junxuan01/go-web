# Go Web 服务

这是一个使用 Gin 框架构建的简单 HTTP Web 服务，支持基本的 GET 和 POST 请求。

## 项目结构

```
go-web/
├── cmd/
│   └── server/
│       └── main.go          # 应用程序入口
├── internal/
│   ├── handlers/
│   │   └── user.go          # HTTP 请求处理器
│   ├── models/
│   │   └── user.go          # 数据模型定义
│   └── routes/
│       └── routes.go        # 路由配置
├── go.mod                   # Go 模块文件
└── README.md               # 项目说明
```

## 功能特性

- ✅ 使用 Gin 框架
- ✅ 支持 GET 和 POST 请求
- ✅ 返回 JSON 格式数据
- ✅ 统一的响应格式
- ✅ 参数验证
- ✅ 错误处理

## API 接口

### 1. 健康检查

- **GET** `/api/ping`
- 响应: 返回服务状态

### 2. 获取所有用户

- **GET** `/api/users`
- 响应: 返回用户列表

### 3. 根据 ID 获取用户

- **GET** `/api/users/:id`
- 参数: `id` (路径参数)
- 响应: 返回指定用户信息

### 4. 创建新用户

- **POST** `/api/users`
- 请求体:

```json
{
  "name": "用户名",
  "email": "邮箱地址"
}
```

- 响应: 返回创建的用户信息

## 安装和运行

1. 确保已安装 Go 1.21+

2. 克隆项目后，进入项目目录：

```bash
cd go-web
```

3. 下载依赖：

```bash
go mod tidy
```

4. 运行服务：

```bash
go run cmd/server/main.go
```

5. 服务将在 `http://localhost:8080` 启动

## 测试接口

### 使用 curl 测试：

```bash
# 健康检查
curl http://localhost:8080/api/ping

# 获取所有用户（初始为空）
curl http://localhost:8080/api/users

# 创建用户
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","email":"zhangsan@example.com"}'

# 获取用户列表（现在有数据了）
curl http://localhost:8080/api/users

# 根据ID获取用户
curl http://localhost:8080/api/users/1
```

## 响应格式

所有 API 响应都遵循统一格式：

```json
{
  "code": 200,
  "message": "操作成功",
  "data": {}
}
```

- `code`: 状态码（200=成功，400=请求错误，404=未找到，etc.）
- `message`: 响应消息
- `data`: 具体数据内容

## 项目特点

1. **标准项目结构**: 遵循 Go 社区推荐的项目布局
2. **分层架构**: 清晰的分层，便于维护和扩展
3. **详细注释**: 每个重要部分都有中文注释说明
4. **错误处理**: 完善的错误处理机制
5. **数据验证**: 自动验证请求参数

## 学习要点

- Gin 框架的基本使用
- Go 项目的标准结构组织
- HTTP 路由和中间件
- JSON 数据绑定和验证
- 错误处理最佳实践
