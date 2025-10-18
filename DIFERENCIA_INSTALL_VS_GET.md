# 🎯 Loom: CLI Tool vs Library - ¿Cuál es la diferencia?

## 🤔 La Confusión Común

**Tu pregunta es válida y muy común:** 
> "Si `go install` es para agregar dependencias a un proyecto, ¿por qué usarlo para Loom si Loom es para CREAR proyectos nuevos?"

## ✅ La Respuesta: Loom NO es una dependencia

### 📦 Dos Tipos de Paquetes Go

```
┌─────────────────────────────────────────────────────────┐
│                    PAQUETES GO                          │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  1️⃣  LIBRARIES (Dependencias de Proyecto)              │
│     ↓                                                   │
│     Se agregan al go.mod de TU proyecto                │
│     Ejemplo: gin, gorm, cobra                          │
│                                                         │
│  2️⃣  CLI TOOLS (Herramientas del Sistema)              │
│     ↓                                                   │
│     Se instalan GLOBALMENTE en tu computadora          │
│     Ejemplo: loom, air, cobra-cli, migrate            │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

---

## 📚 Tipo 1: LIBRARY (Dependencia de Proyecto)

### Ejemplo: Gin (Framework Web)

```bash
# Estás en un proyecto existente
cd ~/mi-proyecto-existente

# Agregas Gin como dependencia
go get -u github.com/gin-gonic/gin

# Resultado: Se agrega a go.mod
```

**Archivo `go.mod`:**
```go
module mi-proyecto

go 1.23

require (
    github.com/gin-gonic/gin v1.9.1  // ← Dependencia del PROYECTO
)
```

**Usas Gin EN TU CÓDIGO:**
```go
package main

import "github.com/gin-gonic/gin"  // ← Importas la librería

func main() {
    r := gin.Default()  // ← Usas el código de Gin
    r.Run()
}
```

### 🎯 Características de una LIBRARY:
- ✅ Se importa en tu código (`import "github.com/..."`)
- ✅ Está en el `go.mod` de tu proyecto
- ✅ Es diferente en cada proyecto
- ✅ Se usa DENTRO del código Go
- ❌ NO ejecutas comandos en la terminal

---

## 🛠️ Tipo 2: CLI TOOL (Herramienta Global)

### Ejemplo: Loom (Generador de Proyectos)

```bash
# NO estás en ningún proyecto
# Puedes estar en cualquier directorio

# Instalas Loom GLOBALMENTE en tu sistema
go install github.com/geomark27/loom-go/cmd/loom@latest

# Resultado: Binario instalado en $GOPATH/bin o $GOBIN
# Típicamente: ~/go/bin/loom (Linux/Mac) o C:\Users\TU_USUARIO\go\bin\loom.exe (Windows)
```

**NO aparece en ningún `go.mod`:**
```go
module mi-proyecto

go 1.23

require (
    // ← Loom NO está aquí
    // ← Nunca estará aquí
)
```

**NO lo importas en tu código:**
```go
package main

// import "github.com/geomark27/loom-go"  ← ❌ NUNCA harás esto

func main() {
    // loom.New()  ← ❌ NO existe
}
```

**Lo ejecutas en la TERMINAL:**
```bash
# Usas Loom como comando del sistema
loom new mi-proyecto
loom --version
loom --help
```

### 🎯 Características de un CLI TOOL:
- ✅ Se ejecuta en la terminal como comando
- ✅ Está instalado GLOBALMENTE (una sola vez)
- ✅ Disponible en CUALQUIER directorio
- ✅ Crea/modifica archivos desde fuera
- ❌ NUNCA se importa en código Go
- ❌ NUNCA aparece en go.mod de tus proyectos

---

## 🔄 Comparación Práctica

### Escenario: Crear una API

#### ❌ **INCORRECTO** - Intentar usar Loom como library:

```bash
# Paso 1: Crear proyecto manualmente
mkdir mi-api
cd mi-api
go mod init mi-api

