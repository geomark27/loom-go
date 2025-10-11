# 🧶 Loom - El Tejedor de Proyectos Go

## 📋 ¿Qué es Loom?

**Loom** es una herramienta CLI (Command Line Interface) escrita en Go que automatiza la creación de proyectos backend siguiendo las mejores prácticas y estándares de la comunidad Go.

Su nombre viene de "tejer" (weaving en inglés), porque literalmente "teje" la estructura completa de un proyecto profesional en segundos.

## 🎯 ¿Por qué se creó Loom?

### El Problema

Los desarrolladores que vienen de ecosistemas como:
- **Laravel** (PHP)
- **NestJS** (Node.js/TypeScript)
- **Spring Boot** (Java)
- **Django** (Python)

Se encuentran con una barrera de entrada significativa al trabajar con Go:

1. **No hay estructura estándar clara** - Cada proyecto tiene su propia organización
2. **Demasiado boilerplate inicial** - Configurar servidor, rutas, middlewares, etc.
3. **Decisiones arquitectónicas** - ¿Dónde va cada cosa? ¿Cómo organizo los handlers?
4. **Tiempo perdido** - Horas configurando antes de escribir la primera línea de lógica de negocio

### La Solución

Loom elimina todo ese trabajo inicial generando automáticamente:

✅ **Estructura de directorios profesional**  
✅ **Arquitectura modular y escalable**  
✅ **Código base funcional (servidor HTTP + CRUD)**  
✅ **Configuración de entorno**  
✅ **Documentación y herramientas de desarrollo**

## 🏗️ ¿Qué hace exactamente Loom?

### Comando Principal: `loom new <nombre-proyecto>`

Al ejecutar este comando, Loom genera instantáneamente:

### 1. **Estructura de Directorios Completa**

```
proyecto/
├── cmd/
│   └── proyecto/
│       └── main.go              # Punto de entrada de la aplicación
├── internal/
│   ├── app/
│   │   ├── handlers/            # Controladores HTTP (como Controllers en NestJS)
│   │   │   ├── health_handler.go
│   │   │   └── user_handler.go
│   │   ├── services/            # Lógica de negocio
│   │   │   └── user_service.go
│   │   ├── dtos/                # Data Transfer Objects (validación)
│   │   │   └── user_dto.go
│   │   ├── models/              # Modelos de datos/entidades
│   │   │   └── user.go
│   │   ├── repositories/        # Capa de persistencia
│   │   │   └── user_repository.go
│   │   └── middleware/          # Middlewares HTTP
│   │       └── cors_middleware.go
│   ├── config/                  # Configuración de la aplicación
│   │   └── config.go
│   └── server/                  # Configuración del servidor HTTP
│       ├── server.go
│       └── routes.go
├── pkg/                         # Código reutilizable público
├── docs/                        # Documentación
│   └── API.md
├── scripts/                     # Scripts de automatización
├── .env.example                 # Plantilla de variables de entorno
├── .gitignore                   # Git ignore completo
├── Makefile                     # Comandos de desarrollo
├── go.mod                       # Dependencias
└── README.md                    # Documentación del proyecto
```

### 2. **API REST Funcional**

El proyecto generado incluye un servidor HTTP completamente funcional con:

#### Endpoints de Salud
- `GET /api/v1/health` - Estado del servicio
- `GET /api/v1/health/ready` - Verificación de preparación

#### CRUD de Usuarios (Ejemplo)
- `GET /api/v1/users` - Listar todos los usuarios
- `GET /api/v1/users/{id}` - Obtener un usuario
- `POST /api/v1/users` - Crear un usuario
- `PUT /api/v1/users/{id}` - Actualizar un usuario
- `DELETE /api/v1/users/{id}` - Eliminar un usuario

### 3. **Arquitectura en Capas (Inspirada en NestJS)**

```
┌─────────────┐
│   Handler   │  ← Maneja HTTP requests/responses
└──────┬──────┘
       │
┌──────▼──────┐
│   Service   │  ← Lógica de negocio
└──────┬──────┘
       │
┌──────▼──────┐
│ Repository  │  ← Persistencia de datos
└─────────────┘
```

**Flujo de una request:**
1. Cliente hace `POST /api/v1/users`
2. **Handler** valida el DTO de entrada
3. **Handler** llama al **Service**
4. **Service** ejecuta lógica de negocio
5. **Service** llama al **Repository**
6. **Repository** guarda en la base de datos
7. Respuesta JSON al cliente

### 4. **Inyección de Dependencias Clara**

```go
// En server.go
func New(cfg *config.Config) *Server {
    // 1. Crear repositorios
    userRepo := repositories.NewUserRepository()
    
    // 2. Inyectar repositorios en servicios
    userService := services.NewUserService(userRepo)
    
    // 3. Inyectar servicios en handlers
    userHandler := handlers.NewUserHandler(userService)
    
    // 4. Registrar rutas
    registerRoutes(router, userHandler)
    
    return &Server{...}
}
```

### 5. **Configuración de Entorno**

