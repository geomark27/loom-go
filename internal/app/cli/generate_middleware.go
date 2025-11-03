package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/generator"
	"github.com/spf13/cobra"
)

var generateMiddlewareCmd = &cobra.Command{
	Use:   "middleware [name]",
	Short: "Generates an HTTP middleware",
	Long: `Generates a middleware to intercept HTTP requests.

The file will be generated in the appropriate location according to the architecture:
  - Layered: internal/app/middleware/{name}.go
  - Modular: internal/middleware/{name}.go

Examples:
  loom generate middleware auth
  loom generate middleware logger --force`,
	Aliases: []string{"mw"},
	Args:    cobra.ExactArgs(1),
	RunE:    runGenerateMiddleware,
}

func init() {
	generateCmd.AddCommand(generateMiddlewareCmd)
}

func runGenerateMiddleware(cmd *cobra.Command, args []string) error {
	name := args[0]
	force, _ := cmd.Flags().GetBool("force")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	projectInfo, err := generator.DetectProject()
	if err != nil {
		return fmt.Errorf("error: no valid Loom project detected. %w", err)
	}

	if err := generator.ValidateComponentName(name); err != nil {
		return fmt.Errorf("invalid middleware name: %w", err)
	}

	fmt.Printf("🔍 Project: %s (%s)\n", projectInfo.Name, projectInfo.Architecture)
	fmt.Printf("📦 Generating middleware: %s\n\n", name)

	gen := generator.NewModuleGenerator(projectInfo)
	files, err := gen.GenerateMiddleware(name, force, dryRun)
	if err != nil {
		return fmt.Errorf("error generating middleware: %w", err)
	}

	if dryRun {
		fmt.Println("📋 File that would be generated:")
		for _, file := range files {
			fmt.Printf("   ✨ %s\n", file)
		}
		fmt.Println("\n💡 Run without --dry-run to create the file")
		return nil
	}

	fmt.Println("✅ Middleware generated successfully!")
	fmt.Println("\n📝 File created:")
	for _, file := range files {
		fmt.Printf("   ✨ %s\n", file)
	}

	fmt.Println("\n📝 Next step:")
	fmt.Println("   Register the middleware in your router or on specific routes")

	return nil
}
