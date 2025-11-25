# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/).

---

## [1.1.0] - 2025-11-24 🚀

### ⚠️ BREAKING CHANGES
- **Default router changed**: Gorilla Mux replaced with **Gin** as the default router
  - All generated projects now use `github.com/gin-gonic/gin v1.10.0`
  - **Both architectures** (Layered and Modular) updated to use Gin
  - Handler signatures changed from `http.ResponseWriter, *http.Request` to `*gin.Context`
  - Route parameters now use `:id` syntax instead of `{id}`
  - Module routes now use `*gin.RouterGroup` instead of `*mux.Router`

### ✨ Added
- **Complete GORM addon** (`loom add orm gorm`):
  - Full GORM v1.25.5 + PostgreSQL driver installation
  - Database connection manager (`internal/database/database.go`)
  - Model registry for auto-migration (`internal/database/models_all.go`)
  - Seeder system with interface and registry (`internal/database/seeders/`)
  - Example UserSeeder with bcrypt password hashing
  - Console CLI (`cmd/console/main.go`) with Cobra commands:
    - `go run cmd/console/main.go migrate` - Run migrations
    - `go run cmd/console/main.go migrate --fresh` - Drop all tables and migrate
    - `go run cmd/console/main.go migrate --seed` - Migrate with seeders
    - `go run cmd/console/main.go seed` - Run seeders only
  - Makefile targets: `make db-migrate`, `make db-seed`, `make db-fresh`
  - Auto-updates `.env.example` with DB variables

- **New `loom make` commands** for database components:
  - `loom make model <Name>` - Generate GORM model with auto-registration in `models_all.go`
  - `loom make seeder <Name>` - Generate seeder with auto-registration in `seeders_all.go`
  - Both commands support Layered and Modular architectures
  - Generated code includes GORM tags, JSON tags, and example fields

- **User model with GORM support**:
  - Uses `gorm.Model` (ID, CreatedAt, UpdatedAt, DeletedAt)
  - GORM tags for validation (size, not null, uniqueIndex)
  - Password field with `json:"-"` for security
  - IsActive boolean field

- **Architecture-aware templates**:
  - Templates detect layered vs modular architecture
  - Correct paths for config and models based on architecture
  - Dynamic module name injection from go.mod

### 🔧 Changed
- **Gin as default router** (both architectures):
  - **Layered architecture**:
    - `layered/server.go.tmpl` - Uses `gin.Engine` and `gin.Default()`
    - `layered/routes.go.tmpl` - Uses `router.Group()` and `api.GET/POST/PUT/DELETE`
    - `layered/user_handler.go.tmpl` - Uses `*gin.Context` and `c.JSON()`
    - `layered/health_handler.go.tmpl` - Uses Gin response methods
  - **Modular architecture**:
    - `modular/server.go.tmpl` - Uses `gin.Engine` with integrated CORS
    - `modular/router.go.tmpl` - Health routes with Gin handlers
    - `modular/handler.go.tmpl` - Uses `*gin.Context` and `c.ShouldBindJSON()`
    - `modular/module.go.tmpl` - Routes use `*gin.RouterGroup`
  - CORS middleware now uses `gin.HandlerFunc` (inline in server.go)

- **go.mod template** - Now includes `github.com/gin-gonic/gin v1.10.0`

### 🏗️ New Project Structure (after `loom add orm gorm`)
```
my-project/
├── cmd/
│   ├── my-project/main.go
│   └── console/main.go          # NEW: Database CLI
├── internal/
│   ├── database/                 # NEW: Database layer
│   │   ├── database.go           # GORM connection
│   │   ├── models_all.go         # Model registry
│   │   └── seeders/
│   │       ├── seeders_all.go    # Seeder interface
│   │       ├── database_seeder.go# Seeder orchestrator
│   │       └── user_seeder.go    # Example seeder
│   └── ...
└── Makefile                      # Updated with db-* commands
```

### 📝 Migration Guide from v1.0.x

