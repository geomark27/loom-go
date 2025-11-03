package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var generateServiceCmd = &cobra.Command{
	Use:   "service [name]",
	Short: "Generate a service with business logic",
	Long: `Generate a service file to contain business logic.

The file will be generated in the appropriate location based on architecture:
  - Layered: internal/app/services/{name}_service.go
  - Modular: internal/modules/{name}/service.go

Examples:
  loom generate service products
  loom generate service email --force`,
	Aliases: []string{"svc", "s"},
	Args:    cobra.ExactArgs(1),
	RunE:    runGenerateService,
}

func init() {
	generateCmd.AddCommand(generateServiceCmd)
}

func runGenerateService(cmd *cobra.Command, args []string) error {
	name := args[0]
	force, _ := cmd.Flags().GetBool("force")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no valid Loom project detected. %w", err)
	}

	if err := generator.ValidateComponentName(name); err != nil {
		return fmt.Errorf("invalid service name: %w", err)
	}

	fmt.Printf("🔍 Project: %s (%s)\n", projectInfo.Name, projectInfo.Architecture)
	fmt.Printf("📦 Generating service: %s\n\n", name)

	gen := generator.NewModuleGenerator(projectInfo)
	files, err := gen.GenerateService(name, force, dryRun)
	if err != nil {
		return fmt.Errorf("error generating service: %w", err)
	}

	if dryRun {
		fmt.Println("📋 File that would be generated:")
		for _, file := range files {
			fmt.Printf("   ✨ %s\n", file)
		}
		fmt.Println("\n💡 Run without --dry-run to create the file")
		return nil
	}

	fmt.Println("✅ Service generated successfully!")
	fmt.Println("\n📝 File created:")
	for _, file := range files {
		fmt.Printf("   ✨ %s\n", file)
	}

	return nil
}
