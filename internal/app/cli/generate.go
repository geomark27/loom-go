package cli

import (
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate components in an existing project",
	Long: `Generate modules, handlers, services, models and other components
in an existing Loom project.

The command will automatically detect if your project uses Layered
or Modular architecture and generate the appropriate code.

Examples:
  loom generate module products
  loom generate handler orders
  loom generate service email
  loom generate model Category
  loom generate middleware auth`,
	Aliases: []string{"gen", "g"},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Global flags for all generate subcommands
	generateCmd.PersistentFlags().Bool("force", false, "Overwrite existing files")
	generateCmd.PersistentFlags().Bool("dry-run", false, "Show what would be generated without creating files")
}