1. **Router change** (if updating existing project):
   - Replace `github.com/gorilla/mux` with `github.com/gin-gonic/gin`
   - Update handler signatures to use `*gin.Context`
   - Replace `mux.Vars(r)["id"]` with `c.Param("id")`
   - Replace `json.NewEncoder(w).Encode()` with `c.JSON()`

2. **Using GORM**:
   - After `loom add orm gorm`, manually add DB fields to your Config struct
   - Add `GetDBConnectionString()` method to Config
   - Load DB environment variables in `config.Load()`

---

## [1.0.6] - 2025-11-03

### 🔧 Changed
- **Complete dynamic version system**: CLI version now also reads from `version.Current`
  - Removed hardcoded version string from `root.go`
  - CLI `--version` now uses `version.Current.String()` dynamically
  - TRUE single source of truth: only `version.go` needs updates

### 🎯 Impact
- **Zero manual updates**: Change version ONCE in `version.go`, everything syncs
  - CLI version ✅ (from `version.Current`)
  - Generated projects ✅ (via `{{.LoomVersion}}`)
  - Perfect synchronization guaranteed

### 📝 Note
This completes the dynamic version system started in v1.0.5.
Now the entire codebase has a true single source of truth for versioning.

---

## [1.0.5] - 2025-11-03

### 🔧 Changed
- **Dynamic version injection**: Template now uses `{{.LoomVersion}}` for automatic version updates
  - Added `LoomVersion` field to `ProjectConfig`
  - Version is now injected automatically from `version.Current`
  - Eliminates manual version updates in templates

### 🐛 Fixed
- Fixed embedded go.mod template referencing outdated version
- Generated projects now correctly use the current Loom version

### 🎯 Impact
- No more manual version updates needed in templates
- Single source of truth for version (`internal/version/version.go`)
- Future releases will automatically use correct version in generated projects

---

## [1.0.4] - 2025-11-03

### 🌍 Changed
- **Complete internationalization**: All text translated from Spanish to English
  - CLI commands and messages
  - Code comments and documentation
  - README.md and DOCS.md
  - Embedded templates
  - Error messages and user-facing strings
- Project is now ready for international community adoption

### 📝 Documentation
- All documentation now in English
- Maintained all examples and code snippets
- Updated for global audience

---

## [1.0.3] - 2025-10-27

### 🔧 Changed
- Version bump for Go proxy compatibility
- Improved module versioning

---

## [1.0.1] - 2025-10-27

### 🔧 Changed
- Version bump for Go proxy compatibility

---

## [1.0.0] - 2025-10-27 🎉

### 🚀 Official Release v1.0.0

**First stable production version** of Loom. This version marks the project's maturity with all core functionalities complete and tested.

### ✨ Added in v1.0.0

- **`loom add` command**: Complete addon system to extend projects
  - **Routers**: `loom add router [gin|chi|echo]` - Replaces Gorilla Mux
  - **ORMs**: `loom add orm [gorm|sqlc]` - Adds ORMs (structure ready, implementation in progress)
  - **Databases**: `loom add database [postgres|mysql|mongodb|redis]` - Configures databases
    - PostgreSQL fully implemented
    - MySQL, MongoDB, Redis with base structure
  - **Auth**: `loom add auth [jwt|oauth2]` - Authentication system (structure ready)
  - **Docker**: `loom add docker` - Adds Dockerfile, docker-compose.yml and .dockerignore
  - `loom add list` - Lists all available addons

- **Addon System**: Extensible architecture with interfaces
  - `AddonManager` with registration and conflict detection
  - `ProjectDetector` to identify installed technologies
  - Helpers to update `go.mod`, `.env.example` and imports

- **Advanced code generation**:
  - Complete templates for Gin, Chi and Echo
  - Intelligent rewriting of `server.go` based on router
  - Optimized multi-stage Dockerfile
  - docker-compose.yml with database detection
  - Automatic Makefile updates with Docker commands