# Paso 2: Intentar agregar Loom como dependencia
go get github.com/geomark27/loom-go  # ❌ ERROR conceptual

# Paso 3: Intentar usarlo en código
# main.go
import "github.com/geomark27/loom-go"  // ❌ No funciona así

func main() {
    loom.GenerateProject("mi-api")  // ❌ No existe esta API
}
```

**Problema:** Loom NO es una librería, es una herramienta.

---

#### ✅ **CORRECTO** - Usar Loom como CLI tool:

```bash
# Paso 1: Instalar Loom GLOBALMENTE (solo una vez en tu vida)
go install github.com/geomark27/loom-go/cmd/loom@latest

# Paso 2: Usar Loom para CREAR el proyecto (desde la terminal)
loom new mi-api

# Paso 3: Loom crea TODO por ti
# ✅ Crea carpetas
# ✅ Genera archivos
# ✅ Crea go.mod
# ✅ Instala dependencias (gin, mux, etc.)

# Paso 4: Tu proyecto YA está listo
cd mi-api
go run cmd/mi-api/main.go
```

**Resultado:** Proyecto completo sin que Loom sea parte del código.

---

## 🎯 Analogía con Otras Herramientas

### Herramientas Globales vs Dependencias de Proyecto

```
┌──────────────────────────────────────────────────────────────────┐
│                    HERRAMIENTAS GLOBALES                         │
│                    (Instalas UNA VEZ)                            │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  🛠️  Loom        → Genera proyectos Go                          │
│  🛠️  npm         → Gestiona paquetes Node.js                    │
│  🛠️  composer    → Gestiona dependencias PHP                    │
│  🛠️  create-react-app → Crea proyectos React                    │
│  🛠️  nest-cli    → Crea proyectos NestJS                        │
│  🛠️  django-admin → Crea proyectos Django                       │
│  🛠️  rails       → Crea proyectos Rails                         │
│  🛠️  git         → Control de versiones                         │
│  🛠️  docker      → Containerización                             │
│                                                                  │
│  Características:                                               │
│  • Se instalan en el SISTEMA OPERATIVO                          │
│  • Disponibles en CUALQUIER carpeta                             │
│  • Ejecutas comandos en la terminal                             │
│  • NO aparecen en package.json/go.mod/composer.json             │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────────┐
│                 DEPENDENCIAS DE PROYECTO                         │
│              (Diferentes en CADA proyecto)                       │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  📦  gin         → Framework web para Go                         │
│  📦  gorm        → ORM para Go                                   │
│  📦  express     → Framework web para Node.js                    │
│  📦  react       → Librería UI para JavaScript                   │
│  📦  axios       → Cliente HTTP para JavaScript                  │
│  📦  lodash      → Utilidades para JavaScript                    │
│  📦  laravel     → Framework para PHP                            │
│  📦  requests    → Cliente HTTP para Python                      │
│                                                                  │
│  Características:                                               │
│  • Están en go.mod/package.json del PROYECTO                    │
│  • Se importan en el CÓDIGO                                     │
│  • Diferentes versiones en cada proyecto                        │
│  • NO se ejecutan como comandos                                 │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

---

## 🎬 Flujo Real de Uso de Loom

### 📅 Timeline de un Developer

#### **Día 1 - UNA SOLA VEZ - Instalación de Loom**

```bash
# En CUALQUIER directorio de tu computadora
cd ~
# o
cd C:\Users\Marcos
# o donde sea

# Instalas Loom GLOBALMENTE
go install github.com/geomark27/loom-go/cmd/loom@latest

# Verificas
loom --version
# Output: loom version 0.1.0

# ✅ Loom ya está disponible en TODA tu computadora
```

**Resultado:**
```
Sistema Operativo
├── Programas Instalados
│   ├── Chrome ✅
│   ├── VS Code ✅
│   ├── Git ✅
│   ├── Docker ✅
│   ├── Node.js ✅
│   └── Loom ✅  ← Instalado como herramienta del sistema
└── Proyectos
    ├── proyecto-1/
    ├── proyecto-2/
    └── proyecto-3/
```

