package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loom",
	Short: "Loom - El tejedor de proyectos Go",
	Long: `Loom es una herramienta CLI que automatiza la creación y extensión 
de proyectos en Go siguiendo las mejores prácticas de la comunidad.

Loom no es un framework en tiempo de ejecución, sino un generador 
de código y un orquestador de la arquitectura del proyecto.`,
	Version: "1.0.3",
}

// Execute ejecuta el comando raíz
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Aquí se pueden agregar flags globales
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Salida detallada")
}
