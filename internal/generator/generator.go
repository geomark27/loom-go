package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// ProjectConfig contiene la configuración para generar un proyecto
type ProjectConfig struct {
	Name         string // Nombre del proyecto
	Path         string // Ruta donde se creará el proyecto
	ModuleName   string // Nombre del módulo Go
	Description  string // Descripción del proyecto
	UseHelpers   bool   // Si true, incluye los helpers de Loom en el proyecto
	IsModular    bool   // Si true, usa arquitectura modular en lugar de por capas
	Architecture string // "layered" o "modular"
}

// Generator es responsable de generar proyectos
type Generator struct {
	templates map[string]string
}

// New crea una nueva instancia del generator
func New() *Generator {
	return &Generator{
		templates: getTemplates(),
	}
}

// GenerateProject genera un nuevo proyecto basado en la configuración
func (g *Generator) GenerateProject(config *ProjectConfig) error {
	// Crear el directorio raíz del proyecto
	if err := os.MkdirAll(config.Path, 0755); err != nil {
		return fmt.Errorf("error creando directorio del proyecto: %w", err)
	}

	// Crear estructura de directorios según arquitectura
	dirs := g.getDirectories(config)

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creando directorio %s: %w", dir, err)
		}
	}

	// Generar archivos desde plantillas según arquitectura
	files := g.getFileMapping(config)

	for filePath, templateName := range files {
		if err := g.generateFile(filePath, templateName, config); err != nil {
			return fmt.Errorf("error generando archivo %s: %w", filePath, err)
		}
	}

	return nil
}

// getDirectories retorna los directorios a crear según la arquitectura
func (g *Generator) getDirectories(config *ProjectConfig) []string {
	if config.IsModular {
		return g.getModularDirectories(config)
	}
	return g.getLayeredDirectories(config)
}

// getLayeredDirectories retorna los directorios para arquitectura por capas
func (g *Generator) getLayeredDirectories(config *ProjectConfig) []string {
	return []string{
		// cmd
		filepath.Join(config.Path, "cmd", config.Name),

		// internal/app - Lógica de negocio
		filepath.Join(config.Path, "internal", "app", "handlers"),
		filepath.Join(config.Path, "internal", "app", "services"),
		filepath.Join(config.Path, "internal", "app", "repositories"),
		filepath.Join(config.Path, "internal", "app", "models"),
		filepath.Join(config.Path, "internal", "app", "dtos"),

		// internal/platform - Infraestructura técnica
		filepath.Join(config.Path, "internal", "platform", "config"),
		filepath.Join(config.Path, "internal", "platform", "server"),
		filepath.Join(config.Path, "internal", "platform", "database"),

		// internal/shared - Utilidades transversales
		filepath.Join(config.Path, "internal", "shared", "middleware"),
		filepath.Join(config.Path, "internal", "shared", "response"),
		filepath.Join(config.Path, "internal", "shared", "errors"),

		// otros
		filepath.Join(config.Path, "pkg"),
		filepath.Join(config.Path, "docs"),
	}
}

// getModularDirectories retorna los directorios para arquitectura modular
func (g *Generator) getModularDirectories(config *ProjectConfig) []string {
	return []string{
		// cmd
		filepath.Join(config.Path, "cmd", config.Name),

		// internal/modules - Módulos de dominio
		filepath.Join(config.Path, "internal", "modules", "users"),

		// internal/platform - Infraestructura técnica
		filepath.Join(config.Path, "internal", "platform", "config"),
		filepath.Join(config.Path, "internal", "platform", "server"),
		filepath.Join(config.Path, "internal", "platform", "events"),
		filepath.Join(config.Path, "internal", "platform", "database"),

		// internal/shared - Utilidades transversales
		filepath.Join(config.Path, "internal", "shared", "middleware"),
		filepath.Join(config.Path, "internal", "shared", "response"),
		filepath.Join(config.Path, "internal", "shared", "errors"),

		// otros
		filepath.Join(config.Path, "pkg"),
		filepath.Join(config.Path, "docs"),
	}
}

// getFileMapping retorna el mapeo de archivos según la arquitectura
func (g *Generator) getFileMapping(config *ProjectConfig) map[string]string {
	if config.IsModular {
		return g.getModularFiles(config)
	}
	return g.getLayeredFiles(config)
}

