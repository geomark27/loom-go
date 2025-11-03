package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var generateHandlerCmd = &cobra.Command{
	Use:   "handler [name]",
	Short: "Generates an HTTP handler",
	Long: `Generates a handler file to handle HTTP requests.

The file will be generated in the appropriate location according to the architecture:
  - Layered: internal/app/handlers/{name}_handler.go
  - Modular: internal/modules/{name}/handler.go

Examples:
  loom generate handler products
  loom generate handler users --force`,
	Aliases: []string{"h"},
	Args:    cobra.ExactArgs(1),
	RunE:    runGenerateHandler,
}

func init() {
	generateCmd.AddCommand(generateHandlerCmd)
}

func runGenerateHandler(cmd *cobra.Command, args []string) error {
	name := args[0]
	force, _ := cmd.Flags().GetBool("force")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no valid Loom project detected. %w", err)
	}

	if err := generator.ValidateComponentName(name); err != nil {
		return fmt.Errorf("invalid handler name: %w", err)
	}

	fmt.Printf("🔍 Project: %s (%s)\n", projectInfo.Name, projectInfo.Architecture)
	fmt.Printf("📦 Generating handler: %s\n\n", name)

	gen := generator.NewModuleGenerator(projectInfo)
	files, err := gen.GenerateHandler(name, force, dryRun)
	if err != nil {
		return fmt.Errorf("error generating handler: %w", err)
	}

	if dryRun {
		fmt.Println("📋 File that would be generated:")
		for _, file := range files {
			fmt.Printf("   ✨ %s\n", file)
		}
		fmt.Println("\n💡 Run without --dry-run to create the file")
		return nil
	}

	fmt.Println("✅ Handler generated successfully!")
	fmt.Println("\n📝 File created:")
	for _, file := range files {
		fmt.Printf("   ✨ %s\n", file)
	}

	return nil
}