---

#### **Día 2 - Proyecto 1: API de Tienda**

```bash
# Usas Loom (ya instalado globalmente)
cd ~/proyectos
loom new tienda-api

# Loom genera:
tienda-api/
├── go.mod              ← Loom lo crea, pero Loom NO está dentro
│   require (
│       github.com/gorilla/mux v1.8.1  ← Estas sí son dependencias
│   )
├── cmd/
├── internal/
└── ...

# Entras al proyecto
cd tienda-api
go run cmd/tienda-api/main.go

# ✅ Proyecto funcionando
# ❌ Loom NO es parte del proyecto
# ✅ Loom solo fue la herramienta que lo creó
```

---

#### **Día 5 - Proyecto 2: API de Blog**

```bash
# Usas Loom OTRA VEZ (ya instalado, no reinstalar)
cd ~/proyectos
loom new blog-api

# Loom genera OTRO proyecto
blog-api/
├── go.mod              ← Loom lo crea, pero Loom NO está dentro
│   require (
│       github.com/gorilla/mux v1.8.1
│   )
├── cmd/
├── internal/
└── ...

# ✅ Segundo proyecto funcionando
# ❌ Loom NO es parte de este proyecto tampoco
```

---

#### **Día 10 - Proyecto 3: Microservicio de Notificaciones**

```bash
# Usas Loom NUEVAMENTE (misma instalación global)
cd ~/empresa/microservicios
loom new notification-service

# ✅ Tercer proyecto funcionando
# ❌ Loom NO es parte de este proyecto tampoco
```

---

### 📊 Resumen Visual

```
┌─────────────────────────────────────────────────────────────────┐
│                         TU COMPUTADORA                          │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  🛠️  LOOM (Instalado 1 vez globalmente)                         │
│     ↓                                                           │
│     C:\Users\Marcos\go\bin\loom.exe                            │
│                                                                 │
│  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━  │
│                                                                 │
│  📁 Proyectos Creados (Loom NO está en go.mod)                  │
│                                                                 │
│  ├── 📂 tienda-api/                                             │
│  │   ├── go.mod  ← require github.com/gorilla/mux              │
│  │   ├── cmd/                                                   │
│  │   └── internal/                                              │
│  │                                                              │
│  ├── 📂 blog-api/                                               │
│  │   ├── go.mod  ← require github.com/gorilla/mux              │
│  │   ├── cmd/                                                   │
│  │   └── internal/                                              │
│  │                                                              │
│  └── 📂 notification-service/                                   │
│      ├── go.mod  ← require github.com/gorilla/mux              │
│      ├── cmd/                                                   │
│      └── internal/                                              │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## 🤔 ¿Por Qué Esta Confusión?

### El comando `go install` tiene dos usos diferentes:

#### 1️⃣ **Instalar CLI Tool (Herramienta Global)**

```bash
# Instala el BINARIO en $GOPATH/bin
go install github.com/geomark27/loom-go/cmd/loom@latest

# Resultado: Ejecutable en el sistema
# Ubicación: ~/go/bin/loom o C:\Users\Marcos\go\bin\loom.exe
# Uso: loom new proyecto
```

#### 2️⃣ **Compilar Proyecto Local**

```bash
# Dentro de un proyecto
cd ~/mi-proyecto
go install

# Resultado: Compila TU proyecto y lo instala en $GOPATH/bin
# Uso: Para instalar tu propio programa
```

### La Clave: El `/cmd/` en la ruta

```bash
# Con /cmd/ → CLI Tool
go install github.com/geomark27/loom-go/cmd/loom@latest
                                         ^^^^
                                      Esto indica que es un COMANDO

# Sin /cmd/ → Library (usualmente con go get)
go get github.com/gin-gonic/gin
       ^^^ 
    Agrega al go.mod del proyecto
