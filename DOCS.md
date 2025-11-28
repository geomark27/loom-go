# 📖 Technical Documentation - Loom

Complete documentation for developers using Loom.

## 📑 Table of Contents

- [Core Concepts](#core-concepts)
- [Architectures](#architectures)
- [Detailed Commands](#detailed-commands)
- [Addon System](#addon-system)
- [Versioning System](#versioning-system)
- [Helpers API](#helpers-api)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)

---

## 🎯 Core Concepts

### What Loom is NOT

Loom **is NOT a framework**. It does not add runtime overhead nor force you to use a specific API.

### What Loom IS

Loom is a **code generator**:
- ✅ Generates idiomatic Go code
- ✅ Creates professional project structure
- ✅ Provides tools to extend the project
- ✅ Does not add mandatory dependencies (with `--standalone`)
- ✅ No magic, no unnecessary reflection, no DSLs

### Philosophy: "Closed for modification, Open for extension"

Once the code is generated:
- It's **YOUR code**, not Loom's
- You can modify it as you wish
- No vendor lock-in
- You can remove Loom from the project if you want

---

## 🏗️ Architectures

Loom supports two different architectures for different needs:

### 1. Layered Architecture

**When to use:**
- Simple REST APIs
- Small microservices
- Single-domain projects
- Small teams (1-3 developers)

**Structure:**

```
project/
├── cmd/
│   └── project/
│       └── main.go
├── internal/
│   ├── app/
│   │   ├── handlers/      # Presentation layer (HTTP)
│   │   ├── services/      # Business layer
│   │   ├── repositories/  # Data layer
│   │   ├── models/        # Domain models
│   │   ├── dtos/          # Transfer objects
│   │   └── middleware/    # HTTP middleware
│   ├── config/            # Configuration
│   └── server/            # HTTP server
└── pkg/
    └── helpers/           # Utilities (optional)
```

**Data flow:**
```
HTTP Request → Handler → Service → Repository → Database
                  ↓
HTTP Response ← Handler ← Service ← Repository ← Database
```

**Advantages:**
- ✅ Simple and straightforward
- ✅ Easy to understand
- ✅ Good separation of concerns
- ✅ Low coupling between layers

**Disadvantages:**
- ❌ Hard to scale with many domains
- ❌ Can become a "big ball of mud" in large projects

### 2. Modular Architecture (By Domains)

**When to use:**
- Large applications
- Multiple business domains
- Large teams (4+ developers)
- Need clear separation by feature/domain

**Structure:**

```
project/
├── cmd/
│   └── project/
│       └── main.go
├── internal/
│   ├── modules/                # Domain modules
│   │   ├── users/
│   │   │   ├── handler.go      # HTTP handler
│   │   │   ├── service.go      # Business logic
│   │   │   ├── repository.go   # Data access
│   │   │   ├── model.go        # Model
│   │   │   ├── dto.go          # DTOs
│   │   │   ├── router.go       # Module routes
│   │   │   ├── validator.go    # Specific validations
│   │   │   ├── errors.go       # Domain errors
│   │   │   └── ports.go        # Interfaces (DIP)
│   │   └── products/
│   │       └── ...             # Same structure
│   └── platform/               # Shared infrastructure
│       ├── server/
│       ├── config/
│       └── eventbus/           # Inter-module communication
└── pkg/
    └── helpers/
```

**Inter-module communication:**

```go
// Use Event Bus (recommended)
eventbus.Publish("user.created", UserCreatedEvent{UserID: id})

// Or interfaces (DIP - Dependency Inversion Principle)
type UserFinder interface {
    FindByID(id int) (*User, error)
}
```

**Advantages:**
- ✅ Horizontal scalability
- ✅ Independent modules
- ✅ Easy to work in teams (each dev a module)
- ✅ Low coupling between domains
- ✅ Easy to extract to microservices

**Disadvantages:**
- ❌ More complex initially
- ❌ Communication overhead between modules
- ❌ More files and structure

### Quick Comparison

| Aspect | Layered | Modular |
|---------|---------|---------|
| **Complexity** | Low | Medium-High |
| **Scalability** | Limited | Excellent |
| **Learning curve** | Gentle | Steep |
| **Ideal size** | Small-Medium | Medium-Large |
| **Team** | 1-3 devs | 4+ devs |
| **Domains** | 1-2 | 3+ |

---

## 🎨 Detailed Commands

### `loom new` - Create Projects

```bash
loom new [name] [flags]
```

**Flags:**
- `--modular` - Use Modular architecture (default: Layered)
- `--standalone` - No helpers (100% independent)
- `-m, --module <name>` - Specify Go module name

**Examples:**

```bash
# Basic Layered project
loom new api

# Modular project
loom new app --modular

# Standalone project (no helpers)
loom new service --standalone

# With custom Go module
loom new api -m github.com/mycompany/my-api
```

**What does it generate?**
- Complete directory structure
- `go.mod` with Go module
- Functional HTTP server (**Gin** - high performance)
- Example user CRUD
- Health checks (/, /health, /ready)
- CORS middleware
- `.env.example` with configuration
- `Makefile` with useful commands
- `README.md` and `docs/API.md`
- Complete `.gitignore`

---

### `loom generate` - Generate Components

```bash
loom generate [type] [name] [flags]
```

**Subcommands:**

#### `loom generate module`

Generates a complete module with all its layers.

```bash
loom generate module products
```

**Layered generates:**
- `internal/app/handlers/products_handler.go`
- `internal/app/services/products_service.go`
- `internal/app/repositories/products_repository.go`
- `internal/app/models/product.go`
- `internal/app/dtos/products_dto.go`

**Modular generates:**
- `internal/modules/products/handler.go`
- `internal/modules/products/service.go`
- `internal/modules/products/repository.go`
- `internal/modules/products/model.go`
- `internal/modules/products/dto.go`
- `internal/modules/products/router.go`
- `internal/modules/products/validator.go`
- `internal/modules/products/errors.go`

#### `loom generate handler`

Generates only an HTTP handler.

```bash
loom generate handler orders
```

#### `loom generate service`

Generates only a service with business logic.

```bash
loom generate service email
```

#### `loom generate model`

Generates only a data model.

```bash
loom generate model Category
```

#### `loom generate middleware`

Generates HTTP middleware.

```bash
loom generate middleware auth
```

**Common Flags:**
- `--force` - Overwrite existing files
- `--dry-run` - Preview without creating files

**Automatic Detection:**
- Detects if you're in a Loom project
- Detects architecture (Layered/Modular)
- Generates appropriate code for the architecture

---

### `loom make` - Database Components

```bash
loom make [type] [name] [flags]
```

Generate database-related components with auto-registration. **Requires** `loom add orm gorm` first.

#### `loom make model`

Generate a GORM model with automatic registration in `models_all.go`.

```bash
loom make model Product
loom make model Category --force
```

**Generates:**
- Layered: `internal/app/models/product.go`
- Modular: `internal/models/product.go`

**Generated code:**
```go
package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name        string `gorm:"size:100;not null" json:"name"`
    Description string `gorm:"type:text" json:"description"`
    // Add your fields here
}
```

**Auto-registers in `models_all.go`:**
```go
var AllModels = []interface{}{
    &models.User{},
    &models.Product{},  // ← Automatically added
}
```

#### `loom make seeder`

Generate a seeder with automatic registration in `seeders_all.go`.

```bash
loom make seeder Product
loom make seeder Category --force
```

**Generates:** `internal/database/seeders/product_seeder.go`

**Generated code:**
```go
package seeders

import (
    "your-module/internal/app/models"
    "gorm.io/gorm"
)

type ProductSeeder struct{}

func (s *ProductSeeder) Run(db *gorm.DB) error {
    products := []models.Product{
        // Add your seed data here
    }
    
    for _, item := range products {
        db.FirstOrCreate(&item, models.Product{Name: item.Name})
    }
    return nil
}
```

**Auto-registers in `seeders_all.go`:**
```go
var AllSeeders = []Seeder{
    &UserSeeder{},
    &ProductSeeder{},  // ← Automatically added
}
```

---

### `loom db:*` - Database Commands (v1.1.2+)

Laravel Artisan-style commands for database management. **Requires** `loom add orm gorm` first.

```bash
loom db:migrate              # Run database migrations
loom db:migrate --seed       # Run migrations + seeders
loom db:fresh                # Drop all tables and re-run migrations
loom db:fresh --seed         # Fresh migration + seeders  
loom db:seed                 # Run seeders only
```

#### `loom db:migrate`

Run all pending database migrations using GORM AutoMigrate.

```bash
loom db:migrate
loom db:migrate --seed    # Also run seeders after migration
```

#### `loom db:fresh`

⚠️ **WARNING**: This is destructive! Drops ALL tables and re-creates them.

```bash
loom db:fresh             # Drop + migrate
loom db:fresh --seed      # Drop + migrate + seed
```

#### `loom db:seed`

Execute all seeders registered in `seeders_all.go`.

```bash
loom db:seed
```

#### Equivalencias Laravel → Loom

| Laravel Artisan | Loom CLI |
|-----------------|----------|
| `php artisan migrate` | `loom db:migrate` |
| `php artisan migrate:fresh` | `loom db:fresh` |
| `php artisan migrate --seed` | `loom db:migrate --seed` |
| `php artisan migrate:fresh --seed` | `loom db:fresh --seed` |
| `php artisan db:seed` | `loom db:seed` |

---

### `loom add` - Addon System

```bash
loom add [category] [name] [flags]
```

**Available categories:**

#### HTTP Routers

> **Note:** Since v1.1.0, **Gin is the default router**. You only need to use `loom add router` if you want to switch to a different router.

```bash
# Chi (lightweight, compatible with net/http)
loom add router chi

# Echo (minimalist)
loom add router echo

# Fiber (fastest, Express-like)
loom add router fiber
```

**What does it do?**
1. Updates `go.mod` with the new dependency
2. Replaces `internal/server/server.go` or `internal/platform/server/server.go`
3. Generates appropriate code for the router
4. Updates handlers to use the new router's API

**Note:** Replaces current router. Use `--force` to confirm.

#### ORMs

```bash
# GORM (full ORM) - Recommended
loom add orm gorm

# sqlc (generator from SQL)
loom add orm sqlc
```

**What does `loom add orm gorm` do?**

1. **Adds dependencies:**
   - `gorm.io/gorm v1.25.5`
   - `gorm.io/driver/postgres v1.5.4`
   - `golang.org/x/crypto v0.17.0` (bcrypt for seeders)
   - `github.com/spf13/cobra v1.9.1` (console CLI)

2. **Creates database layer:**
   ```
   internal/database/
   ├── database.go       # GORM connection manager
   ├── models_all.go     # Model registry for auto-migration
   └── seeders/
       ├── seeders_all.go      # Seeder interface
       ├── database_seeder.go  # Seeder orchestrator
       └── user_seeder.go      # Example seeder with bcrypt
   ```

3. **Creates console CLI:**
   ```
   cmd/console/main.go   # Cobra CLI with migrate/seed commands
   ```

4. **Updates Makefile** with database commands:
   ```bash
   make db-migrate     # Run migrations
   make db-seed        # Run seeders
   make db-fresh       # Drop all + migrate + seed
   ```

5. **Updates `.env.example`** with database variables:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=password
   DB_NAME=myapp
   DB_SSLMODE=disable
   ```

**Console commands:**
```bash
# Run migrations
go run cmd/console/main.go migrate

# Run migrations + seeders
go run cmd/console/main.go migrate --seed

# Fresh database (drop all + migrate + seed)
go run cmd/console/main.go migrate --fresh

# Run only seeders
go run cmd/console/main.go seed
```

**Add new models to migration:**
```go
// internal/database/models_all.go
var AllModels = []interface{}{
    &models.User{},
    &models.Product{},  // Add your models here
    &models.Order{},
}
```

**Add new seeders:**
```go
// internal/database/seeders/seeders_all.go
var AllSeeders = []Seeder{
    &UserSeeder{},
    &ProductSeeder{},  // Add your seeders here
}
```

#### Databases

```bash
# PostgreSQL
loom add database postgres

# MySQL
loom add database mysql

# MongoDB
loom add database mongodb

# Redis (cache)
loom add database redis
```

**What does it do?**
1. Adds database driver
2. Updates `.env.example` with DB_* variables
3. If you added Docker, updates `docker-compose.yml`
4. Creates configuration file in `internal/database/`

#### Authentication

```bash
# JWT
loom add auth jwt

# OAuth2
loom add auth oauth2
```

**What does JWT do?**
1. Adds dependency `github.com/golang-jwt/jwt/v5`
2. Creates `internal/auth/jwt.go` with generation/validation
3. Creates `internal/auth/middleware.go` to protect routes
4. Creates `internal/handlers/auth_handler.go` (login, register, refresh)
5. Updates `.env.example` with `JWT_SECRET`

#### Infrastructure

```bash
# Docker + Docker Compose
loom add docker
```

**What does it do?**
1. Creates optimized multi-stage `Dockerfile`
2. Creates `.dockerignore`
3. Creates `docker-compose.yml` with:
   - App service
   - PostgreSQL (if configured)
   - Volumes and networks
4. Updates `Makefile` with Docker commands

**Flags:**
- `--force` - Force installation (replaces existing)

**View available addons:**
```bash
loom add list
```

---

### `loom self-update` - Update Loom CLI (v1.1.3+)

```bash
loom self-update [version] [flags]
```

Update the Loom CLI itself to the latest or a specific version.

**Examples:**
```bash
loom self-update           # Update to latest version
loom self-update v1.1.3    # Update to specific version
loom self-update --check   # Check for updates without installing
```

**Flags:**
- `--check` - Check for updates without installing

**How it works:**
1. Fetches latest version from GitHub (releases or tags)
2. Compares with current version
3. Runs `go install` with direct proxy to bypass cache
4. Confirms successful update

---

### `loom upgrade` - Upgrade Projects

```bash
loom upgrade [flags]
```

**Flags:**
- `--no-backup` - Don't create backup before upgrading
- `--show-changes` - Show changes without applying
- `--restore <backup>` - Restore a specific backup

**Upgrade flow:**

1. **Detects current version:**
   - Reads `.loom` file
   - Or looks for comments in `go.mod`
   - If not found, assumes v0.1.0

2. **Compares with CLI version:**
   - If project is up to date, exits
   - If CLI is older, warns
   - If upgrade is available, continues

3. **Creates backup (optional):**
   - Folder `.loom-backups/backup-<timestamp>/`
   - Copies `internal/`, `cmd/`, `pkg/`, `go.mod`, `.loom`

4. **Applies incremental migrations:**
   - v0.1.0 → v0.2.0: Adds helpers
   - v0.2.0 → v0.3.0: Updates docs
   - v0.3.0 → v0.4.0: Creates `.loom` file
   - v0.4.0 → v0.5.0: Prepares upgrade system
   - v0.5.0 → v0.6.0: Prepares addon system
   - v1.0.0 → v1.1.0: Gin default + GORM addon

5. **Updates `.loom`:**
   ```
   # Loom Project Configuration
   version=1.1.0
   architecture=layered
   created_with=loom-cli
   ```

**`.loom` file:**

Loom creates this file to track version and configuration:

```
# Loom Project Configuration
version=1.1.0
architecture=layered
created_with=loom-cli
```

**Restore backup:**

```bash
# List available backups
ls .loom-backups/

# Restore
loom upgrade --restore backup-20251027-153045
```

---

### `loom version`

```bash
loom version
```

Shows:
- Loom CLI version
- Current project version (if in a Loom project)
- Update status

**Example output:**

```
🔧 Loom CLI v1.1.0
📦 Current project: v1.0.0

⚠️  Your project uses an old version of Loom
💡 Update with: loom upgrade
```

---

## 📦 Helpers API

If you don't use `--standalone`, your project includes `pkg/helpers`.

### Response Helpers

```go
import "github.com/geomark27/loom-go/pkg/helpers"

// Success (200)
helpers.RespondSuccess(w, data, "Operation successful")
// { "success": true, "message": "...", "data": {...} }

// Created (201)
helpers.RespondCreated(w, user, "User created")

// Error (400, 404, 500, etc.)
helpers.RespondError(w, err, http.StatusBadRequest)
// { "success": false, "error": "...", "code": 400 }

// No Content (204)
helpers.RespondNoContent(w)
```

### Validator

```go
type UserDTO struct {
    Name  string `json:"name" validate:"required,min=3,max=50"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"gte=18,lte=120"`
}

dto := &UserDTO{Name: "Jo", Email: "invalid"}

// Validate
errors := helpers.ValidateStruct(dto)
if len(errors) > 0 {
    helpers.RespondError(w, fmt.Errorf("%v", errors), http.StatusBadRequest)
    return
}
```

**Supported validation tags:**
- `required` - Required field
- `email` - Valid email
- `min=N` - Minimum length/value
- `max=N` - Maximum length/value
- `gte=N` - Greater than or equal to
- `lte=N` - Less than or equal to
- `oneof=val1 val2` - One of the values
- `url` - Valid URL
- `uuid` - Valid UUID

### Logger

```go
logger := helpers.NewLogger()

// Log levels
logger.Info("Server started", "port", 8080)
logger.Warn("High memory usage", "usage", 85)
logger.Error("Database connection failed", "error", err)
logger.Debug("Query executed", "sql", query, "duration", duration)

// Structured context
logger.Info("User created",
    "user_id", user.ID,
    "username", user.Name,
    "ip", r.RemoteAddr,
)
```

### Predefined Errors

```go
// Common HTTP errors
helpers.ErrNotFound          // 404
helpers.ErrBadRequest        // 400
helpers.ErrUnauthorized      // 401
helpers.ErrForbidden         // 403
helpers.ErrInternalServer    // 500
helpers.ErrConflict          // 409

// Usage
if user == nil {
    helpers.RespondError(w, helpers.ErrNotFound, http.StatusNotFound)
    return
}
```

### Context Utilities

```go
// Get values from context
userID := helpers.GetUserIDFromContext(ctx)
requestID := helpers.GetRequestIDFromContext(ctx)

// Add values to context
ctx = helpers.SetUserIDInContext(ctx, userID)
ctx = helpers.SetRequestIDInContext(ctx, requestID)
```

**Update helpers:**

```bash
go get -u github.com/geomark27/loom-go/pkg/helpers
go mod tidy
```

---

## ✅ Best Practices

### 1. Project Structure

**DO:**
- ✅ Use `internal/` for private code
- ✅ Use `pkg/` for reusable code
- ✅ Separate configuration in `internal/config/`
- ✅ One handler per file
- ✅ Dependency injection

**DON'T:**
- ❌ Don't put business logic in handlers
- ❌ Don't access DB directly from handlers
- ❌ Don't use global variables

### 2. Handlers (with Gin)

```go
// ✅ GOOD: Clean handler with Gin
func (h *UserHandler) Create(c *gin.Context) {
    var dto UserDTO
    if err := c.ShouldBindJSON(&dto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status": "error",
            "message": err.Error(),
        })
        return
    }

    user, err := h.service.CreateUser(&dto)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": "error",
            "message": err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "status": "success",
        "message": "User created",
        "data": user,
    })
}

// ❌ BAD: Business logic in handler
func (h *UserHandler) Create(c *gin.Context) {
    // ... complex validations
    // ... DB queries
    // ... sending emails
    // ... business logic
}
```

### 3. Services

```go
// ✅ GOOD: Service with business logic
type UserService struct {
    repo UserRepository
    emailService EmailService
}

func (s *UserService) CreateUser(dto *UserDTO) (*User, error) {
    // Business validation
    if exists := s.repo.ExistsByEmail(dto.Email); exists {
        return nil, ErrEmailAlreadyExists
    }

    // Create user
    user := &User{
        Name: dto.Name,
        Email: dto.Email,
    }

    if err := s.repo.Create(user); err != nil {
        return nil, err
    }

    // Send welcome email
    s.emailService.SendWelcome(user.Email)

    return user, nil
}
```

### 4. Repositories

```go
// ✅ GOOD: Repository with only data access
type UserRepository struct {
    db *sql.DB
}

func (r *UserRepository) Create(user *User) error {
    query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
    return r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

func (r *UserRepository) FindByID(id int) (*User, error) {
    // ... query
}

// ❌ BAD: Repository with business logic
func (r *UserRepository) Create(user *User) error {
    // ❌ Business validations
    // ❌ Sending emails
    // ❌ Calls to other services
}
```

### 5. DTOs vs Models

**Models** - Internal/DB representation:
```go
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`  // Hidden in JSON
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

**DTOs** - HTTP transfer:
```go
type CreateUserDTO struct {
    Name     string `json:"name" validate:"required,min=3"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}

type UpdateUserDTO struct {
    Name  *string `json:"name,omitempty" validate:"omitempty,min=3"`
    Email *string `json:"email,omitempty" validate:"omitempty,email"`
}
```

### 6. Error Handling

```go
// ✅ GOOD: Domain-specific errors
var (
    ErrUserNotFound      = errors.New("user not found")
    ErrEmailAlreadyExists = errors.New("email already exists")
    ErrInvalidPassword    = errors.New("invalid password")
)

// Usage
user, err := s.repo.FindByID(id)
if err != nil {
    if err == sql.ErrNoRows {
        return nil, ErrUserNotFound
    }
    return nil, err
}
```

### 7. Testing

```go
// Use interfaces for testing
type UserRepository interface {
    Create(user *User) error
    FindByID(id int) (*User, error)
}

// Mock in tests
type MockUserRepository struct {
    CreateFunc func(user *User) error
}

func (m *MockUserRepository) Create(user *User) error {
    return m.CreateFunc(user)
}
```

---

## 🔧 Troubleshooting

### Project not detected

**Error:**
```
error: no valid Loom project detected
```

**Solution:**
1. Verify you're in the project directory
2. Verify that `internal/app/` or `internal/modules/` exists
3. Verify that `go.mod` exists

### Cannot change router

**Error:**
```
conflict detected: gin is installed
```

**Solution:**
```bash
loom add router chi --force
```

### Backup won't restore

**Error:**
```
backup not found: backup-xxx
```

**Solution:**
1. List backups: `ls .loom-backups/`
2. Use the full name: `loom upgrade --restore backup-20251027-153045`

### Helpers not found

**Error:**
```
cannot find package "github.com/geomark27/loom-go/pkg/helpers"
```

**Solution:**
```bash
go get github.com/geomark27/loom-go/pkg/helpers
go mod tidy
```

### Port in use

**Error:**
```
listen tcp :8080: bind: address already in use
```

**Solution:**
1. Change port in `.env`: `PORT=8081`
2. Or kill process: `lsof -ti:8080 | xargs kill`

---

## 📞 Support

Problems? Questions?

- **Issues**: [github.com/geomark27/loom-go/issues](https://github.com/geomark27/loom-go/issues)
- **Discussions**: [github.com/geomark27/loom-go/discussions](https://github.com/geomark27/loom-go/discussions)

---

**Last updated:** November 24, 2025 - v1.1.0
