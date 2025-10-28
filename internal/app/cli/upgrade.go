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
	Short: "Actualiza un proyecto Loom a la última versión",
	Long: `Actualiza un proyecto Loom existente a la última versión del CLI.

El comando:
  1. Detecta la versión actual del proyecto
  2. Crea un backup automático (opcional)
  3. Aplica las migraciones necesarias
  4. Actualiza el archivo .loom con la nueva versión

Ejemplos:
  loom upgrade                    # Actualizar con backup
  loom upgrade --no-backup        # Actualizar sin backup
  loom upgrade --show-changes     # Ver cambios sin actualizar
  loom upgrade --restore backup-20231027-153045  # Restaurar backup`,
	RunE: runUpgrade,
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	upgradeCmd.Flags().BoolVar(&noBackup, "no-backup", false, "No crear backup antes de actualizar")
	upgradeCmd.Flags().BoolVar(&showChanges, "show-changes", false, "Mostrar cambios sin actualizar")
	upgradeCmd.Flags().StringVar(&restoreName, "restore", "", "Restaurar un backup específico")
}

func runUpgrade(cmd *cobra.Command, args []string) error {
	// Modo restaurar backup
	if restoreName != "" {
		return restoreBackup(restoreName)
	}

	// Detectar versión del proyecto
	currentVersion, err := version.DetectProjectVersion()
	if err != nil {
		fmt.Println("⚠️  No se pudo detectar la versión del proyecto.")
		fmt.Println("ℹ️  Asumiendo versión 0.1.0")
		currentVersion = version.Version{Major: 0, Minor: 1, Patch: 0}
	}

	targetVersion := version.Current

	fmt.Printf("📊 Versión actual del proyecto: v%s\n", currentVersion.String())
	fmt.Printf("🎯 Versión objetivo: v%s\n\n", targetVersion.String())

	// Mostrar cambios si se solicita
	if showChanges {
		return showVersionChanges(currentVersion, targetVersion)
	}

	// Crear upgrader
	upg := upgrader.NewUpgrader(currentVersion, targetVersion)

	// Verificar si se puede actualizar
	canUpgrade, reason := upg.CanUpgrade()
	if !canUpgrade {
		fmt.Println("ℹ️ ", reason)
		return nil
	}

	// Mostrar cambios
	changelog := version.GetChangelogBetween(currentVersion, targetVersion)
	if changelog != "" {
		fmt.Println("📋 Cambios que se aplicarán:")
		fmt.Println(changelog)
		fmt.Println()
	}

	// Ejecutar upgrade
	createBackup := !noBackup
	if err := upg.Upgrade(createBackup); err != nil {
		return fmt.Errorf("error durante el upgrade: %w", err)
	}

	fmt.Println("\n📝 Próximos pasos:")
	fmt.Println("   1. Ejecuta: go mod tidy")
	fmt.Println("   2. Revisa los cambios con: git diff")
	fmt.Println("   3. Prueba tu proyecto: go build ./cmd/...")

	if createBackup {
		backups, _ := upg.ListBackups()
		if len(backups) > 0 {
			fmt.Printf("\n💡 Si algo salió mal, restaura con: loom upgrade --restore %s\n", backups[len(backups)-1])
		}
	}

	return nil
}

func showVersionChanges(current, target version.Version) error {
	fmt.Println("📋 Cambios entre versiones:")
	fmt.Println()

	changelog := version.GetChangelogBetween(current, target)
	if changelog == "" {
		fmt.Println("No hay cambios registrados entre estas versiones.")
		return nil
	}

	fmt.Println(changelog)
	fmt.Println("\n💡 Ejecuta 'loom upgrade' para aplicar estos cambios")

	return nil
}

func restoreBackup(backupName string) error {
	fmt.Printf("♻️  Restaurando backup: %s\n", backupName)

	upg := upgrader.NewUpgrader(version.Version{}, version.Version{})

	if err := upg.RestoreBackup(backupName); err != nil {
		return fmt.Errorf("error al restaurar backup: %w", err)
	}

	fmt.Println("✅ Backup restaurado exitosamente!")
	fmt.Println("\n📝 Recomendaciones:")
	fmt.Println("   1. Verifica los archivos restaurados")
	fmt.Println("   2. Ejecuta: go mod tidy")
	fmt.Println("   3. Compila tu proyecto: go build ./cmd/...")

	return nil
}
