# FastFiber v2

🚀 基于 Go Fiber 框架的现代化 Web API 脚手架，开箱即用的企业级后端解决方案。

## ✨ 特性

- 🚀 **高性能**: 基于 Fiber 框架，性能优异
- 🏗️ **架构清晰**: 分层架构，代码结构清晰
- 🔐 **认证系统**: 完整的 JWT 认证体系
- 🎯 **验证码**: 内置图形验证码功能
- 📊 **缓存支持**: Redis 缓存集成
- 🗄️ **多数据库**: 支持 MySQL/PostgreSQL
- 📝 **日志系统**: 结构化日志记录
- 🔧 **开发友好**: 热重载，快速开发
- 📦 **开箱即用**: 零配置启动，快速上手

## 技术栈

- **框架**: Go Fiber v2
- **数据库**: MySQL/PostgreSQL (GORM)
- **缓存**: Redis
- **认证**: JWT
- **日志**: Zap Logger
- **验证码**: base64Captcha

## 项目结构

```
FastFiber_v2/
├── api/           # API 接口层
├── config/        # 配置文件
├── core/          # 核心初始化
├── middleware/    # 中间件
├── models/        # 数据模型
├── routers/       # 路由配置
├── service/       # 业务逻辑层
├── utils/         # 工具函数
└── main.go        # 程序入口
```

## 🚀 快速开始

### 1. 克隆项目

```bash
git clone https://github.com/your-username/FastFiber_v2.git
cd FastFiber_v2
```

### 2. 环境要求

- Go 1.23+
- MySQL 5.7+ 或 PostgreSQL 12+
- Redis 6.0+

### 3. 安装依赖

```bash
go mod download
```

### 4. 配置文件

复制并修改配置文件：

```bash
cp settings.yaml settings_dev.yaml
```

编辑 `settings.yaml` 配置数据库和 Redis 连接信息：

```yaml
db:
  mode: mysql
  db_name: "your_database"
  host: 127.0.0.1
  port: 3306
  user: your_username
  password: your_password

redis:
  addr: "127.0.0.1:6379"
  password: "your_redis_password"
  db: 1
```

### 5. 运行项目

#### 开发模式（热重载）

```bash
# 安装 air
go install github.com/cosmtrek/air@latest

# 启动热重载
air
```

#### 生产模式

```bash
go run main.go
```

### 6. 数据库初始化

```bash
go run main.go -db
```

## API 接口

### 用户相关

- `POST /api/user/register` - 用户注册
- `POST /api/user/login` - 用户登录
- `GET /api/user/list` - 用户列表（需认证）

### 验证码

- `GET /api/captcha` - 获取验证码

## 配置说明

### 系统配置

```yaml
system:
  mode: debug          # 运行模式
  host: localhost      # 服务地址
  port: 8000          # 服务端口
  version: 1.0.1      # 版本号
```

### JWT 配置

```yaml
jwt:
  issuer: "FastFiber"
  access_secret: "your_access_secret"
  refresh_secret: "your_refresh_secret"
  access_expire: 2     # 访问令牌过期时间（小时）
  refresh_expire: 168  # 刷新令牌过期时间（小时）
```

## 📖 使用指南

### 添加新功能

1. **API 层**: 在 `api/` 目录下创建处理函数
2. **路由层**: 在 `routers/` 目录下添加路由
3. **模型层**: 在 `models/` 目录下定义数据模型
4. **服务层**: 在 `service/` 目录下编写业务逻辑

### 内置中间件

- `auth.go` - JWT 认证中间件
- `bind.go` - 请求参数绑定中间件

### 自定义配置

所有配置项都在 `settings.yaml` 中，支持环境变量覆盖。

## 🚀 部署

### Docker 部署

```dockerfile
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/settings.yaml .
CMD ["./main"]
```

### 构建二进制文件

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o fastfiber-linux main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o fastfiber.exe main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o fastfiber-mac main.go
```

## 🤝 贡献

欢迎贡献代码！请遵循以下步骤：

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## ⭐ Star History

如果这个项目对你有帮助，请给个 Star ⭐

## 📄 许可证

本项目基于 MIT 许可证开源 - 查看 [LICENSE](LICENSE) 文件了解详情

## 📞 联系

- 项目地址: [https://github.com/your-username/FastFiber_v2](https://github.com/your-username/FastFiber_v2)
- 问题反馈: [Issues](https://github.com/your-username/FastFiber_v2/issues)

---

**如果觉得有用，请给个 ⭐ Star 支持一下！**