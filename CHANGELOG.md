# Changelog# Changelog



Todos los cambios notables de este proyecto serán documentados en este archivo.Todos los cambios notables de este proyecto serán documentados en este archivo.



El formato está basado en [Keep a Changelog](https://keepachangelog.com/es-ES/1.0.0/),El formato está basado en [Keep a Changelog](https://keepachangelog.com/es-ES/1.0.0/),

y este proyecto adhiere a [Semantic Versioning](https://semver.org/lang/es/).y este proyecto adhiere a [Semantic Versioning](https://semver.org/lang/es/).



------



## [1.0.0] - 2025-10-27 🎉## [1.0.0] - 2025-10-27 🎉



### 🚀 Release Oficial v1.0.0### 🚀 Release Oficial v1.0.0



**Primera versión estable de producción** de Loom. Esta versión marca la madurez del proyecto con todas las funcionalidades core completas y testeadas.**Primera versión estable de producción** de Loom. Esta versión marca la madurez del proyecto con todas las funcionalidades core completas y testeadas.



### ✨ Características Principales### ✨ Agregado en v1.0.0

- **Comando `loom add`**: Sistema completo de addons para extender proyectos

#### Comando `loom new` - Generador de Proyectos  - **Routers**: `loom add router [gin|chi|echo]` - Reemplaza Gorilla Mux

- Generación de proyectos Go con arquitectura profesional  - **ORMs**: `loom add orm [gorm|sqlc]` - Añade ORMs (estructura lista, implementación en progreso)

- Soporte para arquitectura **Layered** (por capas) y **Modular** (por dominios)  - **Databases**: `loom add database [postgres|mysql|mongodb|redis]` - Configura bases de datos

- Modelo híbrido: Con helpers opcionales o 100% standalone    - PostgreSQL completamente implementado

- Detección automática del usuario de GitHub desde git config    - MySQL, MongoDB, Redis con estructura base

- Flag `--standalone` para proyectos completamente independientes  - **Auth**: `loom add auth [jwt|oauth2]` - Sistema de autenticación (estructura lista)

- Flag `--modular` para arquitectura basada en dominios  - **Docker**: `loom add docker` - Añade Dockerfile, docker-compose.yml y .dockerignore

- Flag `--module` para especificar nombre del módulo manualmente  - `loom add list` - Lista todos los addons disponibles

- **Sistema de Addons**: Arquitectura extensible con interfaces

#### Comando `loom generate` - Generación de Componentes  - `AddonManager` con registro y detección de conflictos

- Generación de módulos completos (handler, service, repository, model, DTO)  - `ProjectDetector` para identificar tecnologías instaladas

- Generación de componentes individuales:  - Helpers para actualizar `go.mod`, `.env.example` e imports

  - `loom generate handler` - HTTP handlers- **Generación de código avanzada**:

  - `loom generate service` - Servicios de lógica de negocio  - Templates completos para Gin, Chi y Echo

  - `loom generate model` - Modelos de datos  - Reescritura inteligente de `server.go` según router

  - `loom generate middleware` - Middlewares HTTP  - Dockerfile multi-stage optimizado

- Detección automática de arquitectura del proyecto (Layered/Modular)  - docker-compose.yml con detección de base de datos

- Flags `--dry-run` para vista previa y `--force` para sobrescribir  - Actualización automática de Makefile con comandos Docker



#### Comando `loom add` - Sistema de Addons### 📚 Documentación

- **Routers**: Reemplazo de routers HTTP- **Consolidación completa de documentación**:

  - `loom add router gin` - Migrar a Gin  - `README.md` - Documentación principal para usuarios (instalación, quick start, ejemplos)

  - `loom add router chi` - Migrar a Chi  - `DOCS.md` - Documentación técnica completa (arquitecturas, APIs, troubleshooting)

  - `loom add router echo` - Migrar a Echo  - `CHANGELOG.md` - Historial de versiones

- **Databases**: Configuración de bases de datos- Eliminados 13+ archivos .md obsoletos y fragmentados

  - `loom add database postgres` - PostgreSQL con docker-compose- Guías detalladas de arquitecturas Layered vs Modular

  - `loom add database mysql` - MySQL- Documentación completa de Helpers API

  - `loom add database mongodb` - MongoDB- Buenas prácticas y troubleshooting

  - `loom add database redis` - Redis

- **ORMs**: Integración de ORMs### 🔧 Cambiado

  - `loom add orm gorm` - GORM (estructura base)- Versión del CLI actualizada a 1.0.0

- **Auth**: Sistemas de autenticación- Estructura de documentación simplificada (de 30+ archivos a solo 3)

  - `loom add auth jwt` - JWT (estructura base)- Mejoras en mensajes de ayuda y next steps

  - `loom add auth oauth2` - OAuth2 (estructura base)

- **Docker**: Dockerización### 🎯 Features Completas en v1.0.0

  - `loom add docker` - Dockerfile multi-stage + docker-compose.yml1. ✅ **Generación de proyectos** (`loom new`) - Layered y Modular

- `loom add list` - Lista todos los addons disponibles2. ✅ **Generación de componentes** (`loom generate`) - Módulos, handlers, services, etc.

3. ✅ **Sistema de addons** (`loom add`) - Routers, DBs, Docker

#### Comando `loom upgrade` - Actualización de Proyectos4. ✅ **Sistema de actualización** (`loom upgrade`) - Con backups y restore

- Sistema de versionado de proyectos (`.loom/version.json`)5. ✅ **Versionado** (`loom version`) - Tracking de versiones

- Comparación inteligente de versiones6. ✅ **Helpers package** - Response, Validator, Logger, Errors

- Backup automático antes de actualizar7. ✅ **Documentación completa** - User-facing y technical

- `loom upgrade --show-changes` - Vista previa de cambios

- `loom upgrade --no-backup` - Sin backup### 📦 Instalación

- `loom upgrade --restore` - Restaurar desde backup```bash

- Actualización selectiva de componentesgo install github.com/geomark27/loom-go/cmd/loom@latest

```

#### pkg/helpers - Librería de Utilidades

- **HTTP Responses**: Respuestas estandarizadas---

  - `RespondSuccess`, `RespondError`, `RespondCreated`, etc.

- **Validator**: Validación de structs con tags## [0.6.0] - 2025-10-27

  - `ValidateStruct`, `ValidateEmail`, `ValidateMin/Max`

- **Logger**: Logging estructurado### ✨ Agregado

  - Interfaz Logger con implementación DefaultLogger- Sistema base de addons (preparación para v1.0.0)

- **Errors**: Manejo de errores con contexto- Estructura de `internal/addon/`

  - `AppError` type, `Wrap/Unwrap`, errores predefinidos- Interfaces para extensiones futuras

- **Context**: Helpers para context de Go

  - `GetUserID`, `GetRequestID`, `GetTenantID`---



### 🏗️ Arquitectura## [0.5.0] - 2025-10-27



#### Templates Embebidos### ✨ Agregado

- 17 plantillas organizadas en `internal/generator/templates/`- **Comando `loom upgrade`**: Sistema completo de actualización de proyectos

- Soporte para generación condicional (con/sin helpers)  - Detección automática de versión del proyecto

- Templates para ambas arquitecturas (Layered y Modular)  - Backup automático antes de actualizar (con opción `--no-backup`)

  - Aplicación de migraciones incrementales entre versiones

#### Proyectos Generados Incluyen  - Restauración de backups con `--restore`

- ✅ Arquitectura modular (handlers, services, repositories, dtos, models)  - Vista previa de cambios con `--show-changes`

- ✅ API REST funcional con CRUD de ejemplo- **Comando `loom version`**: Muestra versión del CLI y del proyecto actual

- ✅ Servidor HTTP configurado (Gorilla Mux por defecto)- **Sistema de versionado**: Infraestructura completa para gestión de versiones

- ✅ Middlewares (CORS, etc.)  - `internal/version/`: Detección y comparación de versiones

- ✅ Health checks implementados  - `internal/upgrader/`: Sistema de actualización y backups

- ✅ Configuración de entorno (.env.example)  - Archivo `.loom` para tracking de versión del proyecto

- ✅ Makefile con comandos útiles- **Changelog integrado**: Muestra cambios entre versiones durante upgrade

- ✅ Documentación (README.md + API.md)

- ✅ .gitignore completo### 🔧 Cambiado

- Versión del CLI actualizada a 0.5.0

### 📚 Documentación- Proyectos ahora incluyen archivo `.loom` con metadata

- `README.md` - Documentación principal para usuarios

- `DOCS.md` - Documentación técnica completa### 📚 Documentación

- `CHANGELOG.md` - Historial de versiones- Documentación del comando upgrade en help

- Guías detalladas de arquitecturas Layered vs Modular- Ejemplos de uso de upgrade y restore

- Documentación completa de Helpers API

- Buenas prácticas y troubleshooting---



### 🔧 Tecnologías## [0.4.0] - 2025-10-27

- Go 1.23+

- Cobra CLI framework### ✨ Agregado

- Gorilla Mux (router HTTP por defecto)- **Comando `loom generate`**: Generación de componentes individuales en proyectos existentes

- text/template para generación de código  - `loom generate module <name>`: Genera módulo completo (handler, service, repository, model, DTO)

- embed package para templates embebidos  - `loom generate handler <name>`: Genera solo handler

  - `loom generate service <name>`: Genera solo service

### 🎯 Filosofía  - `loom generate model <name>`: Genera solo model

- **No es un framework** - Solo genera código, sin overhead en runtime  - `loom generate middleware <name>`: Genera middleware HTTP

- **Idiomático** - Respeta las convenciones de Go- **Detección automática de arquitectura**: El comando genera código apropiado según Layered/Modular

- **Sin magia** - Código explícito y entendible- **Flags globales para generate**:

- **Extensible** - Fácil agregar nuevas funcionalidades  - `--force`: Sobrescribe archivos existentes

- **Zero vendor lock-in** - Código generado es 100% tuyo  - `--dry-run`: Vista previa sin crear archivos

- **Alias de comandos**: `gen`, `g` para generate; aliases para subcomandos

---- **Validación de nombres**: Verifica que los nombres de componentes sean válidos

- **Templates completos**: Plantillas para todos los tipos de componentes en ambas arquitecturas

## Releases Anteriores

### 🔧 Cambiado

Este es el primer release estable de producción. Las versiones de desarrollo (v0.x.x) fueron iteraciones internas que culminaron en esta versión v1.0.0.- Versión del CLI actualizada a 0.4.0

- Estructura interna mejorada con `internal/generator/`

---

### 📚 Documentación

[1.0.0]: https://github.com/geomark27/loom-go/releases/tag/v1.0.0- Documentación completa del comando generate

- Ejemplos de uso para cada subcomando
- Guía de próximos pasos después de generar componentes

---

## [0.3.0] - 2025-10-27

### ✨ Agregado
- **Arquitectura Dual**: Soporte para arquitectura Layered (por defecto) y Modular
- Flag `--modular` para generar proyectos con arquitectura modular por dominios
- Detección automática de arquitectura en proyectos existentes
- Módulo `users` de ejemplo en arquitectura modular
- Event Bus para comunicación entre módulos
- Archivo `ports.go` con interfaces en módulos
- Plantillas separadas para cada arquitectura (`templates/layered/` y `templates/modular/`)

### 🔧 Cambiado
- Reorganización de estructura interna de plantillas
- Separación clara entre `platform` (infraestructura) y `app`/`modules` (lógica)
- Mejora en detección de usuario GitHub desde config
- Mensajes informativos sobre arquitectura seleccionada

### 📚 Documentación
- Guías sobre cuándo usar cada arquitectura
- Ejemplos de ambas arquitecturas
- Documentación de Event Bus y comunicación entre módulos

---

## [0.2.0] - 2025-10-XX

### ✨ Agregado
- **Helpers Package** (`pkg/helpers/`): Biblioteca de utilidades reutilizables
  - `response.go`: Respuestas HTTP estandarizadas
  - `validator.go`: Validación de structs con tags
  - `logger.go`: Logging estructurado
  - `errors.go`: Manejo de errores con contexto
  - `context.go`: Utilidades de contexto
- Flag `--standalone` para generar proyectos sin helpers
- Campo `UseHelpers` en ProjectConfig
- Soporte para proyectos 100% independientes

### 🔧 Cambiado
- Los proyectos por defecto ahora incluyen helpers
- Estructura más clara con `internal/platform` para infraestructura
- Separación entre `internal/app` (negocio) y `internal/shared` (utilidades)

### 📚 Documentación
- Documentación de helpers disponibles
- Guía de uso de helpers
- Ejemplos de proyectos con y sin helpers
- Sección "Modelo Híbrido" en README

---

## [0.1.0] - 2025-10-XX

### ✨ Agregado - Lanzamiento Inicial
- Comando `loom new <nombre>` para crear proyectos
- Generación de estructura de proyecto completa:
  - `cmd/` - Entry point
  - `internal/app/` - Lógica de negocio (handlers, services, repositories)
  - `internal/config/` - Configuración
  - `internal/server/` - Servidor HTTP
  - `pkg/` - Código público
  - `docs/` - Documentación
- Arquitectura en capas (Layered) como estándar
- API REST funcional con CRUD de usuarios de ejemplo
- Health checks implementados
- Middleware CORS configurado
- Servidor HTTP con Gorilla Mux
- Plantillas embebidas con `text/template`
- Generación de archivos:
  - `go.mod`
  - `README.md`
  - `.gitignore`
  - `.env.example`
  - `Makefile`
  - `docs/API.md`
- Flag `-m, --module` para especificar nombre del módulo Go
- Auto-detección de usuario GitHub desde git config
- Validación de nombres de proyecto
- Mensajes informativos post-generación

### 📚 Documentación
- README principal con guía de instalación
- DESCRIPCION.md con explicación detallada
- INSTALACION.md con guía paso a paso
- FLUJOS_REALES.md con casos de uso
- INSTALL_FROM_GITHUB.md con instalación desde GitHub
- COMPATIBILIDAD_MULTIPLATAFORMA.md
- DISTRIBUCION_GLOBAL.md
- VERIFICACION.md

### 🛠️ Infraestructura
- Configuración de Go modules
- Dependencia en Cobra para CLI
- Estructura de paquetes interna
- Sistema de plantillas

---

## Tipos de Cambios

- **✨ Agregado** (`Added`): Para nuevas funcionalidades
- **🔧 Cambiado** (`Changed`): Para cambios en funcionalidades existentes
- **❌ Obsoleto** (`Deprecated`): Para funcionalidades que serán removidas
- **🗑️ Removido** (`Removed`): Para funcionalidades removidas
- **🐛 Corregido** (`Fixed`): Para corrección de bugs
- **🔒 Seguridad** (`Security`): Para vulnerabilidades de seguridad

---

## Links de Versiones

- [Unreleased]: https://github.com/geomark27/loom-go/compare/v0.5.0...HEAD
- [0.5.0]: https://github.com/geomark27/loom-go/compare/v0.4.0...v0.5.0
- [0.4.0]: https://github.com/geomark27/loom-go/compare/v0.3.0...v0.4.0
- [0.3.0]: https://github.com/geomark27/loom-go/compare/v0.2.0...v0.3.0
- [0.2.0]: https://github.com/geomark27/loom-go/compare/v0.1.0...v0.2.0
- [0.1.0]: https://github.com/geomark27/loom-go/releases/tag/v0.1.0
