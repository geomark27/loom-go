package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var makeSeederCmd = &cobra.Command{
	Use:   "seeder [name]",
	Short: "Generate a seeder with auto-registration",
	Long: `Generate a seeder file and automatically register it in seeders_all.go.

This command requires that you have previously run 'loom add orm gorm'.

The seeder will be generated with:
  - Seeder interface implementation
  - Template for seed data
  - Automatic registration in internal/database/seeders/seeders_all.go

Location:
  internal/database/seeders/{name}_seeder.go

Examples:
  loom make seeder Product
  loom make seeder Category --force`,
	Args: cobra.ExactArgs(1),
	RunE: runMakeSeeder,
}

func init() {
	makeCmd.AddCommand(makeSeederCmd)
	makeSeederCmd.Flags().Bool("force", false, "Overwrite existing files")
}

func runMakeSeeder(cmd *cobra.Command, args []string) error {
	name := args[0]
	force, _ := cmd.Flags().GetBool("force")

	// Detect project
	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no valid Loom project detected. %w", err)
	}

	// Validate name
	if err := generator.ValidateComponentName(name); err != nil {
		return fmt.Errorf("invalid seeder name: %w", err)
	}

	// Capitalize first letter for struct name
	structName := capitalizeFirst(name)
	fileName := strings.ToLower(name)

	// Determine paths
	seederPath := filepath.Join(projectInfo.RootPath, "internal", "database", "seeders", fileName+"_seeder.go")
	seedersAllPath := filepath.Join(projectInfo.RootPath, "internal", "database", "seeders", "seeders_all.go")

	// Check if seeders_all.go exists (GORM addon installed)
	if _, err := os.Stat(seedersAllPath); os.IsNotExist(err) {
		return fmt.Errorf("GORM not installed. Run 'loom add orm gorm' first")
	}

	// Check if seeder file already exists
	if _, err := os.Stat(seederPath); err == nil && !force {
		return fmt.Errorf("seeder %s already exists. Use --force to overwrite", fileName)
	}

	fmt.Printf("🔍 Project: %s (%s)\n", projectInfo.Name, projectInfo.Architecture)
	fmt.Printf("🌱 Creating seeder: %sSeeder\n\n", structName)

	// Determine models path based on architecture
	var modelsPath string
	if projectInfo.Architecture == "modular" {
		modelsPath = projectInfo.ModuleName + "/internal/models"
	} else {
		modelsPath = projectInfo.ModuleName + "/internal/app/models"
	}

	// Generate seeder file
	seederContent := generateSeederContent(structName, projectInfo.ModuleName, modelsPath)

	// Create directory if needed
	if err := os.MkdirAll(filepath.Dir(seederPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Write seeder file
	if err := os.WriteFile(seederPath, []byte(seederContent), 0644); err != nil {
		return fmt.Errorf("failed to write seeder file: %w", err)
	}

	fmt.Printf("   ✅ Created: %s\n", seederPath)

	// Update seeders_all.go
	if err := addSeederToRegistry(seedersAllPath, structName); err != nil {
		fmt.Printf("   ⚠️  Warning: Could not auto-register seeder: %v\n", err)
		fmt.Printf("   💡 Manually add '&%sSeeder{}' to AllSeeders in seeders_all.go\n", structName)
	} else {
		fmt.Printf("   ✅ Registered in: %s\n", seedersAllPath)
	}

	fmt.Println("\n✅ Seeder created successfully!")
	fmt.Println("\n📝 Next steps:")
	fmt.Printf("   1. Edit %s to add your seed data\n", seederPath)
	fmt.Println("   2. Run 'make db-seed' to execute seeders")

	return nil
}

func generateSeederContent(structName, moduleName, modelsPath string) string {
	lowerName := strings.ToLower(structName)

	return fmt.Sprintf(`package seeders

import (
	"%s"
	"fmt"
	"gorm.io/gorm"
)

// %sSeeder seeds the %s table
type %sSeeder struct{}

// Run implements the Seeder interface
func (s *%sSeeder) Run(db *gorm.DB) error {
	%ss := []models.%s{
		// Add your seed data here
		// Example:
		// {Name: "Example 1", Description: "First example"},
		// {Name: "Example 2", Description: "Second example"},
	}

	for _, item := range %ss {
		// Use FirstOrCreate to avoid duplicates
		result := db.FirstOrCreate(&item, models.%s{Name: item.Name})
		if result.Error != nil {
			return fmt.Errorf("failed to seed %s: %%w", result.Error)
		}
	}

	fmt.Printf("✅ %sSeeder: seeded %%d %ss\n", len(%ss))
	return nil
}
`, modelsPath,
		structName, lowerName, structName,
		structName,
		lowerName, structName,
		lowerName, structName,
		lowerName,
		structName, lowerName, lowerName)
}

func addSeederToRegistry(seedersAllPath, structName string) error {
	content, err := os.ReadFile(seedersAllPath)
	if err != nil {
		return err
	}

	contentStr := string(content)
	seederName := structName + "Seeder"

	// Check if already registered
	registryLine := "&" + seederName + "{}"
	if strings.Contains(contentStr, registryLine) && !strings.Contains(contentStr, "// "+registryLine) {
		return nil // Already registered
	}

	// Find the actual seeder entries (not comments)
	lines := strings.Split(contentStr, "\n")
	var insertIndex int = -1

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		// Look for actual seeder entries (start with & and contain Seeder{},)
		if strings.HasPrefix(trimmed, "&") && strings.Contains(trimmed, "Seeder{},") {
			insertIndex = i
		}
	}

	if insertIndex == -1 {
		return fmt.Errorf("could not find AllSeeders entries")
	}

	// Insert new seeder after the last entry
	newLine := fmt.Sprintf("\t&%s{},", seederName)
	newLines := make([]string, 0, len(lines)+1)
	newLines = append(newLines, lines[:insertIndex+1]...)
	newLines = append(newLines, newLine)
	newLines = append(newLines, lines[insertIndex+1:]...)

	newContent := strings.Join(newLines, "\n")
	return os.WriteFile(seedersAllPath, []byte(newContent), 0644)
}
