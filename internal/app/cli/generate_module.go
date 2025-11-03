package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var generateModuleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Generates a complete module (handler, service, repository, model, DTO)",
	Long: `Generates a complete module with all its layers in the current project.

The command will automatically detect your project's architecture
(Layered or Modular) and generate the appropriate files.

For Layered architecture, it generates:
  - internal/app/handlers/{name}_handler.go
  - internal/domain/services/{name}_service.go
  - internal/infrastructure/repositories/{name}_repository.go
  - internal/domain/models/{name}.go
  - internal/domain/dto/{name}_dto.go

For Modular architecture, it generates:
  - internal/modules/{name}/handler.go
  - internal/modules/{name}/service.go
  - internal/modules/{name}/repository.go
  - internal/modules/{name}/model.go
  - internal/modules/{name}/dto.go
  - internal/modules/{name}/router.go
  - internal/modules/{name}/validator.go
  - internal/modules/{name}/errors.go

Examples:
  loom generate module products
  loom generate module users --force
  loom generate module orders --dry-run`,
	Aliases: []string{"mod", "m"},
	Args:    cobra.ExactArgs(1),
	RunE:    runGenerateModule,
}

func init() {
	generateCmd.AddCommand(generateModuleCmd)
}

func runGenerateModule(cmd *cobra.Command, args []string) error {
	moduleName := args[0]
	force, _ := cmd.Flags().GetBool("force")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	// Detect the current project (without arguments)
	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no valid Loom project detected. %w", err)
	}

	// Validate the module name
	if err := generator.ValidateComponentName(moduleName); err != nil {
		return fmt.Errorf("invalid module name: %w", err)
	}

	fmt.Printf("🔍 Project detected: %s\n", projectInfo.Name)
	fmt.Printf("📐 Architecture: %s\n", projectInfo.Architecture)
	fmt.Printf("📦 Generating module: %s\n\n", moduleName)

	// Create the generator
	gen := generator.NewModuleGenerator(projectInfo)

	// Generate the module (returns the list of files)
	files, err := gen.GenerateModule(moduleName, force, dryRun)
	if err != nil {
		return fmt.Errorf("error generating module: %w", err)
	}

	// Dry-run mode
	if dryRun {
		fmt.Println("📋 Files that would be generated:")
		for _, file := range files {
			fmt.Printf("   ✨ %s\n", file)
		}
		fmt.Println("\n💡 Run without --dry-run to create the files")
		return nil
	}

	fmt.Println("✅ Module generated successfully!")
	fmt.Println("\n📝 Files created:")
	for _, file := range files {
		fmt.Printf("   ✨ %s\n", file)
	}

	fmt.Println("\n📝 Next steps:")

	if projectInfo.Architecture == "modular" {
		fmt.Printf("   1. Register the router in cmd/loom/main.go:\n")
		fmt.Printf("      %sRouter := %s.NewRouter()\n", moduleName, moduleName)
		fmt.Printf("      router.PathPrefix(\"/%s\").Handler(%sRouter)\n\n", moduleName, moduleName)
	} else {
		fmt.Printf("   1. Register the routes in cmd/loom/main.go:\n")
		fmt.Printf("      %sHandler := handlers.New%sHandler()\n", moduleName, toPascalCase(moduleName))
		fmt.Printf("      router.HandleFunc(\"/%s\", %sHandler.Create).Methods(\"POST\")\n\n", moduleName, moduleName)
	}

	fmt.Println("   2. Run: go mod tidy")
	fmt.Println("   3. Implement the business logic in the generated files")

	return nil
}

// toPascalCase converts a string to PascalCase
func toPascalCase(s string) string {
	if len(s) == 0 {
		return s
	}
	// Convert first letter to uppercase
	if s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0]-32) + s[1:]
	}
	return s
}
