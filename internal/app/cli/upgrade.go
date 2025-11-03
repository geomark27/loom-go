package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/upgrader"
	"github.com/geomark27/loom-go/internal/version"
	"github.com/spf13/cobra"
)

var (
	noBackup    bool
	showChanges bool
	restoreName string
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade a Loom project to the latest version",
	Long: `Upgrade an existing Loom project to the latest CLI version.

The command:
  1. Detects the current project version
  2. Creates an automatic backup (optional)
  3. Applies necessary migrations
  4. Updates the .loom file with the new version

Examples:
  loom upgrade                    # Upgrade with backup
  loom upgrade --no-backup        # Upgrade without backup
  loom upgrade --show-changes     # Show changes without upgrading
  loom upgrade --restore backup-20231027-153045  # Restore backup`,
	RunE: runUpgrade,
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	upgradeCmd.Flags().BoolVar(&noBackup, "no-backup", false, "Don't create backup before upgrading")
	upgradeCmd.Flags().BoolVar(&showChanges, "show-changes", false, "Show changes without upgrading")
	upgradeCmd.Flags().StringVar(&restoreName, "restore", "", "Restore a specific backup")
}

func runUpgrade(cmd *cobra.Command, args []string) error {
	// Restore backup mode
	if restoreName != "" {
		return restoreBackup(restoreName)
	}

	// Detect project version
	currentVersion, err := version.DetectProjectVersion()
	if err != nil {
		fmt.Println("⚠️  Could not detect project version.")
		fmt.Println("ℹ️  Assuming version 0.1.0")
		currentVersion = version.Version{Major: 0, Minor: 1, Patch: 0}
	}

	targetVersion := version.Current

	fmt.Printf("📊 Current project version: v%s\n", currentVersion.String())
	fmt.Printf("🎯 Target version: v%s\n\n", targetVersion.String())

	// Show changes if requested
	if showChanges {
		return showVersionChanges(currentVersion, targetVersion)
	}

	// Create upgrader
	upg := upgrader.NewUpgrader(currentVersion, targetVersion)

	// Check if upgrade is possible
	canUpgrade, reason := upg.CanUpgrade()
	if !canUpgrade {
		fmt.Println("ℹ️ ", reason)
		return nil
	}

	// Show changes
	changelog := version.GetChangelogBetween(currentVersion, targetVersion)
	if changelog != "" {
		fmt.Println("📋 Changes to be applied:")
		fmt.Println(changelog)
		fmt.Println()
	}

	// Execute upgrade
	createBackup := !noBackup
	if err := upg.Upgrade(createBackup); err != nil {
		return fmt.Errorf("error during upgrade: %w", err)
	}

	fmt.Println("\n📝 Next steps:")
	fmt.Println("   1. Run: go mod tidy")
	fmt.Println("   2. Review changes with: git diff")
	fmt.Println("   3. Test your project: go build ./cmd/...")

	if createBackup {
		backups, _ := upg.ListBackups()
		if len(backups) > 0 {
			fmt.Printf("\n💡 If something went wrong, restore with: loom upgrade --restore %s\n", backups[len(backups)-1])
		}
	}

	return nil
}

func showVersionChanges(current, target version.Version) error {
	fmt.Println("📋 Changes between versions:")
	fmt.Println()

	changelog := version.GetChangelogBetween(current, target)
	if changelog == "" {
		fmt.Println("No changes recorded between these versions.")
		return nil
	}

	fmt.Println(changelog)
	fmt.Println("\n💡 Run 'loom upgrade' to apply these changes")

	return nil
}

func restoreBackup(backupName string) error {
	fmt.Printf("♻️  Restoring backup: %s\n", backupName)

	upg := upgrader.NewUpgrader(version.Version{}, version.Version{})

	if err := upg.RestoreBackup(backupName); err != nil {
		return fmt.Errorf("error restoring backup: %w", err)
	}

	fmt.Println("✅ Backup restored successfully!")
	fmt.Println("\n📝 Recommendations:")
	fmt.Println("   1. Verify the restored files")
	fmt.Println("   2. Run: go mod tidy")
	fmt.Println("   3. Build your project: go build ./cmd/...")

	return nil
}
