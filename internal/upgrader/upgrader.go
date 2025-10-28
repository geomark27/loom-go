package upgrader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/geomark27/loom-go/internal/version"
)

// Upgrader gestiona las actualizaciones del proyecto
type Upgrader struct {
	currentVersion version.Version
	targetVersion  version.Version
	backupManager  *BackupManager
}

// NewUpgrader crea un nuevo upgrader
func NewUpgrader(current, target version.Version) *Upgrader {
	return &Upgrader{
		currentVersion: current,
		targetVersion:  target,
		backupManager:  NewBackupManager(),
	}
}

// CanUpgrade verifica si se puede hacer el upgrade
func (u *Upgrader) CanUpgrade() (bool, string) {
	// No se puede downgrade
	if u.targetVersion.IsOlder(u.currentVersion) {
		return false, "No se puede hacer downgrade. La versión objetivo es más antigua que la actual."
	}

	// Ya está actualizado
	if u.currentVersion.Compare(u.targetVersion) == 0 {
		return false, "El proyecto ya está en la versión más reciente."
	}

	// Verificar compatibilidad major version
	if !u.targetVersion.IsCompatible(u.currentVersion) {
		return false, "El upgrade requiere cambios mayores (major version). Consulta la documentación de migración."
	}

	return true, ""
}

// Upgrade ejecuta el proceso de actualización
func (u *Upgrader) Upgrade(createBackup bool) error {
	// Verificar si se puede actualizar
	canUpgrade, reason := u.CanUpgrade()
	if !canUpgrade {
		return fmt.Errorf(reason)
	}

	// Crear backup si se solicita
	var backupPath string
	if createBackup {
		var err error
		backupPath, err = u.backupManager.CreateBackup()
		if err != nil {
			return fmt.Errorf("error al crear backup: %w", err)
		}
		fmt.Printf("💾 Backup creado en: %s\n", backupPath)
	}

	// Aplicar upgrades incrementales
	if err := u.applyUpgrades(); err != nil {
		if createBackup && backupPath != "" {
			fmt.Println("❌ Error durante el upgrade. Puedes restaurar el backup con:")
			fmt.Printf("   loom upgrade --restore %s\n", backupPath)
		}
		return err
	}

	return nil
}

// applyUpgrades aplica todos los upgrades necesarios
func (u *Upgrader) applyUpgrades() error {
	current := u.currentVersion

	// Aplicar upgrades paso a paso
	for current.Compare(u.targetVersion) < 0 {
		nextVersion := u.getNextVersion(current)

		fmt.Printf("⬆️  Actualizando de v%s a v%s...\n", current.String(), nextVersion.String())

		if err := u.applyVersionUpgrade(current, nextVersion); err != nil {
			return fmt.Errorf("error al actualizar a v%s: %w", nextVersion.String(), err)
		}

		current = nextVersion
	}

	// Actualizar archivo .loom
	architecture := "layered"
	if _, err := os.Stat("internal/modules"); err == nil {
		architecture = "modular"
	}

	if err := version.CreateLoomFile(u.targetVersion, architecture); err != nil {
		return fmt.Errorf("error al actualizar .loom: %w", err)
	}

	fmt.Printf("✅ Proyecto actualizado a v%s exitosamente!\n", u.targetVersion.String())

	return nil
}

// getNextVersion determina la siguiente versión en el camino de upgrade
func (u *Upgrader) getNextVersion(current version.Version) version.Version {
	// Actualizar a la siguiente minor version
	next := current
	next.Minor++
	next.Patch = 0

	// Si excede el target, usar el target
	if next.Compare(u.targetVersion) > 0 {
		return u.targetVersion
	}

	return next
}

// applyVersionUpgrade aplica un upgrade específico
func (u *Upgrader) applyVersionUpgrade(from, to version.Version) error {
	// Ejecutar migraciones específicas según la versión
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
		// No hay migraciones específicas para esta versión
		return nil
	}
}

// upgradeTo020 migra a versión 0.2.0 (añade helpers)
func (u *Upgrader) upgradeTo020() error {
	fmt.Println("   📦 Añadiendo helpers...")

	// Verificar si ya tiene helpers
	if _, err := os.Stat("pkg/helpers"); err == nil {
		fmt.Println("   ℹ️  Los helpers ya existen, omitiendo.")
		return nil
	}

	// Aquí se añadirían los helpers (en un upgrade real)
	// Por ahora solo informamos
	fmt.Println("   ⚠️  Necesitas añadir manualmente los helpers de pkg/helpers/")

	return nil
}

// upgradeTo030 migra a versión 0.3.0
func (u *Upgrader) upgradeTo030() error {
	fmt.Println("   📚 Actualizando documentación...")

	// Añadir comentario en go.mod
	if err := u.addLoomCommentToGoMod(); err != nil {
		return err
	}

	return nil
}

// upgradeTo040 migra a versión 0.4.0
func (u *Upgrader) upgradeTo040() error {
	fmt.Println("   ✨ Preparando soporte para 'loom generate'...")

	// Crear archivo .loom si no existe
	if _, err := os.Stat(".loom"); os.IsNotExist(err) {
		architecture := "layered"
		if _, err := os.Stat("internal/modules"); err == nil {
			architecture = "modular"
		}

		if err := version.CreateLoomFile(version.Version{Major: 0, Minor: 4, Patch: 0}, architecture); err != nil {
			return err
		}
		fmt.Println("   ✅ Archivo .loom creado")
	}

	return nil
}

// upgradeTo050 migra a versión 0.5.0
func (u *Upgrader) upgradeTo050() error {
	fmt.Println("   ⬆️  Preparando soporte para 'loom upgrade'...")

	// Asegurar que existe el archivo .loom
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

// addLoomCommentToGoMod añade un comentario con la versión de Loom
func (u *Upgrader) addLoomCommentToGoMod() error {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return err
	}

	contentStr := string(content)

	// Verificar si ya tiene el comentario
	if strings.Contains(contentStr, "// Generated with Loom") {
		return nil
	}

	// Añadir comentario al principio (después de la línea module)
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

// RestoreBackup restaura un backup específico
func (u *Upgrader) RestoreBackup(backupName string) error {
	backupPath := filepath.Join(u.backupManager.BackupDir, backupName)
	return u.backupManager.RestoreBackup(backupPath)
}

// ListBackups lista los backups disponibles
func (u *Upgrader) ListBackups() ([]string, error) {
	return u.backupManager.ListBackups()
}
