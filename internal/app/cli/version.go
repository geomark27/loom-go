package cli

import (
	"fmt"

	"github.com/geomark27/loom-go/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Muestra la versión de Loom CLI y del proyecto actual",
	Long: `Muestra información de versiones:
  - Versión del CLI de Loom
  - Versión del proyecto actual (si está en un proyecto Loom)
  - Estado de actualización`,
	RunE: runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string) error {
	fmt.Printf("🔧 Loom CLI v%s\n", version.Current.String())

	// Intentar detectar versión del proyecto
	projectVersion, err := version.DetectProjectVersion()
	if err == nil {
		fmt.Printf("📦 Proyecto actual: v%s\n", projectVersion.String())

		// Comparar versiones
		if projectVersion.Compare(version.Current) < 0 {
			fmt.Printf("\n⚠️  Tu proyecto usa una versión antigua de Loom\n")
			fmt.Printf("💡 Actualiza con: loom upgrade\n")
		} else if projectVersion.Compare(version.Current) == 0 {
			fmt.Printf("✅ Tu proyecto está actualizado\n")
		} else {
			fmt.Printf("⚠️  Tu proyecto usa una versión más nueva que el CLI\n")
			fmt.Printf("💡 Actualiza el CLI de Loom\n")
		}
	} else {
		fmt.Printf("\nℹ️  No se detectó un proyecto Loom en el directorio actual\n")
	}

	return nil
}
