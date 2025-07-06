# FastFiber v2

🚀 Modern Web API scaffold based on Go Fiber framework, enterprise-ready backend solution out of the box.

## ✨ Features

- 🚀 **High Performance**: Built on Fiber framework with excellent performance
- 🏗️ **Clean Architecture**: Layered architecture with clear code structure
- 🔐 **Authentication**: Complete JWT authentication system
- 🎯 **Captcha**: Built-in image captcha functionality
- 📊 **Cache Support**: Redis cache integration
- 🗄️ **Multi-Database**: Support for MySQL/PostgreSQL
- 📝 **Logging**: Structured logging system
- 🔧 **Developer Friendly**: Hot reload for rapid development
- 📦 **Ready to Use**: Zero configuration startup, quick to get started

## Tech Stack

- **Framework**: Go Fiber v2
- **Database**: MySQL/PostgreSQL (GORM)
- **Cache**: Redis
- **Authentication**: JWT
- **Logging**: Zap Logger
- **Captcha**: base64Captcha

## Project Structure

```
FastFiber_v2/
├── api/           # API interface layer
├── config/        # Configuration files
├── core/          # Core initialization
├── middleware/    # Middleware
├── models/        # Data models
├── routers/       # Route configuration
├── service/       # Business logic layer
├── utils/         # Utility functions
└── main.go        # Program entry point
```

## 🚀 Quick Start

### 1. Clone Project

```bash
git clone https://github.com/your-username/FastFiber_v2.git
cd FastFiber_v2
```

### 2. Requirements

- Go 1.23+
- MySQL 5.7+ or PostgreSQL 12+
- Redis 6.0+

### 3. Install Dependencies

```bash
go mod download
```

### 4. Configuration

Copy and modify configuration file:

```bash
cp settings.yaml settings_dev.yaml
```

Edit `settings.yaml` to configure database and Redis connection:

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

### 5. Run Project

#### Development Mode (Hot Reload)

```bash
# Install air
go install github.com/cosmtrek/air@latest

# Start with hot reload
air
```

#### Production Mode

```bash
go run main.go
```

### 6. Database Initialization

```bash
go run main.go -db
```

## API Endpoints

### User Related

- `POST /api/user/register` - User registration
- `POST /api/user/login` - User login
- `GET /api/user/list` - User list (authentication required)

### Captcha

- `GET /api/captcha` - Get captcha

## Configuration

### System Configuration

```yaml
system:
  mode: debug          # Run mode
  host: localhost      # Service address
  port: 8000          # Service port
  version: 1.0.1      # Version
```

### JWT Configuration

```yaml
jwt:
  issuer: "FastFiber"
  access_secret: "your_access_secret"
  refresh_secret: "your_refresh_secret"
  access_expire: 2     # Access token expiration (hours)
  refresh_expire: 168  # Refresh token expiration (hours)
```

## 📖 Usage Guide

### Adding New Features

1. **API Layer**: Create handler functions in `api/` directory
2. **Router Layer**: Add routes in `routers/` directory
3. **Model Layer**: Define data models in `models/` directory
4. **Service Layer**: Write business logic in `service/` directory

### Built-in Middleware

- `auth.go` - JWT authentication middleware
- `bind.go` - Request parameter binding middleware

### Custom Configuration

All configuration items are in `settings.yaml`, supports environment variable override.

## 🚀 Deployment

### Docker Deployment

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

### Build Binary

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o fastfiber-linux main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o fastfiber.exe main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o fastfiber-mac main.go
```

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the project
2. Create feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Create Pull Request

## ⭐ Star History

If this project helps you, please give it a Star ⭐

## 📄 License

This project is open source under MIT License - see [LICENSE](LICENSE) file for details

## 📞 Contact

- Project URL: [https://github.com/your-username/FastFiber_v2](https://github.com/your-username/FastFiber_v2)
- Issue Reports: [Issues](https://github.com/your-username/FastFiber_v2/issues)

---

**If you find it useful, please give it a ⭐ Star for support!**