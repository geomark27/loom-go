# 🌍 Loom - Compatibilidad Multiplataforma

## ✅ Sistemas Operativos Soportados

Loom funciona en **TODOS** los sistemas operativos donde Go está disponible:

```
┌─────────────────────────────────────────────────────────┐
│              LOOM ES MULTIPLATAFORMA                    │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  🪟  Windows      ✅ Probado en Windows 10/11          │
│  🍎  macOS        ✅ Compatible (Intel y Apple Silicon) │
│  🐧  Linux        ✅ Todas las distribuciones           │
│  🐳  Docker       ✅ En contenedores                    │
│  ☁️  WSL          ✅ Windows Subsystem for Linux        │
│  📱  FreeBSD      ✅ Compatible                         │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

---

## 🎯 ¿Por Qué Es Multiplataforma?

### 1️⃣ **Go es Multiplataforma**

Go compila a binarios nativos para cada sistema operativo:

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o loom.exe

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o loom

# macOS (Apple Silicon - M1/M2/M3)
GOOS=darwin GOARCH=arm64 go build -o loom

# Linux
GOOS=linux GOARCH=amd64 go build -o loom

# FreeBSD
GOOS=freebsd GOARCH=amd64 go build -o loom
```

### 2️⃣ **Loom Usa Código Portable**

```go
// ✅ Código multiplataforma en Loom
import (
    "os"           // ✅ Cross-platform
    "path/filepath" // ✅ Maneja rutas de Windows, Mac, Linux
    "text/template" // ✅ Cross-platform
)

// Ejemplo: Crear directorios funciona en todos los OS
os.MkdirAll(dirPath, 0755)  // ✅ Windows, Mac, Linux

// Ejemplo: Rutas funcionan en todos los OS
filepath.Join("internal", "app", "handlers")
// Windows:  internal\app\handlers
// Unix:     internal/app/handlers
```

---

## 📦 Instalación por Sistema Operativo

### 🪟 **Windows (PowerShell/CMD)**

```powershell
# Instalar Go (si no lo tienes)
# Descargar de: https://go.dev/dl/

# Verificar instalación de Go
go version
# Output: go version go1.23.4 windows/amd64

# Instalar Loom
go install github.com/geomark27/loom-go/cmd/loom@latest

# Verificar instalación
loom --version
# Output: loom version 0.1.0

# Crear proyecto
loom new mi-proyecto-windows
cd mi-proyecto-windows
go mod tidy
go run cmd/mi-proyecto-windows/main.go
```

**Ubicación del binario en Windows:**
```
C:\Users\TU_USUARIO\go\bin\loom.exe
```

---

### 🍎 **macOS (Terminal/Zsh/Bash)**

```bash
# Instalar Go (si no lo tienes)
# Opción 1: Homebrew
brew install go

# Opción 2: Descargar de https://go.dev/dl/

# Verificar instalación de Go
go version
# Output: go version go1.23.4 darwin/amd64  (Intel)
# Output: go version go1.23.4 darwin/arm64  (M1/M2/M3)

# Instalar Loom
go install github.com/geomark27/loom-go/cmd/loom@latest

# Verificar instalación
loom --version
# Output: loom version 0.1.0

# Crear proyecto
loom new mi-proyecto-mac
cd mi-proyecto-mac
go mod tidy
go run cmd/mi-proyecto-mac/main.go
```

**Ubicación del binario en macOS:**
```
/Users/tu-usuario/go/bin/loom
```

---

### 🐧 **Linux (Bash/Zsh)**

```bash
# Instalar Go (si no lo tienes)
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# Fedora
sudo dnf install golang

# Arch Linux
sudo pacman -S go

# O descargar de: https://go.dev/dl/

# Verificar instalación de Go
go version
# Output: go version go1.23.4 linux/amd64

# Instalar Loom
go install github.com/geomark27/loom-go/cmd/loom@latest

# Verificar instalación
loom --version
# Output: loom version 0.1.0

# Crear proyecto
loom new mi-proyecto-linux
cd mi-proyecto-linux
go mod tidy
go run cmd/mi-proyecto-linux/main.go
```

**Ubicación del binario en Linux:**
```
/home/tu-usuario/go/bin/loom
```

---

### ☁️ **WSL (Windows Subsystem for Linux)**

```bash
# Dentro de WSL (Ubuntu, Debian, etc.)
# Instalar Go
sudo apt update
sudo apt install golang-go

# Instalar Loom
go install github.com/geomark27/loom-go/cmd/loom@latest

# Verificar
loom --version

# Crear proyecto
loom new mi-proyecto-wsl
cd mi-proyecto-wsl
go mod tidy
go run cmd/mi-proyecto-wsl/main.go
```

---

### 🐳 **Docker (Cualquier OS con Docker)**

```dockerfile
# Dockerfile
FROM golang:1.23-alpine

# Instalar Loom
RUN go install github.com/geomark27/loom-go/cmd/loom@latest

# Usar Loom
WORKDIR /workspace
RUN loom new mi-proyecto-docker

# Compilar el proyecto generado
WORKDIR /workspace/mi-proyecto-docker
RUN go mod tidy
RUN go build -o app cmd/mi-proyecto-docker/main.go

CMD ["./app"]
```

