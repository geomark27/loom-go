# 🧶 Loom - El Tejedor de Proyectos Go# 🧶 Loom - El Tejedor de Proyectos Go



[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://golang.org)[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://golang.org)

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

[![Version](https://img.shields.io/badge/version-0.6.0-green.svg)](https://github.com/geomark27/loom-go/releases)[![Version](https://img.shields.io/badge/version-0.3.0-green.svg)](https://github.com/geomark27/loom-go/releases)

[![Go Report Card](https://goreportcard.com/badge/github.com/geomark27/loom-go)](https://goreportcard.com/report/github.com/geomark27/loom-go)[![Go Report Card](https://goreportcard.com/badge/github.com/geomark27/loom-go)](https://goreportcard.com/report/github.com/geomark27/loom-go)

[![Roadmap](https://img.shields.io/badge/roadmap-v1.0.0-blue.svg)](ROADMAP_TO_V1.md)

> **Loom** no es un framework, es un **tejedor de código**. Genera proyectos Go profesionales en segundos y te da las herramientas para extenderlos sin límites.

**Loom** es una herramienta CLI que automatiza la creación de proyectos backend en Go con arquitectura profesional, siguiendo las mejores prácticas de la comunidad.

## 🎯 ¿Qué es Loom?

## 📦 Instalación con un solo comando

**Loom** es una herramienta CLI que automatiza la creación y extensión de proyectos backend en Go con arquitectura profesional. Piensa en él como el `create-react-app` o `nest new` del ecosistema Go.

```bash

### ¿Por qué Loom?go install github.com/geomark27/loom-go/cmd/loom@latest

```

- ⚡ **Crea proyectos completos en 30 segundos**

- 🏗️ **Arquitectura dual**: Layered (simple) o Modular (escalable)## 🎯 ¿Qué hace Loom?

- 🔧 **Genera componentes individuales** en proyectos existentes

- ⬆️ **Actualiza proyectos** sin perder tus cambiosGenera en **30 segundos** un proyecto Go completo con:

- 🎨 **Añade tecnologías** (routers, ORMs, databases) on-the-fly

- 📦 **Helpers opcionales** o código 100% standalone- ✅ **Arquitectura modular** (handlers, services, repositories, dtos, models)

- 🚫 **Sin overhead en runtime** - Solo genera código- ✅ **API REST funcional** con CRUD de ejemplo

- ✅ **Servidor HTTP** configurado (Gorilla Mux)

## 📦 Instalación- ✅ **Middlewares** (CORS, etc.)

- ✅ **Health checks** implementados

```bash- ✅ **Helpers opcionales** 🆕 para desarrollo rápido o código 100% independiente

go install github.com/geomark27/loom-go/cmd/loom@latest- ✅ **Configuración de entorno** (.env.example)

loom --version  # Verificar instalación- ✅ **Makefile** con comandos útiles

```- ✅ **Documentación** (README.md + API.md)

- ✅ **.gitignore** completo

## 🚀 Inicio Rápido

## 🆕 Novedades v0.2.0

### 1. Crear un proyecto nuevo

### Modelo Híbrido: Generador + Helpers Opcionales

```bash

# Proyecto con arquitectura Layered (recomendado para APIs simples)Ahora puedes elegir:

loom new mi-api

**Opción 1: Con Helpers (Por defecto)** 🚀

# O proyecto Modular (recomendado para aplicaciones grandes)```bash

loom new mi-app --modularloom new mi-api

```

# Proyecto sin helpers (100% standalone)- Incluye `pkg/helpers` para desarrollo rápido

loom new mi-api --standalone- Respuestas HTTP estandarizadas

```- Validación automática

- Logging estructurado

### 2. Ejecutar- Actualizable con `go get -u`



```bash**Opción 2: Standalone (100% Independiente)** 🔧

cd mi-api```bash

go mod tidyloom new mi-api --standalone

go run cmd/mi-api/main.go```

# 🚀 Servidor corriendo en http://localhost:8080- Código 100% propio, sin dependencias de Loom

```- Control total

- Cero vendor lock-in

### 3. Probar- Perfecto para puristas de Go



```bash## 🚀 Inicio Rápido

# Health check

curl http://localhost:8080/api/v1/health### Instalación



# API de ejemplo (CRUD de usuarios)#### Opción 1: Instalar desde GitHub (Recomendado)

curl http://localhost:8080/api/v1/users

``````bash

# Instalar con un solo comando

## 🎨 Comandos Principalesgo install github.com/geomark27/loom-go/cmd/loom@latest



### `loom new` - Crear proyectos# Verificar instalación

loom --version

```bash```

loom new mi-api              # Proyecto Layered con helpers

loom new mi-app --modular    # Proyecto Modular por dominios#### Opción 2: Clonar y compilar (Para desarrollo)

loom new api --standalone    # Sin helpers (código 100% propio)

``````bash

# Clonar el repositorio

**Arquitecturas disponibles:**git clone https://github.com/geomark27/loom-go.git

- **Layered** (por defecto): Simple, ideal para APIs REST pequeñascd loom-go

- **Modular**: Escalable, ideal para aplicaciones grandes con múltiples dominios

# Instalar globalmente

### `loom generate` - Generar componentes (v0.4.0)go install ./cmd/loom



```bash# Verificar instalación

# Dentro de un proyecto existenteloom --version

loom generate module products    # Módulo completo (handler, service, repo, model, DTO)```

loom generate handler orders     # Solo handler

loom generate service email      # Solo service> 📖 **Más opciones de instalación:** Ver [INSTALL_FROM_GITHUB.md](INSTALL_FROM_GITHUB.md)

loom generate model Category     # Solo model

loom generate middleware auth    # Middleware HTTP### Crear tu Primer Proyecto



# Flags útiles```bash

loom generate module users --dry-run  # Vista previa sin crear archivos# Crear proyecto (con helpers por defecto)

loom generate handler api --force     # Sobrescribir si existeloom new mi-api

```

# O crear proyecto standalone (sin helpers)

**Detecta automáticamente** tu arquitectura (Layered/Modular) y genera código apropiado.loom new mi-api --standalone



### `loom add` - Añadir tecnologías (v0.6.0) 🆕# Entrar al proyecto

cd mi-api

```bash

# Cambiar router HTTP# Instalar dependencias

loom add router gin          # Reemplazar Gorilla Mux por Gingo mod tidy

loom add router chi          # Reemplazar por Chi

loom add router echo         # Reemplazar por Echo# Ejecutar

go run cmd/mi-api/main.go

# Añadir ORM```

loom add orm gorm            # Configurar GORM

**¡Servidor corriendo en http://localhost:8080!** 🎉

# Configurar base de datos

loom add database postgres   # PostgreSQL con docker-compose### 📦 Helpers Disponibles (v0.2.0)

loom add database mysql      # MySQL

loom add database mongodb    # MongoDBSi usas helpers, tu proyecto incluye:

loom add database redis      # Redis

```go

# Añadir autenticaciónimport "github.com/geomark27/loom-go/pkg/helpers"

loom add auth jwt            # JWT Authentication

loom add auth oauth2         # OAuth 2.0// HTTP Responses estandarizadas

helpers.RespondSuccess(w, data, "Success")

# Infrastructurehelpers.RespondError(w, err, http.StatusBadRequest)

loom add docker              # Dockerfile + docker-compose.ymlhelpers.RespondCreated(w, user, "User created")



# Ver todos los addons disponibles// Validación automática

loom add listerrors := helpers.ValidateStruct(myDTO)

```if len(errors) > 0 {

    // Manejar errores de validación

### `loom upgrade` - Actualizar proyectos (v0.5.0)}



```bash// Logging estructurado

loom version                 # Ver versión actual del proyectologger := helpers.NewLogger()

loom upgrade --show-changes  # Ver qué cambiaríalogger.Info("User created", "user_id", user.ID)

loom upgrade                 # Actualizar (con backup automático)logger.Error("Database error", "error", err)

loom upgrade --no-backup     # Actualizar sin backup

// Errores predefinidos

# Si algo sale malhelpers.ErrNotFound

loom upgrade --restore backup-20251027-153045helpers.ErrBadRequest

```helpers.ErrUnauthorized

```

### Otros comandos

Para actualizar helpers:

```bash```bash

loom version     # Ver versión del CLI y del proyectogo get -u github.com/geomark27/loom-go/pkg/helpers

loom --help      # Ver todos los comandos```

```

## 📖 Documentación Completa

## 🏗️ Estructura de Proyectos

- 📋 [**Descripción Detallada**](DESCRIPCION.md) - ¿Qué es Loom y por qué existe?

### Arquitectura Layered (por capas)- 📦 [**Guía de Instalación**](INSTALACION.md) - Instalación y configuración paso a paso

- 🏗️ [**Arquitectura**](#arquitectura) - Estructura de proyectos generados

```

mi-api/## 🏗️ Estructura Generada

├── cmd/

│   └── mi-api/```

│       └── main.go              # Entry pointmi-api/

├── internal/├── cmd/

│   ├── app/│   └── mi-api/

│   │   ├── handlers/            # HTTP handlers (controladores)│       └── main.go              # Punto de entrada

│   │   ├── services/            # Lógica de negocio├── internal/

│   │   ├── repositories/        # Acceso a datos│   ├── app/

│   │   ├── models/              # Modelos de dominio│   │   ├── handlers/            # HTTP handlers (Controllers)

│   │   ├── dtos/                # Data Transfer Objects│   │   ├── services/            # Lógica de negocio

│   │   └── middleware/          # Middlewares HTTP│   │   ├── dtos/                # Data Transfer Objects

│   ├── config/                  # Configuración│   │   ├── models/              # Modelos de datos

│   └── server/                  # Servidor HTTP│   │   ├── repositories/        # Capa de persistencia

├── pkg/│   │   └── middleware/          # Middlewares HTTP

│   └── helpers/                 # Utilidades (opcional)│   ├── config/                  # Configuración

├── docs/│   └── server/                  # Servidor HTTP

│   └── API.md                   # Documentación├── pkg/                         # Código reutilizable

├── .env.example├── docs/

├── Makefile│   └── API.md                   # Documentación de endpoints

└── README.md├── .env.example                 # Variables de entorno

```├── .gitignore

├── Makefile

### Arquitectura Modular (por dominios)└── README.md

```

```

mi-app/## 🔌 API Endpoints Incluidos

├── cmd/

│   └── mi-app/Los proyectos generados incluyen:

│       └── main.go

├── internal/### Health Checks

│   ├── modules/                 # Módulos de dominio- `GET /api/v1/health` - Estado del servicio

│   │   ├── users/- `GET /api/v1/health/ready` - Verificación de preparación

│   │   │   ├── handler.go

│   │   │   ├── service.go### CRUD de Usuarios (Ejemplo)

│   │   │   ├── repository.go- `GET /api/v1/users` - Listar usuarios

│   │   │   ├── model.go- `GET /api/v1/users/{id}` - Obtener usuario

│   │   │   ├── dto.go- `POST /api/v1/users` - Crear usuario

│   │   │   ├── router.go- `PUT /api/v1/users/{id}` - Actualizar usuario

│   │   │   ├── validator.go- `DELETE /api/v1/users/{id}` - Eliminar usuario

│   │   │   └── errors.go

│   │   └── products/## 🎨 Filosofía

│   │       └── ... (misma estructura)

│   └── platform/                # Infraestructura compartida### "Cerrado para modificación, Abierto para extensión"

│       ├── server/

│       ├── config/- **No es un framework** - Solo genera código, sin overhead en runtime

│       └── eventbus/            # Comunicación entre módulos- **Idiomático** - Respeta las convenciones de Go

├── pkg/- **Sin magia** - Código explícito y entendible

│   └── helpers/- **Extensible** - Fácil agregar nuevas funcionalidades

└── ...

```### Inspiración



## 📚 Ejemplos de UsoLoom lleva la experiencia de frameworks como **NestJS**, **Laravel** y **Spring Boot** al ecosistema Go, manteniendo su simplicidad y rendimiento.



### Crear y extender un e-commerce## 🧪 Ejemplo de Uso



```bash```bash

# 1. Crear proyecto base# Crear proyecto

loom new ecommerce --modular$ loom new blog-api

cd ecommerce✅ Proyecto 'blog-api' creado exitosamente



# 2. Generar módulos de dominio# Navegar y ejecutar

loom generate module products$ cd blog-api

loom generate module orders$ go mod tidy

loom generate module payments$ go run cmd/blog-api/main.go

loom generate module customers🚀 Servidor blog-api iniciado en http://localhost:8080



# 3. Añadir PostgreSQL# Probar endpoints

loom add database postgres$ curl http://localhost:8080/api/v1/health

{

# 4. Añadir autenticación JWT  "status": "healthy",

loom add auth jwt  "service": "blog-api",

  "version": "v1.0.0"

# 5. Añadir Docker}

loom add docker```



# 6. Ejecutar## 💻 Comandos Disponibles

docker-compose up -d

```En los proyectos generados:



### Migrar de Gorilla Mux a Gin```bash

make help          # Ver todos los comandos

```bashmake run           # Ejecutar aplicación

cd mi-proyecto-existentemake build         # Compilar

loom add router gin --force  # Reemplaza el router actualmake test          # Ejecutar tests

go mod tidymake test-coverage # Tests con cobertura

# Actualizar handlers manualmente para usar gin.Contextmake fmt           # Formatear código

```make vet           # Análisis estático

make clean         # Limpiar archivos generados

### Mantener proyecto actualizado```



```bash## 🔮 Roadmap

# Ver qué hay nuevo

loom version### ✅ v0.1.0 - Generador Base (Completado)

loom upgrade --show-changes- CLI funcional

- Generación de proyectos con arquitectura Layered

# Actualizar con backup- Templates embebidos

loom upgrade- Instalación global



# Si hay problemas, restaurar### ✅ v0.2.0 - Helpers Opcionales (Completado)

loom upgrade --restore backup-20251027-153045- `pkg/helpers` con utilidades reutilizables

```- Flag `--standalone` para proyectos independientes

- Response, Validator, Logger, Errors, Context helpers

## 🎯 API Endpoints Generados- Helpers actualizables con `go get -u`



Todos los proyectos incluyen:### ✅ v0.3.0 - Arquitectura Dual (Actual)

- Soporte para arquitectura **Layered** (por capas) y **Modular** (por dominios)

### Health Checks- Flag `--modular` para proyectos modulares

```- Event Bus para comunicación entre módulos

GET /api/v1/health        - Estado del servicio- Detección automática de arquitectura

GET /api/v1/health/ready  - Verificación de preparación- Separación clara platform/shared/app o modules

```

### 🚀 v0.4.0 - Comando Generate (Próximo - Nov 2025)

### CRUD de Ejemplo (Usuarios)```bash

```loom generate module products    # Módulo completo

GET    /api/v1/users      - Listar usuariosloom generate handler orders     # Solo handler

GET    /api/v1/users/{id} - Obtener usuarioloom generate service email      # Solo service

POST   /api/v1/users      - Crear usuarioloom generate model Category     # Solo modelo

PUT    /api/v1/users/{id} - Actualizar usuarioloom generate middleware auth    # Solo middleware

DELETE /api/v1/users/{id} - Eliminar usuario```

```- Detección automática de proyectos Loom existentes

- Adaptación a arquitectura Layered o Modular

## 📦 Helpers Incluidos (v0.2.0)- Flags `--force` y `--dry-run`



Si no usas `--standalone`, tu proyecto incluye helpers reutilizables:### � v0.5.0 - Comando Upgrade (Dic 2025)

```bash

```goloom upgrade --check       # Ver actualizaciones disponibles

import "github.com/geomark27/loom-go/pkg/helpers"loom upgrade --apply       # Aplicar actualizaciones

loom upgrade --helpers     # Actualizar solo helpers

// Respuestas HTTP estandarizadasloom migrate handler user  # Migrar componente específico

helpers.RespondSuccess(w, data, "Success")```

helpers.RespondError(w, err, http.StatusBadRequest)- Sistema de versionado de proyectos (`.loom/version.json`)

helpers.RespondCreated(w, user, "User created")- Comparación inteligente de plantillas

- Backup automático antes de actualizar

// Validación con tags- Actualización selectiva de componentes

type UserDTO struct {

    Name  string `json:"name" validate:"required,min=3"`### 🎉 v1.0.0 - Versión Estable (Dic 2025)

    Email string `json:"email" validate:"required,email"`- Release estable con todas las features core

}- Tests coverage > 80%

errors := helpers.ValidateStruct(userDTO)- Documentación completa EN + ES

- CI/CD automatizado

// Logging estructurado- Multi-platform builds

logger := helpers.NewLogger()

logger.Info("User created", "user_id", user.ID)### 🔮 Futuro (v1.1+)

logger.Error("Database error", "error", err)- `loom add database postgres` - Configurar BD

```- `loom add orm gorm` - Agregar ORM

- `loom add auth jwt` - Agregar autenticación

Actualizar helpers:- `loom add docker` - Dockerfiles y docker-compose

```bash- `loom add ci github-actions` - Configurar CI/CD

go get -u github.com/geomark27/loom-go

```## 🤝 Contribuir



## 🛠️ Makefile IncluidoLas contribuciones son bienvenidas! Por favor:



Todos los proyectos tienen estos comandos:1. Fork el proyecto

2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)

```bash3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)

make run            # Ejecutar aplicación4. Push a la rama (`git push origin feature/AmazingFeature`)

make build          # Compilar5. Abre un Pull Request

make test           # Tests

make test-coverage  # Tests con cobertura## 📝 Licencia

make fmt            # Formatear código

make vet            # Análisis estáticoEste proyecto está bajo la licencia MIT. Ver [LICENSE](LICENSE) para más detalles.

make clean          # Limpiar archivos

make help           # Ver todos los comandos## 🙏 Agradecimientos



# Si añadiste Docker:Inspirado por:

make docker-build   # Construir imagen- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

make docker-up      # Levantar containers- [NestJS CLI](https://nestjs.com/)

make docker-down    # Detener containers- [Laravel Artisan](https://laravel.com/docs/artisan)

make docker-logs    # Ver logs- [Cobra CLI](https://github.com/spf13/cobra)

```

## 📞 Contacto

## 📖 Documentación

- **Autor**: Marcos

- **[DOCS.md](DOCS.md)** - Documentación técnica completa- **GitHub**: [@geomark27](https://github.com/geomark27)

- **[CHANGELOG.md](CHANGELOG.md)** - Historial de versiones y cambios- **Proyecto**: [loom-go](https://github.com/geomark27/loom-go)



## 🗺️ Roadmap---



- ✅ **v0.1.0** - Generador base con arquitectura Layered**¿Te gusta Loom? Dale una ⭐ en GitHub!**

- ✅ **v0.2.0** - Helpers opcionales (response, validator, logger)

- ✅ **v0.3.0** - Arquitectura dual (Layered + Modular)Hecho con ❤️ y mucho ☕ por desarrolladores Go para desarrolladores Go.

- ✅ **v0.4.0** - Comando `generate` para componentes individuales
- ✅ **v0.5.0** - Comando `upgrade` con sistema de versionado
- ✅ **v0.6.0** - Comando `add` para tecnologías (routers, ORMs, etc.)
- 🚧 **v1.0.0** - Release estable (Q1 2026)

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
