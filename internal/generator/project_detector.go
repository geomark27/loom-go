package generator

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// ProjectInfo contiene información sobre el proyecto detectado
type ProjectInfo struct {
	Name         string
	Architecture string // "layered" o "modular"
	HasHelpers   bool
	RootPath     string
	ModuleName   string
}

// DetectProject detecta el tipo de proyecto Loom en el directorio actual
func DetectProject() (*ProjectInfo, error) {
	// Buscar go.mod para confirmar que es un proyecto Go
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		return nil, fmt.Errorf("no se encontró go.mod. ¿Estás en un proyecto Go?")
	}

	info := &ProjectInfo{
		RootPath: ".",
	}

	// Detectar arquitectura
	if _, err := os.Stat("internal/modules"); err == nil {
		info.Architecture = "modular"
	} else if _, err := os.Stat("internal/app"); err == nil {
		info.Architecture = "layered"
	} else {
		return nil, fmt.Errorf("no se detectó un proyecto Loom válido (falta internal/modules o internal/app)")
	}

	// Detectar si tiene helpers
	info.HasHelpers = hasHelpersImport()

	// Leer nombre del proyecto desde go.mod
	info.Name = getProjectNameFromGoMod()
	info.ModuleName = getModuleNameFromGoMod()

	return info, nil
}

// hasHelpersImport busca si el proyecto usa los helpers de Loom
func hasHelpersImport() bool {
	// Buscar import "github.com/geomark27/loom-go/pkg/helpers" en archivos .go
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

// getProjectNameFromGoMod extrae el nombre del proyecto desde go.mod
func getProjectNameFromGoMod() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "unknown"
	}

	// Parsear primera línea: "module <nombre>"
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

// getModuleNameFromGoMod extrae el nombre completo del módulo desde go.mod
func getModuleNameFromGoMod() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return ""
	}

	// Parsear primera línea: "module <nombre>"
	lines := bytes.Split(data, []byte("\n"))
	if len(lines) > 0 {
		line := string(lines[0])
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}

	return ""
}

// ValidateComponentName valida que el nombre de un componente sea válido
func ValidateComponentName(name string) error {
	if name == "" {
		return fmt.Errorf("el nombre no puede estar vacío")
	}

	if strings.Contains(name, " ") {
		return fmt.Errorf("el nombre no puede contener espacios")
	}

	if strings.ContainsAny(name, `<>:"/\|?*`) {
		return fmt.Errorf("el nombre contiene caracteres no válidos")
	}

	return nil
}