```bash
# Construir y ejecutar
docker build -t mi-app-loom .
docker run -p 8080:8080 mi-app-loom
```

---

## 🎯 Comando Universal

**El mismo comando funciona en TODOS los sistemas operativos:**

```bash
# ✅ Windows (PowerShell)
go install github.com/geomark27/loom-go/cmd/loom@latest

# ✅ macOS (Terminal)
go install github.com/geomark27/loom-go/cmd/loom@latest

# ✅ Linux (Bash)
go install github.com/geomark27/loom-go/cmd/loom@latest

# ✅ WSL (Ubuntu)
go install github.com/geomark27/loom-go/cmd/loom@latest
```

---

## 📊 Diferencias por Sistema Operativo

### Solo hay diferencias menores en:

#### 1️⃣ **Extensión del Binario**

```bash
# Windows
C:\Users\Marcos\go\bin\loom.exe  ← .exe

# macOS/Linux/WSL
/Users/marcos/go/bin/loom        ← sin extensión
/home/marcos/go/bin/loom
```

#### 2️⃣ **Variable de Entorno PATH**

```bash
# Windows (PowerShell)
$env:PATH += ";C:\Users\Marcos\go\bin"

# macOS/Linux (Bash/Zsh)
export PATH=$PATH:$HOME/go/bin
```

#### 3️⃣ **Separador de Rutas**

```bash
# Windows
internal\app\handlers  ← Backslash \

# macOS/Linux
internal/app/handlers  ← Forward slash /
```

**Pero Loom maneja esto automáticamente con `filepath.Join()`** ✅

---

## 🧪 Pruebas Multiplataforma

### Proyectos Generados Son Idénticos

```bash
# ═══════════════════════════════════════════════════════
# Windows
# ═══════════════════════════════════════════════════════
C:\Users\Marcos> loom new test-windows
✅ Proyecto 'test-windows' creado exitosamente

test-windows/
├── cmd/
├── internal/
├── go.mod
└── ...

# ═══════════════════════════════════════════════════════
# macOS
# ═══════════════════════════════════════════════════════
~/proyectos $ loom new test-mac
✅ Proyecto 'test-mac' creado exitosamente

test-mac/
├── cmd/
├── internal/
├── go.mod
└── ...

# ═══════════════════════════════════════════════════════
# Linux
# ═══════════════════════════════════════════════════════
~/proyectos $ loom new test-linux
✅ Proyecto 'test-linux' creado exitosamente

test-linux/
├── cmd/
├── internal/
├── go.mod
└── ...
```

**Los 3 proyectos son IDÉNTICOS en estructura y código** ✅

---

## 🌐 Casos de Uso Multiplataforma

### Escenario 1: Equipo Distribuido

```
┌─────────────────────────────────────────────────────┐
│              EQUIPO INTERNACIONAL                   │
├─────────────────────────────────────────────────────┤
│                                                     │
│  👨‍💻 Dev 1 (México)    → Windows 11 + Loom        │
│  👩‍💻 Dev 2 (España)    → macOS M2 + Loom          │
│  👨‍💻 Dev 3 (Argentina) → Ubuntu Linux + Loom      │
│  👩‍💻 Dev 4 (USA)       → WSL2 Ubuntu + Loom       │
│                                                     │
│  ✅ Todos usan:                                     │
│     loom new mi-proyecto                           │
│                                                     │
│  ✅ Todos obtienen:                                 │
│     Estructura idéntica                            │
│     Código idéntico                                │
│     Sin conflictos                                 │
│                                                     │
└─────────────────────────────────────────────────────┘
```

---

### Escenario 2: CI/CD Multiplataforma

```yaml
# .github/workflows/test-multiplataforma.yml
name: Test Multiplataforma

on: [push, pull_request]

jobs:
  test-windows:
    runs-on: windows-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.23'
      - name: Install Loom
        run: go install github.com/geomark27/loom-go/cmd/loom@latest
      - name: Test Loom
        run: |
          loom new test-project
          cd test-project
          go mod tidy
          go build ./...

  test-macos:
    runs-on: macos-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.23'
      - name: Install Loom
        run: go install github.com/geomark27/loom-go/cmd/loom@latest
      - name: Test Loom
        run: |
          loom new test-project
          cd test-project
          go mod tidy
          go build ./...

  test-linux:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.23'
      - name: Install Loom
        run: go install github.com/geomark27/loom-go/cmd/loom@latest
      - name: Test Loom
        run: |
          loom new test-project
          cd test-project
          go mod tidy
          go build ./...
```

---

## 🎯 Arquitecturas Soportadas

Loom soporta todas las arquitecturas que Go soporta:

