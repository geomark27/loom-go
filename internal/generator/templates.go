package generator

// getTemplates retorna un mapa con todas las plantillas embebidas
func getTemplates() map[string]string {
	templates := map[string]string{
		"go.mod.tmpl": `module {{.ModuleName}}

go 1.23

require (
	github.com/gorilla/mux v1.8.1
)
`,
		"README.md.tmpl": `# {{.Name}}

{{.Description}}

## 🚀 Características

- ✅ **Arquitectura modular** inspirada en NestJS
- ✅ **Estructura idiomática de Go** siguiendo golang-standard/project-layout
- ✅ **API REST** con endpoints CRUD de usuarios
- ✅ **Inyección de dependencias** clara y mantenible
- ✅ **Middleware CORS** configurado
- ✅ **Health checks** implementados
- ✅ **Documentación de API** incluida
- ✅ **Makefile** con comandos útiles
- ✅ **Variables de entorno** configuradas

## 📁 Estructura del Proyecto

` + "```" + `
{{.Name}}/
├── cmd/
│   └── {{.Name}}/
│       └── main.go              # Punto de entrada
├── internal/
│   ├── app/
│   │   ├── handlers/            # Controllers (HTTP handlers)
│   │   │   ├── health_handler.go
│   │   │   └── user_handler.go
│   │   ├── services/            # Lógica de negocio
│   │   │   └── user_service.go
│   │   ├── dtos/                # Data Transfer Objects
│   │   │   └── user_dto.go
│   │   ├── models/              # Modelos de datos
│   │   │   └── user.go
│   │   ├── repositories/        # Capa de persistencia
│   │   │   └── user_repository.go
│   │   └── middleware/          # Middlewares HTTP
│   │       └── cors_middleware.go
│   ├── config/                  # Configuración
│   │   └── config.go
│   └── server/                  # Configuración del servidor
│       ├── server.go
│       └── routes.go
├── pkg/                         # Código reutilizable
├── docs/                        # Documentación
│   └── API.md
├── scripts/                     # Scripts de build/deploy
├── .env.example                 # Variables de entorno de ejemplo
├── .gitignore
├── Makefile                     # Comandos de desarrollo
├── go.mod
└── README.md
` + "```" + `

## 🏃‍♂️ Inicio Rápido

### Instalación

` + "```bash" + `
cd {{.Name}}
cp .env.example .env
go mod tidy
` + "```" + `

### Ejecución

` + "```bash" + `
# Usando Go directamente
go run cmd/{{.Name}}/main.go

# O usando Makefile
make run
` + "```" + `

El servidor estará disponible en: **http://localhost:8080**

### Desarrollo

` + "```bash" + `
# Ver todos los comandos disponibles
make help

# Compilar
make build

# Ejecutar tests
make test

# Formatear código
make fmt

# Analizar código
make vet
` + "```" + `

## 🔌 API Endpoints

| Método | Endpoint | Descripción |
|---------|----------|-------------|
| GET | ` + "`/`" + ` | Información general |
| GET | ` + "`/api/v1/health`" + ` | Estado del servicio |
| GET | ` + "`/api/v1/users`" + ` | Obtener todos los usuarios |
| POST | ` + "`/api/v1/users`" + ` | Crear usuario |
| GET | ` + "`/api/v1/users/{id}`" + ` | Obtener usuario por ID |
| PUT | ` + "`/api/v1/users/{id}`" + ` | Actualizar usuario |
| DELETE | ` + "`/api/v1/users/{id}`" + ` | Eliminar usuario |

📖 **Documentación detallada**: [docs/API.md](docs/API.md)

## 🧪 Pruebas Rápidas

` + "```bash" + `
# Obtener información general
curl http://localhost:8080

# Health check
curl http://localhost:8080/api/v1/health

# Obtener usuarios
curl http://localhost:8080/api/v1/users

# Crear usuario
curl -X POST http://localhost:8080/api/v1/users \\
  -H "Content-Type: application/json" \\
  -d '{"name": "Ana García", "email": "ana@example.com", "age": 28}'
` + "```" + `

## 🏗️ Arquitectura

Este proyecto sigue el patrón de **arquitectura por capas** inspirado en frameworks como NestJS:

- **Handlers**: Manejan las peticiones HTTP y las respuestas
- **Services**: Contienen la lógica de negocio
- **Repositories**: Manejan la persistencia de datos  
- **DTOs**: Definen la estructura de datos de entrada/salida
- **Models**: Representan las entidades del dominio
- **Middleware**: Procesan las peticiones de forma transversal

## 🔧 Configuración

Las variables de entorno se definen en ` + "`" + `.env` + "`" + `:

` + "```bash" + `
PORT=8080
ENVIRONMENT=development
LOG_LEVEL=info
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8080
` + "```" + `

## 📚 Próximos Pasos

1. **Agregar base de datos**: Reemplazar el repositorio en memoria
2. **Implementar autenticación**: JWT, OAuth, etc.
3. **Agregar validaciones**: Validador de structs más robusto
4. **Tests**: Crear tests unitarios e integración  
5. **Logging**: Implementar logging estructurado
6. **Métricas**: Prometheus, health checks avanzados

## 🛠️ Generado con

Este proyecto fue generado con [**Loom**](https://github.com/geomark27/loom-go) - El tejedor de proyectos Go.

¡Disfruta desarrollando con Go! 🎉
`,
		"main.go.tmpl": `package main

import (
	"log"

	"{{.ModuleName}}/internal/config"
	"{{.ModuleName}}/internal/server"
)

func main() {
	// Cargar configuración
	cfg := config.Load()

	// Crear servidor
	srv := server.New(cfg)

	// Mensaje de inicio
	log.Printf("🚀 Servidor %s iniciado en http://localhost:%s", "{{.Name}}", cfg.Port)
	log.Printf("✨ Proyecto generado con Loom")
	log.Printf("📖 Documentación disponible en: docs/API.md")

	// Iniciar servidor
	if err := srv.Start(); err != nil {
		log.Fatal("Error iniciando servidor:", err)
	}
}
`,
		".gitignore.tmpl": `# Binarios para programas y plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Archivos de test
*.test

# Archivos de cobertura de salida
*.out
*.html

# Directorios de dependencias (vendor tree)
vendor/

# Go workspace file
go.work

# Variables de entorno y configuración
.env
.env.local
.env.development
.env.production
.env.test
*.env
config.local.*
secrets.json
.secrets/

# Archivos específicos del IDE
.vscode/
.idea/
*.swp
*.swo
*~

# Archivos de log
*.log
logs/
log/

# Archivos de OS
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db

# Archivos de build y temporales
dist/
build/
tmp/
temp/
.tmp/
.cache/

# Archivos de database locales
*.db
*.sqlite
*.sqlite3
database.sql

# Archivos de backup
*.bak
*.backup
*.old

# Archivos de profiling
*.prof
*.pprof

# Coverage reports
coverage.txt
coverage.html
coverage.out

# Air (hot reload)
.air.toml
tmp/

# Delve debugger
__debug_bin

# Certificados y claves
*.pem
*.key
*.crt
*.cert

# Archivos de deployment
.terraform/
terraform.tfstate*
*.tfvars
`,
		".env.example.tmpl": `# Configuración del servidor
PORT=8080
ENVIRONMENT=development

# Base de datos (cuando se implemente)
# DATABASE_URL=postgres://user:password@localhost:5432/dbname

# JWT Secret (cuando se implemente autenticación)
# JWT_SECRET=your-super-secret-jwt-key

# Logging
LOG_LEVEL=info

# CORS
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8080
`,
		"Makefile.tmpl": `# {{.Name}} - Makefile generado por Loom

.PHONY: build run test clean fmt vet deps help

# Variables
APP_NAME={{.Name}}
BUILD_DIR=build
CMD_DIR=cmd/$(APP_NAME)

# Comandos principales
help: ## Muestra esta ayuda
	@echo "Comandos disponibles:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Compila la aplicación
	@echo "🔨 Compilando $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)/main.go
	@echo "✅ Compilación exitosa: $(BUILD_DIR)/$(APP_NAME)"

run: ## Ejecuta la aplicación
	@echo "🚀 Ejecutando $(APP_NAME)..."
	@go run $(CMD_DIR)/main.go

test: ## Ejecuta los tests
	@echo "🧪 Ejecutando tests..."
	@go test -v ./...

test-coverage: ## Ejecuta tests con cobertura
	@echo "🧪 Ejecutando tests con cobertura..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "📊 Reporte de cobertura generado: coverage.html"

fmt: ## Formatea el código
	@echo "🎨 Formateando código..."
	@go fmt ./...

vet: ## Ejecuta go vet
	@echo "🔍 Analizando código..."
	@go vet ./...

lint: ## Ejecuta golangci-lint (requiere instalación)
	@echo "🔍 Ejecutando linter..."
	@golangci-lint run

deps: ## Descarga las dependencias
	@echo "📦 Descargando dependencias..."
	@go mod download
	@go mod tidy

clean: ## Limpia archivos generados
	@echo "🧹 Limpiando archivos generados..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@go clean

dev: ## Modo desarrollo (requiere air para hot reload)
	@echo "🔥 Iniciando en modo desarrollo..."
	@air

install-tools: ## Instala herramientas de desarrollo
	@echo "🛠️  Instalando herramientas de desarrollo..."
	@go install github.com/cosmtrek/air@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Comandos de base de datos (para futuras implementaciones)
db-migrate: ## Ejecuta migraciones (cuando se implemente)
	@echo "🗃️  Migraciones de base de datos no implementadas aún"

db-seed: ## Ejecuta seeders (cuando se implemente)
	@echo "🌱 Seeders de base de datos no implementados aún"
`,
		"config.go.tmpl": `package config

import (
	"os"
	"strings"
)

// Config contiene toda la configuración de la aplicación
type Config struct {
	Port        string
	Environment string
	LogLevel    string
	
	// CORS
	CorsAllowedOrigins []string
	
	// Base de datos (para futuras implementaciones)
	DatabaseURL string
	
	// JWT (para futuras implementaciones)
	JWTSecret string
}

// Load carga la configuración desde variables de entorno
func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		
		// CORS
		CorsAllowedOrigins: parseCorsOrigins(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:8080")),
		
		// Base de datos
		DatabaseURL: getEnv("DATABASE_URL", ""),
		
		// JWT
		JWTSecret: getEnv("JWT_SECRET", "your-default-secret-change-in-production"),
	}
}

// getEnv obtiene una variable de entorno o retorna un valor por defecto
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// parseCorsOrigins parsea la lista de orígenes CORS desde una cadena separada por comas
func parseCorsOrigins(origins string) []string {
	if origins == "" {
		return []string{}
	}
	
	result := make([]string, 0)
	for _, origin := range strings.Split(origins, ",") {
		trimmed := strings.TrimSpace(origin)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	
	return result
}

// IsDevelopment retorna true si el entorno es desarrollo
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

// IsProduction retorna true si el entorno es producción
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}
`,
		"server.go.tmpl": `package server

import (
	"context"
	"net/http"
	"time"

	"{{.ModuleName}}/internal/app/handlers"
	"{{.ModuleName}}/internal/app/middleware"
	"{{.ModuleName}}/internal/app/repositories"
	"{{.ModuleName}}/internal/app/services"
	"{{.ModuleName}}/internal/config"

	"github.com/gorilla/mux"
)

// Server representa el servidor HTTP
type Server struct {
	config     *config.Config
	router     *mux.Router
	httpServer *http.Server
}

// New crea una nueva instancia del servidor con todas las dependencias inyectadas
func New(cfg *config.Config) *Server {
	// Crear repositorios
	userRepo := repositories.NewUserRepository()
	
	// Crear servicios (inyectando repositorios)
	userService := services.NewUserService(userRepo)
	
	// Crear handlers (inyectando servicios)
	healthHandler := handlers.NewHealthHandler()
	userHandler := handlers.NewUserHandler(userService)
	
	// Crear router
	router := mux.NewRouter()
	
	// Configurar middleware global
	corsMiddleware := middleware.NewCORSMiddleware(cfg.CorsAllowedOrigins)
	router.Use(corsMiddleware)
	
	// Registrar rutas
	registerRoutes(router, healthHandler, userHandler)
	
	// Configurar servidor HTTP
	httpServer := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	return &Server{
		config:     cfg,
		router:     router,
		httpServer: httpServer,
	}
}

// Start inicia el servidor HTTP
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown apaga el servidor de forma elegante
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
`,
		"routes.go.tmpl": `package server

import (
	"{{.ModuleName}}/internal/app/handlers"
	"github.com/gorilla/mux"
)

// registerRoutes registra todas las rutas de la aplicación
func registerRoutes(
	router *mux.Router,
	healthHandler *handlers.HealthHandler,
	userHandler *handlers.UserHandler,
) {
	// API v1
	api := router.PathPrefix("/api/v1").Subrouter()
	
	// Rutas de salud
	api.HandleFunc("/health", healthHandler.Health).Methods("GET")
	api.HandleFunc("/health/ready", healthHandler.Ready).Methods("GET")
	
	// Rutas de usuarios
	users := api.PathPrefix("/users").Subrouter()
	users.HandleFunc("", userHandler.GetUsers).Methods("GET")
	users.HandleFunc("", userHandler.CreateUser).Methods("POST")
	users.HandleFunc("/{id}", userHandler.GetUser).Methods("GET")
	users.HandleFunc("/{id}", userHandler.UpdateUser).Methods("PUT")
	users.HandleFunc("/{id}", userHandler.DeleteUser).Methods("DELETE")
	
	// Ruta raíz
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(` + "`" + `{
			"message": "¡Bienvenido a {{.Name}}!",
			"status": "success",
			"version": "v1.0.0",
			"generated_with": "Loom",
			"endpoints": {
				"health": "/api/v1/health",
				"users": "/api/v1/users",
				"docs": "/docs/API.md"
			}
		}` + "`" + `))
	}).Methods("GET")
}
`,
		"health_handler.go.tmpl": `package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthHandler maneja las rutas de salud del sistema
type HealthHandler struct {
	startTime time.Time
}

// NewHealthHandler crea una nueva instancia de HealthHandler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{
		startTime: time.Now(),
	}
}

// HealthResponse representa la respuesta de salud
type HealthResponse struct {
	Status    string    ` + "`" + `json:"status"` + "`" + `
	Timestamp time.Time ` + "`" + `json:"timestamp"` + "`" + `
	Service   string    ` + "`" + `json:"service"` + "`" + `
	Version   string    ` + "`" + `json:"version"` + "`" + `
	Uptime    string    ` + "`" + `json:"uptime"` + "`" + `
}

// Health retorna el estado de salud básico del servicio
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "{{.Name}}",
		Version:   "v1.0.0",
		Uptime:    time.Since(h.startTime).String(),
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Ready retorna el estado de preparación del servicio
func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	// Aquí puedes agregar verificaciones adicionales como:
	// - Conectividad a la base de datos
	// - Servicios externos
	// - Dependencias críticas
	
	response := map[string]interface{}{
		"status": "ready",
		"timestamp": time.Now(),
		"checks": map[string]string{
			"service": "ok",
			// "database": "ok", // Cuando se implemente
			// "cache": "ok",    // Cuando se implemente
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
`,
		"user_handler.go.tmpl": `package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"{{.ModuleName}}/internal/app/dtos"
	"{{.ModuleName}}/internal/app/services"

	"github.com/gorilla/mux"
)

// UserHandler maneja las rutas relacionadas con usuarios
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler crea una nueva instancia de UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUsers obtiene todos los usuarios
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		http.Error(w, "Error obteniendo usuarios", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":    users,
		"count":   len(users),
		"status":  "success",
		"message": "Usuarios obtenidos exitosamente",
	})
}

// GetUser obtiene un usuario por ID
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error obteniendo usuario", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":    user,
		"status":  "success",
		"message": "Usuario obtenido exitosamente",
	})
}

// CreateUser crea un nuevo usuario
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserDTO dtos.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&createUserDTO); err != nil {
		http.Error(w, "Datos de entrada inválidos", http.StatusBadRequest)
		return
	}

	// Validar DTO
	if err := createUserDTO.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userService.CreateUser(createUserDTO)
	if err != nil {
		http.Error(w, "Error creando usuario", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":    user,
		"status":  "success",
		"message": "Usuario creado exitosamente",
	})
}

// UpdateUser actualiza un usuario existente
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	var updateUserDTO dtos.UpdateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&updateUserDTO); err != nil {
		http.Error(w, "Datos de entrada inválidos", http.StatusBadRequest)
		return
	}

	user, err := h.userService.UpdateUser(id, updateUserDTO)
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error actualizando usuario", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":    user,
		"status":  "success",
		"message": "Usuario actualizado exitosamente",
	})
}

// DeleteUser elimina un usuario
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	if err := h.userService.DeleteUser(id); err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error eliminando usuario", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Usuario eliminado exitosamente",
	})
}
`,
		"user_service.go.tmpl": `package services

import (
	"fmt"

	"{{.ModuleName}}/internal/app/dtos"
	"{{.ModuleName}}/internal/app/models"
	"{{.ModuleName}}/internal/app/repositories"
)

// UserService contiene la lógica de negocio para usuarios
type UserService struct {
	userRepo *repositories.UserRepository
}

// NewUserService crea una nueva instancia de UserService
func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// GetAllUsers obtiene todos los usuarios
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAll()
}

// GetUserByID obtiene un usuario por su ID
func (s *UserService) GetUserByID(id int) (*models.User, error) {
	user := s.userRepo.GetByID(id)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

// CreateUser crea un nuevo usuario
func (s *UserService) CreateUser(dto dtos.CreateUserDTO) (*models.User, error) {
	// Aquí puedes agregar lógica de negocio adicional:
	// - Validar que el email no exista
	// - Hashear contraseña
	// - Enviar email de bienvenida
	// - etc.

	user := &models.User{
		Name:  dto.Name,
		Email: dto.Email,
		Age:   dto.Age,
	}

	createdUser := s.userRepo.Create(user)
	return createdUser, nil
}

// UpdateUser actualiza un usuario existente
func (s *UserService) UpdateUser(id int, dto dtos.UpdateUserDTO) (*models.User, error) {
	existingUser := s.userRepo.GetByID(id)
	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	// Actualizar campos si se proporcionan
	if dto.Name != nil {
		existingUser.Name = *dto.Name
	}
	if dto.Email != nil {
		existingUser.Email = *dto.Email
	}
	if dto.Age != nil {
		existingUser.Age = *dto.Age
	}

	updatedUser := s.userRepo.Update(existingUser)
	return updatedUser, nil
}

// DeleteUser elimina un usuario
func (s *UserService) DeleteUser(id int) error {
	existingUser := s.userRepo.GetByID(id)
	if existingUser == nil {
		return fmt.Errorf("user not found")
	}

	s.userRepo.Delete(id)
	return nil
}
`,
		"user_dto.go.tmpl": `package dtos

import (
	"fmt"
	"strings"
)

// CreateUserDTO representa los datos para crear un usuario
type CreateUserDTO struct {
	Name  string ` + "`" + `json:"name"` + "`" + `
	Email string ` + "`" + `json:"email"` + "`" + `
	Age   int    ` + "`" + `json:"age"` + "`" + `
}

// Validate valida los datos del DTO
func (dto CreateUserDTO) Validate() error {
	if strings.TrimSpace(dto.Name) == "" {
		return fmt.Errorf("el nombre es requerido")
	}

	if strings.TrimSpace(dto.Email) == "" {
		return fmt.Errorf("el email es requerido")
	}

	if dto.Age < 0 || dto.Age > 150 {
		return fmt.Errorf("la edad debe estar entre 0 y 150 años")
	}

	// Validación básica de email
	if !strings.Contains(dto.Email, "@") {
		return fmt.Errorf("el email debe tener un formato válido")
	}

	return nil
}

// UpdateUserDTO representa los datos para actualizar un usuario
type UpdateUserDTO struct {
	Name  *string ` + "`" + `json:"name,omitempty"` + "`" + `
	Email *string ` + "`" + `json:"email,omitempty"` + "`" + `
	Age   *int    ` + "`" + `json:"age,omitempty"` + "`" + `
}

// UserResponseDTO representa la respuesta de un usuario (sin datos sensibles)
type UserResponseDTO struct {
	ID    int    ` + "`" + `json:"id"` + "`" + `
	Name  string ` + "`" + `json:"name"` + "`" + `
	Email string ` + "`" + `json:"email"` + "`" + `
	Age   int    ` + "`" + `json:"age"` + "`" + `
}
`,
		"user_model.go.tmpl": `package models

import "time"

// User representa un usuario en el sistema
type User struct {
	ID        int       ` + "`" + `json:"id"` + "`" + `
	Name      string    ` + "`" + `json:"name"` + "`" + `
	Email     string    ` + "`" + `json:"email"` + "`" + `
	Age       int       ` + "`" + `json:"age"` + "`" + `
	CreatedAt time.Time ` + "`" + `json:"created_at"` + "`" + `
	UpdatedAt time.Time ` + "`" + `json:"updated_at"` + "`" + `
}

// NewUser crea una nueva instancia de User con timestamps
func NewUser(name, email string, age int) *User {
	now := time.Now()
	return &User{
		Name:      name,
		Email:     email,
		Age:       age,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// UpdateTimestamp actualiza el timestamp de UpdatedAt
func (u *User) UpdateTimestamp() {
	u.UpdatedAt = time.Now()
}
`,
		"user_repository.go.tmpl": `package repositories

import (
	"sync"
	"time"

	"{{.ModuleName}}/internal/app/models"
)

// UserRepository maneja la persistencia de usuarios
// En este ejemplo usamos una implementación en memoria
// En producción, esto sería reemplazado por una base de datos real
type UserRepository struct {
	users  map[int]*models.User
	nextID int
	mutex  sync.RWMutex
}

// NewUserRepository crea una nueva instancia de UserRepository
func NewUserRepository() *UserRepository {
	repo := &UserRepository{
		users:  make(map[int]*models.User),
		nextID: 1,
	}
	
	// Agregar algunos usuarios de ejemplo
	repo.seedData()
	
	return repo
}

// seedData crea algunos usuarios de ejemplo
func (r *UserRepository) seedData() {
	users := []*models.User{
		{
			ID:        1,
			Name:      "Juan Pérez",
			Email:     "juan@example.com",
			Age:       25,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "María García",
			Email:     "maria@example.com",
			Age:       30,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Name:      "Carlos López",
			Email:     "carlos@example.com",
			Age:       28,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	
	for _, user := range users {
		r.users[user.ID] = user
		if user.ID >= r.nextID {
			r.nextID = user.ID + 1
		}
	}
}

// GetAll obtiene todos los usuarios
func (r *UserRepository) GetAll() ([]models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	users := make([]models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, *user)
	}
	
	return users, nil
}

// GetByID obtiene un usuario por su ID
func (r *UserRepository) GetByID(id int) *models.User {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	if user, exists := r.users[id]; exists {
		// Retornar una copia para evitar modificaciones accidentales
		userCopy := *user
		return &userCopy
	}
	
	return nil
}

// Create crea un nuevo usuario
func (r *UserRepository) Create(user *models.User) *models.User {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	// Asignar ID y timestamps
	user.ID = r.nextID
	r.nextID++
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	
	// Guardar el usuario
	r.users[user.ID] = user
	
	// Retornar una copia
	userCopy := *user
	return &userCopy
}

// Update actualiza un usuario existente
func (r *UserRepository) Update(user *models.User) *models.User {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	if _, exists := r.users[user.ID]; exists {
		user.UpdatedAt = time.Now()
		r.users[user.ID] = user
		
		// Retornar una copia
		userCopy := *user
		return &userCopy
	}
	
	return nil
}

// Delete elimina un usuario por su ID
func (r *UserRepository) Delete(id int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	delete(r.users, id)
}

// GetByEmail obtiene un usuario por su email
func (r *UserRepository) GetByEmail(email string) *models.User {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	for _, user := range r.users {
		if user.Email == email {
			// Retornar una copia
			userCopy := *user
			return &userCopy
		}
	}
	
	return nil
}
`,
	}

	// Agregar plantillas adicionales
	additionalTemplates := getAdditionalTemplates()
	for key, value := range additionalTemplates {
		templates[key] = value
	}

	return templates
}
