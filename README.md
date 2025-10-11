# 🧶 Loom - El Tejedor de Proyectos Go

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Version](https://img.shields.io/badge/version-0.1.0-green.svg)](https://github.com/geomark27/loom-go)

**Loom** es una herramienta CLI que automatiza la creación de proyectos backend en Go con arquitectura profesional, siguiendo las mejores prácticas de la comunidad.

## 🎯 ¿Qué hace Loom?

Genera en **30 segundos** un proyecto Go completo con:

- ✅ **Arquitectura modular** (handlers, services, repositories, dtos, models)
- ✅ **API REST funcional** con CRUD de ejemplo
- ✅ **Servidor HTTP** configurado (Gorilla Mux)
- ✅ **Middlewares** (CORS, etc.)
- ✅ **Health checks** implementados
- ✅ **Configuración de entorno** (.env.example)
- ✅ **Makefile** con comandos útiles
- ✅ **Documentación** (README.md + API.md)
- ✅ **.gitignore** completo

## 🚀 Inicio Rápido

### Instalación

```bash
# Clonar el repositorio
git clone https://github.com/geomark27/loom-go.git
cd loom-go

# Instalar globalmente
go install ./cmd/loom

# Verificar instalación
loom --version
```

### Crear tu Primer Proyecto

```bash
# Crear proyecto
loom new mi-api

# Entrar al proyecto
cd mi-api

# Instalar dependencias
go mod tidy

# Ejecutar
go run cmd/mi-api/main.go
```

**¡Servidor corriendo en http://localhost:8080!** 🎉

## 📖 Documentación Completa

- 📋 [**Descripción Detallada**](DESCRIPCION.md) - ¿Qué es Loom y por qué existe?
- 📦 [**Guía de Instalación**](INSTALACION.md) - Instalación y configuración paso a paso
- 🏗️ [**Arquitectura**](#arquitectura) - Estructura de proyectos generados

## 🏗️ Estructura Generada

```
mi-api/
├── cmd/
│   └── mi-api/
│       └── main.go              # Punto de entrada
├── internal/
│   ├── app/
│   │   ├── handlers/            # HTTP handlers (Controllers)
│   │   ├── services/            # Lógica de negocio
│   │   ├── dtos/                # Data Transfer Objects
│   │   ├── models/              # Modelos de datos
│   │   ├── repositories/        # Capa de persistencia
│   │   └── middleware/          # Middlewares HTTP
│   ├── config/                  # Configuración
│   └── server/                  # Servidor HTTP
├── pkg/                         # Código reutilizable
├── docs/
│   └── API.md                   # Documentación de endpoints
├── .env.example                 # Variables de entorno
├── .gitignore
├── Makefile
└── README.md
```

## 🔌 API Endpoints Incluidos

Los proyectos generados incluyen:

### Health Checks
- `GET /api/v1/health` - Estado del servicio
- `GET /api/v1/health/ready` - Verificación de preparación

### CRUD de Usuarios (Ejemplo)
- `GET /api/v1/users` - Listar usuarios
- `GET /api/v1/users/{id}` - Obtener usuario
- `POST /api/v1/users` - Crear usuario
- `PUT /api/v1/users/{id}` - Actualizar usuario
- `DELETE /api/v1/users/{id}` - Eliminar usuario

## 🎨 Filosofía

### "Cerrado para modificación, Abierto para extensión"

- **No es un framework** - Solo genera código, sin overhead en runtime
- **Idiomático** - Respeta las convenciones de Go
- **Sin magia** - Código explícito y entendible
- **Extensible** - Fácil agregar nuevas funcionalidades

### Inspiración

Loom lleva la experiencia de frameworks como **NestJS**, **Laravel** y **Spring Boot** al ecosistema Go, manteniendo su simplicidad y rendimiento.

## 🧪 Ejemplo de Uso

```bash
# Crear proyecto
$ loom new blog-api
✅ Proyecto 'blog-api' creado exitosamente

# Navegar y ejecutar
$ cd blog-api
$ go mod tidy
$ go run cmd/blog-api/main.go
🚀 Servidor blog-api iniciado en http://localhost:8080

# Probar endpoints
$ curl http://localhost:8080/api/v1/health
{
  "status": "healthy",
  "service": "blog-api",
  "version": "v1.0.0"
}
```

## 💻 Comandos Disponibles

En los proyectos generados:

```bash
make help          # Ver todos los comandos
make run           # Ejecutar aplicación
make build         # Compilar
make test          # Ejecutar tests
make test-coverage # Tests con cobertura
make fmt           # Formatear código
make vet           # Análisis estático
make clean         # Limpiar archivos generados
```

## 🔮 Roadmap

Funcionalidades planeadas:

- [ ] `loom generate module <nombre>` - Generar módulos completos
- [ ] `loom generate handler <nombre>` - Generar handler individual
- [ ] `loom add router <gin|chi|echo>` - Cambiar router HTTP
- [ ] `loom add orm <gorm|sqlc>` - Agregar ORM
- [ ] `loom add database <postgres|mysql|mongo>` - Configurar BD
- [ ] `loom add auth <jwt|oauth>` - Agregar autenticación
- [ ] `loom add docker` - Agregar Dockerfile y docker-compose
- [ ] `loom add ci <github|gitlab>` - Configurar CI/CD

## 🤝 Contribuir

Las contribuciones son bienvenidas! Por favor:

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📝 Licencia

Este proyecto está bajo la licencia MIT. Ver [LICENSE](LICENSE) para más detalles.

## 🙏 Agradecimientos

Inspirado por:
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [NestJS CLI](https://nestjs.com/)
- [Laravel Artisan](https://laravel.com/docs/artisan)
- [Cobra CLI](https://github.com/spf13/cobra)

## 📞 Contacto

- **Autor**: Marcos
- **GitHub**: [@geomark27](https://github.com/geomark27)
- **Proyecto**: [loom-go](https://github.com/geomark27/loom-go)

---

**¿Te gusta Loom? Dale una ⭐ en GitHub!**

Hecho con ❤️ y mucho ☕ por desarrolladores Go para desarrolladores Go.
