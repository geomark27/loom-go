package generator

// getAdditionalTemplates retorna las plantillas adicionales
func getAdditionalTemplates() map[string]string {
	return map[string]string{
		"cors_middleware.go.tmpl": `package middleware

import (
	"net/http"
	"strings"
)

// CORSMiddleware maneja las políticas de CORS
func NewCORSMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			
			// Verificar si el origen está permitido
			if isOriginAllowed(origin, allowedOrigins) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			
			// Manejar preflight requests
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	}
}

// isOriginAllowed verifica si un origen está en la lista de orígenes permitidos
func isOriginAllowed(origin string, allowedOrigins []string) bool {
	if origin == "" {
		return false
	}
	
	for _, allowed := range allowedOrigins {
		if strings.EqualFold(origin, allowed) || allowed == "*" {
			return true
		}
	}
	
	return false
}
`,
		"api_docs.tmpl": `# {{.Name}} - Documentación de API

Esta documentación describe los endpoints disponibles en la API de **{{.Name}}**.

## Información General

- **URL Base**: ` + "`http://localhost:8080/api/v1`" + `
- **Formato de Respuesta**: JSON
- **Autenticación**: No requerida (por ahora)

## Endpoints

### Salud del Sistema

#### GET /health
Verifica el estado básico del servicio.

**Respuesta:**
` + "```json" + `
{
  "status": "healthy",
  "timestamp": "2025-08-26T10:30:00Z",
  "service": "{{.Name}}",
  "version": "v1.0.0",
  "uptime": "2h30m15s"
}
` + "```" + `

### Usuarios

#### GET /users
Obtiene todos los usuarios.

**Respuesta:**
` + "```json" + `
{
  "data": [
    {
      "id": 1,
      "name": "Juan Pérez",
      "email": "juan@example.com",
      "age": 25
    }
  ],
  "count": 1,
  "status": "success",
  "message": "Usuarios obtenidos exitosamente"
}
` + "```" + `

#### POST /users
Crea un nuevo usuario.

**Cuerpo de la petición:**
` + "```json" + `
{
  "name": "Juan Pérez",
  "email": "juan@example.com",
  "age": 25
}
` + "```" + `

## Ejemplos de Uso

### Crear un usuario
` + "```bash" + `
curl -X POST http://localhost:8080/api/v1/users \\
  -H "Content-Type: application/json" \\
  -d '{
    "name": "Ana García",
    "email": "ana@example.com",
    "age": 28
  }'
` + "```" + `

### Obtener todos los usuarios
` + "```bash" + `
curl http://localhost:8080/api/v1/users
` + "```" + `
`,
		"routes.go.tmpl": `package server

import (
	"net/http"
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
	}
}
