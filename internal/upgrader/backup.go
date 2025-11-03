package upgrader

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// BackupManager manages project backups
type BackupManager struct {
	BackupDir string
}

// NewBackupManager creates a new backup manager
func NewBackupManager() *BackupManager {
	return &BackupManager{
		BackupDir: ".loom-backups",
	}
}

// CreateBackup creates a complete project backup
func (bm *BackupManager) CreateBackup() (string, error) {
	// Create backups directory if it doesn't exist
	if err := os.MkdirAll(bm.BackupDir, 0755); err != nil {
		return "", fmt.Errorf("error creating backup directory: %w", err)
	}

	// Backup name with timestamp
	timestamp := time.Now().Format("20060102-150405")
	backupName := fmt.Sprintf("backup-%s", timestamp)
	backupPath := filepath.Join(bm.BackupDir, backupName)

	// Create backup directory
	if err := os.MkdirAll(backupPath, 0755); err != nil {
		return "", fmt.Errorf("error creating backup directory: %w", err)
	}

	// Copy important files
	filesToBackup := []string{
		"go.mod",
		"go.sum",
		"internal/",
		"cmd/",
		"pkg/",
		".loom",
	}

	for _, file := range filesToBackup {
		if _, err := os.Stat(file); err == nil {
			destPath := filepath.Join(backupPath, file)
			if err := copyPath(file, destPath); err != nil {
				return "", fmt.Errorf("error copying %s: %w", file, err)
			}
		}
	}

	return backupPath, nil
}

// RestoreBackup restores a backup
func (bm *BackupManager) RestoreBackup(backupPath string) error {
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		return fmt.Errorf("backup not found: %s", backupPath)
	}

	// List files in the backup
	entries, err := os.ReadDir(backupPath)
	if err != nil {
		return fmt.Errorf("error reading backup: %w", err)
	}

	// Restore each file
	for _, entry := range entries {
		src := filepath.Join(backupPath, entry.Name())
		dest := entry.Name()

		if err := copyPath(src, dest); err != nil {
			return fmt.Errorf("error restoring %s: %w", entry.Name(), err)
		}
	}

	return nil
}

// ListBackups lists available backups
func (bm *BackupManager) ListBackups() ([]string, error) {
	if _, err := os.Stat(bm.BackupDir); os.IsNotExist(err) {
		return []string{}, nil
	}

	entries, err := os.ReadDir(bm.BackupDir)
	if err != nil {
		return nil, err
	}

	backups := []string{}
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != ".gitkeep" {
			backups = append(backups, entry.Name())
		}
	}

	return backups, nil
}

// copyPath copies a file or directory recursively
func copyPath(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return copyDir(src, dst)
	}
	return copyFile(src, dst)
}

// copyFile copies a file
func copyFile(src, dst string) error {
	// Create destination directory if it doesn't exist
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}

	// Read source file
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// Write destination file
	return os.WriteFile(dst, data, 0644)
}

// copyDir copies a directory recursively
func copyDir(src, dst string) error {
	// Create destination directory
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	// List content
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// Copy each entry
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if err := copyPath(srcPath, dstPath); err != nil {
			return err
		}
	}

	return nil
}