### 📚 Documentation

- **Complete documentation consolidation**:
  - `README.md` - Main user documentation (installation, quick start, examples)
  - `DOCS.md` - Complete technical documentation (architectures, APIs, troubleshooting)
  - `CHANGELOG.md` - Version history
- Removed 13+ obsolete and fragmented .md files
- Detailed guides for Layered vs Modular architectures
- Complete Helpers API documentation
- Best practices and troubleshooting

### 🔧 Changed

- CLI version updated to 1.0.0
- Simplified documentation structure (from 30+ files to only 3)
- Improved help messages and next steps

### 🎯 Complete Features in v1.0.0

1. ✅ **Project generation** (`loom new`) - Layered and Modular
2. ✅ **Component generation** (`loom generate`) - Modules, handlers, services, etc.
3. ✅ **Addon system** (`loom add`) - Routers, DBs, Docker
4. ✅ **Update system** (`loom upgrade`) - With backups and restore
5. ✅ **Versioning** (`loom version`) - Version tracking
6. ✅ **Helpers package** - Response, Validator, Logger, Errors
7. ✅ **Complete documentation** - User-facing and technical

### 📦 Installation

```bash
go install github.com/geomark27/loom-go/cmd/loom@latest
```

---

## [0.6.0] - 2025-10-27

### ✨ Added

- Base addon system (preparation for v1.0.0)
- Structure of `internal/addon/`
- Interfaces for future extensions

---

## [0.5.0] - 2025-10-27

### ✨ Added

- **`loom upgrade` command**: Complete project update system
  - Automatic project version detection
  - Automatic backup before updating (with `--no-backup` option)
  - Application of incremental migrations between versions
  - Backup restoration with `--restore`
  - Change preview with `--show-changes`
- **`loom version` command**: Shows CLI and current project version
- **Versioning system**: Complete infrastructure for version management
  - `internal/version/`: Version detection and comparison
  - `internal/upgrader/`: Update and backup system
  - `.loom` file for project version tracking
- **Integrated changelog**: Shows changes between versions during upgrade

### 🔧 Changed

- CLI version updated to 0.5.0
- Projects now include `.loom` file with metadata

### 📚 Documentation

- Documentation of upgrade command in help
- Examples of upgrade and restore usage

---

## [0.4.0] - 2025-10-27

### ✨ Added

- **`loom generate` command**: Individual component generation in existing projects
  - `loom generate module <name>`: Generates complete module (handler, service, repository, model, DTO)
  - `loom generate handler <name>`: Generates handler only
  - `loom generate service <name>`: Generates service only
  - `loom generate model <name>`: Generates model only
  - `loom generate middleware <name>`: Generates HTTP middleware
- **Automatic architecture detection**: Command generates appropriate code based on Layered/Modular
- **Global flags for generate**:
  - `--force`: Overwrites existing files
  - `--dry-run`: Preview without creating files
- **Command aliases**: `gen`, `g` for generate; aliases for subcommands
- **Name validation**: Verifies component names are valid
- **Complete templates**: Templates for all component types in both architectures

### 🔧 Changed

- CLI version updated to 0.4.0
- Improved internal structure with `internal/generator/`

### 📚 Documentation

- Complete documentation of generate command
- Usage examples for each subcommand
- Guide for next steps after generating components

---

## [0.3.0] - 2025-10-27

### ✨ Added

- **Dual Architecture**: Support for Layered (default) and Modular architecture
- `--modular` flag to generate projects with modular architecture by domains
- Automatic architecture detection in existing projects
- Example `users` module in modular architecture
- Event Bus for inter-module communication
- `ports.go` file with interfaces in modules
- Separate templates for each architecture (`templates/layered/` and `templates/modular/`)

### 🔧 Changed

- Reorganization of internal template structure
- Clear separation between `platform` (infrastructure) and `app`/`modules` (logic)
- Improved GitHub user detection from config
- Informative messages about selected architecture

### 📚 Documentation

