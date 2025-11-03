package addon

import (
	"fmt"
	"os"
	"strings"
)

// Addon represents a component that can be added to the project
type Addon interface {
	// Name returns the addon name
	Name() string

	// Description returns the addon description
	Description() string

	// IsInstalled checks if the addon is already installed
	IsInstalled() (bool, error)

	// CanInstall checks if the addon can be installed (dependencies, etc.)
	CanInstall() (bool, string, error)

	// Install installs the addon
	Install(force bool) error

	// GetConflicts returns addons that may conflict
	GetConflicts() []string
}

// AddonManager manages available addons
type AddonManager struct {
	projectRoot  string
	architecture string // "layered" or "modular"
	addons       map[string]Addon
}

// NewAddonManager creates a new addon manager
func NewAddonManager(projectRoot, architecture string) *AddonManager {
	am := &AddonManager{
		projectRoot:  projectRoot,
		architecture: architecture,
		addons:       make(map[string]Addon),
	}

	// Register available addons
	am.registerAddons()

	return am
}

// registerAddons registers all available addons
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

// GetAddon returns an addon by name
func (am *AddonManager) GetAddon(name string) (Addon, error) {
	addon, exists := am.addons[name]
	if !exists {
		return nil, fmt.Errorf("addon '%s' not found", name)
	}
	return addon, nil
}

// ListAddons returns all available addons by category
func (am *AddonManager) ListAddons() map[string][]string {
	return map[string][]string{
		"routers":        {"gin", "chi", "echo"},
		"orms":           {"gorm", "sqlc"},
		"databases":      {"postgres", "mysql", "mongodb", "redis"},
		"authentication": {"jwt", "oauth2"},
		"infrastructure": {"docker"},
	}
}

// InstallAddon installs an addon
func (am *AddonManager) InstallAddon(name string, force bool) error {
	addon, err := am.GetAddon(name)
	if err != nil {
		return err
	}

	// Check if already installed
	installed, err := addon.IsInstalled()
	if err != nil {
		return fmt.Errorf("error checking installation: %w", err)
	}

	if installed && !force {
		return fmt.Errorf("%s is already installed. Use --force to reinstall", addon.Name())
	}

	// Check if it can be installed
	canInstall, reason, err := addon.CanInstall()
	if err != nil {
		return fmt.Errorf("error checking compatibility: %w", err)
	}

	if !canInstall {
		return fmt.Errorf("cannot install %s: %s", addon.Name(), reason)
	}

	// Check conflicts
	conflicts := addon.GetConflicts()
	for _, conflictName := range conflicts {
		conflictAddon, _ := am.GetAddon(conflictName)
		if conflictAddon != nil {
			if conflictInstalled, _ := conflictAddon.IsInstalled(); conflictInstalled {
				if !force {
					return fmt.Errorf("conflict detected: %s is installed. Use --force to replace", conflictName)
				}
				fmt.Printf("⚠️  Replacing %s with %s...\n", conflictName, name)
			}
		}
	}

	// Install
	fmt.Printf("📦 Installing %s...\n", addon.Name())
	if err := addon.Install(force); err != nil {
		return fmt.Errorf("error installing %s: %w", addon.Name(), err)
	}

	fmt.Printf("✅ %s installed successfully!\n", addon.Name())
	return nil
}

// Helper functions

// FileExists checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ReadFile reads the content of a file
func ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// WriteFile writes content to a file
func WriteFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// HasImport checks if a Go file has a specific import
func HasImport(filePath, importPath string) bool {
	content, err := ReadFile(filePath)
	if err != nil {
		return false
	}
	return strings.Contains(content, importPath)
}

// AddImport adds an import to a Go file (simplified)
func AddImport(filePath, importPath string) error {
	content, err := ReadFile(filePath)
	if err != nil {
		return err
	}

	// If it already has the import, do nothing
	if strings.Contains(content, importPath) {
		return nil
	}

	// Search for the import block
	lines := strings.Split(content, "\n")
	newLines := []string{}
	importAdded := false

	for i, line := range lines {
		newLines = append(newLines, line)

		// If we find import ( and haven't added the import
		if strings.Contains(line, "import (") && !importAdded {
			// Add the new import after this line
			newLines = append(newLines, fmt.Sprintf("\t\"%s\"", importPath))
			importAdded = true
		}

		// If there's no import block, create one
		if i == 0 && strings.HasPrefix(line, "package ") && !importAdded {
			// Search if there's a simple import
			for j := i + 1; j < len(lines); j++ {
				if strings.HasPrefix(strings.TrimSpace(lines[j]), "import ") {
					importAdded = true
					break
				}
				if strings.TrimSpace(lines[j]) != "" {
					// We reached the code, insert import here
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

// UpdateGoMod updates go.mod with a new dependency
func UpdateGoMod(module, version string) error {
	// Read current go.mod
	content, err := ReadFile("go.mod")
	if err != nil {
		return err
	}

	// If it already has the dependency, do nothing
	if strings.Contains(content, module) {
		return nil
	}

	// Add to require block
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

	// If there's no require block, add at the end
	if !requireAdded {
		newLines = append(newLines, "")
		newLines = append(newLines, "require (")
		newLines = append(newLines, fmt.Sprintf("\t%s %s", module, version))
		newLines = append(newLines, ")")
	}

	return WriteFile("go.mod", strings.Join(newLines, "\n"))
}

// UpdateEnvExample updates .env.example with new variables
func UpdateEnvExample(variables map[string]string, section string) error {
	content := ""

	// Read existing if it exists
	if FileExists(".env.example") {
		existingContent, err := ReadFile(".env.example")
		if err != nil {
			return err
		}
		content = existingContent
	}

	// Add section if it doesn't exist
	sectionHeader := fmt.Sprintf("\n# %s\n", section)
	if !strings.Contains(content, sectionHeader) {
		content += sectionHeader
		for key, value := range variables {
			content += fmt.Sprintf("%s=%s\n", key, value)
		}
	}

	return WriteFile(".env.example", content)
}
