package generator

import (
"embed"
"fmt"
)

// Embed all templates at compile time
//go:embed templates/*
//go:embed templates/**/*
var templatesFS embed.FS

// getTemplates retorna un mapa con todas las plantillas embebidas
func getTemplates() map[string]string {
templates := make(map[string]string)

// Lista de todos los templates
templateFiles := map[string]string{
// Project files
"go.mod.tmpl":        "templates/project/go.mod.tmpl",
"README.md.tmpl":     "templates/project/README.md.tmpl",
".gitignore.tmpl":    "templates/project/.gitignore.tmpl",
".env.example.tmpl":  "templates/project/.env.example.tmpl",
"main.go.tmpl":       "templates/project/main.go.tmpl",
"Makefile.tmpl":      "templates/project/Makefile.tmpl",

// Config
"config.go.tmpl": "templates/config/config.go.tmpl",

// Server
"server.go.tmpl": "templates/server/server.go.tmpl",
"routes.go.tmpl": "templates/server/routes.go.tmpl",

// Handlers
"health_handler.go.tmpl": "templates/handlers/health_handler.go.tmpl",
"user_handler.go.tmpl":   "templates/handlers/user_handler.go.tmpl",

// Services
"user_service.go.tmpl": "templates/services/user_service.go.tmpl",

// Models
"user_model.go.tmpl": "templates/models/user_model.go.tmpl",

// Repositories
"user_repository.go.tmpl": "templates/repositories/user_repository.go.tmpl",

// DTOs
"user_dto.go.tmpl": "templates/dtos/user_dto.go.tmpl",

// Middleware
"cors_middleware.go.tmpl": "templates/middleware/cors_middleware.go.tmpl",

// Docs
"api_docs.tmpl": "templates/docs/api_docs.tmpl",
}

// Cargar cada template
for key, path := range templateFiles {
content, err := templatesFS.ReadFile(path)
if err != nil {
panic(fmt.Sprintf("Error loading template %s: %v", path, err))
}
templates[key] = string(content)
}

return templates
}