// getLayeredFiles retorna el mapeo de archivos para arquitectura por capas
func (g *Generator) getLayeredFiles(config *ProjectConfig) map[string]string {
	return map[string]string{
		// Archivos raíz del proyecto
		filepath.Join(config.Path, "go.mod"):       "go.mod.tmpl",
		filepath.Join(config.Path, "README.md"):    "README.md.tmpl",
		filepath.Join(config.Path, ".gitignore"):   ".gitignore.tmpl",
		filepath.Join(config.Path, ".env.example"): ".env.example.tmpl",
		filepath.Join(config.Path, "Makefile"):     "Makefile.tmpl",

		// cmd
		filepath.Join(config.Path, "cmd", config.Name, "main.go"): "main.go.tmpl",

		// internal/platform
		filepath.Join(config.Path, "internal", "platform", "config", "config.go"): "config.go.tmpl",
		filepath.Join(config.Path, "internal", "platform", "server", "server.go"): "layered/server.go.tmpl",
		filepath.Join(config.Path, "internal", "platform", "server", "routes.go"): "layered/routes.go.tmpl",

		// internal/app
		filepath.Join(config.Path, "internal", "app", "handlers", "health_handler.go"):      "layered/health_handler.go.tmpl",
		filepath.Join(config.Path, "internal", "app", "handlers", "user_handler.go"):        "layered/user_handler.go.tmpl",
		filepath.Join(config.Path, "internal", "app", "services", "user_service.go"):        "layered/user_service.go.tmpl",
		filepath.Join(config.Path, "internal", "app", "repositories", "user_repository.go"): "layered/user_repository.go.tmpl",
		filepath.Join(config.Path, "internal", "app", "models", "user.go"):                  "user_model.go.tmpl",
		filepath.Join(config.Path, "internal", "app", "dtos", "user_dto.go"):                "user_dto.go.tmpl",

		// internal/shared
		filepath.Join(config.Path, "internal", "shared", "middleware", "cors_middleware.go"): "cors_middleware.go.tmpl",

		// docs
		filepath.Join(config.Path, "docs", "API.md"): "api_docs.tmpl",
	}
}

// getModularFiles retorna el mapeo de archivos para arquitectura modular
func (g *Generator) getModularFiles(config *ProjectConfig) map[string]string {
	return map[string]string{
		// Archivos raíz del proyecto
		filepath.Join(config.Path, "go.mod"):       "go.mod.tmpl",
		filepath.Join(config.Path, "README.md"):    "README.md.tmpl",
		filepath.Join(config.Path, ".gitignore"):   ".gitignore.tmpl",
		filepath.Join(config.Path, ".env.example"): ".env.example.tmpl",
		filepath.Join(config.Path, "Makefile"):     "Makefile.tmpl",

		// cmd
		filepath.Join(config.Path, "cmd", config.Name, "main.go"): "modular/main.go.tmpl",

		// internal/platform
		filepath.Join(config.Path, "internal", "platform", "config", "config.go"):    "config.go.tmpl",
		filepath.Join(config.Path, "internal", "platform", "server", "server.go"):    "modular/server.go.tmpl",
		filepath.Join(config.Path, "internal", "platform", "server", "router.go"):    "modular/router.go.tmpl",
		filepath.Join(config.Path, "internal", "platform", "events", "event_bus.go"): "modular/event_bus.go.tmpl",

		// internal/modules/users
		filepath.Join(config.Path, "internal", "modules", "users", "ports.go"):      "modular/ports.go.tmpl",
		filepath.Join(config.Path, "internal", "modules", "users", "service.go"):    "modular/service.go.tmpl",
		filepath.Join(config.Path, "internal", "modules", "users", "repository.go"): "modular/repository.go.tmpl",
		filepath.Join(config.Path, "internal", "modules", "users", "handler.go"):    "modular/handler.go.tmpl",
		filepath.Join(config.Path, "internal", "modules", "users", "model.go"):      "modular/model.go.tmpl",
		filepath.Join(config.Path, "internal", "modules", "users", "dto.go"):        "modular/dto.go.tmpl",
		filepath.Join(config.Path, "internal", "modules", "users", "module.go"):     "modular/module.go.tmpl",
		filepath.Join(config.Path, "internal", "modules", "users", "errors.go"):     "modular/errors.go.tmpl",

		// internal/shared
		filepath.Join(config.Path, "internal", "shared", "middleware", "cors_middleware.go"): "cors_middleware.go.tmpl",

		// docs
		filepath.Join(config.Path, "docs", "API.md"): "modular/api_docs.tmpl",
	}
}

// generateFile genera un archivo específico usando una plantilla
func (g *Generator) generateFile(filePath, templateName string, config *ProjectConfig) error {
	templateContent, exists := g.templates[templateName]
	if !exists {
		return fmt.Errorf("plantilla %s no encontrada", templateName)
	}

	tmpl, err := template.New(templateName).Parse(templateContent)
	if err != nil {
		return fmt.Errorf("error parseando plantilla %s: %w", templateName, err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creando archivo %s: %w", filePath, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, config); err != nil {
		return fmt.Errorf("error ejecutando plantilla %s: %w", templateName, err)
	}

	return nil
}
