package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// ProjectConfig contains the configuration for generating a project
type ProjectConfig struct {
	Name         string // Project name
	Path         string // Path where the project will be created
	ModuleName   string // Go module name
	Description  string // Project description
	UseHelpers   bool   // If true, includes Loom helpers in the project
	IsModular    bool   // If true, uses modular architecture instead of layered
	Architecture string // "layered" or "modular"
	LoomVersion  string // Loom version to use in go.mod (injected automatically)
}

// Generator is responsible for generating projects
type Generator struct {
	templates map[string]string
}

// New creates a new generator instance
func New() *Generator {
	return &Generator{
		templates: getTemplates(),
	}
}

// GenerateProject generates a new project based on the configuration
func (g *Generator) GenerateProject(config *ProjectConfig) error {
	// Create the project root directory
	if err := os.MkdirAll(config.Path, 0755); err != nil {
		return fmt.Errorf("error creating project directory: %w", err)
	}

	// Create directory structure based on architecture
	dirs := g.getDirectories(config)

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %w", dir, err)
		}
	}

	// Generate files from templates based on architecture
	files := g.getFileMapping(config)

	for filePath, templateName := range files {
		if err := g.generateFile(filePath, templateName, config); err != nil {
			return fmt.Errorf("error generating file %s: %w", filePath, err)
		}
	}

	return nil
}

// getDirectories returns the directories to create based on the architecture
func (g *Generator) getDirectories(config *ProjectConfig) []string {
	if config.IsModular {
		return g.getModularDirectories(config)
	}
	return g.getLayeredDirectories(config)
}

// getLayeredDirectories returns the directories for layered architecture
func (g *Generator) getLayeredDirectories(config *ProjectConfig) []string {
	return []string{
		// cmd
		filepath.Join(config.Path, "cmd", config.Name),

		// internal/app - Business logic
		filepath.Join(config.Path, "internal", "app", "handlers"),
		filepath.Join(config.Path, "internal", "app", "services"),
		filepath.Join(config.Path, "internal", "app", "repositories"),
		filepath.Join(config.Path, "internal", "app", "models"),
		filepath.Join(config.Path, "internal", "app", "dtos"),

		// internal/platform - Technical infrastructure
		filepath.Join(config.Path, "internal", "platform", "config"),
		filepath.Join(config.Path, "internal", "platform", "server"),
		filepath.Join(config.Path, "internal", "platform", "database"),

		// internal/shared - Cross-cutting utilities
		filepath.Join(config.Path, "internal", "shared", "middleware"),
		filepath.Join(config.Path, "internal", "shared", "response"),
		filepath.Join(config.Path, "internal", "shared", "errors"),

		// other
		filepath.Join(config.Path, "pkg"),
		filepath.Join(config.Path, "docs"),
	}
}

// getModularDirectories returns the directories for modular architecture
func (g *Generator) getModularDirectories(config *ProjectConfig) []string {
	return []string{
		// cmd
		filepath.Join(config.Path, "cmd", config.Name),

		// internal/modules - Domain modules
		filepath.Join(config.Path, "internal", "modules", "users"),

		// internal/platform - Technical infrastructure
		filepath.Join(config.Path, "internal", "platform", "config"),
		filepath.Join(config.Path, "internal", "platform", "server"),
		filepath.Join(config.Path, "internal", "platform", "events"),
		filepath.Join(config.Path, "internal", "platform", "database"),

		// internal/shared - Cross-cutting utilities
		filepath.Join(config.Path, "internal", "shared", "middleware"),
		filepath.Join(config.Path, "internal", "shared", "response"),
		filepath.Join(config.Path, "internal", "shared", "errors"),

		// other
		filepath.Join(config.Path, "pkg"),
		filepath.Join(config.Path, "docs"),
	}
}

// getFileMapping returns the file mapping based on the architecture
func (g *Generator) getFileMapping(config *ProjectConfig) map[string]string {
	if config.IsModular {
		return g.getModularFiles(config)
	}
	return g.getLayeredFiles(config)
}

// getLayeredFiles returns the file mapping for layered architecture
func (g *Generator) getLayeredFiles(config *ProjectConfig) map[string]string {
	return map[string]string{
		// Project root files
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

// getModularFiles returns the file mapping for modular architecture
func (g *Generator) getModularFiles(config *ProjectConfig) map[string]string {
	return map[string]string{
		// Project root files
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

// generateFile generates a specific file using a template
func (g *Generator) generateFile(filePath, templateName string, config *ProjectConfig) error {
	templateContent, exists := g.templates[templateName]
	if !exists {
		return fmt.Errorf("template %s not found", templateName)
	}

	tmpl, err := template.New(templateName).Parse(templateContent)
	if err != nil {
		return fmt.Errorf("error parsing template %s: %w", templateName, err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", filePath, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, config); err != nil {
		return fmt.Errorf("error executing template %s: %w", templateName, err)
	}

	return nil
}
