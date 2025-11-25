package addon

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/geomark27/loom-go/internal/generator"
)

// GetModuleName reads the module name from go.mod in the specified project root
func GetModuleName(projectRoot string) (string, error) {
	goModPath := filepath.Join(projectRoot, "go.mod")

	data, err := os.ReadFile(goModPath)
	if err != nil {
		return "", fmt.Errorf("failed to read go.mod: %w", err)
	}

	// Parse first line: "module <name>"
	lines := bytes.Split(data, []byte("\n"))
	if len(lines) > 0 {
		line := string(lines[0])
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	return "", fmt.Errorf("module name not found in go.mod")
}

// GenerateFileFromTemplate generates a file from a template
func GenerateFileFromTemplate(templateName, targetPath string, data map[string]interface{}) error {
	// Get template content
	content, err := generator.GetTemplateContent(templateName)
	if err != nil {
		return fmt.Errorf("failed to get template %s: %w", templateName, err)
	}

	// Parse template
	tmpl, err := template.New(templateName).Parse(content)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templateName, err)
	}

	// Create target file
	file, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", targetPath, err)
	}
	defer file.Close()

	// Execute template
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", templateName, err)
	}

	return nil
}
