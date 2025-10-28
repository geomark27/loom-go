package addon

import (
	"fmt"
	"os"
	"strings"
)

// Addon representa un componente que puede ser añadido al proyecto
type Addon interface {
	// Name retorna el nombre del addon
	Name() string

	// Description retorna la descripción del addon
	Description() string

	// IsInstalled verifica si el addon ya está instalado
	IsInstalled() (bool, error)

	// CanInstall verifica si el addon puede ser instalado (dependencias, etc.)
	CanInstall() (bool, string, error)

	// Install instala el addon
	Install(force bool) error

	// GetConflicts retorna addons que pueden conflictuar
	GetConflicts() []string
}

// AddonManager gestiona los addons disponibles
type AddonManager struct {
	projectRoot  string
	architecture string // "layered" o "modular"
	addons       map[string]Addon
}

// NewAddonManager crea un nuevo gestor de addons
func NewAddonManager(projectRoot, architecture string) *AddonManager {
	am := &AddonManager{
		projectRoot:  projectRoot,
		architecture: architecture,
		addons:       make(map[string]Addon),
	}

	// Registrar addons disponibles
	am.registerAddons()

	return am
}

// registerAddons registra todos los addons disponibles
func (am *AddonManager) registerAddons() {
	// Routers
	am.addons["gin"] = NewRouterAddon(am.projectRoot, am.architecture, "gin")
	am.addons["chi"] = NewRouterAddon(am.projectRoot, am.architecture, "chi")
	am.addons["echo"] = NewRouterAddon(am.projectRoot, am.architecture, "echo")

	// ORMs
	am.addons["gorm"] = NewORMAddon(am.projectRoot, am.architecture, "gorm")
	am.addons["sqlc"] = NewORMAddon(am.projectRoot, am.architecture, "sqlc")

	// Databases
	am.addons["postgres"] = NewDatabaseAddon(am.projectRoot, am.architecture, "postgres")
	am.addons["mysql"] = NewDatabaseAddon(am.projectRoot, am.architecture, "mysql")
	am.addons["mongodb"] = NewDatabaseAddon(am.projectRoot, am.architecture, "mongodb")
	am.addons["redis"] = NewDatabaseAddon(am.projectRoot, am.architecture, "redis")

	// Auth
	am.addons["jwt"] = NewAuthAddon(am.projectRoot, am.architecture, "jwt")
	am.addons["oauth2"] = NewAuthAddon(am.projectRoot, am.architecture, "oauth2")

	// Infrastructure
	am.addons["docker"] = NewDockerAddon(am.projectRoot, am.architecture)
}

// GetAddon retorna un addon por nombre
func (am *AddonManager) GetAddon(name string) (Addon, error) {
	addon, exists := am.addons[name]
	if !exists {
		return nil, fmt.Errorf("addon '%s' no encontrado", name)
	}
	return addon, nil
}

// ListAddons retorna todos los addons disponibles por categoría
func (am *AddonManager) ListAddons() map[string][]string {
	return map[string][]string{
		"routers":        {"gin", "chi", "echo"},
		"orms":           {"gorm", "sqlc"},
		"databases":      {"postgres", "mysql", "mongodb", "redis"},
		"authentication": {"jwt", "oauth2"},
		"infrastructure": {"docker"},
	}
}

// InstallAddon instala un addon
func (am *AddonManager) InstallAddon(name string, force bool) error {
	addon, err := am.GetAddon(name)
	if err != nil {
		return err
	}

	// Verificar si ya está instalado
	installed, err := addon.IsInstalled()
	if err != nil {
		return fmt.Errorf("error al verificar instalación: %w", err)
	}

	if installed && !force {
		return fmt.Errorf("%s ya está instalado. Usa --force para reinstalar", addon.Name())
	}

	// Verificar si puede ser instalado
	canInstall, reason, err := addon.CanInstall()
	if err != nil {
		return fmt.Errorf("error al verificar compatibilidad: %w", err)
	}

	if !canInstall {
		return fmt.Errorf("no se puede instalar %s: %s", addon.Name(), reason)
	}

	// Verificar conflictos
	conflicts := addon.GetConflicts()
	for _, conflictName := range conflicts {
		conflictAddon, _ := am.GetAddon(conflictName)
		if conflictAddon != nil {
			if conflictInstalled, _ := conflictAddon.IsInstalled(); conflictInstalled {
				if !force {
					return fmt.Errorf("conflicto detectado: %s está instalado. Usa --force para reemplazar", conflictName)
				}
				fmt.Printf("⚠️  Reemplazando %s con %s...\n", conflictName, name)
			}
		}
	}

	// Instalar
	fmt.Printf("📦 Instalando %s...\n", addon.Name())
	if err := addon.Install(force); err != nil {
		return fmt.Errorf("error al instalar %s: %w", addon.Name(), err)
	}

	fmt.Printf("✅ %s instalado exitosamente!\n", addon.Name())
	return nil
}

