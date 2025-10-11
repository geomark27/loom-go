# 📦 Instalación de Loom desde GitHub

## 🚀 Instalación Rápida (Recomendado)

Ahora puedes instalar Loom directamente desde GitHub con un solo comando:

```bash
go install github.com/geomark27/loom-go/cmd/loom@latest
```

Esto descargará, compilará e instalará Loom automáticamente en `$GOPATH/bin/loom`.

## 📌 Instalación de Versión Específica

```bash
# Instalar versión específica
go install github.com/geomark27/loom-go/cmd/loom@v0.1.0

# Instalar la última versión (recomendado)
go install github.com/geomark27/loom-go/cmd/loom@latest
```

## ✅ Verificar la Instalación

```bash
# Verificar que Loom está instalado
loom --version

# Ver ayuda
loom --help

# Crear tu primer proyecto
loom new mi-proyecto
```

## 🔧 Requisitos Previos

- **Go 1.21+** instalado
- `$GOPATH/bin` en tu PATH

### Verificar Requisitos

```bash
# Verificar versión de Go
go version

# Verificar GOPATH
go env GOPATH

# En Linux/Mac, agregar GOPATH/bin al PATH (si no está)
export PATH="$PATH:$(go env GOPATH)/bin"

# En Windows PowerShell
$env:PATH += ";$(go env GOPATH)\bin"
```

## 🌍 Instalación en Diferentes Sistemas Operativos

### Linux / macOS

```bash
# 1. Instalar Loom
go install github.com/geomark27/loom-go/cmd/loom@latest

# 2. Verificar instalación
loom --version

# 3. Si no funciona, agregar al PATH (agregar a ~/.bashrc o ~/.zshrc)
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.bashrc
source ~/.bashrc

# 4. Crear proyecto
loom new mi-api
```

### Windows (PowerShell)

```powershell
# 1. Instalar Loom
go install github.com/geomark27/loom-go/cmd/loom@latest

# 2. Verificar instalación
loom --version

# 3. Si no funciona, agregar al PATH permanentemente
# Panel de Control → Sistema → Variables de entorno
# Agregar: C:\Users\TuUsuario\go\bin

# 4. Crear proyecto
loom new mi-api
```

### Windows (Git Bash)

```bash
# 1. Instalar Loom
go install github.com/geomark27/loom-go/cmd/loom@latest

# 2. Verificar instalación
loom --version

# 3. Crear proyecto
loom new mi-api
```

## 🆚 Comparación de Métodos de Instalación

### Método 1: `go install` desde GitHub (Recomendado para usuarios)

```bash
go install github.com/geomark27/loom-go/cmd/loom@latest
```

**Ventajas:**
- ✅ Un solo comando
- ✅ Instala la versión publicada
- ✅ Fácil de actualizar
- ✅ No requiere clonar el repositorio

### Método 2: Clonar y compilar (Recomendado para desarrollo)

```bash
git clone https://github.com/geomark27/loom-go.git
cd loom-go
go install ./cmd/loom
```

**Ventajas:**
- ✅ Acceso al código fuente
- ✅ Puedes modificar y contribuir
- ✅ Útil para desarrollo

## 🔄 Actualización

Para actualizar a la última versión:

```bash
# Actualizar a la última versión
go install github.com/geomark27/loom-go/cmd/loom@latest

# Verificar nueva versión
loom --version
```

## 🗑️ Desinstalación

Para desinstalar Loom:

```bash
# Linux/Mac
rm $(which loom)

# Windows PowerShell
Remove-Item (Get-Command loom).Source

# O manualmente
# Linux/Mac: rm $GOPATH/bin/loom
# Windows: del %GOPATH%\bin\loom.exe
```

## 📝 Ejemplo Completo de Uso

```bash
# 1. Instalar
go install github.com/geomark27/loom-go/cmd/loom@latest

# 2. Verificar
loom --version
# Salida: loom version 0.1.0

# 3. Crear proyecto
loom new mi-super-api

# 4. Entrar al proyecto
cd mi-super-api

# 5. Instalar dependencias
go mod tidy

# 6. Ejecutar
go run cmd/mi-super-api/main.go
# Salida: 🚀 Servidor mi-super-api iniciado en http://localhost:8080

# 7. Probar (en otra terminal)
curl http://localhost:8080/api/v1/health
# Salida: {"status":"healthy"...}
```

## 🐛 Solución de Problemas

### Error: "loom: command not found"

```bash
# Verificar que Go esté instalado
go version

# Verificar GOPATH
go env GOPATH

# Agregar GOPATH/bin al PATH
# Linux/Mac (agregar a ~/.bashrc o ~/.zshrc):
export PATH="$PATH:$(go env GOPATH)/bin"

# Windows (Variables de entorno):
# Agregar: C:\Users\TuUsuario\go\bin al PATH
```

### Error: "cannot find package"

```bash
# Verificar conexión a internet
ping github.com

# Limpiar cache de Go
go clean -modcache

# Intentar nuevamente
go install github.com/geomark27/loom-go/cmd/loom@latest
```

### Error: "go: github.com/geomark27/loom-go/cmd/loom@latest: no matching versions"

```bash
# Usar una versión específica
go install github.com/geomark27/loom-go/cmd/loom@v0.1.0

# O verificar que el repositorio sea público
# https://github.com/geomark27/loom-go
```

## 🌟 Comparación con Otras Herramientas

| Herramienta | Comando de Instalación |
|-------------|----------------------|
| **Loom** | `go install github.com/geomark27/loom-go/cmd/loom@latest` |
| Gin | `go get -u github.com/gin-gonic/gin` |
| Cobra CLI | `go install github.com/spf13/cobra-cli@latest` |
| Air (hot reload) | `go install github.com/cosmtrek/air@latest` |
| golangci-lint | `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` |

## 📚 Próximos Pasos

Después de instalar:

1. ✅ Leer la documentación: [README.md](README.md)
2. ✅ Ver la guía completa: [DESCRIPCION.md](DESCRIPCION.md)
3. ✅ Crear tu primer proyecto: `loom new mi-proyecto`
4. ✅ Explorar los ejemplos generados
5. ✅ Contribuir al proyecto: [CONTRIBUTING.md](CONTRIBUTING.md)

## 🤝 Compartir con tu Equipo

Para que tu equipo use Loom:

```bash
# En el README de tu proyecto, agrega:
# Instalación de herramientas
go install github.com/geomark27/loom-go/cmd/loom@latest

# O en un script de setup:
#!/bin/bash
echo "Instalando herramientas de desarrollo..."
go install github.com/geomark27/loom-go/cmd/loom@latest
echo "✅ Loom instalado"
```

## 📖 Documentación Adicional

- 📦 [Instalación Detallada](INSTALACION.md)
- 📋 [Descripción del Proyecto](DESCRIPCION.md)
- 🔧 [Guía de Contribución](CONTRIBUTING.md)
- 🐛 [Reportar Issues](https://github.com/geomark27/loom-go/issues)

---

**¡Ahora Loom está disponible para todos! Comparte este enlace:**

```
https://github.com/geomark27/loom-go
```

**Instala con:**

```bash
go install github.com/geomark27/loom-go/cmd/loom@latest
```

🧶 ¡Feliz tejido de proyectos!