```

---

## 🎯 Entonces, ¿Cuándo Instalo Loom?

### ✅ **Instalas Loom:**

1. **Primera vez** que quieres usar Loom
2. **Una sola vez** en tu computadora
3. **ANTES de crear** cualquier proyecto
4. **Desde cualquier** directorio

```bash
# No importa dónde estés
cd ~
# o
cd C:\
# o
cd /cualquier/lugar

# Instalas Loom
go install github.com/geomark27/loom-go/cmd/loom@latest
```

### ✅ **Usas Loom:**

1. **Cada vez** que quieres crear un proyecto nuevo
2. **Desde la terminal**, no desde código
3. **En el directorio** donde quieres el proyecto

```bash
# Donde quieres crear proyectos
cd ~/proyectos

# Usas Loom
loom new mi-api-1
loom new mi-api-2
loom new mi-api-3
```

### ❌ **NO instalas Loom:**

1. ❌ En cada proyecto que creas
2. ❌ En el go.mod de tus proyectos
3. ❌ Dentro del código Go
4. ❌ Como dependencia del proyecto

---

## 🆚 Comparación Final: Loom vs Gin

### Gin (Library - Dependencia de Proyecto)

```bash
# Paso 1: Crear proyecto manualmente
mkdir mi-api
cd mi-api
go mod init mi-api

# Paso 2: Agregar Gin al proyecto
go get github.com/gin-gonic/gin

# Paso 3: Usar Gin EN el código
# main.go
import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.Run()
}

# Paso 4: Ejecutar el proyecto
go run main.go
```

**go.mod:**
```
require github.com/gin-gonic/gin v1.9.1  ← Gin ESTÁ aquí
```

---

### Loom (CLI Tool - Herramienta Global)

```bash
# Paso 1: Instalar Loom (una sola vez)
go install github.com/geomark27/loom-go/cmd/loom@latest

# Paso 2: Crear proyecto con Loom
loom new mi-api

# Paso 3: Entrar al proyecto
cd mi-api

# Paso 4: Ejecutar (Loom ya hizo todo)
go run cmd/mi-api/main.go
```

**go.mod:**
```
require github.com/gorilla/mux v1.8.1    ← Loom NO está aquí
                                         ← Loom nunca estará aquí
```

---

## 🎓 Conclusión

### Loom es como:

| Herramienta | Tipo | Instalación | Uso |
|-------------|------|-------------|-----|
| **Loom** | CLI Tool | Una vez global | Crear proyectos desde terminal |
| **npm** | CLI Tool | Una vez global | Gestionar paquetes desde terminal |
| **create-react-app** | CLI Tool | Una vez global | Crear proyectos React desde terminal |
| **git** | CLI Tool | Una vez global | Control de versiones desde terminal |

### Loom NO es como:

| Herramienta | Tipo | Instalación | Uso |
|-------------|------|-------------|-----|
| **Gin** | Library | Por proyecto en go.mod | Importar en código Go |
| **Express** | Library | Por proyecto en package.json | Importar en código Node.js |
| **React** | Library | Por proyecto en package.json | Importar en código JavaScript |

---

## 🚀 Respuesta Corta a Tu Pregunta

> **"¿Cuándo aplica instalar `go install github.com/geomark27/loom-go/cmd/loom@latest`?"**

**Respuesta:** 
- ✅ **Una sola vez** en tu vida (o cuando quieras actualizar Loom)
- ✅ **ANTES de crear** proyectos
- ✅ **Desde cualquier** directorio

> **"¿Y cuándo el uso del mismo?"**

**Respuesta:**
- ✅ **Cada vez** que quieras crear un proyecto nuevo
- ✅ **Desde la terminal**: `loom new nombre-proyecto`
- ✅ **No necesitas reinstalar** Loom cada vez

---

**En resumen:** Loom es una herramienta que installas UNA vez y usas INFINITAS veces. 🧶✨
