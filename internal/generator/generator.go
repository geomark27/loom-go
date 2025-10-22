package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// ProjectConfig contiene la configuración para generar un proyecto
type ProjectConfig struct {
	Name        string // Nombre del proyecto
	Path        string // Ruta donde se creará el proyecto
	ModuleName  string // Nombre del módulo Go
	Description string // Descripción del proyecto
	UseHelpers  bool   // Si true, incluye los helpers de Loom en el proyecto
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

	// Crear estructura de directorios
	dirs := []string{
		filepath.Join(config.Path, "cmd", config.Name),
		filepath.Join(config.Path, "internal", "app", "handlers"),
		filepath.Join(config.Path, "internal", "app", "services"),
		filepath.Join(config.Path, "internal", "app", "dtos"),
		filepath.Join(config.Path, "internal", "app", "models"),
		filepath.Join(config.Path, "internal", "app", "repositories"),
		filepath.Join(config.Path, "internal", "app", "middleware"),
		filepath.Join(config.Path, "internal", "config"),
		filepath.Join(config.Path, "internal", "server"),
		filepath.Join(config.Path, "pkg"),
		filepath.Join(config.Path, "docs"),
		filepath.Join(config.Path, "scripts"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creando directorio %s: %w", dir, err)
		}
	}

	// Generar archivos desde plantillas
	files := map[string]string{
		filepath.Join(config.Path, "go.mod"):                                                "go.mod.tmpl",
		filepath.Join(config.Path, "README.md"):                                             "README.md.tmpl",
		filepath.Join(config.Path, "cmd", config.Name, "main.go"):                           "main.go.tmpl",
		filepath.Join(config.Path, ".gitignore"):                                            ".gitignore.tmpl",
		filepath.Join(config.Path, ".env.example"):                                          ".env.example.tmpl",
		filepath.Join(config.Path, "Makefile"):                                              "Makefile.tmpl",
		filepath.Join(config.Path, "internal", "config", "config.go"):                       "config.go.tmpl",
		filepath.Join(config.Path, "internal", "server", "server.go"):                       "server.go.tmpl",
		filepath.Join(config.Path, "internal", "server", "routes.go"):                       "routes.go.tmpl",
		filepath.Join(config.Path, "internal", "app", "handlers", "health_handler.go"):      "health_handler.go.tmpl",
		filepath.Join(config.Path, "internal", "app", "handlers", "user_handler.go"):        "user_handler.go.tmpl",
		filepath.Join(config.Path, "internal", "app", "services", "user_service.go"):        "user_service.go.tmpl",
		filepath.Join(config.Path, "internal", "app", "dtos", "user_dto.go"):                "user_dto.go.tmpl",
		filepath.Join(config.Path, "internal", "app", "models", "user.go"):                  "user_model.go.tmpl",
		filepath.Join(config.Path, "internal", "app", "repositories", "user_repository.go"): "user_repository.go.tmpl",
		filepath.Join(config.Path, "internal", "app", "middleware", "cors_middleware.go"):   "cors_middleware.go.tmpl",
		filepath.Join(config.Path, "docs", "API.md"):                                        "api_docs.tmpl",
	}

	for filePath, templateName := range files {
		if err := g.generateFile(filePath, templateName, config); err != nil {
			return fmt.Errorf("error generando archivo %s: %w", filePath, err)
		}
	}

	return nil
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
