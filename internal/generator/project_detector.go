package generator

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// ProjectInfo contains information about the detected project
type ProjectInfo struct {
	Name         string
	Architecture string // "layered" or "modular"
	HasHelpers   bool
	RootPath     string
	ModuleName   string
}

// DetectProject detects the type of Loom project in the current directory
func DetectProject() (*ProjectInfo, error) {
	// Look for go.mod to confirm it's a Go project
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		return nil, fmt.Errorf("go.mod not found. Are you in a Go project?")
	}

	info := &ProjectInfo{
		RootPath: ".",
	}

	// Detect architecture
	if _, err := os.Stat("internal/modules"); err == nil {
		info.Architecture = "modular"
	} else if _, err := os.Stat("internal/app"); err == nil {
		info.Architecture = "layered"
	} else {
		return nil, fmt.Errorf("no valid Loom project detected (missing internal/modules or internal/app)")
	}

	// Detect if it has helpers
	info.HasHelpers = hasHelpersImport()

	// Read project name from go.mod
	info.Name = getProjectNameFromGoMod()
	info.ModuleName = getModuleNameFromGoMod()

	return info, nil
}

// hasHelpersImport checks if the project uses Loom helpers
func hasHelpersImport() bool {
	// Search for import "github.com/geomark27/loom-go/pkg/helpers" in .go files
	files := []string{
		"internal/app/handlers/user_handler.go",
		"internal/modules/users/handler.go",
	}

	for _, file := range files {
		if data, err := os.ReadFile(file); err == nil {
			if bytes.Contains(data, []byte("loom-go/pkg/helpers")) {
				return true
			}
		}
	}

	return false
}

// getProjectNameFromGoMod extracts the project name from go.mod
func getProjectNameFromGoMod() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "unknown"
	}

	// Parse first line: "module <name>"
	lines := bytes.Split(data, []byte("\n"))
	if len(lines) > 0 {
		line := string(lines[0])
		if strings.HasPrefix(line, "module ") {
			moduleName := strings.TrimSpace(strings.TrimPrefix(line, "module "))
			parts := strings.Split(moduleName, "/")
			return parts[len(parts)-1]
		}
	}

	return "unknown"
}

// getModuleNameFromGoMod extracts the full module name from go.mod
func getModuleNameFromGoMod() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return ""
	}

	// Parse first line: "module <name>"
	lines := bytes.Split(data, []byte("\n"))
	if len(lines) > 0 {
		line := string(lines[0])
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}

	return ""
}

// ValidateComponentName validates that a component name is valid
func ValidateComponentName(name string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if strings.Contains(name, " ") {
		return fmt.Errorf("name cannot contain spaces")
	}

	if strings.ContainsAny(name, `<>:"/\|?*`) {
		return fmt.Errorf("name contains invalid characters")
	}

	return nil
}