// Helper functions

// FileExists verifica si un archivo existe
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ReadFile lee el contenido de un archivo
func ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// WriteFile escribe contenido en un archivo
func WriteFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// HasImport verifica si un archivo Go tiene un import específico
func HasImport(filePath, importPath string) bool {
	content, err := ReadFile(filePath)
	if err != nil {
		return false
	}
	return strings.Contains(content, importPath)
}

// AddImport añade un import a un archivo Go (simplificado)
func AddImport(filePath, importPath string) error {
	content, err := ReadFile(filePath)
	if err != nil {
		return err
	}

	// Si ya tiene el import, no hacer nada
	if strings.Contains(content, importPath) {
		return nil
	}

	// Buscar el bloque import
	lines := strings.Split(content, "\n")
	newLines := []string{}
	importAdded := false

	for i, line := range lines {
		newLines = append(newLines, line)

		// Si encontramos import ( y no hemos añadido el import
		if strings.Contains(line, "import (") && !importAdded {
			// Añadir el nuevo import después de esta línea
			newLines = append(newLines, fmt.Sprintf("\t\"%s\"", importPath))
			importAdded = true
		}

		// Si no hay bloque import, crear uno
		if i == 0 && strings.HasPrefix(line, "package ") && !importAdded {
			// Buscar si hay import simple
			for j := i + 1; j < len(lines); j++ {
				if strings.HasPrefix(strings.TrimSpace(lines[j]), "import ") {
					importAdded = true
					break
				}
				if strings.TrimSpace(lines[j]) != "" {
					// Llegamos al código, insertar import aquí
					newLines = append(newLines, "")
					newLines = append(newLines, "import (")
					newLines = append(newLines, fmt.Sprintf("\t\"%s\"", importPath))
					newLines = append(newLines, ")")
					importAdded = true
					break
				}
			}
		}
	}

	return WriteFile(filePath, strings.Join(newLines, "\n"))
}

// UpdateGoMod actualiza go.mod con una nueva dependencia
func UpdateGoMod(module, version string) error {
	// Leer go.mod actual
	content, err := ReadFile("go.mod")
	if err != nil {
		return err
	}

	// Si ya tiene la dependencia, no hacer nada
	if strings.Contains(content, module) {
		return nil
	}

	// Añadir al bloque require
	lines := strings.Split(content, "\n")
	newLines := []string{}
	requireAdded := false

	for _, line := range lines {
		newLines = append(newLines, line)

		if strings.Contains(line, "require (") && !requireAdded {
			newLines = append(newLines, fmt.Sprintf("\t%s %s", module, version))
			requireAdded = true
		}
	}

	// Si no hay bloque require, añadir al final
	if !requireAdded {
		newLines = append(newLines, "")
		newLines = append(newLines, "require (")
		newLines = append(newLines, fmt.Sprintf("\t%s %s", module, version))
		newLines = append(newLines, ")")
	}

	return WriteFile("go.mod", strings.Join(newLines, "\n"))
}

// UpdateEnvExample actualiza .env.example con nuevas variables
func UpdateEnvExample(variables map[string]string, section string) error {
	content := ""

	// Leer existente si existe
	if FileExists(".env.example") {
		existingContent, err := ReadFile(".env.example")
		if err != nil {
			return err
		}
		content = existingContent
	}

	// Añadir sección si no existe
	sectionHeader := fmt.Sprintf("\n# %s\n", section)
	if !strings.Contains(content, sectionHeader) {
		content += sectionHeader
		for key, value := range variables {
			content += fmt.Sprintf("%s=%s\n", key, value)
		}
	}

	return WriteFile(".env.example", content)
}
