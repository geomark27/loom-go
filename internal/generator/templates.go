package generator

import (
	"embed"
	"fmt"
)

//go:embed all:templates
var templatesFS embed.FS

// getTemplates retorna un mapa con todas las plantillas embebidas
func getTemplates() map[string]string {
	templates := make(map[string]string)

	// Lista de todos los templates
	templateFiles := map[string]string{
		// ======================================
		// Project files (compartidos)
		// ======================================
		"go.mod.tmpl":       "templates/project/go.mod.tmpl",
		"README.md.tmpl":    "templates/project/README.md.tmpl",
		".gitignore.tmpl":   "templates/project/.gitignore.tmpl",
		".env.example.tmpl": "templates/project/.env.example.tmpl",
		"main.go.tmpl":      "templates/project/main.go.tmpl",
		"Makefile.tmpl":     "templates/project/Makefile.tmpl",

		// ======================================
		// Config (compartido)
		// ======================================
		"config.go.tmpl": "templates/config/config.go.tmpl",

		// ======================================
		// LAYERED Architecture Templates
		// ======================================
		"layered/server.go.tmpl":          "templates/layered/server.go.tmpl",
		"layered/routes.go.tmpl":          "templates/layered/routes.go.tmpl",
		"layered/health_handler.go.tmpl": "templates/layered/health_handler.go.tmpl",
		"layered/user_handler.go.tmpl":   "templates/layered/user_handler.go.tmpl",
		"layered/user_service.go.tmpl":   "templates/layered/user_service.go.tmpl",
		"layered/user_repository.go.tmpl": "templates/layered/user_repository.go.tmpl",

		// ======================================
		// MODULAR Architecture Templates
		// ======================================
		"modular/main.go.tmpl":      "templates/modular/main.go.tmpl",
		"modular/server.go.tmpl":    "templates/modular/server.go.tmpl",
		"modular/router.go.tmpl":    "templates/modular/router.go.tmpl",
		"modular/event_bus.go.tmpl": "templates/modular/event_bus.go.tmpl",
		"modular/ports.go.tmpl":     "templates/modular/ports.go.tmpl",
		"modular/service.go.tmpl":   "templates/modular/service.go.tmpl",
		"modular/repository.go.tmpl": "templates/modular/repository.go.tmpl",
		"modular/handler.go.tmpl":   "templates/modular/handler.go.tmpl",
		"modular/model.go.tmpl":     "templates/modular/model.go.tmpl",
		"modular/dto.go.tmpl":       "templates/modular/dto.go.tmpl",
		"modular/module.go.tmpl":    "templates/modular/module.go.tmpl",
		"modular/errors.go.tmpl":    "templates/modular/errors.go.tmpl",
		"modular/api_docs.tmpl":     "templates/modular/api_docs.tmpl",

		// ======================================
		// Shared Templates (usados por ambas arquitecturas)
		// ======================================
		"user_model.go.tmpl":      "templates/models/user_model.go.tmpl",
		"user_dto.go.tmpl":        "templates/dtos/user_dto.go.tmpl",
		"cors_middleware.go.tmpl": "templates/middleware/cors_middleware.go.tmpl",
		"api_docs.tmpl":           "templates/docs/api_docs.tmpl",
	}

	// Cargar cada template
	for key, path := range templateFiles {
		content, err := templatesFS.ReadFile(path)
		if err != nil {
			// Log el error pero no hacer panic - algunos templates pueden no existir
			// dependiendo de qué carpeta se use
			fmt.Printf("Warning: Template %s not found, skipping\n", path)
			continue
		}
		templates[key] = string(content)
	}

	return templates
}
