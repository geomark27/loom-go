package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Loom CLI and current project version",
	Long: `Show version information:
  - Loom CLI version
  - Current project version (if in a Loom project)
  - Update status`,
	RunE: runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string) error {
	fmt.Printf("🔧 Loom CLI v%s\n", version.Current.String())

	// Try to detect project version
	projectVersion, err := version.DetectProjectVersion()
	if err == nil {
		fmt.Printf("📦 Current project: v%s\n", projectVersion.String())

		// Compare versions
		if projectVersion.Compare(version.Current) < 0 {
			fmt.Printf("\n⚠️  Your project uses an old version of Loom\n")
			fmt.Printf("💡 Update with: loom upgrade\n")
		} else if projectVersion.Compare(version.Current) == 0 {
			fmt.Printf("✅ Your project is up to date\n")
		} else {
			fmt.Printf("⚠️  Your project uses a newer version than the CLI\n")
			fmt.Printf("💡 Update the Loom CLI\n")
		}
	} else {
		fmt.Printf("\nℹ️  No Loom project detected in current directory\n")
	}

	return nil
}
