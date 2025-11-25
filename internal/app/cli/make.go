package cli

import (
	"github.com/spf13/cobra"
)

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Generate database-related components (models, seeders)",
	Long: `Generate database-related components for GORM.

These commands require that you have previously run 'loom add orm gorm'
to set up the database structure.

Available subcommands:
  model   - Generate a GORM model with auto-registration
  seeder  - Generate a seeder with auto-registration

Examples:
  loom make model Product
  loom make seeder Product`,
	Aliases: []string{"mk"},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}