Archivo `.env.example` con:
```bash
PORT=8080
ENVIRONMENT=development
LOG_LEVEL=info
CORS_ALLOWED_ORIGINS=http://localhost:3000
DATABASE_URL=postgres://...
JWT_SECRET=your-secret-key
```

### 6. **Makefile con Comandos Útiles**

```bash
make run           # Ejecutar la aplicación
make build         # Compilar
make test          # Ejecutar tests
make test-coverage # Tests con cobertura
make fmt           # Formatear código
make vet           # Análisis estático
make clean         # Limpiar archivos generados
```

### 7. **Documentación Automática**

- **README.md** - Guía completa del proyecto
- **docs/API.md** - Documentación detallada de endpoints con ejemplos curl

### 8. **.gitignore Robusto**

Incluye ignorar:
- Variables de entorno (`.env`, `.env.local`, etc.)
- Archivos de configuración sensibles
- Binarios y builds
- Archivos de IDE
- Logs y temporales
- Certificados y claves

## 🎨 Filosofía de Diseño

### "Cerrado para modificación, Abierto para extensión"

- **Base sólida**: La estructura generada sigue las mejores prácticas
- **Totalmente extensible**: Puedes agregar nuevos módulos sin modificar la base
- **No es un framework**: No agrega overhead en runtime, solo genera código

### Respeta los Idiomas de Go

- No intenta "mágicamente" hacer Go como JavaScript o Python
- Usa patrones nativos de Go (interfaces, structs, paquetes)
- Código claro, explícito y sin sorpresas
- Compatible con todas las herramientas estándar de Go

## 🚀 Casos de Uso

### 1. **Prototipos Rápidos**
```bash
loom new mi-prototipo
cd mi-prototipo
go run cmd/mi-prototipo/main.go
# ¡API funcionando en 30 segundos!
```

### 2. **Proyectos Profesionales**
- Estructura escalable para equipos
- Separación clara de responsabilidades
- Fácil de mantener y extender

### 3. **Aprendizaje**
- Ver cómo se estructura un proyecto Go profesional
- Entender patrones de diseño (Repository, Service Layer, etc.)
- Código de ejemplo funcional

### 4. **Microservicios**
- Generar múltiples servicios con estructura consistente
- Fácil onboarding para nuevos desarrolladores

## 🔧 Tecnologías Usadas

### En Loom (la herramienta)
- **Go 1.23** - Lenguaje
- **Cobra** - Framework para CLI
- **text/template** - Sistema de plantillas

### En los proyectos generados
- **Go stdlib** - net/http para el servidor
- **Gorilla Mux** - Router HTTP
- **Arquitectura limpia** - Separation of concerns

## 📊 Comparación con Otros Ecosistemas

| Característica | Laravel (PHP) | NestJS (Node) | Spring Boot (Java) | Loom (Go) |
|---------------|---------------|---------------|-------------------|-----------|
| CLI Generador | ✅ `php artisan` | ✅ `nest generate` | ✅ `spring init` | ✅ `loom new` |
| Estructura Opinada | ✅ | ✅ | ✅ | ✅ |
| Arquitectura Modular | ✅ | ✅ | ✅ | ✅ |
| Runtime Framework | ✅ | ✅ | ✅ | ❌ (solo generador) |
| Performance | ⚡ Media | ⚡ Media | ⚡ Media | ⚡⚡⚡ Alta |
| Curva de Aprendizaje | 📚 Alta | 📚 Media-Alta | 📚 Alta | 📚 Baja |

## 🎯 Ventajas de Loom

✅ **Rápido** - Proyecto completo en segundos  
✅ **Idiomático** - Sigue las convenciones de Go  
✅ **Sin magia** - Código explícito y entendible  
✅ **Sin overhead** - No es un framework en runtime  
✅ **Extensible** - Fácil agregar nuevas funcionalidades  
✅ **Educativo** - Aprende buenas prácticas de Go  
✅ **Profesional** - Estructura lista para producción  

## 🔮 Visión Futura

Loom está diseñado para crecer con comandos adicionales:

```bash
loom generate module auth        # Generar módulo de autenticación
loom generate handler product    # Generar handler de productos
loom add orm gorm               # Agregar GORM
loom add router gin             # Cambiar a Gin router
loom add database postgres      # Configurar PostgreSQL
loom add auth jwt               # Agregar JWT authentication
```

## 📝 Resumen

**Loom es a Go lo que:**
- `php artisan` es a Laravel
- `nest generate` es a NestJS  
- `rails new` es a Ruby on Rails
- `django-admin startproject` es a Django

Pero con la filosofía de Go: **simple, explícito y sin magia**.

---

## 🎓 Conclusión

Loom transforma la experiencia de comenzar un proyecto Go de:

❌ **Sin Loom:**
- 2-4 horas configurando estructura
- Decisiones arquitectónicas constantes
- Código boilerplate repetitivo
- Búsqueda de "la forma correcta"

✅ **Con Loom:**
- ⏱️ 30 segundos para tener proyecto completo
- 🏗️ Arquitectura probada y escalable
- 📝 Código de ejemplo funcional
- 🚀 Enfoque inmediato en lógica de negocio

**Loom no reinventa Go, lo hace más accesible y productivo.** 🧶
