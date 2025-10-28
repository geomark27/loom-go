# Changelog

Todos los cambios notables de este proyecto serán documentados en este archivo.

El formato está basado en [Keep a Changelog](https://keepachangelog.com/es-ES/1.0.0/),
y este proyecto adhiere a [Semantic Versioning](https://semver.org/lang/es/).

---

## [1.0.0] - 2025-10-27 🎉

### 🚀 Release Oficial v1.0.0

**Primera versión estable de producción** de Loom. Esta versión marca la madurez del proyecto con todas las funcionalidades core completas y testeadas.

### ✨ Agregado en v1.0.0
- **Comando `loom add`**: Sistema completo de addons para extender proyectos
  - **Routers**: `loom add router [gin|chi|echo]` - Reemplaza Gorilla Mux
  - **ORMs**: `loom add orm [gorm|sqlc]` - Añade ORMs (estructura lista, implementación en progreso)
  - **Databases**: `loom add database [postgres|mysql|mongodb|redis]` - Configura bases de datos
    - PostgreSQL completamente implementado
    - MySQL, MongoDB, Redis con estructura base
  - **Auth**: `loom add auth [jwt|oauth2]` - Sistema de autenticación (estructura lista)
  - **Docker**: `loom add docker` - Añade Dockerfile, docker-compose.yml y .dockerignore
  - `loom add list` - Lista todos los addons disponibles
- **Sistema de Addons**: Arquitectura extensible con interfaces
  - `AddonManager` con registro y detección de conflictos
  - `ProjectDetector` para identificar tecnologías instaladas
  - Helpers para actualizar `go.mod`, `.env.example` e imports
- **Generación de código avanzada**:
  - Templates completos para Gin, Chi y Echo
  - Reescritura inteligente de `server.go` según router
  - Dockerfile multi-stage optimizado
  - docker-compose.yml con detección de base de datos
  - Actualización automática de Makefile con comandos Docker

### 📚 Documentación
- **Consolidación completa de documentación**:
  - `README.md` - Documentación principal para usuarios (instalación, quick start, ejemplos)
  - `DOCS.md` - Documentación técnica completa (arquitecturas, APIs, troubleshooting)
  - `CHANGELOG.md` - Historial de versiones
- Eliminados 13+ archivos .md obsoletos y fragmentados
- Guías detalladas de arquitecturas Layered vs Modular
- Documentación completa de Helpers API
- Buenas prácticas y troubleshooting

### 🔧 Cambiado
- Versión del CLI actualizada a 1.0.0
- Estructura de documentación simplificada (de 30+ archivos a solo 3)
- Mejoras en mensajes de ayuda y next steps

### 🎯 Features Completas en v1.0.0
1. ✅ **Generación de proyectos** (`loom new`) - Layered y Modular
2. ✅ **Generación de componentes** (`loom generate`) - Módulos, handlers, services, etc.
3. ✅ **Sistema de addons** (`loom add`) - Routers, DBs, Docker
4. ✅ **Sistema de actualización** (`loom upgrade`) - Con backups y restore
5. ✅ **Versionado** (`loom version`) - Tracking de versiones
6. ✅ **Helpers package** - Response, Validator, Logger, Errors
7. ✅ **Documentación completa** - User-facing y technical

### 📦 Instalación
```bash
go install github.com/geomark27/loom-go/cmd/loom@latest
```

---

## [0.6.0] - 2025-10-27

### ✨ Agregado
- Sistema base de addons (preparación para v1.0.0)
- Estructura de `internal/addon/`
- Interfaces para extensiones futuras

---

## [0.5.0] - 2025-10-27

### ✨ Agregado
- **Comando `loom upgrade`**: Sistema completo de actualización de proyectos
  - Detección automática de versión del proyecto
  - Backup automático antes de actualizar (con opción `--no-backup`)
  - Aplicación de migraciones incrementales entre versiones
  - Restauración de backups con `--restore`
  - Vista previa de cambios con `--show-changes`
- **Comando `loom version`**: Muestra versión del CLI y del proyecto actual
- **Sistema de versionado**: Infraestructura completa para gestión de versiones
  - `internal/version/`: Detección y comparación de versiones
  - `internal/upgrader/`: Sistema de actualización y backups
  - Archivo `.loom` para tracking de versión del proyecto
- **Changelog integrado**: Muestra cambios entre versiones durante upgrade

### 🔧 Cambiado
- Versión del CLI actualizada a 0.5.0
- Proyectos ahora incluyen archivo `.loom` con metadata

### 📚 Documentación
- Documentación del comando upgrade en help
- Ejemplos de uso de upgrade y restore

---

## [0.4.0] - 2025-10-27

### ✨ Agregado
- **Comando `loom generate`**: Generación de componentes individuales en proyectos existentes
  - `loom generate module <name>`: Genera módulo completo (handler, service, repository, model, DTO)
  - `loom generate handler <name>`: Genera solo handler
  - `loom generate service <name>`: Genera solo service
  - `loom generate model <name>`: Genera solo model
  - `loom generate middleware <name>`: Genera middleware HTTP
- **Detección automática de arquitectura**: El comando genera código apropiado según Layered/Modular
- **Flags globales para generate**:
  - `--force`: Sobrescribe archivos existentes
  - `--dry-run`: Vista previa sin crear archivos
- **Alias de comandos**: `gen`, `g` para generate; aliases para subcomandos
- **Validación de nombres**: Verifica que los nombres de componentes sean válidos
- **Templates completos**: Plantillas para todos los tipos de componentes en ambas arquitecturas

### 🔧 Cambiado
- Versión del CLI actualizada a 0.4.0
- Estructura interna mejorada con `internal/generator/`

### 📚 Documentación
- Documentación completa del comando generate
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
