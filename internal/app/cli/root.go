package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loom",
	Short: "Loom - The Go project weaver",
	Long: `Loom is a CLI tool that automates the creation and extension
of Go projects following community best practices.

Loom is not a runtime framework, but a code generator
and project architecture orchestrator.`,
	Version: "1.0.4",
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags can be added here
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
}
