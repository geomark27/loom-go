# 🧶 Loom - The Go Project Weaver

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Version](https://img.shields.io/badge/version-1.1.1-green.svg)](https://github.com/geomark27/loom-go/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/geomark27/loom-go)](https://goreportcard.com/report/github.com/geomark27/loom-go)

> **Loom** is not a framework, it's a **code weaver**. Generate professional Go projects in seconds and get the tools to extend them without limits.

**Loom** is a CLI tool that automates the creation and extension of backend Go projects with professional architecture. Think of it as the `create-react-app` or `nest new` of the Go ecosystem.

## 🚀 Key Features

- ⚡ **Create complete projects in 30 seconds**
- 🏗️ **Dual architecture**: Layered (simple) or Modular (scalable)
- 🎯 **Gin as default router** - Modern, fast, and popular (v1.1.0)
- 🗃️ **Full GORM integration** - Migrations, seeders, and CLI (v1.1.0)
- 🔧 **Generate individual components** in existing projects
- ⬆️ **Update projects** without losing your changes
- 🎨 **Add technologies** (routers, ORMs, databases) on-the-fly
- 📦 **Optional helpers** or 100% standalone code
- 🚫 **No runtime overhead** - Just generates code

## 📦 Installation

```bash
# Install globally
go install github.com/geomark27/loom-go/cmd/loom@latest

# Verify installation
loom --version
```

## 🎯 Quick Start

### Create a Project

```bash
# With helpers (recommended for rapid development)
loom new my-api

# Or without helpers (100% independent)
loom new my-api --standalone

# With modular architecture
loom new my-app --modular
```

### Add GORM ORM (NEW in v1.1.0!)

```bash
cd my-api
loom add orm gorm

# This generates:
# - internal/database/database.go (GORM connection)
# - internal/database/seeders/ (seeding system)
# - cmd/console/main.go (database CLI)
# - Makefile targets (db-migrate, db-seed, db-fresh)
```

### Run

```bash
cd my-api
go mod tidy
go run cmd/my-api/main.go
# 🚀 Server running at http://localhost:8080
```

### Database Commands (with GORM)

```bash
# Run migrations
go run cmd/console/main.go migrate

# Run migrations with seeders
go run cmd/console/main.go migrate --seed

# Fresh migration (drop all + migrate + seed)
go run cmd/console/main.go migrate --fresh --seed

# Or use Makefile
make db-migrate
make db-seed
make db-fresh
```

### Test

```bash
# Health check
curl http://localhost:8080/api/v1/health

# Users CRUD (included example)
curl http://localhost:8080/api/v1/users
```

## 🎨 Available Commands

### `loom new` - Create projects

```bash
loom new my-api              # Layered project with helpers
loom new my-app --modular    # Modular project by domains
loom new api --standalone    # Without helpers (100% own code)
```

**Available architectures:**
- **Layered** (default): Simple, ideal for REST APIs
- **Modular**: Scalable, ideal for large applications with multiple domains

### `loom generate` - Generate components

```bash
# Inside an existing Loom project
loom generate module products    # Complete module
loom generate handler orders     # Handler only
loom generate service email      # Service only
loom generate model Category     # Model only
loom generate middleware auth    # HTTP middleware

# Useful flags
loom generate module users --dry-run  # Preview
loom generate handler api --force     # Overwrite
```

### `loom add` - Add technologies

```bash
# Add ORM (NEW in v1.1.0! - Full implementation)
loom add orm gorm            # GORM + PostgreSQL driver + migrations + seeders + CLI

# Change HTTP router
loom add router gin          # Replace with Gin (now default!)
loom add router chi          # Or with Chi
loom add router echo         # Or with Echo

# Configure database
loom add database postgres   # PostgreSQL with docker-compose
loom add database mysql      # MySQL
loom add database mongodb    # MongoDB
loom add database redis      # Redis

# Add authentication
loom add auth jwt            # JWT Authentication
loom add auth oauth2         # OAuth 2.0

# Infrastructure
loom add docker              # Dockerfile + docker-compose.yml

# View all available addons
loom add list
```

### `loom upgrade` - Update projects

```bash
loom version                 # View current version
loom upgrade --show-changes  # See what would change
loom upgrade                 # Update (with automatic backup)
loom upgrade --no-backup     # Update without backup

# If something goes wrong
loom upgrade --restore backup-20251027-153045
```

## 📦 Optional Helpers

If you don't use `--standalone`, your project includes reusable helpers:

```go
import "github.com/geomark27/loom-go/pkg/helpers"

// Standardized HTTP responses
helpers.RespondSuccess(w, data, "Success")
helpers.RespondError(w, err, http.StatusBadRequest)
helpers.RespondCreated(w, user, "User created")

// Automatic validation
type UserDTO struct {
    Name  string `json:"name" validate:"required,min=3"`
    Email string `json:"email" validate:"required,email"`
}
errors := helpers.ValidateStruct(userDTO)

// Structured logging
logger := helpers.NewLogger()
logger.Info("User created", "user_id", user.ID)
logger.Error("Database error", "error", err)
```

Update helpers:
```bash
go get -u github.com/geomark27/loom-go/pkg/helpers
```

## 🏗️ Project Structure

### Layered Architecture

```
my-api/
├── cmd/
│   └── my-api/
│       └── main.go              # Entry point
├── internal/
│   ├── app/
│   │   ├── handlers/            # HTTP handlers
│   │   ├── services/            # Business logic
│   │   ├── repositories/        # Data access
│   │   ├── models/              # Domain models
│   │   ├── dtos/                # Data Transfer Objects
│   │   └── middleware/          # HTTP middlewares
│   ├── config/                  # Configuration
│   └── server/                  # HTTP server
├── pkg/
│   └── helpers/                 # Utilities (optional)
├── docs/
│   └── API.md                   # Documentation
├── .env.example
├── Makefile
└── README.md
```

### Modular Architecture

```
my-app/
├── cmd/
│   └── my-app/
│       └── main.go
├── internal/
│   ├── modules/                 # Domain modules
│   │   ├── users/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── repository.go
│   │   │   ├── model.go
│   │   │   ├── dto.go
│   │   │   ├── router.go
│   │   │   ├── validator.go
│   │   │   └── errors.go
│   │   └── products/
│   │       └── ... (same structure)
│   └── platform/                # Shared infrastructure
│       ├── server/
│       ├── config/
│       └── eventbus/            # Inter-module communication
├── pkg/
│   └── helpers/
└── ...
```

## 🔌 Included API Endpoints

All generated projects include:

### Health Checks
- `GET /api/v1/health` - Service status
- `GET /api/v1/health/ready` - Readiness check

### Users CRUD (Example)
- `GET /api/v1/users` - List users
- `GET /api/v1/users/{id}` - Get user
- `POST /api/v1/users` - Create user
- `PUT /api/v1/users/{id}` - Update user
- `DELETE /api/v1/users/{id}` - Delete user

## 💻 Included Makefile

All projects have these commands:

```bash
make help           # View all commands
make run            # Run application
make build          # Compile
make test           # Tests
make test-coverage  # Tests with coverage
make fmt            # Format code
make vet            # Static analysis
make clean          # Clean files

# Database commands (after loom add orm gorm):
make db-migrate     # Run migrations
make db-seed        # Run seeders
make db-fresh       # Drop all + migrate + seed

# If you added Docker:
make docker-build   # Build image
make docker-up      # Start containers
make docker-down    # Stop containers
make docker-logs    # View logs
```

## 🎨 Philosophy

### "Closed for modification, Open for extension"

- **Not a framework** - Just generates code, no runtime overhead
- **Idiomatic** - Respects Go conventions
- **No magic** - Explicit and understandable code
- **Extensible** - Easy to add new features

### Inspiration

Loom brings the experience of frameworks like **NestJS**, **Laravel**, and **Spring Boot** to the Go ecosystem, while maintaining its simplicity and performance.

## 📚 Real-World Examples

### Create an e-commerce

```bash
# 1. Create base project
loom new ecommerce --modular
cd ecommerce

# 2. Generate domain modules
loom generate module products
loom generate module orders
loom generate module payments
loom generate module customers

# 3. Add PostgreSQL
loom add database postgres

# 4. Add JWT authentication
loom add auth jwt

# 5. Add Docker
loom add docker

# 6. Run
docker-compose up -d
```

### Migrate from Gorilla Mux to Gin

```bash
cd my-existing-project
loom add router gin --force  # Replaces current router
go mod tidy
# Manually update handlers to use gin.Context
```

## 📖 Documentation

- **[DOCS.md](DOCS.md)** - Complete technical documentation
- **[CHANGELOG.md](CHANGELOG.md)** - Version history and changes

## 🤝 Contributing

Contributions are welcome!

1. Fork the project
2. Create your branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add AmazingFeature'`)
4. Push (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

MIT License - see [LICENSE](LICENSE) for details.

## 🙏 Inspiration

- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [NestJS CLI](https://nestjs.com/)
- [Laravel Artisan](https://laravel.com/docs/artisan)
- [Spring Boot CLI](https://spring.io/projects/spring-boot)

## 📞 Contact

- **Author**: Marcos
- **GitHub**: [@geomark27](https://github.com/geomark27)
- **Repository**: [loom-go](https://github.com/geomark27/loom-go)

---

**Do you like Loom?** Give it a ⭐ on GitHub!

Made with ❤️ and ☕ by the Go community
