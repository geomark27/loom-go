# 🧶 Loom - El Tejedor de Proyectos Go

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Version](https://img.shields.io/badge/version-1.0.0-green.svg)](https://github.com/geomark27/loom-go/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/geomark27/loom-go)](https://goreportcard.com/report/github.com/geomark27/loom-go)

> **Loom** no es un framework, es un **tejedor de código**. Genera proyectos Go profesionales en segundos y te da las herramientas para extenderlos sin límites.

**Loom** es una herramienta CLI que automatiza la creación y extensión de proyectos backend en Go con arquitectura profesional. Piensa en él como el `create-react-app` o `nest new` del ecosistema Go.

## 🚀 Características Principales

- ⚡ **Crea proyectos completos en 30 segundos**
- 🏗️ **Arquitectura dual**: Layered (simple) o Modular (escalable)
- 🔧 **Genera componentes individuales** en proyectos existentes
- ⬆️ **Actualiza proyectos** sin perder tus cambios
- 🎨 **Añade tecnologías** (routers, ORMs, databases) on-the-fly
- 📦 **Helpers opcionales** o código 100% standalone
- 🚫 **Sin overhead en runtime** - Solo genera código

## 📦 Instalación

```bash
# Instalar globalmente
go install github.com/geomark27/loom-go/cmd/loom@latest

# Verificar instalación
loom --version
```

## 🎯 Inicio Rápido

### Crear un Proyecto

```bash
# Con helpers (recomendado para desarrollo rápido)
loom new mi-api

# O sin helpers (100% independiente)
loom new mi-api --standalone

# Con arquitectura modular
loom new mi-app --modular
```

### Ejecutar

```bash
cd mi-api
go mod tidy
go run cmd/mi-api/main.go
# 🚀 Servidor corriendo en http://localhost:8080
```

### Probar

```bash
# Health check
curl http://localhost:8080/api/v1/health

# CRUD de usuarios (ejemplo incluido)
curl http://localhost:8080/api/v1/users
```

## 🎨 Comandos Disponibles

### `loom new` - Crear proyectos

```bash
loom new mi-api              # Proyecto Layered con helpers
loom new mi-app --modular    # Proyecto Modular por dominios
loom new api --standalone    # Sin helpers (código 100% propio)
```

**Arquitecturas disponibles:**
- **Layered** (por defecto): Simple, ideal para APIs REST
- **Modular**: Escalable, ideal para aplicaciones grandes con múltiples dominios

### `loom generate` - Generar componentes

```bash
# Dentro de un proyecto Loom existente
loom generate module products    # Módulo completo
loom generate handler orders     # Solo handler
loom generate service email      # Solo service
loom generate model Category     # Solo model
loom generate middleware auth    # Middleware HTTP

# Flags útiles
loom generate module users --dry-run  # Vista previa
loom generate handler api --force     # Sobrescribir
```

### `loom add` - Añadir tecnologías

```bash
# Cambiar router HTTP
loom add router gin          # Reemplazar Gorilla Mux por Gin
loom add router chi          # O por Chi
loom add router echo         # O por Echo

# Añadir ORM
loom add orm gorm            # Configurar GORM

# Configurar base de datos
loom add database postgres   # PostgreSQL con docker-compose
loom add database mysql      # MySQL
loom add database mongodb    # MongoDB
loom add database redis      # Redis

# Añadir autenticación
loom add auth jwt            # JWT Authentication
loom add auth oauth2         # OAuth 2.0

# Infrastructure
loom add docker              # Dockerfile + docker-compose.yml

# Ver todos los addons disponibles
loom add list
```

### `loom upgrade` - Actualizar proyectos

```bash
loom version                 # Ver versión actual
loom upgrade --show-changes  # Ver qué cambiaría
loom upgrade                 # Actualizar (con backup automático)
loom upgrade --no-backup     # Actualizar sin backup

# Si algo sale mal
loom upgrade --restore backup-20251027-153045
```

## 📦 Helpers Opcionales

Si no usas `--standalone`, tu proyecto incluye helpers reutilizables:

```go
import "github.com/geomark27/loom-go/pkg/helpers"

// HTTP Responses estandarizadas
helpers.RespondSuccess(w, data, "Success")
helpers.RespondError(w, err, http.StatusBadRequest)
helpers.RespondCreated(w, user, "User created")

// Validación automática
type UserDTO struct {
    Name  string `json:"name" validate:"required,min=3"`
    Email string `json:"email" validate:"required,email"`
}
errors := helpers.ValidateStruct(userDTO)

// Logging estructurado
logger := helpers.NewLogger()
logger.Info("User created", "user_id", user.ID)
logger.Error("Database error", "error", err)
```

Actualizar helpers:
```bash
go get -u github.com/geomark27/loom-go/pkg/helpers
```

## 🏗️ Estructura de Proyectos

### Arquitectura Layered

```
mi-api/
├── cmd/
│   └── mi-api/
│       └── main.go              # Punto de entrada
├── internal/
│   ├── app/
│   │   ├── handlers/            # HTTP handlers
│   │   ├── services/            # Lógica de negocio
│   │   ├── repositories/        # Acceso a datos
│   │   ├── models/              # Modelos de dominio
│   │   ├── dtos/                # Data Transfer Objects
│   │   └── middleware/          # Middlewares HTTP
│   ├── config/                  # Configuración
│   └── server/                  # Servidor HTTP
├── pkg/
│   └── helpers/                 # Utilidades (opcional)
├── docs/
│   └── API.md                   # Documentación
├── .env.example
├── Makefile
└── README.md
```

### Arquitectura Modular

```
mi-app/
├── cmd/
│   └── mi-app/
│       └── main.go
├── internal/
│   ├── modules/                 # Módulos de dominio
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
│   │       └── ... (misma estructura)
│   └── platform/                # Infraestructura compartida
│       ├── server/
│       ├── config/
│       └── eventbus/            # Comunicación entre módulos
├── pkg/
│   └── helpers/
└── ...
```

## 🔌 API Endpoints Incluidos

Todos los proyectos generados incluyen:

### Health Checks
- `GET /api/v1/health` - Estado del servicio
- `GET /api/v1/health/ready` - Verificación de preparación

### CRUD de Usuarios (Ejemplo)
- `GET /api/v1/users` - Listar usuarios
- `GET /api/v1/users/{id}` - Obtener usuario
- `POST /api/v1/users` - Crear usuario
- `PUT /api/v1/users/{id}` - Actualizar usuario
- `DELETE /api/v1/users/{id}` - Eliminar usuario

## 💻 Makefile Incluido

Todos los proyectos tienen estos comandos:

```bash
make help           # Ver todos los comandos
make run            # Ejecutar aplicación
make build          # Compilar
make test           # Tests
make test-coverage  # Tests con cobertura
make fmt            # Formatear código
make vet            # Análisis estático
make clean          # Limpiar archivos

# Si añadiste Docker:
make docker-build   # Construir imagen
make docker-up      # Levantar containers
make docker-down    # Detener containers
make docker-logs    # Ver logs
```

## 🎨 Filosofía

### "Cerrado para modificación, Abierto para extensión"

- **No es un framework** - Solo genera código, sin overhead en runtime
- **Idiomático** - Respeta las convenciones de Go
- **Sin magia** - Código explícito y entendible
- **Extensible** - Fácil agregar nuevas funcionalidades

### Inspiración

Loom lleva la experiencia de frameworks como **NestJS**, **Laravel** y **Spring Boot** al ecosistema Go, manteniendo su simplicidad y rendimiento.

## 📚 Ejemplos Reales

### Crear un e-commerce

```bash
# 1. Crear proyecto base
loom new ecommerce --modular
cd ecommerce

# 2. Generar módulos de dominio
loom generate module products
loom generate module orders
loom generate module payments
loom generate module customers

# 3. Añadir PostgreSQL
loom add database postgres

# 4. Añadir autenticación JWT
loom add auth jwt

# 5. Añadir Docker
loom add docker

# 6. Ejecutar
docker-compose up -d
```

### Migrar de Gorilla Mux a Gin

```bash
cd mi-proyecto-existente
loom add router gin --force  # Reemplaza el router actual
go mod tidy
# Actualizar handlers manualmente para usar gin.Context
```

## 📖 Documentación

- **[DOCS.md](DOCS.md)** - Documentación técnica completa
- **[CHANGELOG.md](CHANGELOG.md)** - Historial de versiones y cambios

## 🤝 Contribuir

¡Las contribuciones son bienvenidas!

1. Fork el proyecto
2. Crea tu branch (`git checkout -b feature/AmazingFeature`)
3. Commit cambios (`git commit -m 'Add AmazingFeature'`)
4. Push (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📝 Licencia

MIT License - ve [LICENSE](LICENSE) para detalles.

## 🙏 Inspiración

- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [NestJS CLI](https://nestjs.com/)
- [Laravel Artisan](https://laravel.com/docs/artisan)
- [Spring Boot CLI](https://spring.io/projects/spring-boot)

## 📞 Contacto

- **Autor**: Marcos
- **GitHub**: [@geomark27](https://github.com/geomark27)
- **Repositorio**: [loom-go](https://github.com/geomark27/loom-go)

---

**¿Te gusta Loom?** Dale una ⭐ en GitHub!

Hecho con ❤️ y ☕ por la comunidad Go