```
┌──────────────────────────────────────────────────────┐
│           ARQUITECTURAS SOPORTADAS                   │
├──────────────────────────────────────────────────────┤
│                                                      │
│  💻 amd64 (x86_64)    ✅ Intel/AMD 64-bit           │
│  🍎 arm64 (aarch64)   ✅ Apple Silicon M1/M2/M3     │
│  📱 arm               ✅ ARM 32-bit                  │
│  🖥️  386              ✅ Intel/AMD 32-bit           │
│  🔧 ppc64le           ✅ PowerPC 64-bit LE          │
│  🔧 s390x             ✅ IBM System z               │
│  🔧 mips              ✅ MIPS                        │
│  🔧 riscv64           ✅ RISC-V 64-bit              │
│                                                      │
└──────────────────────────────────────────────────────┘
```

---

## 🚀 Ventajas de la Multiplataforma

### ✅ **Para Developers**

```
🌍 Flexibilidad
   ├── Trabaja desde cualquier OS
   ├── Cambia de laptop sin problemas
   └── Mismo flujo de trabajo

👥 Colaboración
   ├── Equipos distribuidos
   ├── Sin "funciona en mi máquina"
   └── Onboarding simple

🔄 Portabilidad
   ├── Código portable
   ├── Proyectos portable
   └── Sin dependencias del OS
```

### ✅ **Para Empresas**

```
💰 Ahorro de Costos
   ├── Sin licencias específicas de OS
   ├── Developers eligen su OS favorito
   └── Infraestructura flexible

🎯 Productividad
   ├── Misma herramienta en todos lados
   ├── Sin curva de aprendizaje por OS
   └── Consistencia garantizada

📈 Escalabilidad
   ├── Deploy en cualquier cloud
   ├── CI/CD multiplataforma
   └── Contenedores sin problemas
```

---

## 📝 Guía Rápida por Sistema Operativo

### Si Eres Usuario de Windows 🪟

```powershell
# 1. Instalar Go desde: https://go.dev/dl/
# 2. Abrir PowerShell
# 3. Ejecutar:
go install github.com/geomark27/loom-go/cmd/loom@latest

# 4. Verificar PATH (si no funciona loom):
$env:PATH += ";$env:USERPROFILE\go\bin"

# 5. Usar Loom:
loom new mi-proyecto
cd mi-proyecto
go run cmd/mi-proyecto/main.go
```

---

### Si Eres Usuario de Mac 🍎

```bash
# 1. Instalar Go:
brew install go

# 2. Instalar Loom:
go install github.com/geomark27/loom-go/cmd/loom@latest

# 3. Agregar a PATH (agregar a ~/.zshrc o ~/.bashrc):
export PATH=$PATH:$HOME/go/bin

# 4. Recargar shell:
source ~/.zshrc  # o source ~/.bashrc

# 5. Usar Loom:
loom new mi-proyecto
cd mi-proyecto
go run cmd/mi-proyecto/main.go
```

---

### Si Eres Usuario de Linux 🐧

```bash
# 1. Instalar Go:
# Ubuntu/Debian:
sudo apt install golang-go

# Fedora:
sudo dnf install golang

# Arch:
sudo pacman -S go

# 2. Instalar Loom:
go install github.com/geomark27/loom-go/cmd/loom@latest

# 3. Agregar a PATH (agregar a ~/.bashrc):
export PATH=$PATH:$HOME/go/bin

# 4. Recargar shell:
source ~/.bashrc

# 5. Usar Loom:
loom new mi-proyecto
cd mi-proyecto
go run cmd/mi-proyecto/main.go
```

---

## 🎯 Respuesta Corta

### ¿Loom funciona en tu sistema operativo?

```
┌─────────────────────────────────────────────────┐
│                                                 │
│   ¿Tienes Go instalado?                        │
│                                                 │
│   SÍ  → ✅ Loom funcionará perfectamente       │
│   NO  → 📦 Instala Go, luego usa Loom         │
│                                                 │
│   Go está disponible en:                       │
│   • Windows (todas las versiones)              │
│   • macOS (Intel y Apple Silicon)              │
│   • Linux (todas las distribuciones)           │
│   • FreeBSD, OpenBSD, NetBSD                   │
│   • Solaris                                     │
│   • Android                                     │
│   • iOS                                         │
│   • Plan 9                                      │
│   • Y más...                                    │
│                                                 │
└─────────────────────────────────────────────────┘
```

---

## 🏆 Conclusión

**Loom es 100% multiplataforma** porque:

1. ✅ Está escrito en Go (lenguaje multiplataforma)
2. ✅ Usa solo librerías estándar portables
3. ✅ No tiene dependencias del sistema operativo
4. ✅ Se compila a binarios nativos
5. ✅ El mismo comando funciona en todos los OS
6. ✅ Genera proyectos idénticos en todos los OS
7. ✅ Probado en Windows, macOS, Linux

---

**"Un comando, cualquier sistema operativo, mismos resultados."** 🧶✨🌍

### Instala Loom hoy en tu OS favorito:

```bash
go install github.com/geomark27/loom-go/cmd/loom@latest
```

**Funciona en:** Windows • macOS • Linux • WSL • FreeBSD • Docker • Kubernetes • Y más...
