# 📦 Instalación de Loom

Guía completa para instalar Loom globalmente en tu sistema y usarlo desde cualquier directorio.

## 🚀 Instalación Global

### Requisitos Previos

- **Go 1.21+** instalado
- Variable de entorno `GOPATH` configurada (Go lo hace automáticamente)
- `$GOPATH/bin` agregado al PATH del sistema

### Verificar Configuración de Go

```bash
# Verificar versión de Go
go version

# Verificar GOPATH
go env GOPATH

# Verificar que GOPATH/bin está en el PATH
echo $PATH  # Linux/Mac
echo $env:PATH  # Windows PowerShell
```

## 📥 Método 1: Instalación desde el Código Fuente (Desarrollo)

Si tienes el código fuente de Loom:

```bash
# Navegar al directorio de Loom
cd /ruta/a/loom-go

# Instalar globalmente
go install ./cmd/loom
```

Esto compila el binario y lo coloca en `$GOPATH/bin/loom` (o `$GOPATH/bin/loom.exe` en Windows).

## 📥 Método 2: Instalación desde GitHub (Producción)

Una vez que Loom esté publicado en GitHub:

```bash
# Instalar directamente desde GitHub
go install github.com/geomark27/loom-go/cmd/loom@latest
```

## ✅ Verificar la Instalación

```bash
# Verificar que Loom está instalado
loom --version

# Ver ayuda
loom --help

# Ver comandos disponibles
loom
```

**Salida esperada:**
```
Loom - El tejedor de proyectos Go
version 0.1.0
```

## 🔧 Configuración del PATH (Si es necesario)

### Windows

Si `loom` no se reconoce, agrega `GOPATH\bin` al PATH:

1. Obtener GOPATH:
```powershell
go env GOPATH
# Ejemplo: C:\Users\TuUsuario\go
```

2. Agregar al PATH:
   - Ir a **Panel de Control** → **Sistema** → **Configuración avanzada del sistema**
   - Click en **Variables de entorno**
   - En **Variables del sistema**, buscar **Path** y hacer click en **Editar**
   - Agregar nueva entrada: `C:\Users\TuUsuario\go\bin`
   - Click en **Aceptar** en todas las ventanas
   - Reiniciar la terminal

3. Verificar:
```powershell
loom --version
```

### Linux/Mac

Si `loom` no se reconoce, agrega GOPATH/bin al PATH:

```bash
# Obtener GOPATH
go env GOPATH

# Agregar al PATH (agregar a ~/.bashrc, ~/.zshrc, etc.)
export PATH="$PATH:$(go env GOPATH)/bin"

# Recargar configuración
source ~/.bashrc  # o ~/.zshrc
```

## 🎯 Uso Básico

### Crear un Nuevo Proyecto

```bash
# Navegar al directorio donde quieres crear el proyecto
cd ~/proyectos

# Crear proyecto
loom new mi-api

# Entrar al proyecto
cd mi-api

# Instalar dependencias
go mod tidy

# Ejecutar
go run cmd/mi-api/main.go
```

### Desde Cualquier Directorio

Una vez instalado globalmente, puedes usar `loom` desde **cualquier directorio**:

```bash
# Estás en ~/documentos
cd ~/documentos
loom new proyecto-a

# Te mueves a otro lugar
cd ~/trabajo
loom new proyecto-b

# O incluso en el escritorio
cd ~/Desktop
loom new prototipo-rapido
```

## 🏗️ Ejemplo Completo de Flujo de Trabajo

```bash
# 1. Instalar Loom (solo una vez)
cd ~/go/src/loom-go
go install ./cmd/loom

# 2. Verificar instalación
loom --version

# 3. Ir a tu directorio de trabajo
cd ~/proyectos

# 4. Crear nuevo proyecto
loom new mi-super-api

# 5. Navegar al proyecto
cd mi-super-api

# 6. Ver estructura generada
ls -la

# 7. Instalar dependencias
go mod tidy

# 8. Ejecutar servidor
go run cmd/mi-super-api/main.go

# 9. Probar en otra terminal
curl http://localhost:8080/api/v1/health
```

## 🔄 Actualización de Loom

### Desde el Código Fuente

```bash
cd /ruta/a/loom-go
git pull origin main
go install ./cmd/loom
```

### Desde GitHub

```bash
go install github.com/geomark27/loom-go/cmd/loom@latest
```

## 🗑️ Desinstalación

Para desinstalar Loom:

```bash
# Encontrar la ubicación del binario
which loom  # Linux/Mac
where loom  # Windows

# Eliminar el binario
rm $(which loom)  # Linux/Mac
Remove-Item (Get-Command loom).Source  # Windows PowerShell
```

## 🐛 Solución de Problemas

### "loom: command not found" o "loom no se reconoce"

**Problema:** El PATH no está configurado correctamente.

**Solución:**
1. Verificar que Go esté instalado: `go version`
2. Verificar GOPATH: `go env GOPATH`
3. Agregar `$GOPATH/bin` al PATH (ver sección de configuración arriba)
4. Reiniciar la terminal

### "permission denied" (Linux/Mac)

**Problema:** El binario no tiene permisos de ejecución.

**Solución:**
```bash
chmod +x $(go env GOPATH)/bin/loom
```

### El proyecto generado no compila

**Problema:** Faltan dependencias.

**Solución:**
```bash
cd tu-proyecto
go mod tidy
go mod download
```

### Error al ejecutar `go install`

**Problema:** Módulo no inicializado correctamente.

**Solución:**
```bash
cd loom-go
go mod tidy
go install ./cmd/loom
```

## 📚 Comandos Útiles de Referencia

```bash
# Instalación
go install ./cmd/loom                    # Desde código fuente
go install github.com/user/loom@latest   # Desde GitHub

# Información
loom --version                           # Ver versión
loom --help                              # Ver ayuda
loom new --help                          # Ayuda del comando new

# Uso
loom new proyecto                        # Crear proyecto
loom new mi-api -v                       # Con salida verbose

# Mantenimiento
go clean -i github.com/user/loom/cmd/loom  # Limpiar instalación
```

## 🎓 Comparación con Otras Herramientas

| Herramienta | Comando de Instalación | Comando de Uso |
|-------------|----------------------|----------------|
| **Loom (Go)** | `go install ./cmd/loom` | `loom new proyecto` |
| Laravel (PHP) | `composer global require laravel/installer` | `laravel new proyecto` |
| NestJS (Node) | `npm install -g @nestjs/cli` | `nest new proyecto` |
| Rails (Ruby) | `gem install rails` | `rails new proyecto` |
| Django (Python) | `pip install django` | `django-admin startproject proyecto` |

## ✨ Próximos Pasos

Después de instalar Loom:

1. ✅ Crear tu primer proyecto: `loom new mi-proyecto`
2. 📖 Leer la documentación: `cat mi-proyecto/README.md`
3. 🚀 Ejecutar el servidor: `go run cmd/mi-proyecto/main.go`
4. 🧪 Probar los endpoints: `curl http://localhost:8080/api/v1/health`
5. 📝 Ver la documentación de API: `cat mi-proyecto/docs/API.md`

---

**¡Felicitaciones! Ahora tienes Loom instalado y listo para tejer proyectos Go como un profesional.** 🧶✨
