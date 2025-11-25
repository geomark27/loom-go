package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var makeModelCmd = &cobra.Command{
	Use:   "model [name]",
	Short: "Generate a GORM model with auto-registration",
	Long: `Generate a GORM model file and automatically register it in models_all.go.

This command requires that you have previously run 'loom add orm gorm'.

The model will be generated with:
  - gorm.Model (ID, CreatedAt, UpdatedAt, DeletedAt)
  - JSON tags
  - GORM tags for common fields
  - Automatic registration in internal/database/models_all.go

Location depends on architecture:
  - Layered: internal/app/models/{name}.go
  - Modular: internal/models/{name}.go

Examples:
  loom make model Product
  loom make model Category --force`,
	Args: cobra.ExactArgs(1),
	RunE: runMakeModel,
}

func init() {
	makeCmd.AddCommand(makeModelCmd)
	makeModelCmd.Flags().Bool("force", false, "Overwrite existing files")
}

func runMakeModel(cmd *cobra.Command, args []string) error {
	name := args[0]
	force, _ := cmd.Flags().GetBool("force")

	// Detect project
	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no valid Loom project detected. %w", err)
	}

	// Validate name
	if err := generator.ValidateComponentName(name); err != nil {
		return fmt.Errorf("invalid model name: %w", err)
	}

	// Capitalize first letter for struct name
	structName := capitalizeFirst(name)
	fileName := strings.ToLower(name)

	// Determine paths based on architecture
	var modelPath string
	var modelsAllPath string

	if projectInfo.Architecture == "modular" {
		modelPath = filepath.Join(projectInfo.RootPath, "internal", "models", fileName+".go")
		modelsAllPath = filepath.Join(projectInfo.RootPath, "internal", "database", "models_all.go")
	} else {
		modelPath = filepath.Join(projectInfo.RootPath, "internal", "app", "models", fileName+".go")
		modelsAllPath = filepath.Join(projectInfo.RootPath, "internal", "database", "models_all.go")
	}

	// Check if models_all.go exists (GORM addon installed)
	if _, err := os.Stat(modelsAllPath); os.IsNotExist(err) {
		return fmt.Errorf("GORM not installed. Run 'loom add orm gorm' first")
	}

	// Check if model file already exists
	if _, err := os.Stat(modelPath); err == nil && !force {
		return fmt.Errorf("model %s already exists. Use --force to overwrite", fileName)
	}

	fmt.Printf("🔍 Project: %s (%s)\n", projectInfo.Name, projectInfo.Architecture)
	fmt.Printf("📦 Creating GORM model: %s\n\n", structName)

	// Generate model file
	modelContent := generateGORMModelContent(structName, projectInfo.ModuleName)

	// Create directory if needed
	if err := os.MkdirAll(filepath.Dir(modelPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Write model file
	if err := os.WriteFile(modelPath, []byte(modelContent), 0644); err != nil {
		return fmt.Errorf("failed to write model file: %w", err)
	}

	fmt.Printf("   ✅ Created: %s\n", modelPath)

	// Update models_all.go
	if err := addModelToRegistry(modelsAllPath, structName, projectInfo); err != nil {
		fmt.Printf("   ⚠️  Warning: Could not auto-register model: %v\n", err)
		fmt.Printf("   💡 Manually add '&models.%s{}' to AllModels in models_all.go\n", structName)
	} else {
		fmt.Printf("   ✅ Registered in: %s\n", modelsAllPath)
	}

	fmt.Println("\n✅ Model created successfully!")
	fmt.Println("\n📝 Next steps:")
	fmt.Printf("   1. Edit %s to add your fields\n", modelPath)
	fmt.Println("   2. Run 'make db-migrate' to create the table")

	return nil
}

func generateGORMModelContent(structName, moduleName string) string {
	return fmt.Sprintf(`package models

import "gorm.io/gorm"

// %s represents the %s entity in the database
type %s struct {
	gorm.Model
	Name        string %s
	Description string %s
	// Add your fields here
	// Example:
	// Price    float64 %s
	// Stock    int     %s
	// IsActive bool    %s
}

// TableName overrides the table name (optional)
// func (%s) TableName() string {
//     return "%ss"
// }
`, structName, strings.ToLower(structName), structName,
		"`gorm:\"size:100;not null\" json:\"name\"`",
		"`gorm:\"type:text\" json:\"description\"`",
		"`gorm:\"not null;default:0\" json:\"price\"`",
		"`gorm:\"default:0\" json:\"stock\"`",
		"`gorm:\"default:true\" json:\"is_active\"`",
		structName, strings.ToLower(structName))
}

func addModelToRegistry(modelsAllPath, structName string, projectInfo *generator.ProjectInfo) error {
	content, err := os.ReadFile(modelsAllPath)
	if err != nil {
		return err
	}

	contentStr := string(content)

	// Check if already registered (not in comments)
	registryLine := fmt.Sprintf("&models.%s{}", structName)
	if strings.Contains(contentStr, registryLine) && !strings.Contains(contentStr, "// "+registryLine) {
		return nil // Already registered
	}

	// Find the &models.User{}, line (the actual entry, not a comment)
	lines := strings.Split(contentStr, "\n")
	var insertIndex int = -1

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		// Look for actual model entries (start with &models. and end with {},)
		if strings.HasPrefix(trimmed, "&models.") && strings.HasSuffix(trimmed, "{},") {
			insertIndex = i
		}
	}

	if insertIndex == -1 {
		return fmt.Errorf("could not find AllModels entries")
	}

	// Insert new model after the last entry
	newLine := fmt.Sprintf("\t&models.%s{},", structName)
	newLines := make([]string, 0, len(lines)+1)
	newLines = append(newLines, lines[:insertIndex+1]...)
	newLines = append(newLines, newLine)
	newLines = append(newLines, lines[insertIndex+1:]...)

	newContent := strings.Join(newLines, "\n")
	return os.WriteFile(modelsAllPath, []byte(newContent), 0644)
}

func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
