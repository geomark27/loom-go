package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var generateModelCmd = &cobra.Command{
	Use:   "model [name]",
	Short: "Generate a data model",
	Long: `Generate a model file with the data structure.

The file will be generated in the appropriate location according to the architecture:
  - Layered: internal/app/models/{name}.go
  - Modular: internal/modules/{name}/model.go

Examples:
  loom generate model Product
  loom generate model User --force`,
	Aliases: []string{"mod"},
	Args:    cobra.ExactArgs(1),
	RunE:    runGenerateModel,
}

func init() {
	generateCmd.AddCommand(generateModelCmd)
}

func runGenerateModel(cmd *cobra.Command, args []string) error {
	name := args[0]
	force, _ := cmd.Flags().GetBool("force")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no valid Loom project detected. %w", err)
	}

	if err := generator.ValidateComponentName(name); err != nil {
		return fmt.Errorf("invalid model name: %w", err)
	}

	fmt.Printf("🔍 Project: %s (%s)\n", projectInfo.Name, projectInfo.Architecture)
	fmt.Printf("📦 Generating model: %s\n\n", name)

	gen := generator.NewModuleGenerator(projectInfo)
	files, err := gen.GenerateModel(name, force, dryRun)
	if err != nil {
		return fmt.Errorf("error generating model: %w", err)
	}

	if dryRun {
		fmt.Println("📋 File that would be generated:")
		for _, file := range files {
			fmt.Printf("   ✨ %s\n", file)
		}
		fmt.Println("\n💡 Run without --dry-run to create the file")
		return nil
	}

	fmt.Println("✅ Model generated successfully!")
	fmt.Println("\n📝 File created:")
	for _, file := range files {
		fmt.Printf("   ✨ %s\n", file)
	}

	return nil
}
