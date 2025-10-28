# 📖 Documentación Técnica - Loom

Documentación completa para desarrolladores que usan Loom.

## 📑 Tabla de Contenidos

- [Conceptos Fundamentales](#conceptos-fundamentales)
- [Arquitecturas](#arquitecturas)
- [Comandos Detallados](#comandos-detallados)
- [Sistema de Addons](#sistema-de-addons)
- [Sistema de Versionado](#sistema-de-versionado)
- [Helpers API](#helpers-api)
- [Buenas Prácticas](#buenas-prácticas)
- [Troubleshooting](#troubleshooting)

---

## 🎯 Conceptos Fundamentales

### ¿Qué NO es Loom?

Loom **NO es un framework**. No añade overhead en runtime ni te obliga a usar una API específica.

### ¿Qué SÍ es Loom?

Loom es un **generador de código**:
- ✅ Genera código idiomático de Go
- ✅ Crea estructura de proyecto profesional
- ✅ Proporciona herramientas para extender el proyecto
- ✅ No añade dependencias obligatorias (con `--standalone`)
- ✅ Sin magia, sin reflection innecesario, sin DSLs

### Filosofía: "Cerrado para modificación, Abierto para extensión"

Una vez generado el código:
- Es **TU código**, no de Loom
- Puedes modificarlo como quieras
- No hay vendor lock-in
- Puedes remover Loom del proyecto si quieres

---

## 🏗️ Arquitecturas

Loom soporta dos arquitecturas diferentes para diferentes necesidades:

### 1. Arquitectura Layered (Por Capas)

**Cuándo usar:**
- APIs REST simples
- Microservicios pequeños
- Proyectos con un solo dominio
- Equipos pequeños (1-3 desarrolladores)

**Estructura:**

```
proyecto/
├── cmd/
│   └── proyecto/
│       └── main.go
├── internal/
│   ├── app/
│   │   ├── handlers/      # Capa de presentación (HTTP)
│   │   ├── services/      # Capa de negocio
│   │   ├── repositories/  # Capa de datos
│   │   ├── models/        # Modelos de dominio
│   │   ├── dtos/          # Transfer objects
│   │   └── middleware/    # Middleware HTTP
│   ├── config/            # Configuración
│   └── server/            # Servidor HTTP
└── pkg/
    └── helpers/           # Utilidades (opcional)
```

**Flujo de datos:**
```
HTTP Request → Handler → Service → Repository → Database
                  ↓
HTTP Response ← Handler ← Service ← Repository ← Database
```

**Ventajas:**
- ✅ Simple y directo
- ✅ Fácil de entender
- ✅ Buena separación de responsabilidades
- ✅ Bajo acoplamiento entre capas

**Desventajas:**
- ❌ Difícil de escalar con muchos dominios
- ❌ Puede volverse un "big ball of mud" en proyectos grandes

### 2. Arquitectura Modular (Por Dominios)

**Cuándo usar:**
- Aplicaciones grandes
- Múltiples dominios de negocio
- Equipos grandes (4+ desarrolladores)
- Necesitas separación clara por feature/dominio

**Estructura:**

```
proyecto/
├── cmd/
│   └── proyecto/
│       └── main.go
├── internal/
│   ├── modules/                # Módulos de dominio
│   │   ├── users/
│   │   │   ├── handler.go      # HTTP handler
│   │   │   ├── service.go      # Lógica de negocio
│   │   │   ├── repository.go   # Acceso a datos
│   │   │   ├── model.go        # Modelo
│   │   │   ├── dto.go          # DTOs
│   │   │   ├── router.go       # Rutas del módulo
│   │   │   ├── validator.go    # Validaciones específicas
│   │   │   ├── errors.go       # Errores del dominio
│   │   │   └── ports.go        # Interfaces (DIP)
│   │   └── products/
│   │       └── ...             # Misma estructura
│   └── platform/               # Infraestructura compartida
│       ├── server/
│       ├── config/
│       └── eventbus/           # Comunicación entre módulos
└── pkg/
    └── helpers/
```

**Comunicación entre módulos:**

```go
// Usar Event Bus (recomendado)
eventbus.Publish("user.created", UserCreatedEvent{UserID: id})

// O interfaces (DIP - Dependency Inversion Principle)
type UserFinder interface {
    FindByID(id int) (*User, error)
}
```

**Ventajas:**
- ✅ Escalabilidad horizontal
- ✅ Módulos independientes
- ✅ Fácil trabajar en equipo (cada dev un módulo)
- ✅ Bajo acoplamiento entre dominios
- ✅ Fácil extraer a microservicios

**Desventajas:**
- ❌ Más complejo inicialmente
- ❌ Overhead de comunicación entre módulos
- ❌ Más archivos y estructura

### Comparación Rápida

| Aspecto | Layered | Modular |
|---------|---------|---------|
| **Complejidad** | Baja | Media-Alta |
| **Escalabilidad** | Limitada | Excelente |
| **Curva de aprendizaje** | Suave | Empinada |
| **Tamaño ideal** | Pequeño-Mediano | Mediano-Grande |
| **Equipo** | 1-3 devs | 4+ devs |
| **Dominios** | 1-2 | 3+ |

---

## 🎨 Comandos Detallados

### `loom new` - Crear Proyectos

```bash
loom new [nombre] [flags]
```

**Flags:**
- `--modular` - Usa arquitectura Modular (por defecto: Layered)
- `--standalone` - Sin helpers (100% independiente)
- `-m, --module <name>` - Especifica el nombre del módulo Go

**Ejemplos:**

```bash
# Proyecto básico Layered
loom new api

# Proyecto Modular
loom new app --modular

# Proyecto standalone (sin helpers)
loom new service --standalone

# Con módulo Go personalizado
loom new api -m github.com/miempresa/mi-api
```

**¿Qué genera?**
- Estructura completa de directorios
- `go.mod` con módulo Go
- Servidor HTTP funcional (Gorilla Mux)
- CRUD de usuarios de ejemplo
- Health checks (/, /health, /ready)
- Middleware CORS
- `.env.example` con configuración
- `Makefile` con comandos útiles
- `README.md` y `docs/API.md`
- `.gitignore` completo

---

### `loom generate` - Generar Componentes

```bash
loom generate [tipo] [nombre] [flags]
```

**Subcomandos:**

#### `loom generate module`

Genera un módulo completo con todas sus capas.

```bash
loom generate module products
```

**Layered genera:**
- `internal/app/handlers/products_handler.go`
- `internal/app/services/products_service.go`
- `internal/app/repositories/products_repository.go`
- `internal/app/models/product.go`
- `internal/app/dtos/products_dto.go`

**Modular genera:**
- `internal/modules/products/handler.go`
- `internal/modules/products/service.go`
- `internal/modules/products/repository.go`
- `internal/modules/products/model.go`
- `internal/modules/products/dto.go`
- `internal/modules/products/router.go`
- `internal/modules/products/validator.go`
- `internal/modules/products/errors.go`

#### `loom generate handler`

Genera solo un handler HTTP.

```bash
loom generate handler orders
```

#### `loom generate service`

Genera solo un service con lógica de negocio.

```bash
loom generate service email
```

#### `loom generate model`

Genera solo un modelo de datos.

```bash
loom generate model Category
```

#### `loom generate middleware`

Genera middleware HTTP.

```bash
loom generate middleware auth
```

**Flags Comunes:**
- `--force` - Sobrescribe archivos existentes
- `--dry-run` - Vista previa sin crear archivos

**Detección Automática:**
- Detecta si estás en un proyecto Loom
- Detecta arquitectura (Layered/Modular)
- Genera código apropiado para la arquitectura

---

### `loom add` - Sistema de Addons

```bash
loom add [categoría] [nombre] [flags]
```

**Categorías disponibles:**

#### Routers HTTP

```bash
# Gin (recomendado para performance)
loom add router gin

# Chi (ligero, compatible con net/http)
loom add router chi

# Echo (minimalista)
loom add router echo
```

**¿Qué hace?**
1. Actualiza `go.mod` con la nueva dependencia
2. Reemplaza `internal/server/server.go`
3. Genera código apropiado para el router
4. Te advierte que actualices handlers manualmente

**Nota:** Reemplaza Gorilla Mux por defecto. Usa `--force` para confirmar.

#### ORMs

```bash
# GORM (ORM completo)
loom add orm gorm

# sqlc (generador desde SQL)
loom add orm sqlc
```

**¿Qué hace?**
1. Añade dependencia del ORM
2. Crea `internal/database/` con configuración
3. Actualiza repositories para usar ORM
4. Configura migraciones (GORM)

#### Bases de Datos

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

**¿Qué hace?**
1. Añade driver de la base de datos
2. Actualiza `.env.example` con variables DB_*
3. Si añadiste Docker, actualiza `docker-compose.yml`
4. Crea archivo de configuración en `internal/database/`

#### Autenticación

```bash
# JWT
loom add auth jwt

# OAuth2
loom add auth oauth2
```

**¿Qué hace JWT?**
1. Añade dependencia `github.com/golang-jwt/jwt/v5`
2. Crea `internal/auth/jwt.go` con generación/validación
3. Crea `internal/auth/middleware.go` para proteger rutas
4. Crea `internal/handlers/auth_handler.go` (login, register, refresh)
5. Actualiza `.env.example` con `JWT_SECRET`

#### Infrastructure

```bash
# Docker + Docker Compose
loom add docker
```

**¿Qué hace?**
1. Crea `Dockerfile` multi-stage optimizado
2. Crea `.dockerignore`
3. Crea `docker-compose.yml` con:
   - Servicio de la app
   - PostgreSQL (si está configurado)
   - Volúmenes y networks
4. Actualiza `Makefile` con comandos Docker

**Flags:**
- `--force` - Fuerza instalación (reemplaza existente)

**Ver addons disponibles:**
```bash
loom add list
```

---

### `loom upgrade` - Actualizar Proyectos

```bash
loom upgrade [flags]
```

**Flags:**
- `--no-backup` - No crear backup antes de actualizar
- `--show-changes` - Mostrar cambios sin aplicar
- `--restore <backup>` - Restaurar un backup específico

**Flujo de upgrade:**

1. **Detecta versión actual:**
   - Lee archivo `.loom`
   - O busca comentarios en `go.mod`
   - Si no encuentra, asume v0.1.0

2. **Compara con versión del CLI:**
   - Si proyecto está actualizado, termina
   - Si CLI es más antiguo, advierte
   - Si hay upgrade disponible, continúa

3. **Crea backup (opcional):**
   - Carpeta `.loom-backups/backup-<timestamp>/`
   - Copia `internal/`, `cmd/`, `pkg/`, `go.mod`, `.loom`

4. **Aplica migraciones incrementales:**
   - v0.1.0 → v0.2.0: Añade helpers
   - v0.2.0 → v0.3.0: Actualiza docs
   - v0.3.0 → v0.4.0: Crea archivo `.loom`
   - v0.4.0 → v0.5.0: Prepara sistema upgrade
   - v0.5.0 → v0.6.0: Prepara sistema addons

5. **Actualiza `.loom`:**
   ```
   # Loom Project Configuration
   version=0.6.0
   architecture=layered
   created_with=loom-cli
   ```

**Archivo `.loom`:**

Loom crea este archivo para trackear versión y configuración:

```
# Loom Project Configuration
version=0.6.0
architecture=layered
created_with=loom-cli
```

**Restaurar backup:**

```bash
# Listar backups disponibles
ls .loom-backups/

# Restaurar
loom upgrade --restore backup-20251027-153045
```

---

### `loom version`

```bash
loom version
```

Muestra:
- Versión del CLI de Loom
- Versión del proyecto actual (si está en proyecto Loom)
- Estado de actualización

**Salida ejemplo:**

```
🔧 Loom CLI v0.6.0
📦 Proyecto actual: v0.4.0

⚠️  Tu proyecto usa una versión antigua de Loom
💡 Actualiza con: loom upgrade
```

---

## 📦 Helpers API

Si no usas `--standalone`, tu proyecto incluye `pkg/helpers`.

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

// Validar
errors := helpers.ValidateStruct(dto)
if len(errors) > 0 {
    helpers.RespondError(w, fmt.Errorf("%v", errors), http.StatusBadRequest)
    return
}
```

**Tags de validación soportadas:**
- `required` - Campo obligatorio
- `email` - Email válido
- `min=N` - Longitud/valor mínimo
- `max=N` - Longitud/valor máximo
- `gte=N` - Mayor o igual que
- `lte=N` - Menor o igual que
- `oneof=val1 val2` - Uno de los valores
- `url` - URL válida
- `uuid` - UUID válido

### Logger

```go
logger := helpers.NewLogger()

// Niveles de log
logger.Info("Server started", "port", 8080)
logger.Warn("High memory usage", "usage", 85)
logger.Error("Database connection failed", "error", err)
logger.Debug("Query executed", "sql", query, "duration", duration)

// Contexto estructurado
logger.Info("User created",
    "user_id", user.ID,
    "username", user.Name,
    "ip", r.RemoteAddr,
)
```

### Errores Predefinidos

```go
// Errores HTTP comunes
helpers.ErrNotFound          // 404
helpers.ErrBadRequest        // 400
helpers.ErrUnauthorized      // 401
helpers.ErrForbidden         // 403
helpers.ErrInternalServer    // 500
helpers.ErrConflict          // 409

// Uso
if user == nil {
    helpers.RespondError(w, helpers.ErrNotFound, http.StatusNotFound)
    return
}
```

### Context Utilities

```go
// Obtener valores del contexto
userID := helpers.GetUserIDFromContext(ctx)
requestID := helpers.GetRequestIDFromContext(ctx)

// Añadir valores al contexto
ctx = helpers.SetUserIDInContext(ctx, userID)
ctx = helpers.SetRequestIDInContext(ctx, requestID)
```

**Actualizar helpers:**

```bash
go get -u github.com/geomark27/loom-go/pkg/helpers
go mod tidy
```

---

## ✅ Buenas Prácticas

### 1. Estructura de Proyecto

**DO:**
- ✅ Usa `internal/` para código privado
- ✅ Usa `pkg/` para código reutilizable
- ✅ Separa configuración en `internal/config/`
- ✅ Un handler por archivo
- ✅ Inyección de dependencias

**DON'T:**
- ❌ No pongas lógica de negocio en handlers
- ❌ No accedas a BD directamente desde handlers
- ❌ No uses variables globales

### 2. Handlers

```go
// ✅ BIEN: Handler limpio
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var dto UserDTO
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        helpers.RespondError(w, err, http.StatusBadRequest)
        return
    }
    
    user, err := h.service.CreateUser(&dto)
    if err != nil {
        helpers.RespondError(w, err, http.StatusInternalServerError)
        return
    }
    
    helpers.RespondCreated(w, user, "User created")
}

// ❌ MAL: Lógica de negocio en handler
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    // ... validaciones complejas
    // ... consultas a BD
    // ... envío de emails
    // ... lógica de negocio
}
```

### 3. Services

```go
// ✅ BIEN: Service con lógica de negocio
type UserService struct {
    repo UserRepository
    emailService EmailService
}

func (s *UserService) CreateUser(dto *UserDTO) (*User, error) {
    // Validación de negocio
    if exists := s.repo.ExistsByEmail(dto.Email); exists {
        return nil, ErrEmailAlreadyExists
    }
    
    // Crear usuario
    user := &User{
        Name: dto.Name,
        Email: dto.Email,
    }
    
    if err := s.repo.Create(user); err != nil {
        return nil, err
    }
    
    // Enviar email de bienvenida
    s.emailService.SendWelcome(user.Email)
    
    return user, nil
}
```

### 4. Repositories

```go
// ✅ BIEN: Repository con solo acceso a datos
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

// ❌ MAL: Repository con lógica de negocio
func (r *UserRepository) Create(user *User) error {
    // ❌ Validaciones de negocio
    // ❌ Envío de emails
    // ❌ Llamadas a otros servicios
}
```

### 5. DTOs vs Models

**Models** - Representación interna/BD:
```go
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`  // Oculto en JSON
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

**DTOs** - Transferencia HTTP:
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

### 6. Manejo de Errores

```go
// ✅ BIEN: Errores específicos del dominio
var (
    ErrUserNotFound      = errors.New("user not found")
    ErrEmailAlreadyExists = errors.New("email already exists")
    ErrInvalidPassword    = errors.New("invalid password")
)

// Uso
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
// Usar interfaces para testing
type UserRepository interface {
    Create(user *User) error
    FindByID(id int) (*User, error)
}

// Mock en tests
type MockUserRepository struct {
    CreateFunc func(user *User) error
}

func (m *MockUserRepository) Create(user *User) error {
    return m.CreateFunc(user)
}
```

---

## 🔧 Troubleshooting

### Proyecto no detectado

**Error:**
```
error: no se detectó un proyecto Loom válido
```

**Solución:**
1. Verifica que estés en el directorio del proyecto
2. Verifica que exista `internal/app/` o `internal/modules/`
3. Verifica que exista `go.mod`

### No se puede cambiar router

**Error:**
```
conflicto detectado: gin está instalado
```

**Solución:**
```bash
loom add router chi --force
```

### Backup no restaura

**Error:**
```
backup no encontrado: backup-xxx
```

**Solución:**
1. Lista backups: `ls .loom-backups/`
2. Usa el nombre completo: `loom upgrade --restore backup-20251027-153045`

### Helpers no encontrados

**Error:**
```
cannot find package "github.com/geomark27/loom-go/pkg/helpers"
```

**Solución:**
```bash
go get github.com/geomark27/loom-go/pkg/helpers
go mod tidy
```

### Puerto en uso

**Error:**
```
listen tcp :8080: bind: address already in use
```

**Solución:**
1. Cambiar puerto en `.env`: `PORT=8081`
2. O matar proceso: `lsof -ti:8080 | xargs kill`

---

## 📞 Soporte

¿Problemas? ¿Preguntas?

- **Issues**: [github.com/geomark27/loom-go/issues](https://github.com/geomark27/loom-go/issues)
- **Discussions**: [github.com/geomark27/loom-go/discussions](https://github.com/geomark27/loom-go/discussions)

---

**Última actualización:** 27 de octubre de 2025 - v0.6.0