- Guides on when to use each architecture
- Examples of both architectures
- Event Bus and inter-module communication documentation

---

## [0.2.0] - 2025-10-XX

### ✨ Added

- **Helpers Package** (`pkg/helpers/`): Reusable utilities library
  - `response.go`: Standardized HTTP responses
  - `validator.go`: Struct validation with tags
  - `logger.go`: Structured logging
  - `errors.go`: Error handling with context
  - `context.go`: Context utilities
- `--standalone` flag to generate projects without helpers
- `UseHelpers` field in ProjectConfig
- Support for 100% independent projects

### 🔧 Changed

- Projects now include helpers by default
- Clearer structure with `internal/platform` for infrastructure
- Separation between `internal/app` (business) and `internal/shared` (utilities)

### 📚 Documentation

- Documentation of available helpers
- Helpers usage guide
- Examples of projects with and without helpers
- "Hybrid Model" section in README

---

## [0.1.0] - 2025-10-XX

### ✨ Added - Initial Release

- `loom new <name>` command to create projects
- Complete project structure generation:
  - `cmd/` - Entry point
  - `internal/app/` - Business logic (handlers, services, repositories)
  - `internal/config/` - Configuration
  - `internal/server/` - HTTP server
  - `pkg/` - Public code
  - `docs/` - Documentation
- Layered architecture as standard
- Functional REST API with example user CRUD
- Implemented health checks
- Configured CORS middleware
- HTTP server with Gorilla Mux
- Embedded templates with `text/template`
- File generation:
  - `go.mod`
  - `README.md`
  - `.gitignore`
  - `.env.example`
  - `Makefile`
  - `docs/API.md`
- `-m, --module` flag to specify Go module name
- Auto-detection of GitHub user from git config
- Project name validation
- Post-generation informative messages

### 📚 Documentation

- Main README with installation guide
- DESCRIPCION.md with detailed explanation
- INSTALACION.md with step-by-step guide
- FLUJOS_REALES.md with use cases
- INSTALL_FROM_GITHUB.md with GitHub installation
- COMPATIBILIDAD_MULTIPLATAFORMA.md
- DISTRIBUCION_GLOBAL.md
- VERIFICACION.md

### 🛠️ Infrastructure

- Go modules configuration
- Dependency on Cobra for CLI
- Internal package structure
- Template system

---

## Types of Changes

- **✨ Added** (`Added`): For new features
- **🔧 Changed** (`Changed`): For changes in existing functionality
- **❌ Deprecated** (`Deprecated`): For features that will be removed
- **🗑️ Removed** (`Removed`): For removed features
- **🐛 Fixed** (`Fixed`): For bug fixes
- **🔒 Security** (`Security`): For security vulnerabilities

---

## Version Links

- [1.0.6]: https://github.com/geomark27/loom-go/releases/tag/v1.0.6
- [1.0.5]: https://github.com/geomark27/loom-go/releases/tag/v1.0.5
- [1.0.4]: https://github.com/geomark27/loom-go/releases/tag/v1.0.4
- [1.0.3]: https://github.com/geomark27/loom-go/releases/tag/v1.0.3
- [1.0.1]: https://github.com/geomark27/loom-go/releases/tag/v1.0.1
- [1.0.0]: https://github.com/geomark27/loom-go/releases/tag/v1.0.0
- [0.6.0]: https://github.com/geomark27/loom-go/compare/v0.5.0...v0.6.0
- [0.5.0]: https://github.com/geomark27/loom-go/compare/v0.4.0...v0.5.0
- [0.4.0]: https://github.com/geomark27/loom-go/compare/v0.3.0...v0.4.0
- [0.3.0]: https://github.com/geomark27/loom-go/compare/v0.2.0...v0.3.0
- [0.2.0]: https://github.com/geomark27/loom-go/compare/v0.1.0...v0.2.0
- [0.1.0]: https://github.com/geomark27/loom-go/releases/tag/v0.1.0
