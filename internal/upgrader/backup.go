package upgrader

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// BackupManager gestiona los respaldos del proyecto
type BackupManager struct {
	BackupDir string
}

// NewBackupManager crea un nuevo gestor de respaldos
func NewBackupManager() *BackupManager {
	return &BackupManager{
		BackupDir: ".loom-backups",
	}
}

// CreateBackup crea un respaldo completo del proyecto
func (bm *BackupManager) CreateBackup() (string, error) {
	// Crear directorio de backups si no existe
	if err := os.MkdirAll(bm.BackupDir, 0755); err != nil {
		return "", fmt.Errorf("error al crear directorio de backup: %w", err)
	}

	// Nombre del backup con timestamp
	timestamp := time.Now().Format("20060102-150405")
	backupName := fmt.Sprintf("backup-%s", timestamp)
	backupPath := filepath.Join(bm.BackupDir, backupName)

	// Crear directorio del backup
	if err := os.MkdirAll(backupPath, 0755); err != nil {
		return "", fmt.Errorf("error al crear directorio de backup: %w", err)
	}

	// Copiar archivos importantes
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
				return "", fmt.Errorf("error al copiar %s: %w", file, err)
			}
		}
	}

	return backupPath, nil
}

// RestoreBackup restaura un backup
func (bm *BackupManager) RestoreBackup(backupPath string) error {
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		return fmt.Errorf("backup no encontrado: %s", backupPath)
	}

	// Listar archivos en el backup
	entries, err := os.ReadDir(backupPath)
	if err != nil {
		return fmt.Errorf("error al leer backup: %w", err)
	}

	// Restaurar cada archivo
	for _, entry := range entries {
		src := filepath.Join(backupPath, entry.Name())
		dest := entry.Name()

		if err := copyPath(src, dest); err != nil {
			return fmt.Errorf("error al restaurar %s: %w", entry.Name(), err)
		}
	}

	return nil
}

// ListBackups lista los backups disponibles
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

// copyPath copia un archivo o directorio recursivamente
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

// copyFile copia un archivo
func copyFile(src, dst string) error {
	// Crear directorio destino si no existe
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}

	// Leer archivo fuente
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// Escribir archivo destino
	return os.WriteFile(dst, data, 0644)
}

// copyDir copia un directorio recursivamente
func copyDir(src, dst string) error {
	// Crear directorio destino
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	// Listar contenido
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// Copiar cada entrada
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if err := copyPath(srcPath, dstPath); err != nil {
			return err
		}
	}

	return nil
}
