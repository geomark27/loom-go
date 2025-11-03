package upgrader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/geomark27/loom-go/internal/version"
)

// Upgrader manages project updates
type Upgrader struct {
	currentVersion version.Version
	targetVersion  version.Version
	backupManager  *BackupManager
}

// NewUpgrader creates a new upgrader
func NewUpgrader(current, target version.Version) *Upgrader {
	return &Upgrader{
		currentVersion: current,
		targetVersion:  target,
		backupManager:  NewBackupManager(),
	}
}

// CanUpgrade verifies if upgrade can be performed
func (u *Upgrader) CanUpgrade() (bool, string) {
	// Cannot downgrade
	if u.targetVersion.IsOlder(u.currentVersion) {
		return false, "Cannot downgrade. Target version is older than current version."
	}

	// Already updated
	if u.currentVersion.Compare(u.targetVersion) == 0 {
		return false, "The project is already on the most recent version."
	}

	// Check major version compatibility
	if !u.targetVersion.IsCompatible(u.currentVersion) {
		return false, "The upgrade requires major changes (major version). Consult the migration documentation."
	}

	return true, ""
}

// Upgrade executes the update process
func (u *Upgrader) Upgrade(createBackup bool) error {
	// Check if can be updated
	canUpgrade, reason := u.CanUpgrade()
	if !canUpgrade {
		return fmt.Errorf(reason)
	}

	// Create backup if requested
	var backupPath string
	if createBackup {
		var err error
		backupPath, err = u.backupManager.CreateBackup()
		if err != nil {
			return fmt.Errorf("error creating backup: %w", err)
		}
		fmt.Printf("💾 Backup created at: %s\n", backupPath)
	}

	// Apply incremental upgrades
	if err := u.applyUpgrades(); err != nil {
		if createBackup && backupPath != "" {
			fmt.Println("❌ Error during upgrade. You can restore the backup with:")
			fmt.Printf("   loom upgrade --restore %s\n", backupPath)
		}
		return err
	}

	return nil
}

// applyUpgrades applies all necessary upgrades
func (u *Upgrader) applyUpgrades() error {
	current := u.currentVersion

	// Apply upgrades step by step
	for current.Compare(u.targetVersion) < 0 {
		nextVersion := u.getNextVersion(current)

		fmt.Printf("⬆️  Upgrading from v%s to v%s...\n", current.String(), nextVersion.String())

		if err := u.applyVersionUpgrade(current, nextVersion); err != nil {
			return fmt.Errorf("error upgrading to v%s: %w", nextVersion.String(), err)
		}

		current = nextVersion
	}

	// Update .loom file
	architecture := "layered"
	if _, err := os.Stat("internal/modules"); err == nil {
		architecture = "modular"
	}

	if err := version.CreateLoomFile(u.targetVersion, architecture); err != nil {
		return fmt.Errorf("error updating .loom: %w", err)
	}

	fmt.Printf("✅ Project successfully updated to v%s!\n", u.targetVersion.String())

	return nil
}

// getNextVersion determines the next version in the upgrade path
func (u *Upgrader) getNextVersion(current version.Version) version.Version {
	// Update to the next minor version
	next := current
	next.Minor++
	next.Patch = 0

	// If it exceeds the target, use the target
	if next.Compare(u.targetVersion) > 0 {
		return u.targetVersion
	}

	return next
}

// applyVersionUpgrade applies a specific upgrade
func (u *Upgrader) applyVersionUpgrade(from, to version.Version) error {
	// Execute specific migrations based on version
	switch to.String() {
	case "0.2.0":
		return u.upgradeTo020()
	case "0.3.0":
		return u.upgradeTo030()
	case "0.4.0":
		return u.upgradeTo040()
	case "0.5.0":
		return u.upgradeTo050()
	default:
		// No specific migrations for this version
		return nil
	}
}

// upgradeTo020 migrates to version 0.2.0 (adds helpers)
func (u *Upgrader) upgradeTo020() error {
	fmt.Println("   📦 Adding helpers...")

	// Check if already has helpers
	if _, err := os.Stat("pkg/helpers"); err == nil {
		fmt.Println("   ℹ️  Helpers already exist, skipping.")
		return nil
	}

	// Here the helpers would be added (in a real upgrade)
	// For now we just inform
	fmt.Println("   ⚠️  You need to manually add the helpers from pkg/helpers/")

	return nil
}

// upgradeTo030 migrates to version 0.3.0
func (u *Upgrader) upgradeTo030() error {
	fmt.Println("   📚 Updating documentation...")

	// Add comment in go.mod
	if err := u.addLoomCommentToGoMod(); err != nil {
		return err
	}

	return nil
}

// upgradeTo040 migrates to version 0.4.0
func (u *Upgrader) upgradeTo040() error {
	fmt.Println("   ✨ Preparing support for 'loom generate'...")

	// Create .loom file if it doesn't exist
	if _, err := os.Stat(".loom"); os.IsNotExist(err) {
		architecture := "layered"
		if _, err := os.Stat("internal/modules"); err == nil {
			architecture = "modular"
		}

		if err := version.CreateLoomFile(version.Version{Major: 0, Minor: 4, Patch: 0}, architecture); err != nil {
			return err
		}
		fmt.Println("   ✅ .loom file created")
	}

	return nil
}

// upgradeTo050 migrates to version 0.5.0
func (u *Upgrader) upgradeTo050() error {
	fmt.Println("   ⬆️  Preparing support for 'loom upgrade'...")

	// Ensure .loom file exists
	if _, err := os.Stat(".loom"); os.IsNotExist(err) {
		architecture := "layered"
		if _, err := os.Stat("internal/modules"); err == nil {
			architecture = "modular"
		}

		if err := version.CreateLoomFile(version.Version{Major: 0, Minor: 5, Patch: 0}, architecture); err != nil {
			return err
		}
	}

	return nil
}

// addLoomCommentToGoMod adds a comment with the Loom version
func (u *Upgrader) addLoomCommentToGoMod() error {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return err
	}

	contentStr := string(content)

	// Check if it already has the comment
	if strings.Contains(contentStr, "// Generated with Loom") {
		return nil
	}

	// Add comment at the beginning (after the module line)
	lines := strings.Split(contentStr, "\n")
	newLines := []string{}

	for i, line := range lines {
		newLines = append(newLines, line)
		if i == 0 && strings.HasPrefix(line, "module ") {
			newLines = append(newLines, fmt.Sprintf("// Generated with Loom v%s", u.targetVersion.String()))
		}
	}

	newContent := strings.Join(newLines, "\n")
	return os.WriteFile("go.mod", []byte(newContent), 0644)
}

// RestoreBackup restores a specific backup
func (u *Upgrader) RestoreBackup(backupName string) error {
	backupPath := filepath.Join(u.backupManager.BackupDir, backupName)
	return u.backupManager.RestoreBackup(backupPath)
}

// ListBackups lists available backups
func (u *Upgrader) ListBackups() ([]string, error) {
	return u.backupManager.ListBackups()
}
